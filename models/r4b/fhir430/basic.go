package fhir430

import "encoding/json"

// Basic is documented here http://hl7.org/fhir/StructureDefinition/Basic
// Basic is used for handling concepts not yet defined in FHIR, narrative-only resources that don't map to an existing resource, and custom resources not appropriate for inclusion in the FHIR specification.
type Basic struct {
	ID                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []json.RawMessage `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	Code              CodeableConcept   `json:"code"`
	Subject           *Reference        `json:"subject,omitempty"`
	Created           *string           `json:"created,omitempty"`
	Author            *Reference        `json:"author,omitempty"`
}

// This function returns resource reference information
func (r Basic) ResourceRef() (string, *string) {
	return "Basic", r.ID
}

// This function returns resource reference information
func (r Basic) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherBasic Basic

// MarshalJSON marshals the given Basic as JSON into a byte slice
func (r Basic) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBasic
		ResourceType string `json:"resourceType"`
	}{
		OtherBasic:   OtherBasic(r),
		ResourceType: "Basic",
	})
}

// UnmarshalBasic unmarshals a Basic.
func UnmarshalBasic(b []byte) (Basic, error) {
	var basic Basic
	if err := json.Unmarshal(b, &basic); err != nil {
		return basic, err
	}
	return basic, nil
}
