package models

import (
	"fmt"
	"strings"
)

// OperationOutcome is documented here http://hl7.org/fhir/StructureDefinition/OperationOutcome
// OperationOutcome is a FHIR resource that provides information about the outcome of an operation.
// It is used extensively throughout FHIR APIs to communicate detailed error information, warnings
// and informational messages in a standardized way.
// An example of a validation error would be:
//
//	{
//	    "resourceType": "OperationOutcome",
//	    "issue": [{
//	        "severity": "error",
//	        "code": "required",
//	        "details": {
//	            "text": "Missing required field: status"
//	        },
//	        "diagnostics": "Resource Encounter requires a status element"
//	    }]
//	}.
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
	Code              IssueType            `json:"code,omitempty"`
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

type IssueType string

const (
	IssueTypeInvalid         IssueType = "invalid"
	IssueTypeStructure       IssueType = "structure"
	IssueTypeRequired        IssueType = "required"
	IssueTypeValue           IssueType = "value"
	IssueTypeInvariant       IssueType = "invariant"
	IssueTypeSecurity        IssueType = "security"
	IssueTypeLogin           IssueType = "login"
	IssueTypeUnknown         IssueType = "unknown"
	IssueTypeExpired         IssueType = "expired"
	IssueTypeForbidden       IssueType = "forbidden"
	IssueTypeSuppressed      IssueType = "suppressed"
	IssueTypeProcessing      IssueType = "processing"
	IssueTypeNotSupported    IssueType = "not-supported"
	IssueTypeDuplicate       IssueType = "duplicate"
	IssueTypeMultipleMatches IssueType = "multiple-matches"
	IssueTypeNotFound        IssueType = "not-found"
	IssueTypeDeleted         IssueType = "deleted"
	IssueTypeTooLong         IssueType = "too-long"
	IssueTypeCodeInvalid     IssueType = "code-invalid"
	IssueTypeExtension       IssueType = "extension"
	IssueTypeTooCostly       IssueType = "too-costly"
	IssueTypeBusinessRule    IssueType = "business-rule"
	IssueTypeConflict        IssueType = "conflict"
	IssueTypeLimitedFilter   IssueType = "limited-filter"
	IssueTypeTransient       IssueType = "transient"
	IssueTypeLockError       IssueType = "lock-error"
	IssueTypeNoStore         IssueType = "no-store"
	IssueTypeException       IssueType = "exception"
	IssueTypeTimeout         IssueType = "timeout"
	IssueTypeIncomplete      IssueType = "incomplete"
	IssueTypeThrottled       IssueType = "throttled"
	IssueTypeInformational   IssueType = "informational"
	IssueTypeSuccess         IssueType = "success"
)

func (i IssueType) IsValid() bool {
	switch i {
	case
		IssueTypeInvalid,
		IssueTypeStructure,
		IssueTypeRequired,
		IssueTypeValue,
		IssueTypeInvariant,
		IssueTypeSecurity,
		IssueTypeLogin,
		IssueTypeUnknown,
		IssueTypeExpired,
		IssueTypeForbidden,
		IssueTypeSuppressed,
		IssueTypeProcessing,
		IssueTypeNotSupported,
		IssueTypeDuplicate,
		IssueTypeMultipleMatches,
		IssueTypeNotFound,
		IssueTypeDeleted,
		IssueTypeTooLong,
		IssueTypeCodeInvalid,
		IssueTypeExtension,
		IssueTypeTooCostly,
		IssueTypeBusinessRule,
		IssueTypeConflict,
		IssueTypeLimitedFilter,
		IssueTypeTransient,
		IssueTypeLockError,
		IssueTypeNoStore,
		IssueTypeException,
		IssueTypeTimeout,
		IssueTypeIncomplete,
		IssueTypeThrottled,
		IssueTypeInformational,
		IssueTypeSuccess:
		return true
	default:
		return false
	}
}

func (i IssueType) String() string {
	return string(i)
}

func (o *OperationOutcome) ErrorLogging() string {
	var errors, warnings []string

	for _, issue := range o.Issue {
		message := fmt.Sprintf("- FHIR error (HTTP %s): %s", issue.Code, *issue.Diagnostics)
		if issue.Severity == "error" {
			errors = append(errors, message)
		} else if issue.Severity == "warning" {
			warnings = append(warnings, message)
		}
	}

	var output strings.Builder

	output.WriteString("validation errors:\n\n")

	if len(errors) > 0 {
		output.WriteString("Errors:\n")
		output.WriteString(strings.Join(errors, "\n") + "\n\n")
	}

	if len(warnings) > 0 {
		output.WriteString("Warnings:\n")
		output.WriteString(strings.Join(warnings, "\n") + "\n")
	}

	return output.String()
}
