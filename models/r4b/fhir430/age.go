package fhir430

import "encoding/json"

// Age is documented here http://hl7.org/fhir/StructureDefinition/Age
// Base StructureDefinition for Age Type: A duration of time during which an organism (or a process) has existed.
type Age struct {
	ID         *string             `json:"id,omitempty"`
	Extension  []Extension         `json:"extension,omitempty"`
	Value      *json.Number        `json:"value,omitempty"`
	Comparator *QuantityComparator `json:"comparator,omitempty"`
	Unit       *string             `json:"unit,omitempty"`
	System     *string             `json:"system,omitempty"`
	Code       *string             `json:"code,omitempty"`
}
