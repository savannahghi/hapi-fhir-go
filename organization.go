package hapifhirgo

import (
	"context"
	"fmt"

	"github.com/savannahghi/hapi-fhir-go/models"
)

func (c *Client) CreateFHIROrganization(ctx context.Context, input *models.FHIROrganization) (*models.FHIROrganization, error) {
	payload, err := structToMap(input)
	if err != nil {
		return nil, fmt.Errorf("unable to turn %s input into a map: %w", organizationResource, err)
	}

	resource := &models.FHIROrganization{}

	err = c.CreateFHIRResource(ctx, organizationResource, payload, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to create %s resource: %w", organizationResource, err)
	}

	return resource, nil
}

func (c *Client) GetFHIROrganization(ctx context.Context, id string) (*models.FHIROrganization, error) {
	resource := &models.FHIROrganization{}

	err := c.GetFHIRResource(ctx, organizationResource, id, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to get %s with ID %s, err: %w", organizationResource, id, err)
	}

	return resource, nil
}

func (c *Client) SearchFHIROrganization(ctx context.Context, searchParams map[string]interface{}) (*models.Bundle, error) {
	response, err := c.SearchFHIRResource(ctx, organizationResource, searchParams)
	if err != nil {
		return nil, err
	}

	return response, nil
}
