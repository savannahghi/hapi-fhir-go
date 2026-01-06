package hapifhirgo

// operations.go implements the core FHIR REST operations as defined in the FHIR specification.
// It provides methods for:
//   - Basic CRUD operations (Create, Read, Update, Delete)
//   - Search functionality with comprehensive parameter support
//   - Version management and history tracking
//   - Resource validation
//   - Metadata operations
//   - Extended operations ($everything, $validate, $expand, etc.)
//   - Batch/transaction processing

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/mailgun/errors"
)

// CreateFHIRResource creates a FHIR resource
// POST [base]/[type].
func (c *Client) CreateFHIRResource(ctx context.Context, resourceType string, payload map[string]interface{}, resource interface{}) error {
	payload["resourceType"] = resourceType
	payload["language"] = "en"

	err := c.ValidateResource(ctx, resourceType, payload)
	if err != nil {
		return err
	}

	err = c.makeRequest(ctx, http.MethodPost, resourceType, nil, payload, resource, false)
	if err != nil {
		return err
	}

	return nil
}

// deleteFHIRResource performs a logical delete on a resource instance.
func (c *Client) DeleteFHIRResource(ctx context.Context, resourceType, fhirResourceID string) error {
	deletePath := fmt.Sprintf("%v/%v", resourceType, fhirResourceID)

	err := c.makeRequest(ctx, http.MethodDelete, deletePath, nil, nil, nil, false)
	if err != nil {
		return err
	}

	return nil
}

// GetFHIRResource gets a FHIR resource by its ID.
func (c *Client) GetFHIRResource(ctx context.Context, resourceType, fhirResourceID string, resource interface{}) error {
	searchPath := fmt.Sprintf("%v/%v", resourceType, fhirResourceID)

	err := c.makeRequest(ctx, http.MethodGet, searchPath, nil, nil, resource, false)
	if err != nil {
		return err
	}

	return nil
}

// SearchFHIRResource is used to search for a FHIR resource based on certain parameters.
// bundleID has been added here for pagination purposes to avoid repeating FHIR http request logic.
func (c *Client) SearchFHIRResource(ctx context.Context, bundleID, resourceType string, params map[string]any, bundle interface{}) error {
	if bundleID != "" && resourceType != "" {
		return errors.Errorf("bundleID and resourceType cannot both be provided")
	}

	urlParams := convertMapToURLValues(params)

	var path string
	if resourceType != "" {
		path = resourceType + "/_search"
	} else if bundleID != "" {
		path = ""
	}

	err := c.makeRequest(ctx, http.MethodGet, path, urlParams, nil, bundle, false)
	if err != nil {
		return fmt.Errorf("unable to search: %w", err)
	}

	return nil
}

func convertMapToURLValues(params map[string]any) url.Values {
	urlParams := url.Values{}

	for k, v := range params {
		switch value := v.(type) {
		case string:
			urlParams.Add(k, value)
		case []string:
			for _, i := range value {
				urlParams.Add(k, i)
			}
		default:
			urlParams.Add(k, fmt.Sprint(value))
		}
	}

	return urlParams
}

// validateResource validates a FHIR resource against the server's validation rules.
// POST [base]/[type]/$validate.

func (c *Client) ValidateResource(ctx context.Context, resourceType string, payload map[string]interface{}) error {
	path := fmt.Sprintf("%s%s", resourceType, "/$validate")

	err := c.makeRequest(ctx, http.MethodPost, path, nil, payload, nil, false)
	if err != nil {
		return err
	}

	return nil
}

// getPatientEverything retrieves a patient's complete record including all related resources.
// GET [base]/Patient/[id]/$everything
//
// Returns:
//   - Bundle: Contains all resources related to the patient including:
//   - Demographics
//   - Encounters
//   - Conditions
//   - Observations
//   - Medications
//   - Related resources etc.
//
// This operation returns a patient's entire medical record
// in a single request.
func (c *Client) GetPatientEverything(ctx context.Context, patientFhirID string, searchParams map[string]interface{}, bundle interface{}) error {
	path := fmt.Sprintf("Patient/%v/$everything", patientFhirID)

	urlParams := convertMapToURLValues(searchParams)

	err := c.makeRequest(ctx, http.MethodGet, path, urlParams, nil, &bundle, false)
	if err != nil {
		return fmt.Errorf("unable to search: %w", err)
	}

	return nil
}

// getEncounterEverything retrieves all the resources associated to an encounter.
func (c *Client) GetEncounterEverything(ctx context.Context, encounterID string, searchParams map[string]interface{}, bundle interface{}) error {
	path := fmt.Sprintf("Encounter/%v/$everything", encounterID)

	urlParams := convertMapToURLValues(searchParams)

	err := c.makeRequest(ctx, http.MethodGet, path, urlParams, nil, &bundle, false)
	if err != nil {
		return fmt.Errorf("unable to search: %w", err)
	}

	return nil
}

// PatchFHIRResource updates an instance of a fhir resource.
func (c *Client) PatchFHIRResource(ctx context.Context, resourceType string, resourceID string, payload interface{}, resource interface{}) error {
	updatePath := fmt.Sprintf("%v/%v", resourceType, resourceID)

	err := c.makeRequest(ctx, http.MethodPatch, updatePath, nil, payload, &resource, false)
	if err != nil {
		return fmt.Errorf("unable to patch resource: %w", err)
	}

	return nil
}

// FHIRPathPatch implements a syntax-agnostic patch mechanism where elements to be manipulated by the patch interaction are described
// using their FHIRpath (https://www.hl7.org/fhir/fhirpath.html) names and navigation.
// This specification is published at http://hl7.org/fhirpath icon in order to support wider re-use across multiple specifications.
func (c *Client) FHIRPathPatch(ctx context.Context, resourceType string, resourceID string, payload map[string]interface{}, resource interface{}) error {
	patches := []map[string]interface{}{}

	for key, value := range payload {
		if !reflect.ValueOf(value).IsZero() {
			patches = append(
				patches,
				map[string]interface{}{
					"op":    "replace",
					"path":  fmt.Sprintf("/%s", key), //nolint:perfsprint
					"value": value,
				},
			)
		}
	}

	fhirResource := fmt.Sprintf("%s/%s", resourceType, resourceID)

	err := c.makeRequest(ctx, http.MethodPatch, fhirResource, nil, patches, resource, false)
	if err != nil {
		return fmt.Errorf("unable to patch resource: %w", err)
	}

	return nil
}

func (c *Client) PostFHIRBundle(ctx context.Context, payload interface{}, response interface{}) error {
	err := c.makeRequest(ctx, http.MethodPost, "", nil, payload, &response, false)
	if err != nil {
		return fmt.Errorf("failed to post bundle entry: %w", err)
	}

	return nil
}

func (c *Client) PutFHIRResource(ctx context.Context, resourceType, resourceID string, payload map[string]interface{}, resource interface{}, useCREnabledServer bool) error {
	updatePath := fmt.Sprintf("%v/%v", resourceType, resourceID)

	err := c.makeRequest(ctx, http.MethodPut, updatePath, nil, payload, &resource, useCREnabledServer)
	if err != nil {
		return fmt.Errorf("unable to put resource: %w", err)
	}

	return nil
}

// ExtractFHIRResource extract data from a questionnaire response and returns a bundle resource type that is used to create various FHIR resources
// such as Conditions, Observations, etc.
func (c *Client) ExtractFHIRResource(ctx context.Context, resourceType string, payload map[string]interface{}, resource interface{}) error {
	extractionPath := fmt.Sprintf("%v/%s", resourceType, "$extract")

	err := c.makeRequest(ctx, http.MethodPost, extractionPath, nil, payload, &resource, true)
	if err != nil {
		return fmt.Errorf("unable to extract resource %s with err %w", resourceType, err)
	}

	return nil
}

// ExpandValueSet expands a ValueSet resource to produce a simple collection of enumerated codes.
// The operation can be called at the instance level (with valueSetID) or at the type level (with url or valueSet resource).
//
// Parameters:
//   - valueSetID: Optional. If provided, expands the ValueSet at instance level: ValueSet/[id]/$expand
//   - params: Optional map of expansion parameters. Common parameters include:
//   - url: Canonical reference to a ValueSet (required if valueSetID is not provided and valueSet is nil)
//   - context: Context of the value set for resolution
//   - filter: Filter to apply to the expansion
//   - displayLanguage: Language for display names
//   - count: Number of codes to return (for paging)
//   - offset: Offset for paging
//   - activeOnly: Whether to include only active codes
//   - excludeNested: Whether to exclude nested codes
//   - excludeNotForUI: Whether to exclude codes not intended for UI
//   - excludePostCoordinated: Whether to exclude post-coordinated codes
//   - includeDesignations: Whether to include designations
//   - includeSystem: Whether to include the system in the expansion
//   - limitedExpansion: Whether to limit the expansion
//   - profile: Profile to use for expansion
//   - date: Date for the expansion
//   - valueSet: Optional. The ValueSet resource to expand (used for POST requests when valueSetID is not provided)
//   - expandedValueSet: The expanded ValueSet will be unmarshaled into this parameter
//
// The operation supports both GET and POST methods:
//   - GET: ValueSet/$expand?url=... or ValueSet/[id]/$expand
//   - POST: ValueSet/$expand (with valueSet in body) or ValueSet/[id]/$expand
//
// Usage Examples:
//
// Case 1: Instance-level expansion (ValueSet stored on server)
//
//	// Expand a ValueSet by its ID with optional parameters
//	var expandedValueSet r4b.ValueSet
//	err := client.ExpandValueSet(ctx, "my-valueset-id", map[string]interface{}{
//	    "filter": "example",
//	    "count": 100,
//	}, nil, &expandedValueSet)
//
// Case 2a: Type-level expansion by URL (GET request)
//
//	// Expand a ValueSet by its canonical URL
//	var expandedValueSet r4b.ValueSet
//	err := client.ExpandValueSet(ctx, "", map[string]interface{}{
//	    "url": "http://hl7.org/fhir/ValueSet/example",
//	    "filter": "test",
//	    "displayLanguage": "en",
//	    "count": 50,
//	    "offset": 0,
//	}, nil, &expandedValueSet)
//
// Case 2b: Type-level expansion by POST (ValueSet not stored on server)
//
//	// Expand a ValueSet resource provided directly in the request
//	myValueSet := r4b.ValueSet{
//	    // ... ValueSet definition ...
//	}
//	var expandedValueSet r4b.ValueSet
//	err := client.ExpandValueSet(ctx, "", nil, myValueSet, &expandedValueSet)
//
// Returns an error if the expansion fails.
func (c *Client) ExpandValueSet(
	ctx context.Context,
	valueSetID string,
	params map[string]interface{},
	valueSet interface{},
	expandedValueSet interface{},
) error {
	var path string
	var method string
	var payload interface{}
	var urlParams url.Values

	// Determine the path based on whether valueSetID is provided
	if valueSetID != "" {
		path = fmt.Sprintf("ValueSet/%s/$expand", valueSetID)
	} else {
		path = "ValueSet/$expand"
	}

	// Convert params to URL values for GET requests
	urlParams = convertMapToURLValues(params)

	// Determine method and payload
	// If valueSet is provided and no valueSetID, use POST with valueSet in body
	// Otherwise, use GET with parameters in URL
	if valueSet != nil && valueSetID == "" {
		method = http.MethodPost
		// For POST, we need to create a Parameters resource or include valueSet in the payload
		// HAPI FHIR accepts the valueSet directly in POST body for $expand
		payload = valueSet
	} else {
		method = http.MethodGet
		payload = nil

		if len(params) > 0 && urlParams == nil {
			urlParams = convertMapToURLValues(params)
		}
	}

	err := c.makeRequest(ctx, method, path, urlParams, payload, expandedValueSet, false)
	if err != nil {
		return fmt.Errorf("unable to expand value set: %w", err)
	}

	return nil
}
