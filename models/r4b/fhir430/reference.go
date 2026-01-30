package fhir430

// Reference is documented here http://hl7.org/fhir/StructureDefinition/Reference
// Base StructureDefinition for Reference Type: A reference from one resource to another.
type Reference struct {
	ID         *string     `json:"id,omitempty"`
	Extension  []Extension `json:"extension,omitempty"`
	Reference  *string     `json:"reference,omitempty"`
	Type       *string     `json:"type,omitempty"`
	Identifier *Identifier `json:"identifier,omitempty"`
	Display    *string     `json:"display,omitempty"`
}
