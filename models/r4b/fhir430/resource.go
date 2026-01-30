package fhir430

import "encoding/json"

// Resource is documented here http://hl7.org/fhir/StructureDefinition/Resource
// This is the base resource type for everything.
type Resource struct {
	ID            *string `json:"id,omitempty"`
	Meta          *Meta   `json:"meta,omitempty"`
	ImplicitRules *string `json:"implicitRules,omitempty"`
	Language      *string `json:"language,omitempty"`
}

// This function returns resource reference information
func (r Resource) ResourceRef() (string, *string) {
	return "Resource", r.ID
}

// This function returns resource reference information
func (r Resource) ContainedResources() []json.RawMessage {
	return nil
}

type OtherResource Resource

// MarshalJSON marshals the given Resource as JSON into a byte slice
func (r Resource) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherResource
		ResourceType string `json:"resourceType"`
	}{
		OtherResource: OtherResource(r),
		ResourceType:  "Resource",
	})
}

// UnmarshalResource unmarshals a Resource.
func UnmarshalResource(b []byte) (Resource, error) {
	var resource Resource
	if err := json.Unmarshal(b, &resource); err != nil {
		return resource, err
	}
	return resource, nil
}
