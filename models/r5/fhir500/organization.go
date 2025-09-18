package fhir500

type Organization struct {
	ID         *string            `json:"id,omitempty"`
	Active     *bool              `json:"active,omitempty"`
	Identifier []*Identifier      `json:"identifier,omitempty"`
	Type       []*CodeableConcept `json:"type,omitempty"`
	Name       *string            `json:"name,omitempty"`
	Alias      []string           `json:"alias,omitempty"`
	Telecom    []*ContactPoint    `json:"telecom,omitempty"`
	Address    []*Address         `json:"address,omitempty"`
}

type OrganizationRelayPayload struct {
	Resource *Organization `json:"resource,omitempty"`
}
