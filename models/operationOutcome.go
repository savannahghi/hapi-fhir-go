package models

// OperationOutcome is documented here http://hl7.org/fhir/StructureDefinition/OperationOutcome
type OperationOutcome struct {
	ID                *string                 `json:"id,omitempty"`
	Meta              *FHIRMeta               `json:"meta,omitempty"`
	ImplicitRules     *string                 `json:"implicitRules,omitempty"`
	Language          *string                 `json:"language,omitempty"`
	Text              *FHIRNarrative          `json:"text,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Issue             []OperationOutcomeIssue `json:"issue"`
}

type OperationOutcomeIssue struct {
	ID                *string              `json:"id,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	Severity          IssueSeverity        `json:"severity"`
	Details           *FHIRCodeableConcept `json:"details,omitempty"`
	Diagnostics       *string              `json:"diagnostics,omitempty"`
	Location          []string             `json:"location,omitempty"`
	Expression        []string             `json:"expression,omitempty"`
}

// IssueSeverity is documented here http://hl7.org/fhir/ValueSet/issue-severity
type IssueSeverity string

const (
	IssueSeverityFatal       IssueSeverity = "fatal"
	IssueSeverityError       IssueSeverity = "error"
	IssueSeverityWarning     IssueSeverity = "warning"
	IssueSeverityInformation IssueSeverity = "information"
	IssueSeveritySuccess     IssueSeverity = "severity"
)
