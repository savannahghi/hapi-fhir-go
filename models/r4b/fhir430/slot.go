
package fhir430

import "encoding/json"
// Slot is documented here http://hl7.org/fhir/StructureDefinition/Slot
// A slot of time on a schedule that may be available for booking appointments.
type Slot struct {
	ID                *string           `json:"ID,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []json.RawMessage `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	ServiceCategory   []CodeableConcept `json:"serviceCategory,omitempty"`
	ServiceType       []CodeableConcept `json:"serviceType,omitempty"`
	Specialty         []CodeableConcept `json:"specialty,omitempty"`
	AppointmentType   *CodeableConcept  `json:"appointmentType,omitempty"`
	Schedule          Reference         `json:"schedule"`
	Status            SlotStatus        `json:"status"`
	Start             string            `json:"start"`
	End               string            `json:"end"`
	Overbooked        *bool             `json:"overbooked,omitempty"`
	Comment           *string           `json:"comment,omitempty"`
}

// This function returns resource reference information
func (r Slot) ResourceRef() (string, *string) {
	return "Slot", r.ID
}

// This function returns resource reference information
func (r Slot) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherSlot Slot

// MarshalJSON marshals the given Slot as JSON into a byte slice
func (r Slot) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSlot
		ResourceType string `json:"resourceType"`
	}{
		OtherSlot:    OtherSlot(r),
		ResourceType: "Slot",
	})
}

// UnmarshalSlot unmarshals a Slot.
func UnmarshalSlot(b []byte) (Slot, error) {
	var slot Slot
	if err := json.Unmarshal(b, &slot); err != nil {
		return slot, err
	}
	return slot, nil
}
