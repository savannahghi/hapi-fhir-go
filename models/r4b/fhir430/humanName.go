
package fhir430
// HumanName is documented here http://hl7.org/fhir/StructureDefinition/HumanName
// Base StructureDefinition for HumanName Type: A human's name with the ability to identify parts and usage.
type HumanName struct {
	ID        *string     `json:"ID,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Use       *NameUse    `json:"use,omitempty"`
	Text      *string     `json:"text,omitempty"`
	Family    *string     `json:"family,omitempty"`
	Given     []string    `json:"given,omitempty"`
	Prefix    []string    `json:"prefix,omitempty"`
	Suffix    []string    `json:"suffix,omitempty"`
	Period    *Period     `json:"period,omitempty"`
}
