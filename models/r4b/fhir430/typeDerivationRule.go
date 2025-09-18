
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// TypeDerivationRule is documented here http://hl7.org/fhir/ValueSet/type-derivation-rule
type TypeDerivationRule int

const (
	TypeDerivationRuleSpecialization TypeDerivationRule = iota
	TypeDerivationRuleConstraint
)

func (code TypeDerivationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *TypeDerivationRule) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal TypeDerivationRule code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "specialization":
		*code = TypeDerivationRuleSpecialization
	case "constraint":
		*code = TypeDerivationRuleConstraint
	default:
		return fmt.Errorf("unknown TypeDerivationRule code `%s`", s)
	}
	return nil
}
func (code TypeDerivationRule) String() string {
	return code.Code()
}
func (code TypeDerivationRule) Code() string {
	switch code {
	case TypeDerivationRuleSpecialization:
		return "specialization"
	case TypeDerivationRuleConstraint:
		return "constraint"
	}
	return "<unknown>"
}
func (code TypeDerivationRule) Display() string {
	switch code {
	case TypeDerivationRuleSpecialization:
		return "Specialization"
	case TypeDerivationRuleConstraint:
		return "Constraint"
	}
	return "<unknown>"
}
func (code TypeDerivationRule) Definition() string {
	switch code {
	case TypeDerivationRuleSpecialization:
		return "This definition defines a new type that adds additional elements to the base type."
	case TypeDerivationRuleConstraint:
		return "This definition adds additional rules to an existing concrete type."
	}
	return "<unknown>"
}
