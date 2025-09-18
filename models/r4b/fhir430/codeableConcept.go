
package fhir430
// CodeableConcept is documented here http://hl7.org/fhir/StructureDefinition/CodeableConcept
// Base StructureDefinition for CodeableConcept Type: A concept that may be defined by a formal reference to a terminology or ontology or may be provided by text.
type CodeableConcept struct {
	ID        *string     `json:"ID,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Coding    []Coding    `json:"coding,omitempty"`
	Text      *string     `json:"text,omitempty"`
}
