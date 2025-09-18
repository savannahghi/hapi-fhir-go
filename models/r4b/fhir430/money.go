
package fhir430

import "encoding/json"
// Money is documented here http://hl7.org/fhir/StructureDefinition/Money
// Base StructureDefinition for Money Type: An amount of economic utility in some recognized currency.
type Money struct {
	ID        *string      `json:"ID,omitempty"`
	Extension []Extension  `json:"extension,omitempty"`
	Value     *json.Number `json:"value,omitempty"`
	Currency  *string      `json:"currency,omitempty"`
}
