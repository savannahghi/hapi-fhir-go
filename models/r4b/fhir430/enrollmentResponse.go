package fhir430

import "encoding/json"

// EnrollmentResponse is documented here http://hl7.org/fhir/StructureDefinition/EnrollmentResponse
// This resource provides enrollment and plan details from the processing of an EnrollmentRequest resource.
type EnrollmentResponse struct {
	ID                *string                       `json:"id,omitempty"`
	Meta              *Meta                         `json:"meta,omitempty"`
	ImplicitRules     *string                       `json:"implicitRules,omitempty"`
	Language          *string                       `json:"language,omitempty"`
	Text              *Narrative                    `json:"text,omitempty"`
	Contained         []json.RawMessage             `json:"contained,omitempty"`
	Extension         []Extension                   `json:"extension,omitempty"`
	ModifierExtension []Extension                   `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                  `json:"identifier,omitempty"`
	Status            *FinancialResourceStatusCodes `json:"status,omitempty"`
	Request           *Reference                    `json:"request,omitempty"`
	Outcome           *ClaimProcessingCodes         `json:"outcome,omitempty"`
	Disposition       *string                       `json:"disposition,omitempty"`
	Created           *string                       `json:"created,omitempty"`
	Organization      *Reference                    `json:"organization,omitempty"`
	RequestProvider   *Reference                    `json:"requestProvider,omitempty"`
}

// This function returns resource reference information
func (r EnrollmentResponse) ResourceRef() (string, *string) {
	return "EnrollmentResponse", r.ID
}

// This function returns resource reference information
func (r EnrollmentResponse) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherEnrollmentResponse EnrollmentResponse

// MarshalJSON marshals the given EnrollmentResponse as JSON into a byte slice
func (r EnrollmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEnrollmentResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherEnrollmentResponse: OtherEnrollmentResponse(r),
		ResourceType:            "EnrollmentResponse",
	})
}

// UnmarshalEnrollmentResponse unmarshals a EnrollmentResponse.
func UnmarshalEnrollmentResponse(b []byte) (EnrollmentResponse, error) {
	var enrollmentResponse EnrollmentResponse
	if err := json.Unmarshal(b, &enrollmentResponse); err != nil {
		return enrollmentResponse, err
	}
	return enrollmentResponse, nil
}
