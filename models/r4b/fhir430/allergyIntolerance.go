package fhir430

import "encoding/json"

// AllergyIntolerance is documented here http://hl7.org/fhir/StructureDefinition/AllergyIntolerance
// Risk of harmful or undesirable, physiological response which is unique to an individual and associated with exposure to a substance.
type AllergyIntolerance struct {
	ID                 *string                        `json:"ID,omitempty"`
	Meta               *Meta                          `json:"meta,omitempty"`
	ImplicitRules      *string                        `json:"implicitRules,omitempty"`
	Language           *string                        `json:"language,omitempty"`
	Text               *Narrative                     `json:"text,omitempty"`
	Contained          []json.RawMessage              `json:"contained,omitempty"`
	Extension          []Extension                    `json:"extension,omitempty"`
	ModifierExtension  []Extension                    `json:"modifierExtension,omitempty"`
	Identifier         []Identifier                   `json:"identifier,omitempty"`
	ClinicalStatus     *CodeableConcept               `json:"clinicalStatus,omitempty"`
	VerificationStatus *CodeableConcept               `json:"verificationStatus,omitempty"`
	Type               *AllergyIntoleranceType        `json:"type,omitempty"`
	Category           []AllergyIntoleranceCategory   `json:"category,omitempty"`
	Criticality        *AllergyIntoleranceCriticality `json:"criticality,omitempty"`
	Code               *CodeableConcept               `json:"code,omitempty"`
	Patient            *Reference                     `json:"patient"`
	Encounter          *Reference                     `json:"encounter,omitempty"`
	OnsetDateTime      *string                        `json:"onsetDateTime,omitempty"`
	OnsetAge           *Age                           `json:"onsetAge,omitempty"`
	OnsetPeriod        *Period                        `json:"onsetPeriod,omitempty"`
	OnsetRange         *Range                         `json:"onsetRange,omitempty"`
	OnsetString        *string                        `json:"onsetString,omitempty"`
	RecordedDate       *string                        `json:"recordedDate,omitempty"`
	Recorder           *Reference                     `json:"recorder,omitempty"`
	Asserter           *Reference                     `json:"asserter,omitempty"`
	LastOccurrence     *string                        `json:"lastOccurrence,omitempty"`
	Note               []Annotation                   `json:"note,omitempty"`
	Reaction           []AllergyIntoleranceReaction   `json:"reaction,omitempty"`
}

// Details about each adverse reaction event linked to exposure to the identified substance.
type AllergyIntoleranceReaction struct {
	ID                *string                     `json:"ID,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Substance         *CodeableConcept            `json:"substance,omitempty"`
	Manifestation     []CodeableConcept           `json:"manifestation"`
	Description       *string                     `json:"description,omitempty"`
	Onset             *string                     `json:"onset,omitempty"`
	Severity          *AllergyIntoleranceSeverity `json:"severity,omitempty"`
	ExposureRoute     *CodeableConcept            `json:"exposureRoute,omitempty"`
	Note              []Annotation                `json:"note,omitempty"`
}

// This function returns resource reference information
func (r AllergyIntolerance) ResourceRef() (string, *string) {
	return "AllergyIntolerance", r.ID
}

// This function returns resource reference information
func (r AllergyIntolerance) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherAllergyIntolerance AllergyIntolerance

// MarshalJSON marshals the given AllergyIntolerance as JSON into a byte slice
func (r AllergyIntolerance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAllergyIntolerance
		ResourceType string `json:"resourceType"`
	}{
		OtherAllergyIntolerance: OtherAllergyIntolerance(r),
		ResourceType:            "AllergyIntolerance",
	})
}

// UnmarshalAllergyIntolerance unmarshals a AllergyIntolerance.
func UnmarshalAllergyIntolerance(b []byte) (AllergyIntolerance, error) {
	var allergyIntolerance AllergyIntolerance
	if err := json.Unmarshal(b, &allergyIntolerance); err != nil {
		return allergyIntolerance, err
	}
	return allergyIntolerance, nil
}
