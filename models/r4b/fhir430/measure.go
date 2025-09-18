
package fhir430

import "encoding/json"
// Measure is documented here http://hl7.org/fhir/StructureDefinition/Measure
// The Measure resource provides the definition of a quality measure.
type Measure struct {
	ID                              *string                   `json:"ID,omitempty"`
	Meta                            *Meta                     `json:"meta,omitempty"`
	ImplicitRules                   *string                   `json:"implicitRules,omitempty"`
	Language                        *string                   `json:"language,omitempty"`
	Text                            *Narrative                `json:"text,omitempty"`
	Contained                       []json.RawMessage         `json:"contained,omitempty"`
	Extension                       []Extension               `json:"extension,omitempty"`
	ModifierExtension               []Extension               `json:"modifierExtension,omitempty"`
	Url                             *string                   `json:"url,omitempty"`
	Identifier                      []Identifier              `json:"identifier,omitempty"`
	Version                         *string                   `json:"version,omitempty"`
	Name                            *string                   `json:"name,omitempty"`
	Title                           *string                   `json:"title,omitempty"`
	Subtitle                        *string                   `json:"subtitle,omitempty"`
	Status                          PublicationStatus         `json:"status"`
	Experimental                    *bool                     `json:"experimental,omitempty"`
	SubjectCodeableConcept          *CodeableConcept          `json:"subjectCodeableConcept,omitempty"`
	SubjectReference                *Reference                `json:"subjectReference,omitempty"`
	Date                            *string                   `json:"date,omitempty"`
	Publisher                       *string                   `json:"publisher,omitempty"`
	Contact                         []ContactDetail           `json:"contact,omitempty"`
	Description                     *string                   `json:"description,omitempty"`
	UseContext                      []UsageContext            `json:"useContext,omitempty"`
	Jurisdiction                    []CodeableConcept         `json:"jurisdiction,omitempty"`
	Purpose                         *string                   `json:"purpose,omitempty"`
	Usage                           *string                   `json:"usage,omitempty"`
	ApprovalDate                    *string                   `json:"approvalDate,omitempty"`
	LastReviewDate                  *string                   `json:"lastReviewDate,omitempty"`
	EffectivePeriod                 *Period                   `json:"effectivePeriod,omitempty"`
	Topic                           []CodeableConcept         `json:"topic,omitempty"`
	Author                          []ContactDetail           `json:"author,omitempty"`
	Editor                          []ContactDetail           `json:"editor,omitempty"`
	Reviewer                        []ContactDetail           `json:"reviewer,omitempty"`
	Endorser                        []ContactDetail           `json:"endorser,omitempty"`
	RelatedArtifact                 []RelatedArtifact         `json:"relatedArtifact,omitempty"`
	Library                         []string                  `json:"library,omitempty"`
	Disclaimer                      *string                   `json:"disclaimer,omitempty"`
	Scoring                         *CodeableConcept          `json:"scoring,omitempty"`
	CompositeScoring                *CodeableConcept          `json:"compositeScoring,omitempty"`
	Type                            []CodeableConcept         `json:"type,omitempty"`
	RiskAdjustment                  *string                   `json:"riskAdjustment,omitempty"`
	RateAggregation                 *string                   `json:"rateAggregation,omitempty"`
	Rationale                       *string                   `json:"rationale,omitempty"`
	ClinicalRecommendationStatement *string                   `json:"clinicalRecommendationStatement,omitempty"`
	ImprovementNotation             *CodeableConcept          `json:"improvementNotation,omitempty"`
	Definition                      []string                  `json:"definition,omitempty"`
	Guidance                        *string                   `json:"guidance,omitempty"`
	Group                           []MeasureGroup            `json:"group,omitempty"`
	SupplementalData                []MeasureSupplementalData `json:"supplementalData,omitempty"`
}

// A group of population criteria for the measure.
type MeasureGroup struct {
	ID                *string                  `json:"ID,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept         `json:"code,omitempty"`
	Description       *string                  `json:"description,omitempty"`
	Population        []MeasureGroupPopulation `json:"population,omitempty"`
	Stratifier        []MeasureGroupStratifier `json:"stratifier,omitempty"`
}

// A population criteria for the measure.
type MeasureGroupPopulation struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Criteria          Expression       `json:"criteria"`
}

// The stratifier criteria for the measure report, specified as either the name of a valid CQL expression defined within a referenced library or a valid FHIR Resource Path.
type MeasureGroupStratifier struct {
	ID                *string                           `json:"ID,omitempty"`
	Extension         []Extension                       `json:"extension,omitempty"`
	ModifierExtension []Extension                       `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept                  `json:"code,omitempty"`
	Description       *string                           `json:"description,omitempty"`
	Criteria          *Expression                       `json:"criteria,omitempty"`
	Component         []MeasureGroupStratifierComponent `json:"component,omitempty"`
}

// A component of the stratifier criteria for the measure report, specified as either the name of a valid CQL expression defined within a referenced library or a valid FHIR Resource Path.
// Stratifiers are defined either as a single criteria, or as a set of component criteria.
type MeasureGroupStratifierComponent struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `json:"code,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Criteria          Expression       `json:"criteria"`
}

// The supplemental data criteria for the measure report, specified as either the name of a valid CQL expression within a referenced library, or a valid FHIR Resource Path.
// Note that supplemental data are reported as observations for each patient and included in the evaluatedResources bundle. See the MeasureReport resource or the Quality Reporting topic for more information.
type MeasureSupplementalData struct {
	ID                *string           `json:"ID,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              *CodeableConcept  `json:"code,omitempty"`
	Usage             []CodeableConcept `json:"usage,omitempty"`
	Description       *string           `json:"description,omitempty"`
	Criteria          Expression        `json:"criteria"`
}

// This function returns resource reference information
func (r Measure) ResourceRef() (string, *string) {
	return "Measure", r.ID
}

// This function returns resource reference information
func (r Measure) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMeasure Measure

// MarshalJSON marshals the given Measure as JSON into a byte slice
func (r Measure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMeasure
		ResourceType string `json:"resourceType"`
	}{
		OtherMeasure: OtherMeasure(r),
		ResourceType: "Measure",
	})
}

// UnmarshalMeasure unmarshals a Measure.
func UnmarshalMeasure(b []byte) (Measure, error) {
	var measure Measure
	if err := json.Unmarshal(b, &measure); err != nil {
		return measure, err
	}
	return measure, nil
}
