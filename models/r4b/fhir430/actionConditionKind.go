
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ActionConditionKind is documented here http://hl7.org/fhir/ValueSet/action-condition-kind
type ActionConditionKind int

const (
	ActionConditionKindApplicability ActionConditionKind = iota
	ActionConditionKindStart
	ActionConditionKindStop
)

func (code ActionConditionKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ActionConditionKind) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ActionConditionKind code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "applicability":
		*code = ActionConditionKindApplicability
	case "start":
		*code = ActionConditionKindStart
	case "stop":
		*code = ActionConditionKindStop
	default:
		return fmt.Errorf("unknown ActionConditionKind code `%s`", s)
	}
	return nil
}
func (code ActionConditionKind) String() string {
	return code.Code()
}
func (code ActionConditionKind) Code() string {
	switch code {
	case ActionConditionKindApplicability:
		return "applicability"
	case ActionConditionKindStart:
		return "start"
	case ActionConditionKindStop:
		return "stop"
	}
	return "<unknown>"
}
func (code ActionConditionKind) Display() string {
	switch code {
	case ActionConditionKindApplicability:
		return "Applicability"
	case ActionConditionKindStart:
		return "Start"
	case ActionConditionKindStop:
		return "Stop"
	}
	return "<unknown>"
}
func (code ActionConditionKind) Definition() string {
	switch code {
	case ActionConditionKindApplicability:
		return "The condition describes whether or not a given action is applicable."
	case ActionConditionKindStart:
		return "The condition is a starting condition for the action."
	case ActionConditionKindStop:
		return "The condition is a stop, or exit condition for the action."
	}
	return "<unknown>"
}
