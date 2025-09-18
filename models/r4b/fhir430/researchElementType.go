
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ResearchElementType is documented here http://hl7.org/fhir/ValueSet/research-element-type
type ResearchElementType int

const (
	ResearchElementTypePopulation ResearchElementType = iota
	ResearchElementTypeExposure
	ResearchElementTypeOutcome
)

func (code ResearchElementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ResearchElementType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ResearchElementType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "population":
		*code = ResearchElementTypePopulation
	case "exposure":
		*code = ResearchElementTypeExposure
	case "outcome":
		*code = ResearchElementTypeOutcome
	default:
		return fmt.Errorf("unknown ResearchElementType code `%s`", s)
	}
	return nil
}
func (code ResearchElementType) String() string {
	return code.Code()
}
func (code ResearchElementType) Code() string {
	switch code {
	case ResearchElementTypePopulation:
		return "population"
	case ResearchElementTypeExposure:
		return "exposure"
	case ResearchElementTypeOutcome:
		return "outcome"
	}
	return "<unknown>"
}
func (code ResearchElementType) Display() string {
	switch code {
	case ResearchElementTypePopulation:
		return "Population"
	case ResearchElementTypeExposure:
		return "Exposure"
	case ResearchElementTypeOutcome:
		return "Outcome"
	}
	return "<unknown>"
}
func (code ResearchElementType) Definition() string {
	switch code {
	case ResearchElementTypePopulation:
		return "The element defines the population that forms the basis for research."
	case ResearchElementTypeExposure:
		return "The element defines an exposure within the population that is being researched."
	case ResearchElementTypeOutcome:
		return "The element defines an outcome within the population that is being researched."
	}
	return "<unknown>"
}
