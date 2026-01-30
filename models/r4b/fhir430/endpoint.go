package fhir430

import "encoding/json"

// Endpoint is documented here http://hl7.org/fhir/StructureDefinition/Endpoint
// The technical details of an endpoint that can be used for electronic services, such as for web services providing XDS.b or a REST endpoint for another FHIR server. This may include any security context information.
type Endpoint struct {
	ID                   *string           `json:"id,omitempty"`
	Meta                 *Meta             `json:"meta,omitempty"`
	ImplicitRules        *string           `json:"implicitRules,omitempty"`
	Language             *string           `json:"language,omitempty"`
	Text                 *Narrative        `json:"text,omitempty"`
	Contained            []json.RawMessage `json:"contained,omitempty"`
	Extension            []Extension       `json:"extension,omitempty"`
	ModifierExtension    []Extension       `json:"modifierExtension,omitempty"`
	Identifier           []Identifier      `json:"identifier,omitempty"`
	Status               EndpointStatus    `json:"status"`
	ConnectionType       Coding            `json:"connectionType"`
	Name                 *string           `json:"name,omitempty"`
	ManagingOrganization *Reference        `json:"managingOrganization,omitempty"`
	Contact              []ContactPoint    `json:"contact,omitempty"`
	Period               *Period           `json:"period,omitempty"`
	PayloadType          []CodeableConcept `json:"payloadType"`
	PayloadMimeType      []string          `json:"payloadMimeType,omitempty"`
	Address              string            `json:"address"`
	Header               []string          `json:"header,omitempty"`
}

// This function returns resource reference information
func (r Endpoint) ResourceRef() (string, *string) {
	return "Endpoint", r.ID
}

// This function returns resource reference information
func (r Endpoint) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherEndpoint Endpoint

// MarshalJSON marshals the given Endpoint as JSON into a byte slice
func (r Endpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEndpoint
		ResourceType string `json:"resourceType"`
	}{
		OtherEndpoint: OtherEndpoint(r),
		ResourceType:  "Endpoint",
	})
}

// UnmarshalEndpoint unmarshals a Endpoint.
func UnmarshalEndpoint(b []byte) (Endpoint, error) {
	var endpoint Endpoint
	if err := json.Unmarshal(b, &endpoint); err != nil {
		return endpoint, err
	}
	return endpoint, nil
}
