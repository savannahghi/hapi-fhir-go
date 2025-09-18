
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// EnableWhenBehavior is documented here http://hl7.org/fhir/ValueSet/questionnaire-enable-behavior
type EnableWhenBehavior int

const (
	EnableWhenBehaviorAll EnableWhenBehavior = iota
	EnableWhenBehaviorAny
)

func (code EnableWhenBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *EnableWhenBehavior) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal EnableWhenBehavior code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "all":
		*code = EnableWhenBehaviorAll
	case "any":
		*code = EnableWhenBehaviorAny
	default:
		return fmt.Errorf("unknown EnableWhenBehavior code `%s`", s)
	}
	return nil
}
func (code EnableWhenBehavior) String() string {
	return code.Code()
}
func (code EnableWhenBehavior) Code() string {
	switch code {
	case EnableWhenBehaviorAll:
		return "all"
	case EnableWhenBehaviorAny:
		return "any"
	}
	return "<unknown>"
}
func (code EnableWhenBehavior) Display() string {
	switch code {
	case EnableWhenBehaviorAll:
		return "All"
	case EnableWhenBehaviorAny:
		return "Any"
	}
	return "<unknown>"
}
func (code EnableWhenBehavior) Definition() string {
	switch code {
	case EnableWhenBehaviorAll:
		return "Enable the question when all the enableWhen criteria are satisfied."
	case EnableWhenBehaviorAny:
		return "Enable the question when any of the enableWhen criteria are satisfied."
	}
	return "<unknown>"
}
