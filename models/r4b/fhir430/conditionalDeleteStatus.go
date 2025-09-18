
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ConditionalDeleteStatus is documented here http://hl7.org/fhir/ValueSet/conditional-delete-status
type ConditionalDeleteStatus int

const (
	ConditionalDeleteStatusNotSupported ConditionalDeleteStatus = iota
	ConditionalDeleteStatusSingle
	ConditionalDeleteStatusMultiple
)

func (code ConditionalDeleteStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ConditionalDeleteStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ConditionalDeleteStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "not-supported":
		*code = ConditionalDeleteStatusNotSupported
	case "single":
		*code = ConditionalDeleteStatusSingle
	case "multiple":
		*code = ConditionalDeleteStatusMultiple
	default:
		return fmt.Errorf("unknown ConditionalDeleteStatus code `%s`", s)
	}
	return nil
}
func (code ConditionalDeleteStatus) String() string {
	return code.Code()
}
func (code ConditionalDeleteStatus) Code() string {
	switch code {
	case ConditionalDeleteStatusNotSupported:
		return "not-supported"
	case ConditionalDeleteStatusSingle:
		return "single"
	case ConditionalDeleteStatusMultiple:
		return "multiple"
	}
	return "<unknown>"
}
func (code ConditionalDeleteStatus) Display() string {
	switch code {
	case ConditionalDeleteStatusNotSupported:
		return "Not Supported"
	case ConditionalDeleteStatusSingle:
		return "Single Deletes Supported"
	case ConditionalDeleteStatusMultiple:
		return "Multiple Deletes Supported"
	}
	return "<unknown>"
}
func (code ConditionalDeleteStatus) Definition() string {
	switch code {
	case ConditionalDeleteStatusNotSupported:
		return "No support for conditional deletes."
	case ConditionalDeleteStatusSingle:
		return "Conditional deletes are supported, but only single resources at a time."
	case ConditionalDeleteStatusMultiple:
		return "Conditional deletes are supported, and multiple resources can be deleted in a single interaction."
	}
	return "<unknown>"
}
