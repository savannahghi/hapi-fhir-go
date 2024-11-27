package hapifhirgo

import (
	"context"
	"fmt"

	"github.com/savannahghi/hapi-fhir-go/models"
)

func (c *Client) CreateFHIRServiceRequest(ctx context.Context, input *models.FHIRServiceRequest) (*models.FHIRServiceRequest, error) {
	payload, err := structToMap(input)
	if err != nil {
		return nil, fmt.Errorf("unable to turn %s input into a map: %w", serviceRequestResourceType, err)
	}

	resource := &models.FHIRServiceRequest{}

	err = c.CreateFHIRResource(ctx, serviceRequestResourceType, payload, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to create %s resource: %w", serviceRequestResourceType, err)
	}

	return resource, nil
}

func (c *Client) GetFHIRServiceRequest(ctx context.Context, id string) (*models.FHIRServiceRequest, error) {
	resource := &models.FHIRServiceRequest{}

	err := c.GetFHIRResource(ctx, serviceRequestResourceType, id, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to get %s with ID %s, err: %w", serviceRequestResourceType, id, err)
	}

	return resource, nil
}

func (c *Client) SearchFHIRServiceRequest(ctx context.Context, searchParams map[string]interface{}) (*models.Bundle, error) {
	response, err := c.SearchFHIRResource(ctx, serviceRequestResourceType, searchParams)
	if err != nil {
		return nil, err
	}

	return response, nil
}
