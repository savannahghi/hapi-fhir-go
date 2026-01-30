
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// AdverseEventActuality is documented here http://hl7.org/fhir/ValueSet/adverse-event-actuality
type AdverseEventActuality int

const (
	AdverseEventActualityActual AdverseEventActuality = iota
	AdverseEventActualityPotential
)

func (code AdverseEventActuality) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AdverseEventActuality) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal AdverseEventActuality code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "actual":
		*code = AdverseEventActualityActual
	case "potential":
		*code = AdverseEventActualityPotential
	default:
		return fmt.Errorf("unknown AdverseEventActuality code `%s`", s)
	}
	return nil
}
func (code AdverseEventActuality) String() string {
	return code.Code()
}
func (code AdverseEventActuality) Code() string {
	switch code {
	case AdverseEventActualityActual:
		return "actual"
	case AdverseEventActualityPotential:
		return "potential"
	}
	return "<unknown>"
}
func (code AdverseEventActuality) Display() string {
	switch code {
	case AdverseEventActualityActual:
		return "Adverse Event"
	case AdverseEventActualityPotential:
		return "Potential Adverse Event"
	}
	return "<unknown>"
}
func (code AdverseEventActuality) Definition() string {
	switch code {
	case AdverseEventActualityActual:
		return "The adverse event actually happened regardless of whether anyone was affected or harmed."
	case AdverseEventActualityPotential:
		return "A potential adverse event."
	}
	return "<unknown>"
}
