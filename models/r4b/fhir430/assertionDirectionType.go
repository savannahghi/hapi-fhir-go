
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// AssertionDirectionType is documented here http://hl7.org/fhir/ValueSet/assert-direction-codes
type AssertionDirectionType int

const (
	AssertionDirectionTypeResponse AssertionDirectionType = iota
	AssertionDirectionTypeRequest
)

func (code AssertionDirectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AssertionDirectionType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal AssertionDirectionType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "response":
		*code = AssertionDirectionTypeResponse
	case "request":
		*code = AssertionDirectionTypeRequest
	default:
		return fmt.Errorf("unknown AssertionDirectionType code `%s`", s)
	}
	return nil
}
func (code AssertionDirectionType) String() string {
	return code.Code()
}
func (code AssertionDirectionType) Code() string {
	switch code {
	case AssertionDirectionTypeResponse:
		return "response"
	case AssertionDirectionTypeRequest:
		return "request"
	}
	return "<unknown>"
}
func (code AssertionDirectionType) Display() string {
	switch code {
	case AssertionDirectionTypeResponse:
		return "response"
	case AssertionDirectionTypeRequest:
		return "request"
	}
	return "<unknown>"
}
func (code AssertionDirectionType) Definition() string {
	switch code {
	case AssertionDirectionTypeResponse:
		return "The assertion is evaluated on the response. This is the default value."
	case AssertionDirectionTypeRequest:
		return "The assertion is evaluated on the request."
	}
	return "<unknown>"
}
