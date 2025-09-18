
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// DetectedIssueSeverity is documented here http://hl7.org/fhir/ValueSet/detectedissue-severity
type DetectedIssueSeverity int

const (
	DetectedIssueSeverityHigh DetectedIssueSeverity = iota
	DetectedIssueSeverityModerate
	DetectedIssueSeverityLow
)

func (code DetectedIssueSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DetectedIssueSeverity) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal DetectedIssueSeverity code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "high":
		*code = DetectedIssueSeverityHigh
	case "moderate":
		*code = DetectedIssueSeverityModerate
	case "low":
		*code = DetectedIssueSeverityLow
	default:
		return fmt.Errorf("unknown DetectedIssueSeverity code `%s`", s)
	}
	return nil
}
func (code DetectedIssueSeverity) String() string {
	return code.Code()
}
func (code DetectedIssueSeverity) Code() string {
	switch code {
	case DetectedIssueSeverityHigh:
		return "high"
	case DetectedIssueSeverityModerate:
		return "moderate"
	case DetectedIssueSeverityLow:
		return "low"
	}
	return "<unknown>"
}
func (code DetectedIssueSeverity) Display() string {
	switch code {
	case DetectedIssueSeverityHigh:
		return "High"
	case DetectedIssueSeverityModerate:
		return "Moderate"
	case DetectedIssueSeverityLow:
		return "Low"
	}
	return "<unknown>"
}
func (code DetectedIssueSeverity) Definition() string {
	switch code {
	case DetectedIssueSeverityHigh:
		return "Indicates the issue may be life-threatening or has the potential to cause permanent injury."
	case DetectedIssueSeverityModerate:
		return "Indicates the issue may result in noticeable adverse consequences but is unlikely to be life-threatening or cause permanent injury."
	case DetectedIssueSeverityLow:
		return "Indicates the issue may result in some adverse consequences but is unlikely to substantially affect the situation of the subject."
	}
	return "<unknown>"
}
