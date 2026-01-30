
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// OperationParameterUse is documented here http://hl7.org/fhir/ValueSet/operation-parameter-use
type OperationParameterUse int

const (
	OperationParameterUseIn OperationParameterUse = iota
	OperationParameterUseOut
)

func (code OperationParameterUse) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *OperationParameterUse) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal OperationParameterUse code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "in":
		*code = OperationParameterUseIn
	case "out":
		*code = OperationParameterUseOut
	default:
		return fmt.Errorf("unknown OperationParameterUse code `%s`", s)
	}
	return nil
}
func (code OperationParameterUse) String() string {
	return code.Code()
}
func (code OperationParameterUse) Code() string {
	switch code {
	case OperationParameterUseIn:
		return "in"
	case OperationParameterUseOut:
		return "out"
	}
	return "<unknown>"
}
func (code OperationParameterUse) Display() string {
	switch code {
	case OperationParameterUseIn:
		return "In"
	case OperationParameterUseOut:
		return "Out"
	}
	return "<unknown>"
}
func (code OperationParameterUse) Definition() string {
	switch code {
	case OperationParameterUseIn:
		return "This is an input parameter."
	case OperationParameterUseOut:
		return "This is an output parameter."
	}
	return "<unknown>"
}
