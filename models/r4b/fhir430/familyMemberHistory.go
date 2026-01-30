package fhir430

import "encoding/json"

// FamilyMemberHistory is documented here http://hl7.org/fhir/StructureDefinition/FamilyMemberHistory
// Significant health conditions for a person related to the patient relevant in the context of care for the patient.
type FamilyMemberHistory struct {
	ID                    *string                        `json:"id,omitempty"`
	Meta                  *Meta                          `json:"meta,omitempty"`
	ImplicitRules         *string                        `json:"implicitRules,omitempty"`
	Language              *string                        `json:"language,omitempty"`
	Text                  *Narrative                     `json:"text,omitempty"`
	Contained             []json.RawMessage              `json:"contained,omitempty"`
	Extension             []Extension                    `json:"extension,omitempty"`
	ModifierExtension     []Extension                    `json:"modifierExtension,omitempty"`
	Identifier            []Identifier                   `json:"identifier,omitempty"`
	InstantiatesCanonical []string                       `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                       `json:"instantiatesUri,omitempty"`
	Status                FamilyHistoryStatus            `json:"status"`
	DataAbsentReason      *CodeableConcept               `json:"dataAbsentReason,omitempty"`
	Patient               Reference                      `json:"patient"`
	Date                  *string                        `json:"date,omitempty"`
	Name                  *string                        `json:"name,omitempty"`
	Relationship          CodeableConcept                `json:"relationship"`
	Sex                   *CodeableConcept               `json:"sex,omitempty"`
	BornPeriod            *Period                        `json:"bornPeriod,omitempty"`
	BornDate              *string                        `json:"bornDate,omitempty"`
	BornString            *string                        `json:"bornString,omitempty"`
	AgeAge                *Age                           `json:"ageAge,omitempty"`
	AgeRange              *Range                         `json:"ageRange,omitempty"`
	AgeString             *string                        `json:"ageString,omitempty"`
	EstimatedAge          *bool                          `json:"estimatedAge,omitempty"`
	DeceasedBoolean       *bool                          `json:"deceasedBoolean,omitempty"`
	DeceasedAge           *Age                           `json:"deceasedAge,omitempty"`
	DeceasedRange         *Range                         `json:"deceasedRange,omitempty"`
	DeceasedDate          *string                        `json:"deceasedDate,omitempty"`
	DeceasedString        *string                        `json:"deceasedString,omitempty"`
	ReasonCode            []CodeableConcept              `json:"reasonCode,omitempty"`
	ReasonReference       []Reference                    `json:"reasonReference,omitempty"`
	Note                  []Annotation                   `json:"note,omitempty"`
	Condition             []FamilyMemberHistoryCondition `json:"condition,omitempty"`
}

// The significant Conditions (or condition) that the family member had. This is a repeating section to allow a system to represent more than one condition per resource, though there is nothing stopping multiple resources - one per condition.
type FamilyMemberHistoryCondition struct {
	ID                 *string          `json:"id,omitempty"`
	Extension          []Extension      `json:"extension,omitempty"`
	ModifierExtension  []Extension      `json:"modifierExtension,omitempty"`
	Code               CodeableConcept  `json:"code"`
	Outcome            *CodeableConcept `json:"outcome,omitempty"`
	ContributedToDeath *bool            `json:"contributedToDeath,omitempty"`
	OnsetAge           *Age             `json:"onsetAge,omitempty"`
	OnsetRange         *Range           `json:"onsetRange,omitempty"`
	OnsetPeriod        *Period          `json:"onsetPeriod,omitempty"`
	OnsetString        *string          `json:"onsetString,omitempty"`
	Note               []Annotation     `json:"note,omitempty"`
}

// This function returns resource reference information
func (r FamilyMemberHistory) ResourceRef() (string, *string) {
	return "FamilyMemberHistory", r.ID
}

// This function returns resource reference information
func (r FamilyMemberHistory) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherFamilyMemberHistory FamilyMemberHistory

// MarshalJSON marshals the given FamilyMemberHistory as JSON into a byte slice
func (r FamilyMemberHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFamilyMemberHistory
		ResourceType string `json:"resourceType"`
	}{
		OtherFamilyMemberHistory: OtherFamilyMemberHistory(r),
		ResourceType:             "FamilyMemberHistory",
	})
}

// UnmarshalFamilyMemberHistory unmarshals a FamilyMemberHistory.
func UnmarshalFamilyMemberHistory(b []byte) (FamilyMemberHistory, error) {
	var familyMemberHistory FamilyMemberHistory
	if err := json.Unmarshal(b, &familyMemberHistory); err != nil {
		return familyMemberHistory, err
	}
	return familyMemberHistory, nil
}
