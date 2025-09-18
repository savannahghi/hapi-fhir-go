
package fhir430
// Narrative is documented here http://hl7.org/fhir/StructureDefinition/Narrative
// Base StructureDefinition for Narrative Type: A human-readable summary of the resource conveying the essential clinical and business information for the resource.
type Narrative struct {
	ID        *string         `json:"ID,omitempty"`
	Extension []Extension     `json:"extension,omitempty"`
	Status    NarrativeStatus `json:"status"`
	Div       string          `json:"div"`
}
