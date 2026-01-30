package fhir430

import "encoding/json"

// OperationOutcome is documented here http://hl7.org/fhir/StructureDefinition/OperationOutcome
// A collection of error, warning, or information messages that result from a system action.
type OperationOutcome struct {
	ID                *string                 `json:"id,omitempty"`
	Meta              *Meta                   `json:"meta,omitempty"`
	ImplicitRules     *string                 `json:"implicitRules,omitempty"`
	Language          *string                 `json:"language,omitempty"`
	Text              *Narrative              `json:"text,omitempty"`
	Contained         []json.RawMessage       `json:"contained,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Issue             []OperationOutcomeIssue `json:"issue"`
}

// An error, warning, or information message that results from a system action.
type OperationOutcomeIssue struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Severity          IssueSeverity    `json:"severity"`
	Code              IssueType        `json:"code"`
	Details           *CodeableConcept `json:"details,omitempty"`
	Diagnostics       *string          `json:"diagnostics,omitempty"`
	Location          []string         `json:"location,omitempty"`
	Expression        []string         `json:"expression,omitempty"`
}

// This function returns resource reference information
func (r OperationOutcome) ResourceRef() (string, *string) {
	return "OperationOutcome", r.ID
}

// This function returns resource reference information
func (r OperationOutcome) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherOperationOutcome OperationOutcome

// MarshalJSON marshals the given OperationOutcome as JSON into a byte slice
func (r OperationOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOperationOutcome
		ResourceType string `json:"resourceType"`
	}{
		OtherOperationOutcome: OtherOperationOutcome(r),
		ResourceType:          "OperationOutcome",
	})
}

// UnmarshalOperationOutcome unmarshals a OperationOutcome.
func UnmarshalOperationOutcome(b []byte) (OperationOutcome, error) {
	var operationOutcome OperationOutcome
	if err := json.Unmarshal(b, &operationOutcome); err != nil {
		return operationOutcome, err
	}
	return operationOutcome, nil
}
