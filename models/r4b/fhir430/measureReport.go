package fhir430

import "encoding/json"

// MeasureReport is documented here http://hl7.org/fhir/StructureDefinition/MeasureReport
// The MeasureReport resource contains the results of the calculation of a measure; and optionally a reference to the resources involved in that calculation.
type MeasureReport struct {
	ID                  *string              `json:"id,omitempty"`
	Meta                *Meta                `json:"meta,omitempty"`
	ImplicitRules       *string              `json:"implicitRules,omitempty"`
	Language            *string              `json:"language,omitempty"`
	Text                *Narrative           `json:"text,omitempty"`
	Contained           []json.RawMessage    `json:"contained,omitempty"`
	Extension           []Extension          `json:"extension,omitempty"`
	ModifierExtension   []Extension          `json:"modifierExtension,omitempty"`
	Identifier          []Identifier         `json:"identifier,omitempty"`
	Status              MeasureReportStatus  `json:"status"`
	Type                MeasureReportType    `json:"type"`
	Measure             string               `json:"measure"`
	Subject             *Reference           `json:"subject,omitempty"`
	Date                *string              `json:"date,omitempty"`
	Reporter            *Reference           `json:"reporter,omitempty"`
	Period              Period               `json:"period"`
	ImprovementNotation *CodeableConcept     `json:"improvementNotation,omitempty"`
	Group               []MeasureReportGroup `json:"group,omitempty"`
	EvaluatedResource   []Reference          `json:"evaluatedResource,omitempty"`
}

// The results of the calculation, one for each population group in the measure.
type MeasureReportGroup struct {
	ID                *string                        `json:"id,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept               `json:"code,omitempty"`
	Population        []MeasureReportGroupPopulation `json:"population,omitempty"`
	MeasureScore      *Quantity                      `json:"measureScore,omitempty"`
	Stratifier        []MeasureReportGroupStratifier `json:"stratifier,omitempty"`
}

// The populations that make up the population group, one for each type of population appropriate for the measure.
type MeasureReportGroupPopulation struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Count             *int             `json:"count,omitempty"`
	SubjectResults    *Reference       `json:"subjectResults,omitempty"`
}

// When a measure includes multiple stratifiers, there will be a stratifier group for each stratifier defined by the measure.
type MeasureReportGroupStratifier struct {
	ID                *string                               `json:"id,omitempty"`
	Extension         []Extension                           `json:"extension,omitempty"`
	ModifierExtension []Extension                           `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept                     `json:"code,omitempty"`
	Stratum           []MeasureReportGroupStratifierStratum `json:"stratum,omitempty"`
}

// This element contains the results for a single stratum within the stratifier. For example, when stratifying on administrative gender, there will be four strata, one for each possible gender value.
type MeasureReportGroupStratifierStratum struct {
	ID                *string                                         `json:"id,omitempty"`
	Extension         []Extension                                     `json:"extension,omitempty"`
	ModifierExtension []Extension                                     `json:"modifierExtension,omitempty"`
	Value             *CodeableConcept                                `json:"value,omitempty"`
	Component         []MeasureReportGroupStratifierStratumComponent  `json:"component,omitempty"`
	Population        []MeasureReportGroupStratifierStratumPopulation `json:"population,omitempty"`
	MeasureScore      *Quantity                                       `json:"measureScore,omitempty"`
}

// A stratifier component value.
type MeasureReportGroupStratifierStratumComponent struct {
	ID                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
	Value             CodeableConcept `json:"value"`
}

// The populations that make up the stratum, one for each type of population appropriate to the measure.
type MeasureReportGroupStratifierStratumPopulation struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Count             *int             `json:"count,omitempty"`
	SubjectResults    *Reference       `json:"subjectResults,omitempty"`
}

// This function returns resource reference information
func (r MeasureReport) ResourceRef() (string, *string) {
	return "MeasureReport", r.ID
}

// This function returns resource reference information
func (r MeasureReport) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMeasureReport MeasureReport

// MarshalJSON marshals the given MeasureReport as JSON into a byte slice
func (r MeasureReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMeasureReport
		ResourceType string `json:"resourceType"`
	}{
		OtherMeasureReport: OtherMeasureReport(r),
		ResourceType:       "MeasureReport",
	})
}

// UnmarshalMeasureReport unmarshals a MeasureReport.
func UnmarshalMeasureReport(b []byte) (MeasureReport, error) {
	var measureReport MeasureReport
	if err := json.Unmarshal(b, &measureReport); err != nil {
		return measureReport, err
	}
	return measureReport, nil
}
