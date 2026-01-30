
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// EvidenceVariableType is documented here http://hl7.org/fhir/ValueSet/variable-type
type EvidenceVariableType int

const (
	EvidenceVariableTypeDichotomous EvidenceVariableType = iota
	EvidenceVariableTypeContinuous
	EvidenceVariableTypeDescriptive
)

func (code EvidenceVariableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *EvidenceVariableType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal EvidenceVariableType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "dichotomous":
		*code = EvidenceVariableTypeDichotomous
	case "continuous":
		*code = EvidenceVariableTypeContinuous
	case "descriptive":
		*code = EvidenceVariableTypeDescriptive
	default:
		return fmt.Errorf("unknown EvidenceVariableType code `%s`", s)
	}
	return nil
}
func (code EvidenceVariableType) String() string {
	return code.Code()
}
func (code EvidenceVariableType) Code() string {
	switch code {
	case EvidenceVariableTypeDichotomous:
		return "dichotomous"
	case EvidenceVariableTypeContinuous:
		return "continuous"
	case EvidenceVariableTypeDescriptive:
		return "descriptive"
	}
	return "<unknown>"
}
func (code EvidenceVariableType) Display() string {
	switch code {
	case EvidenceVariableTypeDichotomous:
		return "Dichotomous"
	case EvidenceVariableTypeContinuous:
		return "Continuous"
	case EvidenceVariableTypeDescriptive:
		return "Descriptive"
	}
	return "<unknown>"
}
func (code EvidenceVariableType) Definition() string {
	switch code {
	case EvidenceVariableTypeDichotomous:
		return "The variable is dichotomous, such as present or absent."
	case EvidenceVariableTypeContinuous:
		return "The variable is a continuous result such as a quantity."
	case EvidenceVariableTypeDescriptive:
		return "The variable is described narratively rather than quantitatively."
	}
	return "<unknown>"
}
