package hapifhirgo

import (
	"context"
	"fmt"

	"github.com/savannahghi/hapi-fhir-go/models"
)

func (c *Client) CreateFHIRConsent(ctx context.Context, input *models.Consent) (*models.Consent, error) {
	payload, err := structToMap(input)
	if err != nil {
		return nil, fmt.Errorf("unable to turn %s input into a map: %w", consentResourceType, err)
	}

	resource := &models.Consent{}

	err = c.CreateFHIRResource(ctx, consentResourceType, payload, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to create %s resource: %w", consentResourceType, err)
	}

	return resource, nil
}

func (c *Client) GetFHIRConsent(ctx context.Context, id string) (*models.Consent, error) {
	resource := &models.Consent{}

	err := c.GetFHIRResource(ctx, consentResourceType, id, resource)
	if err != nil {
		return nil, fmt.Errorf("unable to get %s with ID %s, err: %w", consentResourceType, id, err)
	}

	return resource, nil
}

func (c *Client) SearchFHIRConsent(ctx context.Context, searchParams map[string]interface{}) (*models.Bundle, error) {
	response, err := c.SearchFHIRResource(ctx, "", consentResourceType, searchParams)
	if err != nil {
		return nil, err
	}

	return response, nil
}
