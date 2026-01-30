package fhir430

import "encoding/json"

// Distance is documented here http://hl7.org/fhir/StructureDefinition/Distance
// Base StructureDefinition for Distance Type: A length - a value with a unit that is a physical distance.
type Distance struct {
	ID         *string             `json:"id,omitempty"`
	Extension  []Extension         `json:"extension,omitempty"`
	Value      *json.Number        `json:"value,omitempty"`
	Comparator *QuantityComparator `json:"comparator,omitempty"`
	Unit       *string             `json:"unit,omitempty"`
	System     *string             `json:"system,omitempty"`
	Code       *string             `json:"code,omitempty"`
}
