package hapifhirgo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	// fhirContentType is the media type for FHIR resource bodies.
	fhirContentType = "application/fhir+json"

	// jsonPatchContentType is the media type HAPI dispatches on to read a request
	// body as a JSON Patch (RFC 6902) document rather than a FHIR resource.
	jsonPatchContentType = "application/json-patch+json"

	// maxErrorBody bounds how much of an error answer is kept on the APIError.
	maxErrorBody = 1 << 20
)

// Issue is one issue of the OperationOutcome a server answers an error with.
type Issue struct {
	Severity    string       `json:"severity,omitempty"`
	Code        string       `json:"code,omitempty"`
	Diagnostics string       `json:"diagnostics,omitempty"`
	Details     IssueDetails `json:"details,omitempty"`
	Location    []string     `json:"location,omitempty"`
	Expression  []string     `json:"expression,omitempty"`
}

// IssueDetails is the human-readable part of an issue's details.
type IssueDetails struct {
	Text string `json:"text,omitempty"`
}

// APIError is any answer of 400 or above. When the server answered an OperationOutcome, as a
// FHIR server does, OperationOutcome holds it decoded and Issues holds its issues parsed. When
// it answered something else, such as a gateway's HTML page, both are empty and Body holds what
// came back, so the status is always there to act on.
type APIError struct {
	StatusCode       int         `json:"statusCode,omitempty"`
	OperationOutcome interface{} `json:"operationOutcome,omitempty"`
	Issues           []Issue     `json:"issues,omitempty"`
	Body             []byte      `json:"-"`
}

// Diagnostics joins the text of the error and fatal issues, which is what a caller shows or
// logs. It is empty when the server answered something other than an OperationOutcome.
func (a APIError) Diagnostics() string {
	var parts []string

	for _, issue := range a.Issues {
		if issue.Severity != "error" && issue.Severity != "fatal" {
			continue
		}

		text := strings.TrimSpace(issue.Diagnostics)
		if text == "" {
			text = strings.TrimSpace(issue.Details.Text)
		}

		if text != "" {
			parts = append(parts, text)
		}
	}

	return strings.Join(parts, "; ")
}

func (a APIError) Error() string {
	if a.OperationOutcome == nil {
		return fmt.Sprintf("FHIR error (HTTP %d)", a.StatusCode)
	}

	outcomeStr := a.formatOperationOutcome()
	if outcomeStr != "" {
		return fmt.Sprintf("FHIR error (HTTP %d): %s", a.StatusCode, outcomeStr)
	}

	outcomeJSON, err := json.Marshal(a.OperationOutcome)
	if err != nil {
		return fmt.Sprintf("FHIR error (HTTP %d): unable to format OperationOutcome", a.StatusCode)
	}

	if len(outcomeJSON) > 0 {
		return fmt.Sprintf("FHIR error (HTTP %d): %s", a.StatusCode, string(outcomeJSON))
	}

	return fmt.Sprintf("FHIR error (HTTP %d)", a.StatusCode)
}

// formatOperationOutcome formats the OperationOutcome into a human-readable string.
func (a APIError) formatOperationOutcome() string {
	if a.OperationOutcome == nil {
		return ""
	}

	outcomeMap, ok := a.OperationOutcome.(map[string]interface{})
	if !ok {
		return ""
	}

	issuesRaw, ok := outcomeMap["issue"].([]interface{})
	if !ok {
		return ""
	}

	if len(issuesRaw) == 0 {
		return ""
	}

	var issues []string
	for _, issueRaw := range issuesRaw {
		if issueRaw == nil {
			continue
		}

		issue, ok := issueRaw.(map[string]interface{})
		if !ok {
			continue
		}

		var parts []string

		severity, ok := issue["severity"].(string)
		if ok && severity != "" {
			parts = append(parts, fmt.Sprintf("severity: %s", severity))
		}

		code, ok := issue["code"].(string)
		if ok && code != "" {
			parts = append(parts, fmt.Sprintf("code: %s", code))
		}

		details, ok := issue["details"].(map[string]interface{})
		if ok && details != nil {
			text, ok := details["text"].(string)
			if ok && text != "" {
				parts = append(parts, fmt.Sprintf("details: %s", text))
			}
		}

		diagnostics, ok := issue["diagnostics"].(string)
		if ok && diagnostics != "" {
			parts = append(parts, fmt.Sprintf("diagnostics: %s", diagnostics))
		}

		location, ok := issue["location"].([]interface{})
		if ok && len(location) > 0 {
			var locs []string
			for _, loc := range location {
				if loc == nil {
					continue
				}
				locStr, ok := loc.(string)
				if ok && locStr != "" {
					locs = append(locs, locStr)
				}
			}
			if len(locs) > 0 {
				parts = append(parts, fmt.Sprintf("location: %s", strings.Join(locs, ", ")))
			}
		}

		if len(parts) > 0 {
			issues = append(issues, strings.Join(parts, "; "))
		}
	}

	if len(issues) == 0 {
		return ""
	}

	return strings.Join(issues, " | ")
}

// GetOperationOutcome returns the OperationOutcome as a map for programmatic access.
func (a APIError) GetOperationOutcome() map[string]interface{} {
	if outcomeMap, ok := a.OperationOutcome.(map[string]interface{}); ok {
		return outcomeMap
	}
	return nil
}

func (c *Client) newRequest(
	ctx context.Context,
	method, path string,
	params url.Values,
	data interface{},
	useCREnabledServer bool,
) (*http.Request, error) {
	return c.buildRequest(ctx, method, path, params, data, useCREnabledServer, true)
}

// buildRequest is newRequest with a say over credentials: the capability statement is read
// without them, since a server publishes it to anyone.
func (c *Client) buildRequest(
	ctx context.Context,
	method, path string,
	params url.Values,
	data interface{},
	useCREnabledServer bool,
	withAuth bool,
) (*http.Request, error) {

	reqUrl, err := c.composeRequestURL(path, params, useCREnabledServer)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, method, reqUrl, http.NoBody)
	if err != nil {
		return nil, err
	}

	if withAuth {
		if err := c.applyAuth(request); err != nil {
			return nil, err
		}
	}

	c.setHeaders(request)

	switch payload := data.(type) {
	case nil:
		request.Body = nil
	case io.ReadCloser:
		// Caller owns lifecycle; we cannot replay this on retry.
		request.Body = payload
	case io.Reader:
		request.Body = io.NopCloser(payload)
	default:
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		request.Body = io.NopCloser(bytes.NewReader(b))
		request.ContentLength = int64(len(b))
		request.GetBody = func() (io.ReadCloser, error) {
			return io.NopCloser(bytes.NewReader(b)), nil
		}
	}

	return request, nil
}

// applyAuth sets the request's Authorization. A configured token provider takes
// precedence over basic auth: it is called with the request context so it can
// relay a per-request token or hand back a cached service-account one.
func (c *Client) applyAuth(r *http.Request) error {
	if c.tokenProvider != nil {
		token, err := c.tokenProvider(r.Context())
		if err != nil {
			return fmt.Errorf("hapifhirgo: token provider failed: %w", err)
		}
		if token != "" {
			r.Header.Set("Authorization", "Bearer "+token)
		}
		return nil
	}

	if c.authCreds != nil {
		r.SetBasicAuth(c.authCreds.username, c.authCreds.password)
	}

	return nil
}

func (c *Client) setHeaders(r *http.Request) {
	r.Header.Set("Content-Type", fhirContentType)
	r.Header.Set("Accept", fhirContentType)
	if !c.omitCacheControl {
		r.Header.Set("Cache-Control", "no-cache")
	}

	for k, v := range c.defaultHeaders {
		r.Header.Set(k, v)
	}
}

func (c *Client) composeRequestURL(path string, params url.Values, useCREnabledServer bool) (string, error) {
	baseURL := c.baseURL

	if useCREnabledServer {
		baseURL = c.CREnabledHAPIFHIRBaseURL
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	if path != "" {
		u.Path, err = url.JoinPath(u.Path, path)
		if err != nil {
			return "", err
		}
	}

	q := u.Query()

	for k, vs := range params {
		for _, v := range vs {
			q.Add(k, v)
		}
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}

func (c *Client) readResponse(response *http.Response, path string, result interface{}) error {
	if response.Body == nil {
		return errors.New("response body is nil")
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return apiError(response)
	}

	if isValidateInPath(path) && response.StatusCode == http.StatusOK {
		respBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return handleValidationResponse(respBytes, response.StatusCode)
	}

	// Happy path: stream-decode straight into the destination. Saves a
	// full copy of the response body for every large Bundle/$everything.
	if result == nil {
		// Drain to allow conn reuse even when caller doesn't want the body.
		_, _ = io.Copy(io.Discard, response.Body)
		return nil
	}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to unmarshall body: %w", err)
	}

	return nil
}

func (c *Client) makeRequest(
	ctx context.Context,
	method, path string,
	params url.Values,
	data, result interface{},
	useCREnabledServer bool,
) error {
	request, err := c.newRequest(ctx, method, path, params, data, useCREnabledServer)
	if err != nil {
		return err
	}

	resp, err := c.HTTP.Do(request)
	if err != nil {
		return err
	}

	return c.readResponse(resp, path, result)
}

// apiError turns an error answer into an APIError, whatever the body is.
func apiError(response *http.Response) error {
	body, err := io.ReadAll(io.LimitReader(response.Body, maxErrorBody))
	if err != nil {
		return fmt.Errorf("failed to read the error body (HTTP %d): %w", response.StatusCode, err)
	}

	apiErr := APIError{StatusCode: response.StatusCode, Body: body}

	var outcome map[string]interface{}
	if json.Unmarshal(body, &outcome) != nil || outcome == nil {
		return apiErr
	}

	apiErr.OperationOutcome = outcome

	var parsed struct {
		Issue []Issue `json:"issue"`
	}

	// The map already decoded, so the only way this fails is an issue whose shape is not an
	// issue's, which leaves Issues empty and OperationOutcome still there.
	_ = json.Unmarshal(body, &parsed)
	apiErr.Issues = parsed.Issue

	return apiErr
}

// isValidSeverity returns true if the severity does not indicate a failure.
// Only "error" and "fatal" severities cause validation to fail.
// "warning", "information", and "success" are considered non-failing.
func isValidSeverity(severity string) bool {
	return severity == "success" || severity == "information" || severity == "warning"
}

/*
isValidateInPath checks whether the request is meant for validating a resource.
This check allows ValidateResource to share the readResponse, inside makeRequest, without confusion.

Note: Validation responses from HAPI FHIR, whether failed or successful, always returns http status code of 200.
This is counterintuitive considering that makeRequest is also used for fetching resources from HAPI FHIR server
and can also return a http status code of 200 if successful.
*/
func isValidateInPath(path string) bool {
	return strings.Contains(path, "$validate")
}

// handleValidationResponse is helper function that handles validation outcome response.
func handleValidationResponse(resBytes []byte, statusCode int) error {
	if len(resBytes) == 0 {
		return fmt.Errorf("empty validation response body")
	}

	var outCome map[string]interface{}

	err := json.Unmarshal(resBytes, &outCome)
	if err != nil {
		return fmt.Errorf("failed to unmarshal validation OperationOutcome: %w", err)
	}

	if outCome == nil {
		return nil
	}

	// Check if there are any issues with severity other than success/information
	issues, ok := outCome["issue"].([]interface{})
	if !ok {
		return nil
	}

	if len(issues) == 0 {
		return nil
	}

	var results []interface{}
	for _, issue := range issues {
		if issue == nil {
			continue
		}

		issueMap, ok := issue.(map[string]interface{})
		if !ok {
			continue
		}

		severity, ok := issueMap["severity"].(string)
		if !ok {
			continue
		}

		if !isValidSeverity(severity) {
			results = append(results, issue)
		}
	}

	if len(results) > 0 {
		outCome["issue"] = results
		return APIError{
			StatusCode:       statusCode,
			OperationOutcome: outCome,
		}
	}

	return nil
}
