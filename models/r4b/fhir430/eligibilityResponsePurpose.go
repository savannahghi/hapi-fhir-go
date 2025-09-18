
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// EligibilityResponsePurpose is documented here http://hl7.org/fhir/ValueSet/eligibilityresponse-purpose
type EligibilityResponsePurpose int

const (
	EligibilityResponsePurposeAuthRequirements EligibilityResponsePurpose = iota
	EligibilityResponsePurposeBenefits
	EligibilityResponsePurposeDiscovery
	EligibilityResponsePurposeValidation
)

func (code EligibilityResponsePurpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *EligibilityResponsePurpose) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal EligibilityResponsePurpose code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "auth-requirements":
		*code = EligibilityResponsePurposeAuthRequirements
	case "benefits":
		*code = EligibilityResponsePurposeBenefits
	case "discovery":
		*code = EligibilityResponsePurposeDiscovery
	case "validation":
		*code = EligibilityResponsePurposeValidation
	default:
		return fmt.Errorf("unknown EligibilityResponsePurpose code `%s`", s)
	}
	return nil
}
func (code EligibilityResponsePurpose) String() string {
	return code.Code()
}
func (code EligibilityResponsePurpose) Code() string {
	switch code {
	case EligibilityResponsePurposeAuthRequirements:
		return "auth-requirements"
	case EligibilityResponsePurposeBenefits:
		return "benefits"
	case EligibilityResponsePurposeDiscovery:
		return "discovery"
	case EligibilityResponsePurposeValidation:
		return "validation"
	}
	return "<unknown>"
}
func (code EligibilityResponsePurpose) Display() string {
	switch code {
	case EligibilityResponsePurposeAuthRequirements:
		return "Coverage auth-requirements"
	case EligibilityResponsePurposeBenefits:
		return "Coverage benefits"
	case EligibilityResponsePurposeDiscovery:
		return "Coverage Discovery"
	case EligibilityResponsePurposeValidation:
		return "Coverage Validation"
	}
	return "<unknown>"
}
func (code EligibilityResponsePurpose) Definition() string {
	switch code {
	case EligibilityResponsePurposeAuthRequirements:
		return "The prior authorization requirements for the listed, or discovered if specified, converages for the categories of service and/or specifed biling codes are requested."
	case EligibilityResponsePurposeBenefits:
		return "The plan benefits and optionally benefits consumed  for the listed, or discovered if specified, converages are requested."
	case EligibilityResponsePurposeDiscovery:
		return "The insurer is requested to report on any coverages which they are aware of in addition to any specifed."
	case EligibilityResponsePurposeValidation:
		return "A check that the specified coverages are in-force is requested."
	}
	return "<unknown>"
}
