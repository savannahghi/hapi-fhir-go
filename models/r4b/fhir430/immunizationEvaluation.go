
package fhir430

import "encoding/json"
// ImmunizationEvaluation is documented here http://hl7.org/fhir/StructureDefinition/ImmunizationEvaluation
// Describes a comparison of an immunization event against published recommendations to determine if the administration is "valid" in relation to those  recommendations.
type ImmunizationEvaluation struct {
	ID                     *string                           `json:"ID,omitempty"`
	Meta                   *Meta                             `json:"meta,omitempty"`
	ImplicitRules          *string                           `json:"implicitRules,omitempty"`
	Language               *string                           `json:"language,omitempty"`
	Text                   *Narrative                        `json:"text,omitempty"`
	Contained              []json.RawMessage                 `json:"contained,omitempty"`
	Extension              []Extension                       `json:"extension,omitempty"`
	ModifierExtension      []Extension                       `json:"modifierExtension,omitempty"`
	Identifier             []Identifier                      `json:"identifier,omitempty"`
	Status                 ImmunizationEvaluationStatusCodes `json:"status"`
	Patient                Reference                         `json:"patient"`
	Date                   *string                           `json:"date,omitempty"`
	Authority              *Reference                        `json:"authority,omitempty"`
	TargetDisease          CodeableConcept                   `json:"targetDisease"`
	ImmunizationEvent      Reference                         `json:"immunizationEvent"`
	DoseStatus             CodeableConcept                   `json:"doseStatus"`
	DoseStatusReason       []CodeableConcept                 `json:"doseStatusReason,omitempty"`
	Description            *string                           `json:"description,omitempty"`
	Series                 *string                           `json:"series,omitempty"`
	DoseNumberPositiveInt  *int                              `json:"doseNumberPositiveInt,omitempty"`
	DoseNumberString       *string                           `json:"doseNumberString,omitempty"`
	SeriesDosesPositiveInt *int                              `json:"seriesDosesPositiveInt,omitempty"`
	SeriesDosesString      *string                           `json:"seriesDosesString,omitempty"`
}

// This function returns resource reference information
func (r ImmunizationEvaluation) ResourceRef() (string, *string) {
	return "ImmunizationEvaluation", r.ID
}

// This function returns resource reference information
func (r ImmunizationEvaluation) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherImmunizationEvaluation ImmunizationEvaluation

// MarshalJSON marshals the given ImmunizationEvaluation as JSON into a byte slice
func (r ImmunizationEvaluation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherImmunizationEvaluation
		ResourceType string `json:"resourceType"`
	}{
		OtherImmunizationEvaluation: OtherImmunizationEvaluation(r),
		ResourceType:                "ImmunizationEvaluation",
	})
}

// UnmarshalImmunizationEvaluation unmarshals a ImmunizationEvaluation.
func UnmarshalImmunizationEvaluation(b []byte) (ImmunizationEvaluation, error) {
	var immunizationEvaluation ImmunizationEvaluation
	if err := json.Unmarshal(b, &immunizationEvaluation); err != nil {
		return immunizationEvaluation, err
	}
	return immunizationEvaluation, nil
}
