
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ObservationRangeCategory is documented here http://hl7.org/fhir/ValueSet/observation-range-category
type ObservationRangeCategory int

const (
	ObservationRangeCategoryReference ObservationRangeCategory = iota
	ObservationRangeCategoryCritical
	ObservationRangeCategoryAbsolute
)

func (code ObservationRangeCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ObservationRangeCategory) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ObservationRangeCategory code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "reference":
		*code = ObservationRangeCategoryReference
	case "critical":
		*code = ObservationRangeCategoryCritical
	case "absolute":
		*code = ObservationRangeCategoryAbsolute
	default:
		return fmt.Errorf("unknown ObservationRangeCategory code `%s`", s)
	}
	return nil
}
func (code ObservationRangeCategory) String() string {
	return code.Code()
}
func (code ObservationRangeCategory) Code() string {
	switch code {
	case ObservationRangeCategoryReference:
		return "reference"
	case ObservationRangeCategoryCritical:
		return "critical"
	case ObservationRangeCategoryAbsolute:
		return "absolute"
	}
	return "<unknown>"
}
func (code ObservationRangeCategory) Display() string {
	switch code {
	case ObservationRangeCategoryReference:
		return "reference range"
	case ObservationRangeCategoryCritical:
		return "critical range"
	case ObservationRangeCategoryAbsolute:
		return "absolute range"
	}
	return "<unknown>"
}
func (code ObservationRangeCategory) Definition() string {
	switch code {
	case ObservationRangeCategoryReference:
		return "Reference (Normal) Range for Ordinal and Continuous Observations."
	case ObservationRangeCategoryCritical:
		return "Critical Range for Ordinal and Continuous Observations."
	case ObservationRangeCategoryAbsolute:
		return "Absolute Range for Ordinal and Continuous Observations. Results outside this range are not possible."
	}
	return "<unknown>"
}
