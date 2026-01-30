package fhir430

import "encoding/json"

// DomainResource is documented here http://hl7.org/fhir/StructureDefinition/DomainResource
// A resource that includes narrative, extensions, and contained resources.
type DomainResource struct {
	ID                *string           `json:"id,omitempty"`
	Meta              *Meta             `json:"meta,omitempty"`
	ImplicitRules     *string           `json:"implicitRules,omitempty"`
	Language          *string           `json:"language,omitempty"`
	Text              *Narrative        `json:"text,omitempty"`
	Contained         []json.RawMessage `json:"contained,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
}

// This function returns resource reference information
func (r DomainResource) ResourceRef() (string, *string) {
	return "DomainResource", r.ID
}

// This function returns resource reference information
func (r DomainResource) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherDomainResource DomainResource

// MarshalJSON marshals the given DomainResource as JSON into a byte slice
func (r DomainResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDomainResource
		ResourceType string `json:"resourceType"`
	}{
		OtherDomainResource: OtherDomainResource(r),
		ResourceType:        "DomainResource",
	})
}

// UnmarshalDomainResource unmarshals a DomainResource.
func UnmarshalDomainResource(b []byte) (DomainResource, error) {
	var domainResource DomainResource
	if err := json.Unmarshal(b, &domainResource); err != nil {
		return domainResource, err
	}
	return domainResource, nil
}
