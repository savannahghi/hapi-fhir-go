
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ActionCardinalityBehavior is documented here http://hl7.org/fhir/ValueSet/action-cardinality-behavior
type ActionCardinalityBehavior int

const (
	ActionCardinalityBehaviorSingle ActionCardinalityBehavior = iota
	ActionCardinalityBehaviorMultiple
)

func (code ActionCardinalityBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ActionCardinalityBehavior) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ActionCardinalityBehavior code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "single":
		*code = ActionCardinalityBehaviorSingle
	case "multiple":
		*code = ActionCardinalityBehaviorMultiple
	default:
		return fmt.Errorf("unknown ActionCardinalityBehavior code `%s`", s)
	}
	return nil
}
func (code ActionCardinalityBehavior) String() string {
	return code.Code()
}
func (code ActionCardinalityBehavior) Code() string {
	switch code {
	case ActionCardinalityBehaviorSingle:
		return "single"
	case ActionCardinalityBehaviorMultiple:
		return "multiple"
	}
	return "<unknown>"
}
func (code ActionCardinalityBehavior) Display() string {
	switch code {
	case ActionCardinalityBehaviorSingle:
		return "Single"
	case ActionCardinalityBehaviorMultiple:
		return "Multiple"
	}
	return "<unknown>"
}
func (code ActionCardinalityBehavior) Definition() string {
	switch code {
	case ActionCardinalityBehaviorSingle:
		return "The action may only be selected one time."
	case ActionCardinalityBehaviorMultiple:
		return "The action may be selected multiple times."
	}
	return "<unknown>"
}
