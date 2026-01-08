package fhir430

import "encoding/json"

// EventDefinition is documented here http://hl7.org/fhir/StructureDefinition/EventDefinition
// The EventDefinition resource provides a reusable description of when a particular event can occur.
type EventDefinition struct {
	ID                     *string             `json:"id,omitempty"`
	Meta                   *Meta               `json:"meta,omitempty"`
	ImplicitRules          *string             `json:"implicitRules,omitempty"`
	Language               *string             `json:"language,omitempty"`
	Text                   *Narrative          `json:"text,omitempty"`
	Contained              []json.RawMessage   `json:"contained,omitempty"`
	Extension              []Extension         `json:"extension,omitempty"`
	ModifierExtension      []Extension         `json:"modifierExtension,omitempty"`
	Url                    *string             `json:"url,omitempty"`
	Identifier             []Identifier        `json:"identifier,omitempty"`
	Version                *string             `json:"version,omitempty"`
	Name                   *string             `json:"name,omitempty"`
	Title                  *string             `json:"title,omitempty"`
	Subtitle               *string             `json:"subtitle,omitempty"`
	Status                 PublicationStatus   `json:"status"`
	Experimental           *bool               `json:"experimental,omitempty"`
	SubjectCodeableConcept *CodeableConcept    `json:"subjectCodeableConcept,omitempty"`
	SubjectReference       *Reference          `json:"subjectReference,omitempty"`
	Date                   *string             `json:"date,omitempty"`
	Publisher              *string             `json:"publisher,omitempty"`
	Contact                []ContactDetail     `json:"contact,omitempty"`
	Description            *string             `json:"description,omitempty"`
	UseContext             []UsageContext      `json:"useContext,omitempty"`
	Jurisdiction           []CodeableConcept   `json:"jurisdiction,omitempty"`
	Purpose                *string             `json:"purpose,omitempty"`
	Usage                  *string             `json:"usage,omitempty"`
	ApprovalDate           *string             `json:"approvalDate,omitempty"`
	LastReviewDate         *string             `json:"lastReviewDate,omitempty"`
	EffectivePeriod        *Period             `json:"effectivePeriod,omitempty"`
	Topic                  []CodeableConcept   `json:"topic,omitempty"`
	Author                 []ContactDetail     `json:"author,omitempty"`
	Editor                 []ContactDetail     `json:"editor,omitempty"`
	Reviewer               []ContactDetail     `json:"reviewer,omitempty"`
	Endorser               []ContactDetail     `json:"endorser,omitempty"`
	RelatedArtifact        []RelatedArtifact   `json:"relatedArtifact,omitempty"`
	Trigger                []TriggerDefinition `json:"trigger"`
}

// This function returns resource reference information
func (r EventDefinition) ResourceRef() (string, *string) {
	return "EventDefinition", r.ID
}

// This function returns resource reference information
func (r EventDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherEventDefinition EventDefinition

// MarshalJSON marshals the given EventDefinition as JSON into a byte slice
func (r EventDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEventDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherEventDefinition: OtherEventDefinition(r),
		ResourceType:         "EventDefinition",
	})
}

// UnmarshalEventDefinition unmarshals a EventDefinition.
func UnmarshalEventDefinition(b []byte) (EventDefinition, error) {
	var eventDefinition EventDefinition
	if err := json.Unmarshal(b, &eventDefinition); err != nil {
		return eventDefinition, err
	}
	return eventDefinition, nil
}
