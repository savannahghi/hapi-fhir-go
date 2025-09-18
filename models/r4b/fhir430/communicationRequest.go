
package fhir430

import "encoding/json"
// CommunicationRequest is documented here http://hl7.org/fhir/StructureDefinition/CommunicationRequest
// A request to convey information; e.g. the CDS system proposes that an alert be sent to a responsible provider, the CDS system proposes that the public health agency be notified about a reportable condition.
type CommunicationRequest struct {
	ID                 *string                       `json:"ID,omitempty"`
	Meta               *Meta                         `json:"meta,omitempty"`
	ImplicitRules      *string                       `json:"implicitRules,omitempty"`
	Language           *string                       `json:"language,omitempty"`
	Text               *Narrative                    `json:"text,omitempty"`
	Contained          []json.RawMessage             `json:"contained,omitempty"`
	Extension          []Extension                   `json:"extension,omitempty"`
	ModifierExtension  []Extension                   `json:"modifierExtension,omitempty"`
	Identifier         []Identifier                  `json:"identifier,omitempty"`
	BasedOn            []Reference                   `json:"basedOn,omitempty"`
	Replaces           []Reference                   `json:"replaces,omitempty"`
	GroupIdentifier    *Identifier                   `json:"groupIdentifier,omitempty"`
	Status             RequestStatus                 `json:"status"`
	StatusReason       *CodeableConcept              `json:"statusReason,omitempty"`
	Category           []CodeableConcept             `json:"category,omitempty"`
	Priority           *RequestPriority              `json:"priority,omitempty"`
	DoNotPerform       *bool                         `json:"doNotPerform,omitempty"`
	Medium             []CodeableConcept             `json:"medium,omitempty"`
	Subject            *Reference                    `json:"subject,omitempty"`
	About              []Reference                   `json:"about,omitempty"`
	Encounter          *Reference                    `json:"encounter,omitempty"`
	Payload            []CommunicationRequestPayload `json:"payload,omitempty"`
	OccurrenceDateTime *string                       `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                       `json:"occurrencePeriod,omitempty"`
	AuthoredOn         *string                       `json:"authoredOn,omitempty"`
	Requester          *Reference                    `json:"requester,omitempty"`
	Recipient          []Reference                   `json:"recipient,omitempty"`
	Sender             *Reference                    `json:"sender,omitempty"`
	ReasonCode         []CodeableConcept             `json:"reasonCode,omitempty"`
	ReasonReference    []Reference                   `json:"reasonReference,omitempty"`
	Note               []Annotation                  `json:"note,omitempty"`
}

// Text, attachment(s), or resource(s) to be communicated to the recipient.
type CommunicationRequestPayload struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	ContentString     string      `json:"contentString"`
	ContentAttachment Attachment  `json:"contentAttachment"`
	ContentReference  Reference   `json:"contentReference"`
}

// This function returns resource reference information
func (r CommunicationRequest) ResourceRef() (string, *string) {
	return "CommunicationRequest", r.ID
}

// This function returns resource reference information
func (r CommunicationRequest) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCommunicationRequest CommunicationRequest

// MarshalJSON marshals the given CommunicationRequest as JSON into a byte slice
func (r CommunicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCommunicationRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherCommunicationRequest: OtherCommunicationRequest(r),
		ResourceType:              "CommunicationRequest",
	})
}

// UnmarshalCommunicationRequest unmarshals a CommunicationRequest.
func UnmarshalCommunicationRequest(b []byte) (CommunicationRequest, error) {
	var communicationRequest CommunicationRequest
	if err := json.Unmarshal(b, &communicationRequest); err != nil {
		return communicationRequest, err
	}
	return communicationRequest, nil
}
