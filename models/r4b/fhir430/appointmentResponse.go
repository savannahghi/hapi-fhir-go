package fhir430

import "encoding/json"

// AppointmentResponse is documented here http://hl7.org/fhir/StructureDefinition/AppointmentResponse
// A reply to an appointment request for a patient and/or practitioner(s), such as a confirmation or rejection.
type AppointmentResponse struct {
	ID                *string             `json:"id,omitempty"`
	Meta              *Meta               `json:"meta,omitempty"`
	ImplicitRules     *string             `json:"implicitRules,omitempty"`
	Language          *string             `json:"language,omitempty"`
	Text              *Narrative          `json:"text,omitempty"`
	Contained         []json.RawMessage   `json:"contained,omitempty"`
	Extension         []Extension         `json:"extension,omitempty"`
	ModifierExtension []Extension         `json:"modifierExtension,omitempty"`
	Identifier        []Identifier        `json:"identifier,omitempty"`
	Appointment       Reference           `json:"appointment"`
	Start             *string             `json:"start,omitempty"`
	End               *string             `json:"end,omitempty"`
	ParticipantType   []CodeableConcept   `json:"participantType,omitempty"`
	Actor             *Reference          `json:"actor,omitempty"`
	ParticipantStatus ParticipationStatus `json:"participantStatus"`
	Comment           *string             `json:"comment,omitempty"`
}

// This function returns resource reference information
func (r AppointmentResponse) ResourceRef() (string, *string) {
	return "AppointmentResponse", r.ID
}

// This function returns resource reference information
func (r AppointmentResponse) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherAppointmentResponse AppointmentResponse

// MarshalJSON marshals the given AppointmentResponse as JSON into a byte slice
func (r AppointmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAppointmentResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherAppointmentResponse: OtherAppointmentResponse(r),
		ResourceType:             "AppointmentResponse",
	})
}

// UnmarshalAppointmentResponse unmarshals a AppointmentResponse.
func UnmarshalAppointmentResponse(b []byte) (AppointmentResponse, error) {
	var appointmentResponse AppointmentResponse
	if err := json.Unmarshal(b, &appointmentResponse); err != nil {
		return appointmentResponse, err
	}
	return appointmentResponse, nil
}
