package fhir430

import "encoding/json"

// BodyStructure is documented here http://hl7.org/fhir/StructureDefinition/BodyStructure
// Record details about an anatomical structure.  This resource may be used when a coded concept does not provide the necessary detail needed for the use case.
type BodyStructure struct {
	ID                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []json.RawMessage `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Identifier        []Identifier      `json:"identifier,omitempty"`
	Active            *bool             `json:"active,omitempty"`
	Morphology        *CodeableConcept  `json:"morphology,omitempty"`
	Location          *CodeableConcept  `json:"location,omitempty"`
	LocationQualifier []CodeableConcept `json:"locationQualifier,omitempty"`
	Description       *string           `json:"description,omitempty"`
	Image             []Attachment      `json:"image,omitempty"`
	Patient           Reference         `json:"patient"`
}

// This function returns resource reference information
func (r BodyStructure) ResourceRef() (string, *string) {
	return "BodyStructure", r.ID
}

// This function returns resource reference information
func (r BodyStructure) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherBodyStructure BodyStructure

// MarshalJSON marshals the given BodyStructure as JSON into a byte slice
func (r BodyStructure) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBodyStructure
		ResourceType string `json:"resourceType"`
	}{
		OtherBodyStructure: OtherBodyStructure(r),
		ResourceType:       "BodyStructure",
	})
}

// UnmarshalBodyStructure unmarshals a BodyStructure.
func UnmarshalBodyStructure(b []byte) (BodyStructure, error) {
	var bodyStructure BodyStructure
	if err := json.Unmarshal(b, &bodyStructure); err != nil {
		return bodyStructure, err
	}
	return bodyStructure, nil
}
