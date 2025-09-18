
package fhir430
// Period is documented here http://hl7.org/fhir/StructureDefinition/Period
// Base StructureDefinition for Period Type: A time period defined by a start and end date and optionally time.
type Period struct {
	ID        *string     `json:"ID,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Start     *string     `json:"start,omitempty"`
	End       *string     `json:"end,omitempty"`
}
