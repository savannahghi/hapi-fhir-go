
package fhir430

import "encoding/json"
// Evidence is documented here http://hl7.org/fhir/StructureDefinition/Evidence
// The Evidence resource describes the conditional state (population and any exposures being compared within the population) and outcome (if specified) that the knowledge (evidence, assertion, recommendation) is about.
type Evidence struct {
	ID                 *string           `json:"ID,omitempty"`
	Meta               *Meta             `json:"meta,omitempty"`
	ImplicitRules      *string           `json:"implicitRules,omitempty"`
	Language           *string           `json:"language,omitempty"`
	Text               *Narrative        `json:"text,omitempty"`
	Contained          []json.RawMessage `json:"contained,omitempty"`
	Extension          []Extension       `json:"extension,omitempty"`
	ModifierExtension  []Extension       `json:"modifierExtension,omitempty"`
	Url                *string           `json:"url,omitempty"`
	Identifier         []Identifier      `json:"identifier,omitempty"`
	Version            *string           `json:"version,omitempty"`
	Name               *string           `json:"name,omitempty"`
	Title              *string           `json:"title,omitempty"`
	ShortTitle         *string           `json:"shortTitle,omitempty"`
	Subtitle           *string           `json:"subtitle,omitempty"`
	Status             PublicationStatus `json:"status"`
	Date               *string           `json:"date,omitempty"`
	Publisher          *string           `json:"publisher,omitempty"`
	Contact            []ContactDetail   `json:"contact,omitempty"`
	Description        *string           `json:"description,omitempty"`
	Note               []Annotation      `json:"note,omitempty"`
	UseContext         []UsageContext    `json:"useContext,omitempty"`
	Jurisdiction       []CodeableConcept `json:"jurisdiction,omitempty"`
	ApprovalDate       *string           `json:"approvalDate,omitempty"`
	LastReviewDate     *string           `json:"lastReviewDate,omitempty"`
	EffectivePeriod    *Period           `json:"effectivePeriod,omitempty"`
	Topic              []CodeableConcept `json:"topic,omitempty"`
	Author             []ContactDetail   `json:"author,omitempty"`
	Editor             []ContactDetail   `json:"editor,omitempty"`
	Reviewer           []ContactDetail   `json:"reviewer,omitempty"`
	Endorser           []ContactDetail   `json:"endorser,omitempty"`
	RelatedArtifact    []RelatedArtifact `json:"relatedArtifact,omitempty"`
	ExposureBackground Reference         `json:"exposureBackground"`
	ExposureVariant    []Reference       `json:"exposureVariant,omitempty"`
	Outcome            []Reference       `json:"outcome,omitempty"`
}

// This function returns resource reference information
func (r Evidence) ResourceRef() (string, *string) {
	return "Evidence", r.ID
}

// This function returns resource reference information
func (r Evidence) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherEvidence Evidence

// MarshalJSON marshals the given Evidence as JSON into a byte slice
func (r Evidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidence
		ResourceType string `json:"resourceType"`
	}{
		OtherEvidence: OtherEvidence(r),
		ResourceType:  "Evidence",
	})
}

// UnmarshalEvidence unmarshals a Evidence.
func UnmarshalEvidence(b []byte) (Evidence, error) {
	var evidence Evidence
	if err := json.Unmarshal(b, &evidence); err != nil {
		return evidence, err
	}
	return evidence, nil
}
