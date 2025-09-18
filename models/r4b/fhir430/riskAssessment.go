
package fhir430

import "encoding/json"
// RiskAssessment is documented here http://hl7.org/fhir/StructureDefinition/RiskAssessment
// An assessment of the likely outcome(s) for a patient or other subject as well as the likelihood of each outcome.
type RiskAssessment struct {
	ID                 *string                    `json:"ID,omitempty"`
	Meta               *Meta                      `json:"meta,omitempty"`
	ImplicitRules      *string                    `json:"implicitRules,omitempty"`
	Language           *string                    `json:"language,omitempty"`
	Text               *Narrative                 `json:"text,omitempty"`
	Contained          []json.RawMessage          `json:"contained,omitempty"`
	Extension          []Extension                `json:"extension,omitempty"`
	ModifierExtension  []Extension                `json:"modifierExtension,omitempty"`
	Identifier         []Identifier               `json:"identifier,omitempty"`
	BasedOn            *Reference                 `json:"basedOn,omitempty"`
	Parent             *Reference                 `json:"parent,omitempty"`
	Status             ObservationStatus          `json:"status"`
	Method             *CodeableConcept           `json:"method,omitempty"`
	Code               *CodeableConcept           `json:"code,omitempty"`
	Subject            Reference                  `json:"subject"`
	Encounter          *Reference                 `json:"encounter,omitempty"`
	OccurrenceDateTime *string                    `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                    `json:"occurrencePeriod,omitempty"`
	Condition          *Reference                 `json:"condition,omitempty"`
	Performer          *Reference                 `json:"performer,omitempty"`
	ReasonCode         []CodeableConcept          `json:"reasonCode,omitempty"`
	ReasonReference    []Reference                `json:"reasonReference,omitempty"`
	Basis              []Reference                `json:"basis,omitempty"`
	Prediction         []RiskAssessmentPrediction `json:"prediction,omitempty"`
	Mitigation         *string                    `json:"mitigation,omitempty"`
	Note               []Annotation               `json:"note,omitempty"`
}

// Describes the expected outcome for the subject.
// Multiple repetitions can be used to identify the same type of outcome in different timeframes as well as different types of outcomes.
type RiskAssessmentPrediction struct {
	ID                 *string          `json:"ID,omitempty"`
	Extension          []Extension      `json:"extension,omitempty"`
	ModifierExtension  []Extension      `json:"modifierExtension,omitempty"`
	Outcome            *CodeableConcept `json:"outcome,omitempty"`
	ProbabilityDecimal *json.Number     `json:"probabilityDecimal,omitempty"`
	ProbabilityRange   *Range           `json:"probabilityRange,omitempty"`
	QualitativeRisk    *CodeableConcept `json:"qualitativeRisk,omitempty"`
	RelativeRisk       *json.Number     `json:"relativeRisk,omitempty"`
	WhenPeriod         *Period          `json:"whenPeriod,omitempty"`
	WhenRange          *Range           `json:"whenRange,omitempty"`
	Rationale          *string          `json:"rationale,omitempty"`
}

// This function returns resource reference information
func (r RiskAssessment) ResourceRef() (string, *string) {
	return "RiskAssessment", r.ID
}

// This function returns resource reference information
func (r RiskAssessment) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherRiskAssessment RiskAssessment

// MarshalJSON marshals the given RiskAssessment as JSON into a byte slice
func (r RiskAssessment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRiskAssessment
		ResourceType string `json:"resourceType"`
	}{
		OtherRiskAssessment: OtherRiskAssessment(r),
		ResourceType:        "RiskAssessment",
	})
}

// UnmarshalRiskAssessment unmarshals a RiskAssessment.
func UnmarshalRiskAssessment(b []byte) (RiskAssessment, error) {
	var riskAssessment RiskAssessment
	if err := json.Unmarshal(b, &riskAssessment); err != nil {
		return riskAssessment, err
	}
	return riskAssessment, nil
}
