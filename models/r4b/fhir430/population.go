
package fhir430
// Population is documented here http://hl7.org/fhir/StructureDefinition/Population
// Base StructureDefinition for Population Type: A populatioof people with some set of grouping criteria.
type Population struct {
	ID                     *string          `json:"ID,omitempty"`
	Extension              []Extension      `json:"extension,omitempty"`
	ModifierExtension      []Extension      `json:"modifierExtension,omitempty"`
	AgeRange               *Range           `json:"ageRange,omitempty"`
	AgeCodeableConcept     *CodeableConcept `json:"ageCodeableConcept,omitempty"`
	Gender                 *CodeableConcept `json:"gender,omitempty"`
	Race                   *CodeableConcept `json:"race,omitempty"`
	PhysiologicalCondition *CodeableConcept `json:"physiologicalCondition,omitempty"`
}
