package fhir430

import "encoding/json"

// Appointment is documented here http://hl7.org/fhir/StructureDefinition/Appointment
// A booking of a healthcare event among patient(s), practitioner(s), related person(s) and/or device(s) for a specific date/time. This may result in one or more Encounter(s).
type Appointment struct {
	ID                    *string                  `json:"id,omitempty"`
	Meta                  *Meta                    `json:"meta,omitempty"`
	ImplicitRules         *string                  `json:"implicitRules,omitempty"`
	Language              *string                  `json:"language,omitempty"`
	Text                  *Narrative               `json:"text,omitempty"`
	Contained             []json.RawMessage        `json:"contained,omitempty"`
	Extension             []Extension              `json:"extension,omitempty"`
	ModifierExtension     []Extension              `json:"modifierExtension,omitempty"`
	Identifier            []Identifier             `json:"identifier,omitempty"`
	Status                AppointmentStatus        `json:"status"`
	CancelationReason     *CodeableConcept         `json:"cancelationReason,omitempty"`
	ServiceCategory       []CodeableConcept        `json:"serviceCategory,omitempty"`
	ServiceType           []CodeableConcept        `json:"serviceType,omitempty"`
	Specialty             []CodeableConcept        `json:"specialty,omitempty"`
	AppointmentType       *CodeableConcept         `json:"appointmentType,omitempty"`
	ReasonCode            []CodeableConcept        `json:"reasonCode,omitempty"`
	ReasonReference       []Reference              `json:"reasonReference,omitempty"`
	Priority              *int                     `json:"priority,omitempty"`
	Description           *string                  `json:"description,omitempty"`
	SupportingInformation []Reference              `json:"supportingInformation,omitempty"`
	Start                 *string                  `json:"start,omitempty"`
	End                   *string                  `json:"end,omitempty"`
	MinutesDuration       *int                     `json:"minutesDuration,omitempty"`
	Slot                  []Reference              `json:"slot,omitempty"`
	Created               *string                  `json:"created,omitempty"`
	Comment               *string                  `json:"comment,omitempty"`
	PatientInstruction    *string                  `json:"patientInstruction,omitempty"`
	BasedOn               []Reference              `json:"basedOn,omitempty"`
	Participant           []AppointmentParticipant `json:"participant"`
	RequestedPeriod       []Period                 `json:"requestedPeriod,omitempty"`
}

// List of participants involved in the appointment.
type AppointmentParticipant struct {
	ID                *string              `json:"id,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	Type              []CodeableConcept    `json:"type,omitempty"`
	Actor             *Reference           `json:"actor,omitempty"`
	Required          *ParticipantRequired `json:"required,omitempty"`
	Status            ParticipationStatus  `json:"status"`
	Period            *Period              `json:"period,omitempty"`
}

// This function returns resource reference information
func (r Appointment) ResourceRef() (string, *string) {
	return "Appointment", r.ID
}

// This function returns resource reference information
func (r Appointment) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherAppointment Appointment

// MarshalJSON marshals the given Appointment as JSON into a byte slice
func (r Appointment) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAppointment
		ResourceType string `json:"resourceType"`
	}{
		OtherAppointment: OtherAppointment(r),
		ResourceType:     "Appointment",
	})
}

// UnmarshalAppointment unmarshals a Appointment.
func UnmarshalAppointment(b []byte) (Appointment, error) {
	var appointment Appointment
	if err := json.Unmarshal(b, &appointment); err != nil {
		return appointment, err
	}
	return appointment, nil
}
