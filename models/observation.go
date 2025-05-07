package models

import (
	"encoding/json"
	"time"

	"github.com/savannahghi/scalarutils"
)

type Observation struct {
	ID                   *string                      `json:"id,omitempty"`
	Text                 *Narrative                   `json:"text,omitempty"`
	Identifier           []*Identifier                `json:"identifier,omitempty"`
	BasedOn              []*Reference                 `json:"basedOn,omitempty"`
	PartOf               []*Reference                 `json:"partOf,omitempty"`
	Status               *ObservationStatus           `json:"status,omitempty"`
	Category             []*CodeableConcept           `json:"category,omitempty"`
	Code                 *CodeableConcept             `json:"code,omitempty"`
	Subject              *Reference                   `json:"subject,omitempty"`
	Focus                []*Reference                 `json:"focus,omitempty"`
	Encounter            *Reference                   `json:"encounter,omitempty"`
	EffectiveDateTime    *scalarutils.Date            `json:"effectiveDateTime,omitempty"`
	EffectivePeriod      *Period                      `json:"effectivePeriod,omitempty"`
	EffectiveTiming      *Timing                      `json:"effectiveTiming,omitempty"`
	EffectiveInstant     *scalarutils.Instant         `json:"effectiveInstant,omitempty"`
	Issued               *scalarutils.Instant         `json:"issued,omitempty"`
	Performer            []*Reference                 `json:"performer,omitempty"`
	ValueQuantity        *Quantity                    `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *string                      `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                      `json:"valueString,omitempty"`
	ValueBoolean         *bool                        `json:"valueBoolean,omitempty"`
	ValueInteger         *string                      `json:"valueInteger,omitempty"`
	ValueRange           *Range                       `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                       `json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                 `json:"valueSampledData,omitempty"`
	ValueTime            *time.Time                   `json:"valueTime,omitempty"`
	ValueDateTime        *scalarutils.Date            `json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                      `json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept             `json:"dataAbsentReason,omitempty"`
	Interpretation       []*CodeableConcept           `json:"interpretation,omitempty"`
	Note                 []*Annotation                `json:"note,omitempty"`
	BodySite             *CodeableConcept             `json:"bodySite,omitempty"`
	Method               *CodeableConcept             `json:"method,omitempty"`
	Specimen             *Reference                   `json:"specimen,omitempty"`
	Device               *Reference                   `json:"device,omitempty"`
	ReferenceRange       []*ObservationReferencerange `json:"referenceRange,omitempty"`
	HasMember            []*Reference                 `json:"hasMember,omitempty"`
	DerivedFrom          []*Reference                 `json:"derivedFrom,omitempty"`
	Component            []*ObservationComponent      `json:"component,omitempty"`
	Meta                 *Meta                        `json:"meta,omitempty"`
	Extension            []*Extension                 `json:"extension,omitempty"`
}

type ObservationComponent struct {
	ID                   *string                      `json:"id,omitempty"`
	Code                 CodeableConcept              `json:"code,omitempty"`
	ValueQuantity        *Quantity                    `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *string                      `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                      `json:"valueString,omitempty"`
	ValueBoolean         *bool                        `json:"valueBoolean,omitempty"`
	ValueInteger         *string                      `json:"valueInteger,omitempty"`
	ValueRange           *Range                       `json:"valueRange,omitempty"`
	ValueRatio           *Ratio                       `json:"valueRatio,omitempty"`
	ValueSampledData     *SampledData                 `json:"valueSampledData,omitempty"`
	ValueTime            *time.Time                   `json:"valueTime,omitempty"`
	ValueDateTime        *scalarutils.Date            `json:"valueDateTime,omitempty"`
	ValuePeriod          *Period                      `json:"valuePeriod,omitempty"`
	DataAbsentReason     *CodeableConcept             `json:"dataAbsentReason,omitempty"`
	Interpretation       []*CodeableConcept           `json:"interpretation,omitempty"`
	ReferenceRange       []*ObservationReferencerange `json:"referenceRange,omitempty"`
}

type ObservationReferencerange struct {
	ID        *string            `json:"id,omitempty"`
	Low       *Quantity          `json:"low,omitempty"`
	High      *Quantity          `json:"high,omitempty"`
	Type      *CodeableConcept   `json:"type,omitempty"`
	AppliesTo []*CodeableConcept `json:"appliesTo,omitempty"`
	Age       *Range             `json:"age,omitempty"`
	Text      *string            `json:"text,omitempty"`
}

type ObservationRelayPayload struct {
	Resource *Observation `json:"resource,omitempty"`
}

type SampledData struct {
	ID         *string              `json:"id,omitempty"`
	Origin     *Quantity            `json:"origin,omitempty"`
	Period     *scalarutils.Decimal `json:"period,omitempty"`
	Factor     *scalarutils.Decimal `json:"factor,omitempty"`
	LowerLimit *scalarutils.Decimal `json:"lowerLimit,omitempty"`
	UpperLimit *scalarutils.Decimal `json:"upperLimit,omitempty"`
	Dimensions *string              `json:"dimensions,omitempty"`
	Data       *string              `json:"data,omitempty"`
}

type Timing struct {
	ID     *string                 `json:"id,omitempty"`
	Event  []*scalarutils.DateTime `json:"event,omitempty"`
	Repeat *TimingRepeat           `json:"repeat,omitempty"`
	Code   string                  `json:"code,omitempty"`
}

type TimingRepeat struct {
	ID             *string               `json:"id,omitempty"`
	BoundsDuration *Duration             `json:"boundsDuration,omitempty"`
	BoundsRange    *Range                `json:"boundsRange,omitempty"`
	BoundsPeriod   *Period               `json:"boundsPeriod,omitempty"`
	Count          *string               `json:"count,omitempty"`
	CountMax       *string               `json:"countMax,omitempty"`
	Duration       *json.Number          `json:"duration,omitempty"`
	DurationMax    *scalarutils.Decimal  `json:"durationMax,omitempty"`
	DurationUnit   *UnitsOfTime          `json:"durationUnit,omitempty"`
	Frequency      *int                  `json:"frequency,omitempty"`
	FrequencyMax   *string               `json:"frequencyMax,omitempty"`
	Period         *json.Number          `json:"period,omitempty"`
	PeriodMax      *scalarutils.Decimal  `json:"periodMax,omitempty"`
	PeriodUnit     *UnitsOfTime          `json:"periodUnit,omitempty"`
	DayOfWeek      []*string             `json:"dayOfWeek,omitempty"`
	TimeOfDay      *time.Time            `json:"timeOfDay,omitempty"`
	When           *TimingRepeatWhenEnum `json:"when,omitempty"`
	Offset         *int                  `json:"offset,omitempty"`
}
