package fhir430

// ContactDetail is documented here http://hl7.org/fhir/StructureDefinition/ContactDetail
// Base StructureDefinition for ContactDetail Type: Specifies contact information for a person or organization.
type ContactDetail struct {
	ID        *string        `json:"id,omitempty"`
	Extension []Extension    `json:"extension,omitempty"`
	Name      *string        `json:"name,omitempty"`
	Telecom   []ContactPoint `json:"telecom,omitempty"`
}
