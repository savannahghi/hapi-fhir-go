package fhir430

import "encoding/json"

// MessageDefinition is documented here http://hl7.org/fhir/StructureDefinition/MessageDefinition
// Defines the characteristics of a message that can be shared between systems, including the type of event that initiates the message, the content to be transmitted and what response(s), if any, are permitted.
type MessageDefinition struct {
	ID                *string                            `json:"id,omitempty"`
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
	Replaces          []string                           `json:"replaces,omitempty"`
	Status            PublicationStatus                  `json:"status"`
	Experimental      *bool                              `json:"experimental,omitempty"`
	Date              string                             `json:"date"`
	Publisher         *string                            `json:"publisher,omitempty"`
	Contact           []ContactDetail                    `json:"contact,omitempty"`
	Description       *string                            `json:"description,omitempty"`
	UseContext        []UsageContext                     `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                  `json:"jurisdiction,omitempty"`
	Purpose           *string                            `json:"purpose,omitempty"`
	Base              *string                            `json:"base,omitempty"`
	Parent            []string                           `json:"parent,omitempty"`
	EventCoding       Coding                             `json:"eventCoding"`
	EventUri          string                             `json:"eventUri"`
	Category          *MessageSignificanceCategory       `json:"category,omitempty"`
	Focus             []MessageDefinitionFocus           `json:"focus,omitempty"`
	ResponseRequired  *string                            `json:"responseRequired,omitempty"`
	AllowedResponse   []MessageDefinitionAllowedResponse `json:"allowedResponse,omitempty"`
	Graph             []string                           `json:"graph,omitempty"`
}

// Identifies the resource (or resources) that are being addressed by the event.  For example, the Encounter for an admit message or two Account records for a merge.
type MessageDefinitionFocus struct {
	ID                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Code              ResourceType `json:"code"`
	Profile           *string      `json:"profile,omitempty"`
	Min               int          `json:"min"`
	Max               *string      `json:"max,omitempty"`
}

// Indicates what types of messages may be sent as an application-level response to this message.
// This indicates an application level response to "close" a transaction implicit in a particular request message.  To define a complete workflow scenario, look to the [[PlanDefinition]] resource which allows the definition of complex orchestrations, conditionality, etc.
type MessageDefinitionAllowedResponse struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Message           string      `json:"message"`
	Situation         *string     `json:"situation,omitempty"`
}

// This function returns resource reference information
func (r MessageDefinition) ResourceRef() (string, *string) {
	return "MessageDefinition", r.ID
}

// This function returns resource reference information
func (r MessageDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMessageDefinition MessageDefinition

// MarshalJSON marshals the given MessageDefinition as JSON into a byte slice
func (r MessageDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMessageDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherMessageDefinition: OtherMessageDefinition(r),
		ResourceType:           "MessageDefinition",
	})
}

// UnmarshalMessageDefinition unmarshals a MessageDefinition.
func UnmarshalMessageDefinition(b []byte) (MessageDefinition, error) {
	var messageDefinition MessageDefinition
	if err := json.Unmarshal(b, &messageDefinition); err != nil {
		return messageDefinition, err
	}
	return messageDefinition, nil
}
