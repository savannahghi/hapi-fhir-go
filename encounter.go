package hapifhirgo

import (
	"context"
	"fmt"

	"github.com/savannahghi/hapi-fhir-go/models"
)

func (c *Client) CreateFHIREncounter(ctx context.Context, input *models.FHIREncounter) (*models.FHIREncounter, error) {
	payload, err := structToMap(input)
	if err != nil {
		return nil, fmt.Errorf("unable to turn %s input into a map: %w", encounterResourceType, err)
	}

	resource := &models.FHIREncounter{}

	err = c.createFHIRResource(ctx, encounterResourceType, payload, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to create %s resource: %w", encounterResourceType, err)
	}

	return resource, nil
}

func (c *Client) GetFHIREncounter(ctx context.Context, id string) (*models.FHIREncounter, error) {
	resource := &models.FHIREncounter{}

	err := c.getFHIRResource(ctx, encounterResourceType, id, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to get %s with ID %s, err: %w", encounterResourceType, id, err)
	}

	return resource, nil
}

func (c *Client) GetFHIREncounterAllData(ctx context.Context, id string, params map[string]interface{}) (*models.Bundle, error) {
	response, err := c.getEncounterEverything(ctx, id, params)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) SearchFHIREncounter(ctx context.Context, searchParams map[string]interface{}) (*models.Bundle, error) {
	response, err := c.searchFHIRResource(ctx, encounterResourceType, searchParams)
	if err != nil {
		return nil, err
	}

	return response, nil
}