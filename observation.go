package hapifhirgo

import (
	"context"
	"fmt"

	"github.com/savannahghi/hapi-fhir-go/models"
)

func (c *Client) CreateFHIRObservation(ctx context.Context, input *models.FHIRObservation) (*models.FHIRObservation, error) {
	payload, err := structToMap(input)
	if err != nil {
		return nil, fmt.Errorf("unable to turn %s input into a map: %w", observationResourceType, err)
	}

	resource := &models.FHIRObservation{}

	err = c.createFHIRResource(ctx, observationResourceType, payload, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to create %s resource: %w", observationResourceType, err)
	}

	return resource, nil
}

func (c *Client) GetFHIRObservation(ctx context.Context, id string) (*models.FHIRObservation, error) {
	resource := &models.FHIRObservation{}

	err := c.getFHIRResource(ctx, observationResourceType, id, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to get %s with ID %s, err: %w", observationResourceType, id, err)
	}

	return resource, nil
}

func (c *Client) SearchFHIRObservation(ctx context.Context, searchParams map[string]interface{}) (*models.Bundle, error) {
	response, err := c.searchFHIRResource(ctx, observationResourceType, searchParams)
	if err != nil {
		return nil, err
	}

	return response, nil
}
