
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// FHIRSubstanceStatus is documented here http://hl7.org/fhir/ValueSet/substance-status
type FHIRSubstanceStatus int

const (
	FHIRSubstanceStatusActive FHIRSubstanceStatus = iota
	FHIRSubstanceStatusInactive
	FHIRSubstanceStatusEnteredInError
)

func (code FHIRSubstanceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *FHIRSubstanceStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal FHIRSubstanceStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "active":
		*code = FHIRSubstanceStatusActive
	case "inactive":
		*code = FHIRSubstanceStatusInactive
	case "entered-in-error":
		*code = FHIRSubstanceStatusEnteredInError
	default:
		return fmt.Errorf("unknown FHIRSubstanceStatus code `%s`", s)
	}
	return nil
}
func (code FHIRSubstanceStatus) String() string {
	return code.Code()
}
func (code FHIRSubstanceStatus) Code() string {
	switch code {
	case FHIRSubstanceStatusActive:
		return "active"
	case FHIRSubstanceStatusInactive:
		return "inactive"
	case FHIRSubstanceStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code FHIRSubstanceStatus) Display() string {
	switch code {
	case FHIRSubstanceStatusActive:
		return "Active"
	case FHIRSubstanceStatusInactive:
		return "Inactive"
	case FHIRSubstanceStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code FHIRSubstanceStatus) Definition() string {
	switch code {
	case FHIRSubstanceStatusActive:
		return "The substance is considered for use or reference."
	case FHIRSubstanceStatusInactive:
		return "The substance is considered for reference, but not for use."
	case FHIRSubstanceStatusEnteredInError:
		return "The substance was entered in error."
	}
	return "<unknown>"
}
