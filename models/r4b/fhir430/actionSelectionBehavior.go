
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ActionSelectionBehavior is documented here http://hl7.org/fhir/ValueSet/action-selection-behavior
type ActionSelectionBehavior int

const (
	ActionSelectionBehaviorAny ActionSelectionBehavior = iota
	ActionSelectionBehaviorAll
	ActionSelectionBehaviorAllOrNone
	ActionSelectionBehaviorExactlyOne
	ActionSelectionBehaviorAtMostOne
	ActionSelectionBehaviorOneOrMore
)

func (code ActionSelectionBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ActionSelectionBehavior) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ActionSelectionBehavior code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "any":
		*code = ActionSelectionBehaviorAny
	case "all":
		*code = ActionSelectionBehaviorAll
	case "all-or-none":
		*code = ActionSelectionBehaviorAllOrNone
	case "exactly-one":
		*code = ActionSelectionBehaviorExactlyOne
	case "at-most-one":
		*code = ActionSelectionBehaviorAtMostOne
	case "one-or-more":
		*code = ActionSelectionBehaviorOneOrMore
	default:
		return fmt.Errorf("unknown ActionSelectionBehavior code `%s`", s)
	}
	return nil
}
func (code ActionSelectionBehavior) String() string {
	return code.Code()
}
func (code ActionSelectionBehavior) Code() string {
	switch code {
	case ActionSelectionBehaviorAny:
		return "any"
	case ActionSelectionBehaviorAll:
		return "all"
	case ActionSelectionBehaviorAllOrNone:
		return "all-or-none"
	case ActionSelectionBehaviorExactlyOne:
		return "exactly-one"
	case ActionSelectionBehaviorAtMostOne:
		return "at-most-one"
	case ActionSelectionBehaviorOneOrMore:
		return "one-or-more"
	}
	return "<unknown>"
}
func (code ActionSelectionBehavior) Display() string {
	switch code {
	case ActionSelectionBehaviorAny:
		return "Any"
	case ActionSelectionBehaviorAll:
		return "All"
	case ActionSelectionBehaviorAllOrNone:
		return "All Or None"
	case ActionSelectionBehaviorExactlyOne:
		return "Exactly One"
	case ActionSelectionBehaviorAtMostOne:
		return "At Most One"
	case ActionSelectionBehaviorOneOrMore:
		return "One Or More"
	}
	return "<unknown>"
}
func (code ActionSelectionBehavior) Definition() string {
	switch code {
	case ActionSelectionBehaviorAny:
		return "Any number of the actions in the group may be chosen, from zero to all."
	case ActionSelectionBehaviorAll:
		return "All the actions in the group must be selected as a single unit."
	case ActionSelectionBehaviorAllOrNone:
		return "All the actions in the group are meant to be chosen as a single unit: either all must be selected by the end user, or none may be selected."
	case ActionSelectionBehaviorExactlyOne:
		return "The end user must choose one and only one of the selectable actions in the group. The user SHALL NOT choose none of the actions in the group."
	case ActionSelectionBehaviorAtMostOne:
		return "The end user may choose zero or at most one of the actions in the group."
	case ActionSelectionBehaviorOneOrMore:
		return "The end user must choose a minimum of one, and as many additional as desired."
	}
	return "<unknown>"
}
