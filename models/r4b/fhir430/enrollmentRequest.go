package fhir430

import "encoding/json"

// EnrollmentRequest is documented here http://hl7.org/fhir/StructureDefinition/EnrollmentRequest
// This resource provides the insurance enrollment details to the insurer regarding a specified coverage.
type EnrollmentRequest struct {
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
	Created           *string                       `json:"created,omitempty"`
	Insurer           *Reference                    `json:"insurer,omitempty"`
	Provider          *Reference                    `json:"provider,omitempty"`
	Candidate         *Reference                    `json:"candidate,omitempty"`
	Coverage          *Reference                    `json:"coverage,omitempty"`
}

// This function returns resource reference information
func (r EnrollmentRequest) ResourceRef() (string, *string) {
	return "EnrollmentRequest", r.ID
}

// This function returns resource reference information
func (r EnrollmentRequest) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherEnrollmentRequest EnrollmentRequest

// MarshalJSON marshals the given EnrollmentRequest as JSON into a byte slice
func (r EnrollmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEnrollmentRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherEnrollmentRequest: OtherEnrollmentRequest(r),
		ResourceType:           "EnrollmentRequest",
	})
}

// UnmarshalEnrollmentRequest unmarshals a EnrollmentRequest.
func UnmarshalEnrollmentRequest(b []byte) (EnrollmentRequest, error) {
	var enrollmentRequest EnrollmentRequest
	if err := json.Unmarshal(b, &enrollmentRequest); err != nil {
		return enrollmentRequest, err
	}
	return enrollmentRequest, nil
}
