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

	"github.com/savannahghi/hapi-fhir-go/models"
)

// CreateFHIRResource creates a FHIR resource
// POST [base]/[type].
func (c *Client) CreateFHIRResource(ctx context.Context, resourceType string, payload map[string]interface{}, resource interface{}) error {
	payload["resourceType"] = resourceType
	payload["language"] = "EN"

	err := c.makeRequest(ctx, http.MethodPost, resourceType, nil, payload, resource)
	if err != nil {
		return err
	}

	return nil
}

// deleteFHIRResource performs a logical delete on a resource instance.
func (c *Client) DeleteFHIRResource(ctx context.Context, resourceType, fhirResourceID string) error {
	deletePath := fmt.Sprintf("%v/%v", resourceType, fhirResourceID)

	err := c.makeRequest(ctx, http.MethodDelete, deletePath, nil, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetFHIRResource gets a FHIR resource by its ID.
func (c *Client) GetFHIRResource(ctx context.Context, resourceType, fhirResourceID string, resource interface{}) error {
	searchPath := fmt.Sprintf("%v/%v", resourceType, fhirResourceID)

	err := c.makeRequest(ctx, http.MethodGet, searchPath, nil, nil, resource)
	if err != nil {
		return err
	}

	return nil
}

// SearchFHIRResource is used to search for a FHIR resource based on certain parameters.
func (c *Client) SearchFHIRResource(ctx context.Context, resourceType string, params map[string]interface{}) (*models.Bundle, error) {
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
			return nil, fmt.Errorf("the search/filter param: %s should all be sent as strings", k)
		}
	}

	path := fmt.Sprintf("%v/_search", resourceType)
	bundle := models.Bundle{}

	err := c.makeRequest(ctx, http.MethodGet, path, urlParams, nil, bundle)
	if err != nil {
		return nil, fmt.Errorf("unable to search: %w", err)
	}

	return &bundle, nil
}

// validateResource validates a FHIR resource against the server's validation rules.
// POST [base]/[type]/$validate.
func (c *Client) ValidateResource(ctx context.Context, resourceType string, payload map[string]interface{}) error {
	path := fmt.Sprintf("%v/$validate", resourceType)

	err := c.makeRequest(ctx, http.MethodPost, path, nil, payload, nil)
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
func (c *Client) GetPatientEverything(ctx context.Context, patientFhirID string, searchParams map[string]interface{}) (*models.Bundle, error) {
	path := fmt.Sprintf("Patient/%v/$everything", patientFhirID)
	bundle := models.Bundle{}

	err := c.makeRequest(ctx, http.MethodGet, path, nil, searchParams, bundle)
	if err != nil {
		return nil, fmt.Errorf("unable to search: %w", err)
	}

	return &bundle, nil
}

// getEncounterEverything retrieves all the resources associated to an encounter.
func (c *Client) GetEncounterEverything(ctx context.Context, encounterID string, searchParams map[string]interface{}) (*models.Bundle, error) {
	path := fmt.Sprintf("Encounter/%v/$everything", encounterID)
	bundle := models.Bundle{}

	err := c.makeRequest(ctx, http.MethodGet, path, nil, searchParams, bundle)
	if err != nil {
		return nil, fmt.Errorf("unable to search: %w", err)
	}

	return &bundle, nil
}
