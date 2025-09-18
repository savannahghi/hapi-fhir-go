
package fhir430

import "encoding/json"
// MedicinalProductUndesirableEffect is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductUndesirableEffect
// Describe the undesirable effects of the medicinal product.
type MedicinalProductUndesirableEffect struct {
	ID                     *string           `json:"ID,omitempty"`
	Meta                   *Meta             `json:"meta,omitempty"`
	ImplicitRules          *string           `json:"implicitRules,omitempty"`
	Language               *string           `json:"language,omitempty"`
	Text                   *Narrative        `json:"text,omitempty"`
	Contained              []json.RawMessage `json:"contained,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Subject                []Reference       `json:"subject,omitempty"`
	SymptomConditionEffect *CodeableConcept  `json:"symptomConditionEffect,omitempty"`
	Classification         *CodeableConcept  `json:"classification,omitempty"`
	FrequencyOfOccurrence  *CodeableConcept  `json:"frequencyOfOccurrence,omitempty"`
	Population             []Population      `json:"population,omitempty"`
}

// This function returns resource reference information
func (r MedicinalProductUndesirableEffect) ResourceRef() (string, *string) {
	return "MedicinalProductUndesirableEffect", r.ID
}

// This function returns resource reference information
func (r MedicinalProductUndesirableEffect) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedicinalProductUndesirableEffect MedicinalProductUndesirableEffect

// MarshalJSON marshals the given MedicinalProductUndesirableEffect as JSON into a byte slice
func (r MedicinalProductUndesirableEffect) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductUndesirableEffect
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductUndesirableEffect: OtherMedicinalProductUndesirableEffect(r),
		ResourceType:                           "MedicinalProductUndesirableEffect",
	})
}

// UnmarshalMedicinalProductUndesirableEffect unmarshals a MedicinalProductUndesirableEffect.
func UnmarshalMedicinalProductUndesirableEffect(b []byte) (MedicinalProductUndesirableEffect, error) {
	var medicinalProductUndesirableEffect MedicinalProductUndesirableEffect
	if err := json.Unmarshal(b, &medicinalProductUndesirableEffect); err != nil {
		return medicinalProductUndesirableEffect, err
	}
	return medicinalProductUndesirableEffect, nil
}
