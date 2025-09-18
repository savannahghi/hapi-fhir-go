
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ExposureState is documented here http://hl7.org/fhir/ValueSet/exposure-state
type ExposureState int

const (
	ExposureStateExposure ExposureState = iota
	ExposureStateExposureAlternative
)

func (code ExposureState) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ExposureState) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ExposureState code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "exposure":
		*code = ExposureStateExposure
	case "exposure-alternative":
		*code = ExposureStateExposureAlternative
	default:
		return fmt.Errorf("unknown ExposureState code `%s`", s)
	}
	return nil
}
func (code ExposureState) String() string {
	return code.Code()
}
func (code ExposureState) Code() string {
	switch code {
	case ExposureStateExposure:
		return "exposure"
	case ExposureStateExposureAlternative:
		return "exposure-alternative"
	}
	return "<unknown>"
}
func (code ExposureState) Display() string {
	switch code {
	case ExposureStateExposure:
		return "Exposure"
	case ExposureStateExposureAlternative:
		return "Exposure Alternative"
	}
	return "<unknown>"
}
func (code ExposureState) Definition() string {
	switch code {
	case ExposureStateExposure:
		return "used when the results by exposure is describing the results for the primary exposure of interest."
	case ExposureStateExposureAlternative:
		return "used when the results by exposure is describing the results for the alternative exposure state, control state or comparator state."
	}
	return "<unknown>"
}
