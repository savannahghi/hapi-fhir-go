
package fhir430

import "encoding/json"
// RiskEvidenceSynthesis is documented here http://hl7.org/fhir/StructureDefinition/RiskEvidenceSynthesis
// The RiskEvidenceSynthesis resource describes the likelihood of an outcome in a population plus exposure state where the risk estimate is derived from a combination of research studies.
type RiskEvidenceSynthesis struct {
	ID                *string                            `json:"ID,omitempty"`
	Meta              *Meta                              `json:"meta,omitempty"`
	ImplicitRules     *string                            `json:"implicitRules,omitempty"`
	Language          *string                            `json:"language,omitempty"`
	Text              *Narrative                         `json:"text,omitempty"`
	Contained         []json.RawMessage                  `json:"contained,omitempty"`
	Extension         []Extension                        `json:"extension,omitempty"`
	ModifierExtension []Extension                        `json:"modifierExtension,omitempty"`
	Url               *string                            `json:"url,omitempty"`
	Identifier        []Identifier                       `json:"identifier,omitempty"`
	Version           *string                            `json:"version,omitempty"`
	Name              *string                            `json:"name,omitempty"`
	Title             *string                            `json:"title,omitempty"`
	Status            PublicationStatus                  `json:"status"`
	Date              *string                            `json:"date,omitempty"`
	Publisher         *string                            `json:"publisher,omitempty"`
	Contact           []ContactDetail                    `json:"contact,omitempty"`
	Description       *string                            `json:"description,omitempty"`
	Note              []Annotation                       `json:"note,omitempty"`
	UseContext        []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                  `json:"jurisdiction,omitempty"`
	ApprovalDate      *string                            `json:"approvalDate,omitempty"`
	LastReviewDate    *string                            `json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                            `json:"effectivePeriod,omitempty"`
	Topic             []CodeableConcept                  `json:"topic,omitempty"`
	Author            []ContactDetail                    `json:"author,omitempty"`
	Editor            []ContactDetail                    `json:"editor,omitempty"`
	Reviewer          []ContactDetail                    `json:"reviewer,omitempty"`
	Endorser          []ContactDetail                    `json:"endorser,omitempty"`
	RelatedArtifact   []RelatedArtifact                  `json:"relatedArtifact,omitempty"`
	SynthesisType     *CodeableConcept                   `json:"synthesisType,omitempty"`
	StudyType         *CodeableConcept                   `json:"studyType,omitempty"`
	Population        Reference                          `json:"population"`
	Exposure          *Reference                         `json:"exposure,omitempty"`
	Outcome           Reference                          `json:"outcome"`
	SampleSize        *RiskEvidenceSynthesisSampleSize   `json:"sampleSize,omitempty"`
	RiskEstimate      *RiskEvidenceSynthesisRiskEstimate `json:"riskEstimate,omitempty"`
	Certainty         []RiskEvidenceSynthesisCertainty   `json:"certainty,omitempty"`
}

// A description of the size of the sample involved in the synthesis.
type RiskEvidenceSynthesisSampleSize struct {
	ID                   *string     `json:"ID,omitempty"`
	Extension            []Extension `json:"extension,omitempty"`
	ModifierExtension    []Extension `json:"modifierExtension,omitempty"`
	Description          *string     `json:"description,omitempty"`
	NumberOfStudies      *int        `json:"numberOfStudies,omitempty"`
	NumberOfParticipants *int        `json:"numberOfParticipants,omitempty"`
}

// The estimated risk of the outcome.
type RiskEvidenceSynthesisRiskEstimate struct {
	ID                *string                                              `json:"ID,omitempty"`
	Extension         []Extension                                          `json:"extension,omitempty"`
	ModifierExtension []Extension                                          `json:"modifierExtension,omitempty"`
	Description       *string                                              `json:"description,omitempty"`
	Type              *CodeableConcept                                     `json:"type,omitempty"`
	Value             *json.Number                                         `json:"value,omitempty"`
	UnitOfMeasure     *CodeableConcept                                     `json:"unitOfMeasure,omitempty"`
	DenominatorCount  *int                                                 `json:"denominatorCount,omitempty"`
	NumeratorCount    *int                                                 `json:"numeratorCount,omitempty"`
	PrecisionEstimate []RiskEvidenceSynthesisRiskEstimatePrecisionEstimate `json:"precisionEstimate,omitempty"`
}

// A description of the precision of the estimate for the effect.
type RiskEvidenceSynthesisRiskEstimatePrecisionEstimate struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	Level             *json.Number     `json:"level,omitempty"`
	From              *json.Number     `json:"from,omitempty"`
	To                *json.Number     `json:"to,omitempty"`
}

// A description of the certainty of the risk estimate.
type RiskEvidenceSynthesisCertainty struct {
	ID                    *string                                               `json:"ID,omitempty"`
	Extension             []Extension                                           `json:"extension,omitempty"`
	ModifierExtension     []Extension                                           `json:"modifierExtension,omitempty"`
	Rating                []CodeableConcept                                     `json:"rating,omitempty"`
	Note                  []Annotation                                          `json:"note,omitempty"`
	CertaintySubcomponent []RiskEvidenceSynthesisCertaintyCertaintySubcomponent `json:"certaintySubcomponent,omitempty"`
}

// A description of a component of the overall certainty.
type RiskEvidenceSynthesisCertaintyCertaintySubcomponent struct {
	ID                *string           `json:"ID,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `json:"type,omitempty"`
	Rating            []CodeableConcept `json:"rating,omitempty"`
	Note              []Annotation      `json:"note,omitempty"`
}

// This function returns resource reference information
func (r RiskEvidenceSynthesis) ResourceRef() (string, *string) {
	return "RiskEvidenceSynthesis", r.ID
}

// This function returns resource reference information
func (r RiskEvidenceSynthesis) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherRiskEvidenceSynthesis RiskEvidenceSynthesis

// MarshalJSON marshals the given RiskEvidenceSynthesis as JSON into a byte slice
func (r RiskEvidenceSynthesis) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherRiskEvidenceSynthesis
		ResourceType string `json:"resourceType"`
	}{
		OtherRiskEvidenceSynthesis: OtherRiskEvidenceSynthesis(r),
		ResourceType:               "RiskEvidenceSynthesis",
	})
}

// UnmarshalRiskEvidenceSynthesis unmarshals a RiskEvidenceSynthesis.
func UnmarshalRiskEvidenceSynthesis(b []byte) (RiskEvidenceSynthesis, error) {
	var riskEvidenceSynthesis RiskEvidenceSynthesis
	if err := json.Unmarshal(b, &riskEvidenceSynthesis); err != nil {
		return riskEvidenceSynthesis, err
	}
	return riskEvidenceSynthesis, nil
}
