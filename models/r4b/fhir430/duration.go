package fhir430

import "encoding/json"

// Duration is documented here http://hl7.org/fhir/StructureDefinition/Duration
// Base StructureDefinition for Duration Type: A length of time.
type Duration struct {
	ID         *string             `json:"id,omitempty"`
	Extension  []Extension         `json:"extension,omitempty"`
	Value      *json.Number        `json:"value,omitempty"`
	Comparator *QuantityComparator `json:"comparator,omitempty"`
	Unit       *string             `json:"unit,omitempty"`
	System     *string             `json:"system,omitempty"`
	Code       *string             `json:"code,omitempty"`
}
