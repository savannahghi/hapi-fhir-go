package fhir430

import "encoding/json"

// ResearchDefinition is documented here http://hl7.org/fhir/StructureDefinition/ResearchDefinition
// The ResearchDefinition resource describes the conditional state (population and any exposures being compared within the population) and outcome (if specified) that the knowledge (evidence, assertion, recommendation) is about.
type ResearchDefinition struct {
	ID                     *string           `json:"id,omitempty"`
	Meta                   *Meta             `json:"meta,omitempty"`
	ImplicitRules          *string           `json:"implicitRules,omitempty"`
	Language               *string           `json:"language,omitempty"`
	Text                   *Narrative        `json:"text,omitempty"`
	Contained              []json.RawMessage `json:"contained,omitempty"`
	Extension              []Extension       `json:"extension,omitempty"`
	ModifierExtension      []Extension       `json:"modifierExtension,omitempty"`
	Url                    *string           `json:"url,omitempty"`
	Identifier             []Identifier      `json:"identifier,omitempty"`
	Version                *string           `json:"version,omitempty"`
	Name                   *string           `json:"name,omitempty"`
	Title                  *string           `json:"title,omitempty"`
	ShortTitle             *string           `json:"shortTitle,omitempty"`
	Subtitle               *string           `json:"subtitle,omitempty"`
	Status                 PublicationStatus `json:"status"`
	Experimental           *bool             `json:"experimental,omitempty"`
	SubjectCodeableConcept *CodeableConcept  `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference        `json:"subjectReference,omitempty"`
	Date                   *string           `json:"date,omitempty"`
	Publisher              *string           `json:"publisher,omitempty"`
	Contact                []ContactDetail   `json:"contact,omitempty"`
	Description            *string           `json:"description,omitempty"`
	Comment                []string          `json:"comment,omitempty"`
	UseContext             []UsageContext    `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept `json:"jurisdiction,omitempty"`
	Purpose                *string           `json:"purpose,omitempty"`
	Usage                  *string           `json:"usage,omitempty"`
	ApprovalDate           *string           `json:"approvalDate,omitempty"`
	LastReviewDate         *string           `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period           `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept `json:"topic,omitempty"`
	Author                 []ContactDetail   `json:"author,omitempty"`
	Editor                 []ContactDetail   `json:"editor,omitempty"`
	Reviewer               []ContactDetail   `json:"reviewer,omitempty"`
	Endorser               []ContactDetail   `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact `json:"relatedArtifact,omitempty"`
	Library                []string          `json:"library,omitempty"`
	Population             Reference         `json:"population"`
	Exposure               *Reference        `json:"exposure,omitempty"`
	ExposureAlternative    *Reference        `json:"exposureAlternative,omitempty"`
	Outcome                *Reference        `json:"outcome,omitempty"`
}

// This function returns resource reference information
func (r ResearchDefinition) ResourceRef() (string, *string) {
	return "ResearchDefinition", r.ID
}

// This function returns resource reference information
func (r ResearchDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherResearchDefinition ResearchDefinition

// MarshalJSON marshals the given ResearchDefinition as JSON into a byte slice
func (r ResearchDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchDefinition: OtherResearchDefinition(r),
		ResourceType:            "ResearchDefinition",
	})
}

// UnmarshalResearchDefinition unmarshals a ResearchDefinition.
func UnmarshalResearchDefinition(b []byte) (ResearchDefinition, error) {
	var researchDefinition ResearchDefinition
	if err := json.Unmarshal(b, &researchDefinition); err != nil {
		return researchDefinition, err
	}
	return researchDefinition, nil
}
