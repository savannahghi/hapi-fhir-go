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

	"github.com/savannahghi/hapi-fhir-go/models"
)

// APIError represents a FHIR specific error with operation outcome.
type APIError struct {
	StatusCode       int                      `json:"statusCode,omitempty"`
	OperationOutcome *models.OperationOutcome `json:"operationOutcome,omitempty"`
}

func (a APIError) Error() string {
	if a.OperationOutcome != nil && len(a.OperationOutcome.Issue) > 0 {
		return a.OperationOutcome.ErrorLogging()
	}

	return fmt.Sprintf("FHIR error (HTTP %d)", a.StatusCode)
}

func (c *Client) newRequest(
	ctx context.Context,
	method, path string,
	params url.Values,
	data interface{},
) (*http.Request, error) {
	url, err := c.composeRequestURL(path, params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequestWithContext(ctx, method, url, http.NoBody)
	if err != nil {
		return nil, err
	}

	c.setHeaders(request)

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

func (c *Client) composeRequestURL(path string, params url.Values) (string, error) {
	u, err := url.Parse(c.baseURL + "/" + path)
	if err != nil {
		return "", errors.New("url parse: " + err.Error())
	}

	u.RawQuery = params.Encode()

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
		var outcome models.OperationOutcome

		err = json.Unmarshal(respBytes, &outcome)
		if err != nil {
			return err
		}

		return APIError{
			StatusCode:       response.StatusCode,
			OperationOutcome: &outcome,
		}
	}

	// A Specific case for validation responses.
	// Validation of the resource is considered valid only when the severity is either "success" or "information"
	// All validation responses, whether valid or invalid, returns a http status code of 200 https://www.hl7.org/fhir/resource-operation-validate.html
	if isValidateInPath(path) && response.StatusCode == http.StatusOK {
		return handleValidationResponse(respBytes, response.StatusCode)
	}

	err = json.Unmarshal(respBytes, result)
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
) error {
	request, err := c.newRequest(ctx, method, path, params, data)
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
	var results []models.OperationOutcomeIssue

	var outCome models.OperationOutcome

	err := json.Unmarshal(resBytes, &outCome)
	if err != nil {
		return err
	}

	for _, item := range outCome.Issue {
		if !isValidSeverity(string(item.Severity)) {
			results = append(results, item)
		}
	}

	if len(results) > 0 {
		outCome.Issue = results

		return APIError{
			StatusCode:       statusCode,
			OperationOutcome: &outCome,
		}
	}

	return nil
}
