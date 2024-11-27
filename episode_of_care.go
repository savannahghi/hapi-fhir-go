package hapifhirgo

import (
	"context"
	"fmt"

	"github.com/savannahghi/hapi-fhir-go/models"
)

func (c *Client) CreateFHIREpisodeOfCare(ctx context.Context, input *models.FHIREpisodeOfCare) (*models.FHIREpisodeOfCare, error) {
	payload, err := structToMap(input)
	if err != nil {
		return nil, fmt.Errorf("unable to turn %s input into a map: %w", episodeOfCareResourceType, err)
	}

	resource := &models.FHIREpisodeOfCare{}

	err = c.CreateFHIRResource(ctx, episodeOfCareResourceType, payload, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to create %s resource: %w", episodeOfCareResourceType, err)
	}

	return resource, nil
}

func (c *Client) GetFHIREpisodeOfCare(ctx context.Context, id string) (*models.FHIREpisodeOfCare, error) {
	resource := &models.FHIREpisodeOfCare{}

	err := c.GetFHIRResource(ctx, episodeOfCareResourceType, id, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to get %s with ID %s, err: %w", episodeOfCareResourceType, id, err)
	}

	return resource, nil
}

func (c *Client) SearchFHIREpisodeOfCare(ctx context.Context, searchParams map[string]interface{}) (*models.Bundle, error) {
	response, err := c.SearchFHIRResource(ctx, episodeOfCareResourceType, searchParams)
	if err != nil {
		return nil, err
	}

	return response, nil
}
