package hapifhirgo

import (
	"context"
	"fmt"

	"github.com/savannahghi/hapi-fhir-go/models"
)

func (c *Client) CreateFHIRTask(ctx context.Context, input *models.Task) (*models.Task, error) {
	payload, err := structToMap(input)
	if err != nil {
		return nil, fmt.Errorf("unable to turn %s input into a map: %w", taskResourceType, err)
	}

	resource := &models.Task{}

	err = c.CreateFHIRResource(ctx, taskResourceType, payload, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to create %s resource: %w", taskResourceType, err)
	}

	return resource, nil
}

func (c *Client) GetFHIRTask(ctx context.Context, id string) (*models.Task, error) {
	resource := &models.Task{}

	err := c.GetFHIRResource(ctx, taskResourceType, id, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to get %s with ID %s, err: %w", taskResourceType, id, err)
	}

	return resource, nil
}

func (c *Client) SearchFHIRTask(ctx context.Context, searchParams map[string]interface{}) (*models.Bundle, error) {
	response, err := c.SearchFHIRResource(ctx, taskResourceType, searchParams)
	if err != nil {
		return nil, err
	}

	return response, nil
}
