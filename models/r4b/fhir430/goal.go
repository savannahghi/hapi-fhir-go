
package fhir430

import "encoding/json"
// Goal is documented here http://hl7.org/fhir/StructureDefinition/Goal
// Describes the intended objective(s) for a patient, group or organization care, for example, weight loss, restoring an activity of daily living, obtaining herd immunity via immunization, meeting a process improvement objective, etc.
type Goal struct {
	ID                   *string             `json:"ID,omitempty"`
	Meta                 *Meta               `json:"meta,omitempty"`
	ImplicitRules        *string             `json:"implicitRules,omitempty"`
	Language             *string             `json:"language,omitempty"`
	Text                 *Narrative          `json:"text,omitempty"`
	Contained            []json.RawMessage   `json:"contained,omitempty"`
	Extension            []Extension         `json:"extension,omitempty"`
	ModifierExtension    []Extension         `json:"modifierExtension,omitempty"`
	Identifier           []Identifier        `json:"identifier,omitempty"`
	LifecycleStatus      GoalLifecycleStatus `json:"lifecycleStatus"`
	AchievementStatus    *CodeableConcept    `json:"achievementStatus,omitempty"`
	Category             []CodeableConcept   `json:"category,omitempty"`
	Priority             *CodeableConcept    `json:"priority,omitempty"`
	Description          CodeableConcept     `json:"description"`
	Subject              Reference           `json:"subject"`
	StartDate            *string             `json:"startDate,omitempty"`
	StartCodeableConcept *CodeableConcept    `json:"startCodeableConcept,omitempty"`
	Target               []GoalTarget        `json:"target,omitempty"`
	StatusDate           *string             `json:"statusDate,omitempty"`
	StatusReason         *string             `json:"statusReason,omitempty"`
	ExpressedBy          *Reference          `json:"expressedBy,omitempty"`
	Addresses            []Reference         `json:"addresses,omitempty"`
	Note                 []Annotation        `json:"note,omitempty"`
	OutcomeCode          []CodeableConcept   `json:"outcomeCode,omitempty"`
	OutcomeReference     []Reference         `json:"outcomeReference,omitempty"`
}

// Indicates what should be done by when.
// When multiple targets are present for a single goal instance, all targets must be met for the overall goal to be met.
type GoalTarget struct {
	ID                    *string          `json:"ID,omitempty"`
	Extension             []Extension      `json:"extension,omitempty"`
	ModifierExtension     []Extension      `json:"modifierExtension,omitempty"`
	Measure               *CodeableConcept `json:"measure,omitempty"`
	DetailQuantity        *Quantity        `json:"detailQuantity,omitempty"`
	DetailRange           *Range           `json:"detailRange,omitempty"`
	DetailCodeableConcept *CodeableConcept `json:"detailCodeableConcept,omitempty"`
	DetailString          *string          `json:"detailString,omitempty"`
	DetailBoolean         *bool            `json:"detailBoolean,omitempty"`
	DetailInteger         *int             `json:"detailInteger,omitempty"`
	DetailRatio           *Ratio           `json:"detailRatio,omitempty"`
	DueDate               *string          `json:"dueDate,omitempty"`
	DueDuration           *Duration        `json:"dueDuration,omitempty"`
}

// This function returns resource reference information
func (r Goal) ResourceRef() (string, *string) {
	return "Goal", r.ID
}

// This function returns resource reference information
func (r Goal) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherGoal Goal

// MarshalJSON marshals the given Goal as JSON into a byte slice
func (r Goal) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGoal
		ResourceType string `json:"resourceType"`
	}{
		OtherGoal:    OtherGoal(r),
		ResourceType: "Goal",
	})
}

// UnmarshalGoal unmarshals a Goal.
func UnmarshalGoal(b []byte) (Goal, error) {
	var goal Goal
	if err := json.Unmarshal(b, &goal); err != nil {
		return goal, err
	}
	return goal, nil
}
