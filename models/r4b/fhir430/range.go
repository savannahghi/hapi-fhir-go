
package fhir430
// Range is documented here http://hl7.org/fhir/StructureDefinition/Range
// Base StructureDefinition for Range Type: A set of ordered Quantities defined by a low and high limit.
type Range struct {
	ID        *string     `json:"ID,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Low       *Quantity   `json:"low,omitempty"`
	High      *Quantity   `json:"high,omitempty"`
}
