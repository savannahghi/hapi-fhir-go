package models

type FHIROrganization struct {
	ID         *string                `json:"id,omitempty"`
	Active     *bool                  `json:"active,omitempty"`
	Identifier []*FHIRIdentifier      `json:"identifier,omitempty"`
	Type       []*FHIRCodeableConcept `json:"type,omitempty"`
	Name       *string                `json:"name,omitempty"`
	Alias      []string               `json:"alias,omitempty"`
	Telecom    []*FHIRContactPoint    `json:"telecom,omitempty"`
	Address    []*FHIRAddress         `json:"address,omitempty"`
}

type FHIROrganizationRelayPayload struct {
	Resource *FHIROrganization `json:"resource,omitempty"`
}
