
package fhir430

import "encoding/json"
// ResearchElementDefinition is documented here http://hl7.org/fhir/StructureDefinition/ResearchElementDefinition
// The ResearchElementDefinition resource describes a "PICO" element that knowledge (evidence, assertion, recommendation) is about.
type ResearchElementDefinition struct {
	ID                     *string                                   `json:"ID,omitempty"`
	Meta                   *Meta                                     `json:"meta,omitempty"`
	ImplicitRules          *string                                   `json:"implicitRules,omitempty"`
	Language               *string                                   `json:"language,omitempty"`
	Text                   *Narrative                                `json:"text,omitempty"`
	Contained              []json.RawMessage                         `json:"contained,omitempty"`
	Extension              []Extension                               `json:"extension,omitempty"`
	ModifierExtension      []Extension                               `json:"modifierExtension,omitempty"`
	Url                    *string                                   `json:"url,omitempty"`
	Identifier             []Identifier                              `json:"identifier,omitempty"`
	Version                *string                                   `json:"version,omitempty"`
	Name                   *string                                   `json:"name,omitempty"`
	Title                  *string                                   `json:"title,omitempty"`
	ShortTitle             *string                                   `json:"shortTitle,omitempty"`
	Subtitle               *string                                   `json:"subtitle,omitempty"`
	Status                 PublicationStatus                         `json:"status"`
	Experimental           *bool                                     `json:"experimental,omitempty"`
	SubjectCodeableConcept *CodeableConcept                          `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference                                `json:"subjectReference,omitempty"`
	Date                   *string                                   `json:"date,omitempty"`
	Publisher              *string                                   `json:"publisher,omitempty"`
	Contact                []ContactDetail                           `json:"contact,omitempty"`
	Description            *string                                   `json:"description,omitempty"`
	Comment                []string                                  `json:"comment,omitempty"`
	UseContext             []UsageContext                            `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept                         `json:"jurisdiction,omitempty"`
	Purpose                *string                                   `json:"purpose,omitempty"`
	Usage                  *string                                   `json:"usage,omitempty"`
	ApprovalDate           *string                                   `json:"approvalDate,omitempty"`
	LastReviewDate         *string                                   `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period                                   `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept                         `json:"topic,omitempty"`
	Author                 []ContactDetail                           `json:"author,omitempty"`
	Editor                 []ContactDetail                           `json:"editor,omitempty"`
	Reviewer               []ContactDetail                           `json:"reviewer,omitempty"`
	Endorser               []ContactDetail                           `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact                         `json:"relatedArtifact,omitempty"`
	Library                []string                                  `json:"library,omitempty"`
	Type                   ResearchElementType                       `json:"type"`
	VariableType           *EvidenceVariableType                     `json:"variableType,omitempty"`
	Characteristic         []ResearchElementDefinitionCharacteristic `json:"characteristic"`
}

// A characteristic that defines the members of the research element. Multiple characteristics are applied with "and" semantics.
// Characteristics can be defined flexibly to accommodate different use cases for membership criteria, ranging from simple codes, all the way to using an expression language to express the criteria.
type ResearchElementDefinitionCharacteristic struct {
	ID                                *string          `json:"ID,omitempty"`
	Extension                         []Extension      `json:"extension,omitempty"`
	ModifierExtension                 []Extension      `json:"modifierExtension,omitempty"`
	DefinitionCodeableConcept         CodeableConcept  `json:"definitionCodeableConcept"`
	DefinitionCanonical               string           `json:"definitionCanonical"`
	DefinitionExpression              Expression       `json:"definitionExpression"`
	DefinitionDataRequirement         DataRequirement  `json:"definitionDataRequirement"`
	UsageContext                      []UsageContext   `json:"usageContext,omitempty"`
	Exclude                           *bool            `json:"exclude,omitempty"`
	UnitOfMeasure                     *CodeableConcept `json:"unitOfMeasure,omitempty"`
	StudyEffectiveDescription         *string          `json:"studyEffectiveDescription,omitempty"`
	StudyEffectiveDateTime            *string          `json:"studyEffectiveDateTime,omitempty"`
	StudyEffectivePeriod              *Period          `json:"studyEffectivePeriod,omitempty"`
	StudyEffectiveDuration            *Duration        `json:"studyEffectiveDuration,omitempty"`
	StudyEffectiveTiming              *Timing          `json:"studyEffectiveTiming,omitempty"`
	StudyEffectiveTimeFromStart       *Duration        `json:"studyEffectiveTimeFromStart,omitempty"`
	StudyEffectiveGroupMeasure        *GroupMeasure    `json:"studyEffectiveGroupMeasure,omitempty"`
	ParticipantEffectiveDescription   *string          `json:"participantEffectiveDescription,omitempty"`
	ParticipantEffectiveDateTime      *string          `json:"participantEffectiveDateTime,omitempty"`
	ParticipantEffectivePeriod        *Period          `json:"participantEffectivePeriod,omitempty"`
	ParticipantEffectiveDuration      *Duration        `json:"participantEffectiveDuration,omitempty"`
	ParticipantEffectiveTiming        *Timing          `json:"participantEffectiveTiming,omitempty"`
	ParticipantEffectiveTimeFromStart *Duration        `json:"participantEffectiveTimeFromStart,omitempty"`
	ParticipantEffectiveGroupMeasure  *GroupMeasure    `json:"participantEffectiveGroupMeasure,omitempty"`
}

// This function returns resource reference information
func (r ResearchElementDefinition) ResourceRef() (string, *string) {
	return "ResearchElementDefinition", r.ID
}

// This function returns resource reference information
func (r ResearchElementDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherResearchElementDefinition ResearchElementDefinition

// MarshalJSON marshals the given ResearchElementDefinition as JSON into a byte slice
func (r ResearchElementDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResearchElementDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherResearchElementDefinition: OtherResearchElementDefinition(r),
		ResourceType:                   "ResearchElementDefinition",
	})
}

// UnmarshalResearchElementDefinition unmarshals a ResearchElementDefinition.
func UnmarshalResearchElementDefinition(b []byte) (ResearchElementDefinition, error) {
	var researchElementDefinition ResearchElementDefinition
	if err := json.Unmarshal(b, &researchElementDefinition); err != nil {
		return researchElementDefinition, err
	}
	return researchElementDefinition, nil
}
