
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// GraphCompartmentUse is documented here http://hl7.org/fhir/ValueSet/graph-compartment-use
type GraphCompartmentUse int

const (
	GraphCompartmentUseCondition GraphCompartmentUse = iota
	GraphCompartmentUseRequirement
)

func (code GraphCompartmentUse) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *GraphCompartmentUse) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal GraphCompartmentUse code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "condition":
		*code = GraphCompartmentUseCondition
	case "requirement":
		*code = GraphCompartmentUseRequirement
	default:
		return fmt.Errorf("unknown GraphCompartmentUse code `%s`", s)
	}
	return nil
}
func (code GraphCompartmentUse) String() string {
	return code.Code()
}
func (code GraphCompartmentUse) Code() string {
	switch code {
	case GraphCompartmentUseCondition:
		return "condition"
	case GraphCompartmentUseRequirement:
		return "requirement"
	}
	return "<unknown>"
}
func (code GraphCompartmentUse) Display() string {
	switch code {
	case GraphCompartmentUseCondition:
		return "Condition"
	case GraphCompartmentUseRequirement:
		return "Requirement"
	}
	return "<unknown>"
}
func (code GraphCompartmentUse) Definition() string {
	switch code {
	case GraphCompartmentUseCondition:
		return "This compartment rule is a condition for whether the rule applies."
	case GraphCompartmentUseRequirement:
		return "This compartment rule is enforced on any relationships that meet the conditions."
	}
	return "<unknown>"
}
