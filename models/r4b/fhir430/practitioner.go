
package fhir430

import "encoding/json"
// Practitioner is documented here http://hl7.org/fhir/StructureDefinition/Practitioner
// A person who is directly or indirectly involved in the provisioning of healthcare.
type Practitioner struct {
	ID                *string                     `json:"ID,omitempty"`
	Meta              *Meta                       `json:"meta,omitempty"`
	ImplicitRules     *string                     `json:"implicitRules,omitempty"`
	Language          *string                     `json:"language,omitempty"`
	Text              *Narrative                  `json:"text,omitempty"`
	Contained         []json.RawMessage           `json:"contained,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `json:"identifier,omitempty"`
	Active            *bool                       `json:"active,omitempty"`
	Name              []HumanName                 `json:"name,omitempty"`
	Telecom           []ContactPoint              `json:"telecom,omitempty"`
	Address           []Address                   `json:"address,omitempty"`
	Gender            *AdministrativeGender       `json:"gender,omitempty"`
	BirthDate         *string                     `json:"birthDate,omitempty"`
	Photo             []Attachment                `json:"photo,omitempty"`
	Qualification     []PractitionerQualification `json:"qualification,omitempty"`
	Communication     []CodeableConcept           `json:"communication,omitempty"`
}

// The official certifications, training, and licenses that authorize or otherwise pertain to the provision of care by the practitioner.  For example, a medical license issued by a medical board authorizing the practitioner to practice medicine within a certian locality.
type PractitionerQualification struct {
	ID                *string         `json:"ID,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Identifier        []Identifier    `json:"identifier,omitempty"`
	Code              CodeableConcept `json:"code"`
	Period            *Period         `json:"period,omitempty"`
	Issuer            *Reference      `json:"issuer,omitempty"`
}

// This function returns resource reference information
func (r Practitioner) ResourceRef() (string, *string) {
	return "Practitioner", r.ID
}

// This function returns resource reference information
func (r Practitioner) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherPractitioner Practitioner

// MarshalJSON marshals the given Practitioner as JSON into a byte slice
func (r Practitioner) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPractitioner
		ResourceType string `json:"resourceType"`
	}{
		OtherPractitioner: OtherPractitioner(r),
		ResourceType:      "Practitioner",
	})
}

// UnmarshalPractitioner unmarshals a Practitioner.
func UnmarshalPractitioner(b []byte) (Practitioner, error) {
	var practitioner Practitioner
	if err := json.Unmarshal(b, &practitioner); err != nil {
		return practitioner, err
	}
	return practitioner, nil
}
