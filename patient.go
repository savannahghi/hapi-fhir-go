package hapifhirgo

import (
	"context"
	"fmt"

	"github.com/savannahghi/hapi-fhir-go/models"
)

// CreateFHIRPatient creates a new FHIR Patient resource.
func (c *Client) CreateFHIRPatient(ctx context.Context, input *models.FHIRPatient) (*models.FHIRPatient, error) {
	payload, err := structToMap(input)
	if err != nil {
		return nil, fmt.Errorf("unable to turn %s input into a map: %w", patientResourceType, err)
	}

	resource := &models.FHIRPatient{}

	err = c.createFHIRResource(ctx, patientResourceType, payload, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to create %s resource: %w", patientResourceType, err)
	}

	return resource, nil
}

// GetFHIRPatient retrieves instances of FHIRPatient by ID.
func (c *Client) GetFHIRPatient(ctx context.Context, id string) (*models.FHIRPatient, error) {
	resource := &models.FHIRPatient{}

	err := c.getFHIRResource(ctx, patientResourceType, id, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to get %s with ID %s, err: %w", patientResourceType, id, err)
	}

	return resource, nil
}

// GetFHIRPatientEverything is used to retrieve all patient related information.
func (c *Client) GetFHIRPatientAllData(ctx context.Context, id string, params map[string]interface{}) (*models.Bundle, error) {
	response, err := c.getPatientEverything(ctx, id, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// SearchFHIRPatient searches for a FHIR patient.
func (c *Client) SearchFHIRPatient(ctx context.Context, searchParams map[string]interface{}) (*models.Bundle, error) {
	response, err := c.searchFHIRResource(ctx, patientResourceType, searchParams)
	if err != nil {
		return nil, err
	}

	return response, nil
}
