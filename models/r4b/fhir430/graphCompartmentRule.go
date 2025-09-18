
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// GraphCompartmentRule is documented here http://hl7.org/fhir/ValueSet/graph-compartment-rule
type GraphCompartmentRule int

const (
	GraphCompartmentRuleIdentical GraphCompartmentRule = iota
	GraphCompartmentRuleMatching
	GraphCompartmentRuleDifferent
	GraphCompartmentRuleCustom
)

func (code GraphCompartmentRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *GraphCompartmentRule) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal GraphCompartmentRule code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "identical":
		*code = GraphCompartmentRuleIdentical
	case "matching":
		*code = GraphCompartmentRuleMatching
	case "different":
		*code = GraphCompartmentRuleDifferent
	case "custom":
		*code = GraphCompartmentRuleCustom
	default:
		return fmt.Errorf("unknown GraphCompartmentRule code `%s`", s)
	}
	return nil
}
func (code GraphCompartmentRule) String() string {
	return code.Code()
}
func (code GraphCompartmentRule) Code() string {
	switch code {
	case GraphCompartmentRuleIdentical:
		return "identical"
	case GraphCompartmentRuleMatching:
		return "matching"
	case GraphCompartmentRuleDifferent:
		return "different"
	case GraphCompartmentRuleCustom:
		return "custom"
	}
	return "<unknown>"
}
func (code GraphCompartmentRule) Display() string {
	switch code {
	case GraphCompartmentRuleIdentical:
		return "Identical"
	case GraphCompartmentRuleMatching:
		return "Matching"
	case GraphCompartmentRuleDifferent:
		return "Different"
	case GraphCompartmentRuleCustom:
		return "Custom"
	}
	return "<unknown>"
}
func (code GraphCompartmentRule) Definition() string {
	switch code {
	case GraphCompartmentRuleIdentical:
		return "The compartment must be identical (the same literal reference)."
	case GraphCompartmentRuleMatching:
		return "The compartment must be the same - the record must be about the same patient, but the reference may be different."
	case GraphCompartmentRuleDifferent:
		return "The compartment must be different."
	case GraphCompartmentRuleCustom:
		return "The compartment rule is defined in the accompanying FHIRPath expression."
	}
	return "<unknown>"
}
