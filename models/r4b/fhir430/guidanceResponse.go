
package fhir430

import "encoding/json"
// GuidanceResponse is documented here http://hl7.org/fhir/StructureDefinition/GuidanceResponse
// A guidance response is the formal response to a guidance request, including any output parameters returned by the evaluation, as well as the description of any proposed actions to be taken.
type GuidanceResponse struct {
	ID                    *string                `json:"ID,omitempty"`
	Meta                  *Meta                  `json:"meta,omitempty"`
	ImplicitRules         *string                `json:"implicitRules,omitempty"`
	Language              *string                `json:"language,omitempty"`
	Text                  *Narrative             `json:"text,omitempty"`
	Contained             []json.RawMessage      `json:"contained,omitempty"`
	Extension             []Extension            `json:"extension,omitempty"`
	ModifierExtension     []Extension            `json:"modifierExtension,omitempty"`
	RequestIdentifier     *Identifier            `json:"requestIdentifier,omitempty"`
	Identifier            []Identifier           `json:"identifier,omitempty"`
	ModuleUri             string                 `json:"moduleUri"`
	ModuleCanonical       string                 `json:"moduleCanonical"`
	ModuleCodeableConcept CodeableConcept        `json:"moduleCodeableConcept"`
	Status                GuidanceResponseStatus `json:"status"`
	Subject               *Reference             `json:"subject,omitempty"`
	Encounter             *Reference             `json:"encounter,omitempty"`
	OccurrenceDateTime    *string                `json:"occurrenceDateTime,omitempty"`
	Performer             *Reference             `json:"performer,omitempty"`
	ReasonCode            []CodeableConcept      `json:"reasonCode,omitempty"`
	ReasonReference       []Reference            `json:"reasonReference,omitempty"`
	Note                  []Annotation           `json:"note,omitempty"`
	EvaluationMessage     []Reference            `json:"evaluationMessage,omitempty"`
	OutputParameters      *Reference             `json:"outputParameters,omitempty"`
	Result                *Reference             `json:"result,omitempty"`
	DataRequirement       []DataRequirement      `json:"dataRequirement,omitempty"`
}

// This function returns resource reference information
func (r GuidanceResponse) ResourceRef() (string, *string) {
	return "GuidanceResponse", r.ID
}

// This function returns resource reference information
func (r GuidanceResponse) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherGuidanceResponse GuidanceResponse

// MarshalJSON marshals the given GuidanceResponse as JSON into a byte slice
func (r GuidanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherGuidanceResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherGuidanceResponse: OtherGuidanceResponse(r),
		ResourceType:          "GuidanceResponse",
	})
}

// UnmarshalGuidanceResponse unmarshals a GuidanceResponse.
func UnmarshalGuidanceResponse(b []byte) (GuidanceResponse, error) {
	var guidanceResponse GuidanceResponse
	if err := json.Unmarshal(b, &guidanceResponse); err != nil {
		return guidanceResponse, err
	}
	return guidanceResponse, nil
}
