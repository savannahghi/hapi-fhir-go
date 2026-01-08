package fhir430

import "encoding/json"

// ClinicalImpression is documented here http://hl7.org/fhir/StructureDefinition/ClinicalImpression
// A record of a clinical assessment performed to determine what problem(s) may affect the patient and before planning the treatments or management strategies that are best to manage a patient's condition. Assessments are often 1:1 with a clinical consultation / encounter,  but this varies greatly depending on the clinical workflow. This resource is called "ClinicalImpression" rather than "ClinicalAssessment" to avoid confusion with the recording of assessment tools such as Apgar score.
type ClinicalImpression struct {
	ID                       *string                           `json:"id,omitempty"`
	Meta                     *Meta                             `json:"meta,omitempty"`
	ImplicitRules            *string                           `json:"implicitRules,omitempty"`
	Language                 *string                           `json:"language,omitempty"`
	Text                     *Narrative                        `json:"text,omitempty"`
	Contained                []json.RawMessage                 `json:"contained,omitempty"`
	Extension                []Extension                       `json:"extension,omitempty"`
	ModifierExtension        []Extension                       `json:"modifierExtension,omitempty"`
	Identifier               []Identifier                      `json:"identifier,omitempty"`
	Status                   ClinicalImpressionStatus          `json:"status"`
	StatusReason             *CodeableConcept                  `json:"statusReason,omitempty"`
	Code                     *CodeableConcept                  `json:"code,omitempty"`
	Description              *string                           `json:"description,omitempty"`
	Subject                  Reference                         `json:"subject"`
	Encounter                *Reference                        `json:"encounter,omitempty"`
	EffectiveDateTime        *string                           `json:"effectiveDateTime,omitempty"`
	EffectivePeriod          *Period                           `json:"effectivePeriod,omitempty"`
	Date                     *string                           `json:"date,omitempty"`
	Assessor                 *Reference                        `json:"assessor,omitempty"`
	Previous                 *Reference                        `json:"previous,omitempty"`
	Problem                  []Reference                       `json:"problem,omitempty"`
	Investigation            []ClinicalImpressionInvestigation `json:"investigation,omitempty"`
	Protocol                 []string                          `json:"protocol,omitempty"`
	Summary                  *string                           `json:"summary,omitempty"`
	Finding                  []ClinicalImpressionFinding       `json:"finding,omitempty"`
	PrognosisCodeableConcept []CodeableConcept                 `json:"prognosisCodeableConcept,omitempty"`
	PrognosisReference       []Reference                       `json:"prognosisReference,omitempty"`
	SupportingInfo           []Reference                       `json:"supportingInfo,omitempty"`
	Note                     []Annotation                      `json:"note,omitempty"`
}

// One or more sets of investigations (signs, symptoms, etc.). The actual grouping of investigations varies greatly depending on the type and context of the assessment. These investigations may include data generated during the assessment process, or data previously generated and recorded that is pertinent to the outcomes.
type ClinicalImpressionInvestigation struct {
	ID                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Code              CodeableConcept `json:"code"`
	Item              []Reference     `json:"item,omitempty"`
}

// Specific findings or diagnoses that were considered likely or relevant to ongoing treatment.
type ClinicalImpressionFinding struct {
	ID                  *string          `json:"id,omitempty"`
	Extension           []Extension      `json:"extension,omitempty"`
	ModifierExtension   []Extension      `json:"modifierExtension,omitempty"`
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `json:"itemReference,omitempty"`
	Basis               *string          `json:"basis,omitempty"`
}

// This function returns resource reference information
func (r ClinicalImpression) ResourceRef() (string, *string) {
	return "ClinicalImpression", r.ID
}

// This function returns resource reference information
func (r ClinicalImpression) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherClinicalImpression ClinicalImpression

// MarshalJSON marshals the given ClinicalImpression as JSON into a byte slice
func (r ClinicalImpression) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherClinicalImpression
		ResourceType string `json:"resourceType"`
	}{
		OtherClinicalImpression: OtherClinicalImpression(r),
		ResourceType:            "ClinicalImpression",
	})
}

// UnmarshalClinicalImpression unmarshals a ClinicalImpression.
func UnmarshalClinicalImpression(b []byte) (ClinicalImpression, error) {
	var clinicalImpression ClinicalImpression
	if err := json.Unmarshal(b, &clinicalImpression); err != nil {
		return clinicalImpression, err
	}
	return clinicalImpression, nil
}
