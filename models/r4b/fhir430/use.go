
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// Use is documented here http://hl7.org/fhir/ValueSet/claim-use
type Use int

const (
	UseClaim Use = iota
	UsePreauthorization
	UsePredetermination
)

func (code Use) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *Use) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal Use code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "claim":
		*code = UseClaim
	case "preauthorization":
		*code = UsePreauthorization
	case "predetermination":
		*code = UsePredetermination
	default:
		return fmt.Errorf("unknown Use code `%s`", s)
	}
	return nil
}
func (code Use) String() string {
	return code.Code()
}
func (code Use) Code() string {
	switch code {
	case UseClaim:
		return "claim"
	case UsePreauthorization:
		return "preauthorization"
	case UsePredetermination:
		return "predetermination"
	}
	return "<unknown>"
}
func (code Use) Display() string {
	switch code {
	case UseClaim:
		return "Claim"
	case UsePreauthorization:
		return "Preauthorization"
	case UsePredetermination:
		return "Predetermination"
	}
	return "<unknown>"
}
func (code Use) Definition() string {
	switch code {
	case UseClaim:
		return "The treatment is complete and this represents a Claim for the services."
	case UsePreauthorization:
		return "The treatment is proposed and this represents a Pre-authorization for the services."
	case UsePredetermination:
		return "The treatment is proposed and this represents a Pre-determination for the services."
	}
	return "<unknown>"
}
