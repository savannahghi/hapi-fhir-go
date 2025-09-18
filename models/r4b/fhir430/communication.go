
package fhir430

import "encoding/json"
// Communication is documented here http://hl7.org/fhir/StructureDefinition/Communication
// An occurrence of information being transmitted; e.g. an alert that was sent to a responsible provider, a public health agency that was notified about a reportable condition.
type Communication struct {
	ID                    *string                `json:"ID,omitempty"`
	Meta                  *Meta                  `json:"meta,omitempty"`
	ImplicitRules         *string                `json:"implicitRules,omitempty"`
	Language              *string                `json:"language,omitempty"`
	Text                  *Narrative             `json:"text,omitempty"`
	Contained             []json.RawMessage      `json:"contained,omitempty"`
	Extension             []Extension            `json:"extension,omitempty"`
	ModifierExtension     []Extension            `json:"modifierExtension,omitempty"`
	Identifier            []Identifier           `json:"identifier,omitempty"`
	InstantiatesCanonical []string               `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string               `json:"instantiatesUri,omitempty"`
	BasedOn               []Reference            `json:"basedOn,omitempty"`
	PartOf                []Reference            `json:"partOf,omitempty"`
	InResponseTo          []Reference            `json:"inResponseTo,omitempty"`
	Status                EventStatus            `json:"status"`
	StatusReason          *CodeableConcept       `json:"statusReason,omitempty"`
	Category              []CodeableConcept      `json:"category,omitempty"`
	Priority              *RequestPriority       `json:"priority,omitempty"`
	Medium                []CodeableConcept      `json:"medium,omitempty"`
	Subject               *Reference             `json:"subject,omitempty"`
	Topic                 *CodeableConcept       `json:"topic,omitempty"`
	About                 []Reference            `json:"about,omitempty"`
	Encounter             *Reference             `json:"encounter,omitempty"`
	Sent                  *string                `json:"sent,omitempty"`
	Received              *string                `json:"received,omitempty"`
	Recipient             []Reference            `json:"recipient,omitempty"`
	Sender                *Reference             `json:"sender,omitempty"`
	ReasonCode            []CodeableConcept      `json:"reasonCode,omitempty"`
	ReasonReference       []Reference            `json:"reasonReference,omitempty"`
	Payload               []CommunicationPayload `json:"payload,omitempty"`
	Note                  []Annotation           `json:"note,omitempty"`
}

// Text, attachment(s), or resource(s) that was communicated to the recipient.
type CommunicationPayload struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentString     string      `json:"contentString"`
	ContentAttachment Attachment  `json:"contentAttachment"`
	ContentReference  Reference   `json:"contentReference"`
}

// This function returns resource reference information
func (r Communication) ResourceRef() (string, *string) {
	return "Communication", r.ID
}

// This function returns resource reference information
func (r Communication) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCommunication Communication

// MarshalJSON marshals the given Communication as JSON into a byte slice
func (r Communication) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCommunication
		ResourceType string `json:"resourceType"`
	}{
		OtherCommunication: OtherCommunication(r),
		ResourceType:       "Communication",
	})
}

// UnmarshalCommunication unmarshals a Communication.
func UnmarshalCommunication(b []byte) (Communication, error) {
	var communication Communication
	if err := json.Unmarshal(b, &communication); err != nil {
		return communication, err
	}
	return communication, nil
}
