
package fhir430
// Contributor is documented here http://hl7.org/fhir/StructureDefinition/Contributor
// Base StructureDefinition for Contributor Type: A contributor to the content of a knowledge asset, including authors, editors, reviewers, and endorsers.
type Contributor struct {
	ID        *string         `json:"ID,omitempty"`
	Extension []Extension     `json:"extension,omitempty"`
	Type      ContributorType `json:"type"`
	Name      string          `json:"name"`
	Contact   []ContactDetail `json:"contact,omitempty"`
}
