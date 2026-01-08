package fhir430

import "encoding/json"

// Binary is documented here http://hl7.org/fhir/StructureDefinition/Binary
// A resource that represents the data of a single raw artifact as digital content accessible in its native format.  A Binary resource can contain any content, whether text, image, pdf, zip archive, etc.
type Binary struct {
	ID              *string    `json:"id,omitempty"`
	Meta            *Meta      `json:"meta,omitempty"`
	ImplicitRules   *string    `json:"implicitRules,omitempty"`
	Language        *string    `json:"language,omitempty"`
	ContentType     string     `json:"contentType"`
	SecurityContext *Reference `json:"securityContext,omitempty"`
	Data            *string    `json:"data,omitempty"`
}

// This function returns resource reference information
func (r Binary) ResourceRef() (string, *string) {
	return "Binary", r.ID
}

// This function returns resource reference information
func (r Binary) ContainedResources() []json.RawMessage {
	return nil
}

type OtherBinary Binary

// MarshalJSON marshals the given Binary as JSON into a byte slice
func (r Binary) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBinary
		ResourceType string `json:"resourceType"`
	}{
		OtherBinary:  OtherBinary(r),
		ResourceType: "Binary",
	})
}

// UnmarshalBinary unmarshals a Binary.
func UnmarshalBinary(b []byte) (Binary, error) {
	var binary Binary
	if err := json.Unmarshal(b, &binary); err != nil {
		return binary, err
	}
	return binary, nil
}
