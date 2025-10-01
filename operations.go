package hapifhirgo

// operations.go implements the core FHIR REST operations as defined in the FHIR specification.
// It provides methods for:
//   - Basic CRUD operations (Create, Read, Update, Delete)
//   - Search functionality with comprehensive parameter support
//   - Version management and history tracking
//   - Resource validation
//   - Metadata operations
//   - Extended operations ($everything, $validate, etc.)
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
	payload["language"] = "EN"

	err := c.makeRequest(ctx, http.MethodPost, resourceType, nil, payload, resource, false)
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

func (c *Client) PutFHIRResource(
	ctx context.Context, resourceType, resourceID string,
	payload map[string]interface{}, resource interface{}, useCREnabledServer bool,
) error {
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
