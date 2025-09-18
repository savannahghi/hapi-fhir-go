
package fhir430

import "encoding/json"
// Condition is documented here http://hl7.org/fhir/StructureDefinition/Condition
// A clinical condition, problem, diagnosis, or other event, situation, issue, or clinical concept that has risen to a level of concern.
type Condition struct {
	ID                 *string             `json:"ID,omitempty"`
	Meta               *Meta               `json:"meta,omitempty"`
	ImplicitRules      *string             `json:"implicitRules,omitempty"`
	Language           *string             `json:"language,omitempty"`
	Text               *Narrative          `json:"text,omitempty"`
	Contained          []json.RawMessage   `json:"contained,omitempty"`
	Extension          []Extension         `json:"extension,omitempty"`
	ModifierExtension  []Extension         `json:"modifierExtension,omitempty"`
	Identifier         []Identifier        `json:"identifier,omitempty"`
	ClinicalStatus     *CodeableConcept    `json:"clinicalStatus,omitempty"`
	VerificationStatus *CodeableConcept    `json:"verificationStatus,omitempty"`
	Category           []CodeableConcept   `json:"category,omitempty"`
	Severity           *CodeableConcept    `json:"severity,omitempty"`
	Code               *CodeableConcept    `json:"code,omitempty"`
	BodySite           []CodeableConcept   `json:"bodySite,omitempty"`
	Subject            Reference           `json:"subject"`
	Encounter          *Reference          `json:"encounter,omitempty"`
	OnsetDateTime      *string             `json:"onsetDateTime,omitempty"`
	OnsetAge           *Age                `json:"onsetAge,omitempty"`
	OnsetPeriod        *Period             `json:"onsetPeriod,omitempty"`
	OnsetRange         *Range              `json:"onsetRange,omitempty"`
	OnsetString        *string             `json:"onsetString,omitempty"`
	AbatementDateTime  *string             `json:"abatementDateTime,omitempty"`
	AbatementAge       *Age                `json:"abatementAge,omitempty"`
	AbatementPeriod    *Period             `json:"abatementPeriod,omitempty"`
	AbatementRange     *Range              `json:"abatementRange,omitempty"`
	AbatementString    *string             `json:"abatementString,omitempty"`
	RecordedDate       *string             `json:"recordedDate,omitempty"`
	Recorder           *Reference          `json:"recorder,omitempty"`
	Asserter           *Reference          `json:"asserter,omitempty"`
	Stage              []ConditionStage    `json:"stage,omitempty"`
	Evidence           []ConditionEvidence `json:"evidence,omitempty"`
	Note               []Annotation        `json:"note,omitempty"`
}

// Clinical stage or grade of a condition. May include formal severity assessments.
type ConditionStage struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Summary           *CodeableConcept `json:"summary,omitempty"`
	Assessment        []Reference      `json:"assessment,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
}

// Supporting evidence / manifestations that are the basis of the Condition's verification status, such as evidence that confirmed or refuted the condition.
// The evidence may be a simple list of coded symptoms/manifestations, or references to observations or formal assessments, or both.
type ConditionEvidence struct {
	ID                *string           `json:"ID,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Detail            []Reference       `json:"detail,omitempty"`
}

// This function returns resource reference information
func (r Condition) ResourceRef() (string, *string) {
	return "Condition", r.ID
}

// This function returns resource reference information
func (r Condition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCondition Condition

// MarshalJSON marshals the given Condition as JSON into a byte slice
func (r Condition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCondition
		ResourceType string `json:"resourceType"`
	}{
		OtherCondition: OtherCondition(r),
		ResourceType:   "Condition",
	})
}

// UnmarshalCondition unmarshals a Condition.
func UnmarshalCondition(b []byte) (Condition, error) {
	var condition Condition
	if err := json.Unmarshal(b, &condition); err != nil {
		return condition, err
	}
	return condition, nil
}
