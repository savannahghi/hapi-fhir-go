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
	"net/http"
)

// CreateFHIRResource creates a FHIR resource
// POST [base]/[type].
func (c *Client) createFHIRResource(ctx context.Context, resourceType string, payload map[string]interface{}, resource interface{}) error {
	payload["resourceType"] = resourceType
	payload["language"] = "EN"

	err := c.makeRequest(ctx, http.MethodPost, resourceType, nil, payload, resource)
	if err != nil {
		return err
	}

	return nil
}
