
package fhir430

import "encoding/json"
// Timing is documented here http://hl7.org/fhir/StructureDefinition/Timing
// Base StructureDefinition for Timing Type: Specifies an event that may occur multiple times. Timing schedules are used to record when things are planned, expected or requested to occur. The most common usage is in dosage instructions for medications. They are also used when planning care of various kinds, and may be used for reporting the schedule to which past regular activities were carried out.
type Timing struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Event             []string         `json:"event,omitempty"`
	Repeat            *TimingRepeat    `json:"repeat,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
}

// A set of rules that describe when the event is scheduled.
type TimingRepeat struct {
	ID             *string      `json:"ID,omitempty"`
	Extension      []Extension  `json:"extension,omitempty"`
	BoundsDuration *Duration    `json:"boundsDuration,omitempty"`
	BoundsRange    *Range       `json:"boundsRange,omitempty"`
	BoundsPeriod   *Period      `json:"boundsPeriod,omitempty"`
	Count          *int         `json:"count,omitempty"`
	CountMax       *int         `json:"countMax,omitempty"`
	Duration       *json.Number `json:"duration,omitempty"`
	DurationMax    *json.Number `json:"durationMax,omitempty"`
	DurationUnit   *string      `json:"durationUnit,omitempty"`
	Frequency      *int         `json:"frequency,omitempty"`
	FrequencyMax   *int         `json:"frequencyMax,omitempty"`
	Period         *json.Number `json:"period,omitempty"`
	PeriodMax      *json.Number `json:"periodMax,omitempty"`
	PeriodUnit     *string      `json:"periodUnit,omitempty"`
	DayOfWeek      []DaysOfWeek `json:"dayOfWeek,omitempty"`
	TimeOfDay      []string     `json:"timeOfDay,omitempty"`
	When           []string     `json:"when,omitempty"`
	Offset         *int         `json:"offset,omitempty"`
}
