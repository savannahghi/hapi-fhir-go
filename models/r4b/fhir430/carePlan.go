package fhir430

import "encoding/json"

// CarePlan is documented here http://hl7.org/fhir/StructureDefinition/CarePlan
// Describes the intention of how one or more practitioners intend to deliver care for a particular patient, group or community for a period of time, possibly limited to care for a specific condition or set of conditions.
type CarePlan struct {
	ID                    *string            `json:"id,omitempty"`
	Meta                  *Meta              `json:"meta,omitempty"`
	ImplicitRules         *string            `json:"implicitRules,omitempty"`
	Language              *string            `json:"language,omitempty"`
	Text                  *Narrative         `json:"text,omitempty"`
	Contained             []json.RawMessage  `json:"contained,omitempty"`
	Extension             []Extension        `json:"extension,omitempty"`
	ModifierExtension     []Extension        `json:"modifierExtension,omitempty"`
	Identifier            []Identifier       `json:"identifier,omitempty"`
	InstantiatesCanonical []string           `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string           `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference        `json:"basedOn,omitempty"`
	Replaces              []Reference        `json:"replaces,omitempty"`
	PartOf                []Reference        `json:"partOf,omitempty"`
	Status                RequestStatus      `json:"status"`
	Intent                CarePlanIntent     `json:"intent"`
	Category              []CodeableConcept  `json:"category,omitempty"`
	Title                 *string            `json:"title,omitempty"`
	Description           *string            `json:"description,omitempty"`
	Subject               Reference          `json:"subject"`
	Encounter             *Reference         `json:"encounter,omitempty"`
	Period                *Period            `json:"period,omitempty"`
	Created               *string            `json:"created,omitempty"`
	Author                *Reference         `json:"author,omitempty"`
	Contributor           []Reference        `json:"contributor,omitempty"`
	CareTeam              []Reference        `json:"careTeam,omitempty"`
	Addresses             []Reference        `json:"addresses,omitempty"`
	SupportingInfo        []Reference        `json:"supportingInfo,omitempty"`
	Goal                  []Reference        `json:"goal,omitempty"`
	Activity              []CarePlanActivity `json:"activity,omitempty"`
	Note                  []Annotation       `json:"note,omitempty"`
}

// Identifies a planned action to occur as part of the plan.  For example, a medication to be used, lab tests to perform, self-monitoring, education, etc.
type CarePlanActivity struct {
	ID                     *string                 `json:"id,omitempty"`
	Extension              []Extension             `json:"extension,omitempty"`
	ModifierExtension      []Extension             `json:"modifierExtension,omitempty"`
	OutcomeCodeableConcept []CodeableConcept       `json:"outcomeCodeableConcept,omitempty"`
	OutcomeReference       []Reference             `json:"outcomeReference,omitempty"`
	Progress               []Annotation            `json:"progress,omitempty"`
	Reference              *Reference              `json:"reference,omitempty"`
	Detail                 *CarePlanActivityDetail `json:"detail,omitempty"`
}

// A simple summary of a planned activity suitable for a general care plan system (e.g. form driven) that doesn't know about specific resources such as procedure etc.
type CarePlanActivityDetail struct {
	ID                     *string                `json:"id,omitempty"`
	Extension              []Extension            `json:"extension,omitempty"`
	ModifierExtension      []Extension            `json:"modifierExtension,omitempty"`
	Kind                   *CarePlanActivityKind  `json:"kind,omitempty"`
	InstantiatesCanonical  []string               `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri        []string               `json:"instantiatesUri,omitempty"`
	Code                   *CodeableConcept       `json:"code,omitempty"`
	ReasonCode             []CodeableConcept      `json:"reasonCode,omitempty"`
	ReasonReference        []Reference            `json:"reasonReference,omitempty"`
	Goal                   []Reference            `json:"goal,omitempty"`
	Status                 CarePlanActivityStatus `json:"status"`
	StatusReason           *CodeableConcept       `json:"statusReason,omitempty"`
	DoNotPerform           *bool                  `json:"doNotPerform,omitempty"`
	ScheduledTiming        *Timing                `json:"scheduledTiming,omitempty"`
	ScheduledPeriod        *Period                `json:"scheduledPeriod,omitempty"`
	ScheduledString        *string                `json:"scheduledString,omitempty"`
	Location               *Reference             `json:"location,omitempty"`
	Performer              []Reference            `json:"performer,omitempty"`
	ProductCodeableConcept *CodeableConcept       `json:"productCodeableConcept,omitempty"`
	ProductReference       *Reference             `json:"productReference,omitempty"`
	DailyAmount            *Quantity              `json:"dailyAmount,omitempty"`
	Quantity               *Quantity              `json:"quantity,omitempty"`
	Description            *string                `json:"description,omitempty"`
}

// This function returns resource reference information
func (r CarePlan) ResourceRef() (string, *string) {
	return "CarePlan", r.ID
}

// This function returns resource reference information
func (r CarePlan) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCarePlan CarePlan

// MarshalJSON marshals the given CarePlan as JSON into a byte slice
func (r CarePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCarePlan
		ResourceType string `json:"resourceType"`
	}{
		OtherCarePlan: OtherCarePlan(r),
		ResourceType:  "CarePlan",
	})
}

// UnmarshalCarePlan unmarshals a CarePlan.
func UnmarshalCarePlan(b []byte) (CarePlan, error) {
	var carePlan CarePlan
	if err := json.Unmarshal(b, &carePlan); err != nil {
		return carePlan, err
	}
	return carePlan, nil
}
