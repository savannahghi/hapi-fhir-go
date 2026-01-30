
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// AggregationMode is documented here http://hl7.org/fhir/ValueSet/resource-aggregation-mode
type AggregationMode int

const (
	AggregationModeContained AggregationMode = iota
	AggregationModeReferenced
	AggregationModeBundled
)

func (code AggregationMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AggregationMode) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal AggregationMode code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "contained":
		*code = AggregationModeContained
	case "referenced":
		*code = AggregationModeReferenced
	case "bundled":
		*code = AggregationModeBundled
	default:
		return fmt.Errorf("unknown AggregationMode code `%s`", s)
	}
	return nil
}
func (code AggregationMode) String() string {
	return code.Code()
}
func (code AggregationMode) Code() string {
	switch code {
	case AggregationModeContained:
		return "contained"
	case AggregationModeReferenced:
		return "referenced"
	case AggregationModeBundled:
		return "bundled"
	}
	return "<unknown>"
}
func (code AggregationMode) Display() string {
	switch code {
	case AggregationModeContained:
		return "Contained"
	case AggregationModeReferenced:
		return "Referenced"
	case AggregationModeBundled:
		return "Bundled"
	}
	return "<unknown>"
}
func (code AggregationMode) Definition() string {
	switch code {
	case AggregationModeContained:
		return "The reference is a local reference to a contained resource."
	case AggregationModeReferenced:
		return "The reference to a resource that has to be resolved externally to the resource that includes the reference."
	case AggregationModeBundled:
		return "The resource the reference points to will be found in the same bundle as the resource that includes the reference."
	}
	return "<unknown>"
}
