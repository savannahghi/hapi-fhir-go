
package fhir430

import "encoding/json"
// Schedule is documented here http://hl7.org/fhir/StructureDefinition/Schedule
// A container for slots of time that may be available for booking appointments.
type Schedule struct {
	ID                *string           `json:"ID,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []json.RawMessage `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	Active            *bool             `json:"active,omitempty"`
	ServiceCategory   []CodeableConcept `json:"serviceCategory,omitempty"`
	ServiceType       []CodeableConcept `json:"serviceType,omitempty"`
	Specialty         []CodeableConcept `json:"specialty,omitempty"`
	Actor             []Reference       `json:"actor"`
	PlanningHorizon   *Period           `json:"planningHorizon,omitempty"`
	Comment           *string           `json:"comment,omitempty"`
}

// This function returns resource reference information
func (r Schedule) ResourceRef() (string, *string) {
	return "Schedule", r.ID
}

// This function returns resource reference information
func (r Schedule) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherSchedule Schedule

// MarshalJSON marshals the given Schedule as JSON into a byte slice
func (r Schedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSchedule
		ResourceType string `json:"resourceType"`
	}{
		OtherSchedule: OtherSchedule(r),
		ResourceType:  "Schedule",
	})
}

// UnmarshalSchedule unmarshals a Schedule.
func UnmarshalSchedule(b []byte) (Schedule, error) {
	var schedule Schedule
	if err := json.Unmarshal(b, &schedule); err != nil {
		return schedule, err
	}
	return schedule, nil
}
