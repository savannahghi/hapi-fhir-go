
package fhir430

import "encoding/json"
// Flag is documented here http://hl7.org/fhir/StructureDefinition/Flag
// Prospective warnings of potential issues when providing care to the patient.
type Flag struct {
	ID                *string           `json:"ID,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []json.RawMessage `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	Status            FlagStatus        `json:"status"`
	Category          []CodeableConcept `json:"category,omitempty"`
	Code              CodeableConcept   `json:"code"`
	Subject           Reference         `json:"subject"`
	Period            *Period           `json:"period,omitempty"`
	Encounter         *Reference        `json:"encounter,omitempty"`
	Author            *Reference        `json:"author,omitempty"`
}

// This function returns resource reference information
func (r Flag) ResourceRef() (string, *string) {
	return "Flag", r.ID
}

// This function returns resource reference information
func (r Flag) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherFlag Flag

// MarshalJSON marshals the given Flag as JSON into a byte slice
func (r Flag) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFlag
		ResourceType string `json:"resourceType"`
	}{
		OtherFlag:    OtherFlag(r),
		ResourceType: "Flag",
	})
}

// UnmarshalFlag unmarshals a Flag.
func UnmarshalFlag(b []byte) (Flag, error) {
	var flag Flag
	if err := json.Unmarshal(b, &flag); err != nil {
		return flag, err
	}
	return flag, nil
}
