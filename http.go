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

// APIError represents a FHIR specific error with operation outcome.
type APIError struct {
	StatusCode       int         `json:"statusCode,omitempty"`
	OperationOutcome interface{} `json:"operationOutcome,omitempty"`
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

	reqUrl, err := c.composeRequestURL(path, params, useCREnabledServer)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, method, reqUrl, http.NoBody)
	if err != nil {
		return nil, err
	}

	if c.authCreds != nil {
		request.SetBasicAuth(c.authCreds.username, c.authCreds.password)
	}

	switch method {
	case http.MethodPatch:
		request.Header.Set("Content-Type", "application/json-patch+json")
	default:
		c.setHeaders(request)
	}

	switch payload := data.(type) {
	case nil:
		request.Body = nil
	case io.ReadCloser:
		request.Body = payload
	case io.Reader:
		request.Body = io.NopCloser(payload)
	default:
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		request.Body = io.NopCloser(bytes.NewReader(b))
	}

	return request, nil
}

func (c *Client) setHeaders(r *http.Request) {
	r.Header.Set("Content-Type", "application/fhir+json")
	r.Header.Set("Accept", "application/fhir+json")
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

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode >= 400 {
		var outcome map[string]interface{}

		err := json.Unmarshal(respBytes, &outcome)
		if err != nil {
			return fmt.Errorf("failed to unmarshal OperationOutcome (HTTP %d): %w", response.StatusCode, err)
		}

		return APIError{
			StatusCode:       response.StatusCode,
			OperationOutcome: outcome,
		}
	}

	// A Specific case for validation responses.
	// Validation of the resource is considered valid only when the severity is either "success" or "information"
	// All validation responses, whether valid or invalid, returns a http status code of 200 https://www.hl7.org/fhir/resource-operation-validate.html
	if isValidateInPath(path) && response.StatusCode == http.StatusOK {
		return handleValidationResponse(respBytes, response.StatusCode)
	}

	err = json.Unmarshal(respBytes, &result)
	if err != nil {
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

func isValidSeverity(severity string) bool {
	return severity == "success" || severity == "information"
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
