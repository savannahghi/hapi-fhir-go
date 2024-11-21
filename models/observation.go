package models

import (
	"encoding/json"
	"time"

	"github.com/savannahghi/scalarutils"
)

type FHIRObservation struct {
	ID                   *string                          `json:"id,omitempty"`
	Text                 *FHIRNarrative                   `json:"text,omitempty"`
	Identifier           []*FHIRIdentifier                `json:"identifier,omitempty"`
	BasedOn              []*FHIRReference                 `json:"basedOn,omitempty"`
	PartOf               []*FHIRReference                 `json:"partOf,omitempty"`
	Status               *ObservationStatus               `json:"status,omitempty"`
	Category             []*FHIRCodeableConcept           `json:"category,omitempty"`
	Code                 *FHIRCodeableConcept             `json:"code,omitempty"`
	Subject              *FHIRReference                   `json:"subject,omitempty"`
	Focus                []*FHIRReference                 `json:"focus,omitempty"`
	Encounter            *FHIRReference                   `json:"encounter,omitempty"`
	EffectiveDateTime    *scalarutils.Date                `json:"effectiveDateTime,omitempty"`
	EffectivePeriod      *FHIRPeriod                      `json:"effectivePeriod,omitempty"`
	EffectiveTiming      *FHIRTiming                      `json:"effectiveTiming,omitempty"`
	EffectiveInstant     *scalarutils.Instant             `json:"effectiveInstant,omitempty"`
	Issued               *scalarutils.Instant             `json:"issued,omitempty"`
	Performer            []*FHIRReference                 `json:"performer,omitempty"`
	ValueQuantity        *FHIRQuantity                    `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *scalarutils.Code                `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                          `json:"valueString,omitempty"`
	ValueBoolean         *bool                            `json:"valueBoolean,omitempty"`
	ValueInteger         *string                          `json:"valueInteger,omitempty"`
	ValueRange           *FHIRRange                       `json:"valueRange,omitempty"`
	ValueRatio           *FHIRRatio                       `json:"valueRatio,omitempty"`
	ValueSampledData     *FHIRSampledData                 `json:"valueSampledData,omitempty"`
	ValueTime            *time.Time                       `json:"valueTime,omitempty"`
	ValueDateTime        *scalarutils.Date                `json:"valueDateTime,omitempty"`
	ValuePeriod          *FHIRPeriod                      `json:"valuePeriod,omitempty"`
	DataAbsentReason     *FHIRCodeableConcept             `json:"dataAbsentReason,omitempty"`
	Interpretation       []*FHIRCodeableConcept           `json:"interpretation,omitempty"`
	Note                 []*FHIRAnnotation                `json:"note,omitempty"`
	BodySite             *FHIRCodeableConcept             `json:"bodySite,omitempty"`
	Method               *FHIRCodeableConcept             `json:"method,omitempty"`
	Specimen             *FHIRReference                   `json:"specimen,omitempty"`
	Device               *FHIRReference                   `json:"device,omitempty"`
	ReferenceRange       []*FHIRObservationReferencerange `json:"referenceRange,omitempty"`
	HasMember            []*FHIRReference                 `json:"hasMember,omitempty"`
	DerivedFrom          []*FHIRReference                 `json:"derivedFrom,omitempty"`
	Component            []*FHIRObservationComponent      `json:"component,omitempty"`
	Meta                 *FHIRMeta                        `json:"meta,omitempty"`
	Extension            []*FHIRExtension                 `json:"extension,omitempty"`
}

type FHIRObservationComponent struct {
	ID                   *string                          `json:"id,omitempty"`
	Code                 FHIRCodeableConcept              `json:"code,omitempty"`
	ValueQuantity        *FHIRQuantity                    `json:"valueQuantity,omitempty"`
	ValueCodeableConcept *scalarutils.Code                `json:"valueCodeableConcept,omitempty"`
	ValueString          *string                          `json:"valueString,omitempty"`
	ValueBoolean         *bool                            `json:"valueBoolean,omitempty"`
	ValueInteger         *string                          `json:"valueInteger,omitempty"`
	ValueRange           *FHIRRange                       `json:"valueRange,omitempty"`
	ValueRatio           *FHIRRatio                       `json:"valueRatio,omitempty"`
	ValueSampledData     *FHIRSampledData                 `json:"valueSampledData,omitempty"`
	ValueTime            *time.Time                       `json:"valueTime,omitempty"`
	ValueDateTime        *scalarutils.Date                `json:"valueDateTime,omitempty"`
	ValuePeriod          *FHIRPeriod                      `json:"valuePeriod,omitempty"`
	DataAbsentReason     *FHIRCodeableConcept             `json:"dataAbsentReason,omitempty"`
	Interpretation       []*FHIRCodeableConcept           `json:"interpretation,omitempty"`
	ReferenceRange       []*FHIRObservationReferencerange `json:"referenceRange,omitempty"`
}

type FHIRObservationReferencerange struct {
	ID        *string                `json:"id,omitempty"`
	Low       *FHIRQuantity          `json:"low,omitempty"`
	High      *FHIRQuantity          `json:"high,omitempty"`
	Type      *FHIRCodeableConcept   `json:"type,omitempty"`
	AppliesTo []*FHIRCodeableConcept `json:"appliesTo,omitempty"`
	Age       *FHIRRange             `json:"age,omitempty"`
	Text      *string                `json:"text,omitempty"`
}

type FHIRObservationRelayPayload struct {
	Resource *FHIRObservation `json:"resource,omitempty"`
}

type FHIRSampledData struct {
	ID         *string              `json:"id,omitempty"`
	Origin     *FHIRQuantity        `json:"origin,omitempty"`
	Period     *scalarutils.Decimal `json:"period,omitempty"`
	Factor     *scalarutils.Decimal `json:"factor,omitempty"`
	LowerLimit *scalarutils.Decimal `json:"lowerLimit,omitempty"`
	UpperLimit *scalarutils.Decimal `json:"upperLimit,omitempty"`
	Dimensions *string              `json:"dimensions,omitempty"`
	Data       *string              `json:"data,omitempty"`
}

type FHIRTiming struct {
	ID     *string                 `json:"id,omitempty"`
	Event  []*scalarutils.DateTime `json:"event,omitempty"`
	Repeat *FHIRTimingRepeat       `json:"repeat,omitempty"`
	Code   scalarutils.Code        `json:"code,omitempty"`
}

type FHIRTimingRepeat struct {
	ID             *string               `json:"id,omitempty"`
	BoundsDuration *FHIRDuration         `json:"boundsDuration,omitempty"`
	BoundsRange    *FHIRRange            `json:"boundsRange,omitempty"`
	BoundsPeriod   *FHIRPeriod           `json:"boundsPeriod,omitempty"`
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
	DayOfWeek      []*scalarutils.Code   `json:"dayOfWeek,omitempty"`
	TimeOfDay      *time.Time            `json:"timeOfDay,omitempty"`
	When           *TimingRepeatWhenEnum `json:"when,omitempty"`
	Offset         *int                  `json:"offset,omitempty"`
}
