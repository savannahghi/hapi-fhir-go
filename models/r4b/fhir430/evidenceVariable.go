package fhir430

import "encoding/json"

// EvidenceVariable is documented here http://hl7.org/fhir/StructureDefinition/EvidenceVariable
// The EvidenceVariable resource describes a "PICO" element that knowledge (evidence, assertion, recommendation) is about.
type EvidenceVariable struct {
	ID                *string                          `json:"id,omitempty"`
	Meta              *Meta                            `json:"meta,omitempty"`
	ImplicitRules     *string                          `json:"implicitRules,omitempty"`
	Language          *string                          `json:"language,omitempty"`
	Text              *Narrative                       `json:"text,omitempty"`
	Contained         []json.RawMessage                `json:"contained,omitempty"`
	Extension         []Extension                      `json:"extension,omitempty"`
	ModifierExtension []Extension                      `json:"modifierExtension,omitempty"`
	Url               *string                          `json:"url,omitempty"`
	Identifier        []Identifier                     `json:"identifier,omitempty"`
	Version           *string                          `json:"version,omitempty"`
	Name              *string                          `json:"name,omitempty"`
	Title             *string                          `json:"title,omitempty"`
	ShortTitle        *string                          `json:"shortTitle,omitempty"`
	Subtitle          *string                          `json:"subtitle,omitempty"`
	Status            PublicationStatus                `json:"status"`
	Date              *string                          `json:"date,omitempty"`
	Publisher         *string                          `json:"publisher,omitempty"`
	Contact           []ContactDetail                  `json:"contact,omitempty"`
	Description       *string                          `json:"description,omitempty"`
	Note              []Annotation                     `json:"note,omitempty"`
	UseContext        []UsageContext                   `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                `json:"jurisdiction,omitempty"`
	ApprovalDate      *string                          `json:"approvalDate,omitempty"`
	LastReviewDate    *string                          `json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                          `json:"effectivePeriod,omitempty"`
	Topic             []CodeableConcept                `json:"topic,omitempty"`
	Author            []ContactDetail                  `json:"author,omitempty"`
	Editor            []ContactDetail                  `json:"editor,omitempty"`
	Reviewer          []ContactDetail                  `json:"reviewer,omitempty"`
	Endorser          []ContactDetail                  `json:"endorser,omitempty"`
	RelatedArtifact   []RelatedArtifact                `json:"relatedArtifact,omitempty"`
	Type              *EvidenceVariableType            `json:"type,omitempty"`
	Characteristic    []EvidenceVariableCharacteristic `json:"characteristic"`
}

// A characteristic that defines the members of the evidence element. Multiple characteristics are applied with "and" semantics.
// Characteristics can be defined flexibly to accommodate different use cases for membership criteria, ranging from simple codes, all the way to using an expression language to express the criteria.
type EvidenceVariableCharacteristic struct {
	ID                           *string           `json:"id,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Description                  *string           `json:"description,omitempty"`
	DefinitionReference          Reference         `json:"definitionReference"`
	DefinitionCanonical          string            `json:"definitionCanonical"`
	DefinitionCodeableConcept    CodeableConcept   `json:"definitionCodeableConcept"`
	DefinitionExpression         Expression        `json:"definitionExpression"`
	DefinitionDataRequirement    DataRequirement   `json:"definitionDataRequirement"`
	DefinitionTriggerDefinition  TriggerDefinition `json:"definitionTriggerDefinition"`
	UsageContext                 []UsageContext    `json:"usageContext,omitempty"`
	Exclude                      *bool             `json:"exclude,omitempty"`
	ParticipantEffectiveDateTime *string           `json:"participantEffectiveDateTime,omitempty"`
	ParticipantEffectivePeriod   *Period           `json:"participantEffectivePeriod,omitempty"`
	ParticipantEffectiveDuration *Duration         `json:"participantEffectiveDuration,omitempty"`
	ParticipantEffectiveTiming   *Timing           `json:"participantEffectiveTiming,omitempty"`
	TimeFromStart                *Duration         `json:"timeFromStart,omitempty"`
	GroupMeasure                 *GroupMeasure     `json:"groupMeasure,omitempty"`
}

// This function returns resource reference information
func (r EvidenceVariable) ResourceRef() (string, *string) {
	return "EvidenceVariable", r.ID
}

// This function returns resource reference information
func (r EvidenceVariable) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherEvidenceVariable EvidenceVariable

// MarshalJSON marshals the given EvidenceVariable as JSON into a byte slice
func (r EvidenceVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEvidenceVariable
		ResourceType string `json:"resourceType"`
	}{
		OtherEvidenceVariable: OtherEvidenceVariable(r),
		ResourceType:          "EvidenceVariable",
	})
}

// UnmarshalEvidenceVariable unmarshals a EvidenceVariable.
func UnmarshalEvidenceVariable(b []byte) (EvidenceVariable, error) {
	var evidenceVariable EvidenceVariable
	if err := json.Unmarshal(b, &evidenceVariable); err != nil {
		return evidenceVariable, err
	}
	return evidenceVariable, nil
}
