
package fhir430
// Coding is documented here http://hl7.org/fhir/StructureDefinition/Coding
// Base StructureDefinition for Coding Type: A reference to a code defined by a terminology system.
type Coding struct {
	ID           *string     `json:"ID,omitempty"`
	Extension    []Extension `json:"extension,omitempty"`
	System       *string     `json:"system,omitempty"`
	Version      *string     `json:"version,omitempty"`
	Code         *string     `json:"code,omitempty"`
	Display      *string     `json:"display,omitempty"`
	UserSelected *bool       `json:"userSelected,omitempty"`
}
