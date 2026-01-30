
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// AllergyIntoleranceSeverity is documented here http://hl7.org/fhir/ValueSet/reaction-event-severity
type AllergyIntoleranceSeverity int

const (
	AllergyIntoleranceSeverityMild AllergyIntoleranceSeverity = iota
	AllergyIntoleranceSeverityModerate
	AllergyIntoleranceSeveritySevere
)

func (code AllergyIntoleranceSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AllergyIntoleranceSeverity) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal AllergyIntoleranceSeverity code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "mild":
		*code = AllergyIntoleranceSeverityMild
	case "moderate":
		*code = AllergyIntoleranceSeverityModerate
	case "severe":
		*code = AllergyIntoleranceSeveritySevere
	default:
		return fmt.Errorf("unknown AllergyIntoleranceSeverity code `%s`", s)
	}
	return nil
}
func (code AllergyIntoleranceSeverity) String() string {
	return code.Code()
}
func (code AllergyIntoleranceSeverity) Code() string {
	switch code {
	case AllergyIntoleranceSeverityMild:
		return "mild"
	case AllergyIntoleranceSeverityModerate:
		return "moderate"
	case AllergyIntoleranceSeveritySevere:
		return "severe"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceSeverity) Display() string {
	switch code {
	case AllergyIntoleranceSeverityMild:
		return "Mild"
	case AllergyIntoleranceSeverityModerate:
		return "Moderate"
	case AllergyIntoleranceSeveritySevere:
		return "Severe"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceSeverity) Definition() string {
	switch code {
	case AllergyIntoleranceSeverityMild:
		return "Causes mild physiological effects."
	case AllergyIntoleranceSeverityModerate:
		return "Causes moderate physiological effects."
	case AllergyIntoleranceSeveritySevere:
		return "Causes severe physiological effects."
	}
	return "<unknown>"
}
