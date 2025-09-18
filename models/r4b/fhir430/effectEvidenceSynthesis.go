
package fhir430

import "encoding/json"
// EffectEvidenceSynthesis is documented here http://hl7.org/fhir/StructureDefinition/EffectEvidenceSynthesis
// The EffectEvidenceSynthesis resource describes the difference in an outcome between exposures states in a population where the effect estimate is derived from a combination of research studies.
type EffectEvidenceSynthesis struct {
	ID                  *string                                    `json:"ID,omitempty"`
	Meta                *Meta                                      `json:"meta,omitempty"`
	ImplicitRules       *string                                    `json:"implicitRules,omitempty"`
	Language            *string                                    `json:"language,omitempty"`
	Text                *Narrative                                 `json:"text,omitempty"`
	Contained           []json.RawMessage                          `json:"contained,omitempty"`
	Extension           []Extension                                `json:"extension,omitempty"`
	ModifierExtension   []Extension                                `json:"modifierExtension,omitempty"`
	Url                 *string                                    `json:"url,omitempty"`
	Identifier          []Identifier                               `json:"identifier,omitempty"`
	Version             *string                                    `json:"version,omitempty"`
	Name                *string                                    `json:"name,omitempty"`
	Title               *string                                    `json:"title,omitempty"`
	Status              PublicationStatus                          `json:"status"`
	Date                *string                                    `json:"date,omitempty"`
	Publisher           *string                                    `json:"publisher,omitempty"`
	Contact             []ContactDetail                            `json:"contact,omitempty"`
	Description         *string                                    `json:"description,omitempty"`
	Note                []Annotation                               `json:"note,omitempty"`
	UseContext          []UsageContext                             `json:"useContext,omitempty"`
	Jurisdiction        []CodeableConcept                          `json:"jurisdiction,omitempty"`
	ApprovalDate        *string                                    `json:"approvalDate,omitempty"`
	LastReviewDate      *string                                    `json:"lastReviewDate,omitempty"`
	EffectivePeriod     *Period                                    `json:"effectivePeriod,omitempty"`
	Topic               []CodeableConcept                          `json:"topic,omitempty"`
	Author              []ContactDetail                            `json:"author,omitempty"`
	Editor              []ContactDetail                            `json:"editor,omitempty"`
	Reviewer            []ContactDetail                            `json:"reviewer,omitempty"`
	Endorser            []ContactDetail                            `json:"endorser,omitempty"`
	RelatedArtifact     []RelatedArtifact                          `json:"relatedArtifact,omitempty"`
	SynthesisType       *CodeableConcept                           `json:"synthesisType,omitempty"`
	StudyType           *CodeableConcept                           `json:"studyType,omitempty"`
	Population          Reference                                  `json:"population"`
	Exposure            Reference                                  `json:"exposure"`
	ExposureAlternative Reference                                  `json:"exposureAlternative"`
	Outcome             Reference                                  `json:"outcome"`
	SampleSize          *EffectEvidenceSynthesisSampleSize         `json:"sampleSize,omitempty"`
	ResultsByExposure   []EffectEvidenceSynthesisResultsByExposure `json:"resultsByExposure,omitempty"`
	EffectEstimate      []EffectEvidenceSynthesisEffectEstimate    `json:"effectEstimate,omitempty"`
	Certainty           []EffectEvidenceSynthesisCertainty         `json:"certainty,omitempty"`
}

// A description of the size of the sample involved in the synthesis.
type EffectEvidenceSynthesisSampleSize struct {
	ID                   *string     `json:"ID,omitempty"`
	Extension            []Extension `json:"extension,omitempty"`
	ModifierExtension    []Extension `json:"modifierExtension,omitempty"`
	Description          *string     `json:"description,omitempty"`
	NumberOfStudies      *int        `json:"numberOfStudies,omitempty"`
	NumberOfParticipants *int        `json:"numberOfParticipants,omitempty"`
}

// A description of the results for each exposure considered in the effect estimate.
type EffectEvidenceSynthesisResultsByExposure struct {
	ID                    *string          `json:"ID,omitempty"`
	Extension             []Extension      `json:"extension,omitempty"`
	ModifierExtension     []Extension      `json:"modifierExtension,omitempty"`
	Description           *string          `json:"description,omitempty"`
	ExposureState         *ExposureState   `json:"exposureState,omitempty"`
	VariantState          *CodeableConcept `json:"variantState,omitempty"`
	RiskEvidenceSynthesis Reference        `json:"riskEvidenceSynthesis"`
}

// The estimated effect of the exposure variant.
type EffectEvidenceSynthesisEffectEstimate struct {
	ID                *string                                                  `json:"ID,omitempty"`
	Extension         []Extension                                              `json:"extension,omitempty"`
	ModifierExtension []Extension                                              `json:"modifierExtension,omitempty"`
	Description       *string                                                  `json:"description,omitempty"`
	Type              *CodeableConcept                                         `json:"type,omitempty"`
	VariantState      *CodeableConcept                                         `json:"variantState,omitempty"`
	Value             *json.Number                                             `json:"value,omitempty"`
	UnitOfMeasure     *CodeableConcept                                         `json:"unitOfMeasure,omitempty"`
	PrecisionEstimate []EffectEvidenceSynthesisEffectEstimatePrecisionEstimate `json:"precisionEstimate,omitempty"`
}

// A description of the precision of the estimate for the effect.
type EffectEvidenceSynthesisEffectEstimatePrecisionEstimate struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Level             *json.Number     `json:"level,omitempty"`
	From              *json.Number     `json:"from,omitempty"`
	To                *json.Number     `json:"to,omitempty"`
}

// A description of the certainty of the effect estimate.
type EffectEvidenceSynthesisCertainty struct {
	ID                    *string                                                 `json:"ID,omitempty"`
	Extension             []Extension                                             `json:"extension,omitempty"`
	ModifierExtension     []Extension                                             `json:"modifierExtension,omitempty"`
	Rating                []CodeableConcept                                       `json:"rating,omitempty"`
	Note                  []Annotation                                            `json:"note,omitempty"`
	CertaintySubcomponent []EffectEvidenceSynthesisCertaintyCertaintySubcomponent `json:"certaintySubcomponent,omitempty"`
}

// A description of a component of the overall certainty.
type EffectEvidenceSynthesisCertaintyCertaintySubcomponent struct {
	ID                *string           `json:"ID,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Rating            []CodeableConcept `json:"rating,omitempty"`
	Note              []Annotation      `json:"note,omitempty"`
}

// This function returns resource reference information
func (r EffectEvidenceSynthesis) ResourceRef() (string, *string) {
	return "EffectEvidenceSynthesis", r.ID
}

// This function returns resource reference information
func (r EffectEvidenceSynthesis) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherEffectEvidenceSynthesis EffectEvidenceSynthesis

// MarshalJSON marshals the given EffectEvidenceSynthesis as JSON into a byte slice
func (r EffectEvidenceSynthesis) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEffectEvidenceSynthesis
		ResourceType string `json:"resourceType"`
	}{
		OtherEffectEvidenceSynthesis: OtherEffectEvidenceSynthesis(r),
		ResourceType:                 "EffectEvidenceSynthesis",
	})
}

// UnmarshalEffectEvidenceSynthesis unmarshals a EffectEvidenceSynthesis.
func UnmarshalEffectEvidenceSynthesis(b []byte) (EffectEvidenceSynthesis, error) {
	var effectEvidenceSynthesis EffectEvidenceSynthesis
	if err := json.Unmarshal(b, &effectEvidenceSynthesis); err != nil {
		return effectEvidenceSynthesis, err
	}
	return effectEvidenceSynthesis, nil
}
