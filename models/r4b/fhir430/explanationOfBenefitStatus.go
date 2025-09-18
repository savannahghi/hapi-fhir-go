
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ExplanationOfBenefitStatus is documented here http://hl7.org/fhir/ValueSet/explanationofbenefit-status
type ExplanationOfBenefitStatus int

const (
	ExplanationOfBenefitStatusActive ExplanationOfBenefitStatus = iota
	ExplanationOfBenefitStatusCancelled
	ExplanationOfBenefitStatusDraft
	ExplanationOfBenefitStatusEnteredInError
)

func (code ExplanationOfBenefitStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ExplanationOfBenefitStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ExplanationOfBenefitStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "active":
		*code = ExplanationOfBenefitStatusActive
	case "cancelled":
		*code = ExplanationOfBenefitStatusCancelled
	case "draft":
		*code = ExplanationOfBenefitStatusDraft
	case "entered-in-error":
		*code = ExplanationOfBenefitStatusEnteredInError
	default:
		return fmt.Errorf("unknown ExplanationOfBenefitStatus code `%s`", s)
	}
	return nil
}
func (code ExplanationOfBenefitStatus) String() string {
	return code.Code()
}
func (code ExplanationOfBenefitStatus) Code() string {
	switch code {
	case ExplanationOfBenefitStatusActive:
		return "active"
	case ExplanationOfBenefitStatusCancelled:
		return "cancelled"
	case ExplanationOfBenefitStatusDraft:
		return "draft"
	case ExplanationOfBenefitStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code ExplanationOfBenefitStatus) Display() string {
	switch code {
	case ExplanationOfBenefitStatusActive:
		return "Active"
	case ExplanationOfBenefitStatusCancelled:
		return "Cancelled"
	case ExplanationOfBenefitStatusDraft:
		return "Draft"
	case ExplanationOfBenefitStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}
func (code ExplanationOfBenefitStatus) Definition() string {
	switch code {
	case ExplanationOfBenefitStatusActive:
		return "The resource instance is currently in-force."
	case ExplanationOfBenefitStatusCancelled:
		return "The resource instance is withdrawn, rescinded or reversed."
	case ExplanationOfBenefitStatusDraft:
		return "A new resource instance the contents of which is not complete."
	case ExplanationOfBenefitStatusEnteredInError:
		return "The resource instance was entered in error."
	}
	return "<unknown>"
}
