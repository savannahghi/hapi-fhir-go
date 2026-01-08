package fhir430

import "encoding/json"

// Count is documented here http://hl7.org/fhir/StructureDefinition/Count
// Base StructureDefinition for Count Type: A measured amount (or an amount that can potentially be measured). Note that measured amounts include amounts that are not precisely quantified, including amounts involving arbitrary units and floating currencies.
type Count struct {
	ID         *string             `json:"id,omitempty"`
	Extension  []Extension         `json:"extension,omitempty"`
	Value      *json.Number        `json:"value,omitempty"`
	Comparator *QuantityComparator `json:"comparator,omitempty"`
	Unit       *string             `json:"unit,omitempty"`
	System     *string             `json:"system,omitempty"`
	Code       *string             `json:"code,omitempty"`
}
