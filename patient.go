package hapifhirgo

import (
	"context"
	"fmt"

	"github.com/savannahghi/hapi-fhir-go/models"
)

const (
	patientResourceType = "Patient"
)

// CreateFHIRPatient creates a new FHIR Patient resource.
func (c *Client) CreateFHIRPatient(ctx context.Context, input *models.FHIRPatient) (*models.PatientPayload, error) {
	payload, err := structToMap(input)
	if err != nil {
		return nil, fmt.Errorf("unable to turn %s input into a map: %w", patientResourceType, err)
	}

	resource := &models.FHIRPatient{}

	err = c.createFHIRResource(ctx, patientResourceType, payload, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to create %s resource: %w", patientResourceType, err)
	}

	output := &models.PatientPayload{
		PatientRecord: resource,
	}

	return output, nil
}
