
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ConstraintSeverity is documented here http://hl7.org/fhir/ValueSet/constraint-severity
type ConstraintSeverity int

const (
	ConstraintSeverityError ConstraintSeverity = iota
	ConstraintSeverityWarning
)

func (code ConstraintSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ConstraintSeverity) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ConstraintSeverity code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "error":
		*code = ConstraintSeverityError
	case "warning":
		*code = ConstraintSeverityWarning
	default:
		return fmt.Errorf("unknown ConstraintSeverity code `%s`", s)
	}
	return nil
}
func (code ConstraintSeverity) String() string {
	return code.Code()
}
func (code ConstraintSeverity) Code() string {
	switch code {
	case ConstraintSeverityError:
		return "error"
	case ConstraintSeverityWarning:
		return "warning"
	}
	return "<unknown>"
}
func (code ConstraintSeverity) Display() string {
	switch code {
	case ConstraintSeverityError:
		return "Error"
	case ConstraintSeverityWarning:
		return "Warning"
	}
	return "<unknown>"
}
func (code ConstraintSeverity) Definition() string {
	switch code {
	case ConstraintSeverityError:
		return "If the constraint is violated, the resource is not conformant."
	case ConstraintSeverityWarning:
		return "If the constraint is violated, the resource is conformant, but it is not necessarily following best practice."
	}
	return "<unknown>"
}
