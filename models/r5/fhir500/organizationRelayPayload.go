package fhir500

// OrganizationRelayPayload is used to return single instances of Organization.
type OrganizationRelayPayload struct {
	Resource *Organization `json:"resource,omitempty"`
}
