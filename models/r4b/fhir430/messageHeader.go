package fhir430

import "encoding/json"

// MessageHeader is documented here http://hl7.org/fhir/StructureDefinition/MessageHeader
// The header for a message exchange that is either requesting or responding to an action.  The reference(s) that are the subject of the action as well as other information related to the action are typically transmitted in a bundle in which the MessageHeader resource instance is the first resource in the bundle.
type MessageHeader struct {
	ID                *string                    `json:"id,omitempty"`
	Meta              *Meta                      `json:"meta,omitempty"`
	ImplicitRules     *string                    `json:"implicitRules,omitempty"`
	Language          *string                    `json:"language,omitempty"`
	Text              *Narrative                 `json:"text,omitempty"`
	Contained         []json.RawMessage          `json:"contained,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	EventCoding       Coding                     `json:"eventCoding"`
	EventUri          string                     `json:"eventUri"`
	Destination       []MessageHeaderDestination `json:"destination,omitempty"`
	Sender            *Reference                 `json:"sender,omitempty"`
	Enterer           *Reference                 `json:"enterer,omitempty"`
	Author            *Reference                 `json:"author,omitempty"`
	Source            MessageHeaderSource        `json:"source"`
	Responsible       *Reference                 `json:"responsible,omitempty"`
	Reason            *CodeableConcept           `json:"reason,omitempty"`
	Response          *MessageHeaderResponse     `json:"response,omitempty"`
	Focus             []Reference                `json:"focus,omitempty"`
	Definition        *string                    `json:"definition,omitempty"`
}

// The destination application which the message is intended for.
// There SHOULD be at least one destination, but in some circumstances, the source system is unaware of any particular destination system.
type MessageHeaderDestination struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              *string     `json:"name,omitempty"`
	Target            *Reference  `json:"target,omitempty"`
	Endpoint          string      `json:"endpoint"`
	Receiver          *Reference  `json:"receiver,omitempty"`
}

// The source application from which this message originated.
type MessageHeaderSource struct {
	ID                *string       `json:"id,omitempty"`
	Extension         []Extension   `json:"extension,omitempty"`
	ModifierExtension []Extension   `json:"modifierExtension,omitempty"`
	Name              *string       `json:"name,omitempty"`
	Software          *string       `json:"software,omitempty"`
	Version           *string       `json:"version,omitempty"`
	Contact           *ContactPoint `json:"contact,omitempty"`
	Endpoint          string        `json:"endpoint"`
}

// Information about the message that this message is a response to.  Only present if this message is a response.
type MessageHeaderResponse struct {
	ID                *string      `json:"id,omitempty"`
	Extension         []Extension  `json:"extension,omitempty"`
	ModifierExtension []Extension  `json:"modifierExtension,omitempty"`
	Identifier        string       `json:"identifier"`
	Code              ResponseType `json:"code"`
	Details           *Reference   `json:"details,omitempty"`
}

// This function returns resource reference information
func (r MessageHeader) ResourceRef() (string, *string) {
	return "MessageHeader", r.ID
}

// This function returns resource reference information
func (r MessageHeader) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMessageHeader MessageHeader

// MarshalJSON marshals the given MessageHeader as JSON into a byte slice
func (r MessageHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMessageHeader
		ResourceType string `json:"resourceType"`
	}{
		OtherMessageHeader: OtherMessageHeader(r),
		ResourceType:       "MessageHeader",
	})
}

// UnmarshalMessageHeader unmarshals a MessageHeader.
func UnmarshalMessageHeader(b []byte) (MessageHeader, error) {
	var messageHeader MessageHeader
	if err := json.Unmarshal(b, &messageHeader); err != nil {
		return messageHeader, err
	}
	return messageHeader, nil
}
