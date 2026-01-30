
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ConceptMapGroupUnmappedMode is documented here http://hl7.org/fhir/ValueSet/conceptmap-unmapped-mode
type ConceptMapGroupUnmappedMode int

const (
	ConceptMapGroupUnmappedModeProvided ConceptMapGroupUnmappedMode = iota
	ConceptMapGroupUnmappedModeFixed
	ConceptMapGroupUnmappedModeOtherMap
)

func (code ConceptMapGroupUnmappedMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ConceptMapGroupUnmappedMode) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ConceptMapGroupUnmappedMode code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "provided":
		*code = ConceptMapGroupUnmappedModeProvided
	case "fixed":
		*code = ConceptMapGroupUnmappedModeFixed
	case "other-map":
		*code = ConceptMapGroupUnmappedModeOtherMap
	default:
		return fmt.Errorf("unknown ConceptMapGroupUnmappedMode code `%s`", s)
	}
	return nil
}
func (code ConceptMapGroupUnmappedMode) String() string {
	return code.Code()
}
func (code ConceptMapGroupUnmappedMode) Code() string {
	switch code {
	case ConceptMapGroupUnmappedModeProvided:
		return "provided"
	case ConceptMapGroupUnmappedModeFixed:
		return "fixed"
	case ConceptMapGroupUnmappedModeOtherMap:
		return "other-map"
	}
	return "<unknown>"
}
func (code ConceptMapGroupUnmappedMode) Display() string {
	switch code {
	case ConceptMapGroupUnmappedModeProvided:
		return "Provided Code"
	case ConceptMapGroupUnmappedModeFixed:
		return "Fixed Code"
	case ConceptMapGroupUnmappedModeOtherMap:
		return "Other Map"
	}
	return "<unknown>"
}
func (code ConceptMapGroupUnmappedMode) Definition() string {
	switch code {
	case ConceptMapGroupUnmappedModeProvided:
		return "Use the code as provided in the $translate request."
	case ConceptMapGroupUnmappedModeFixed:
		return "Use the code explicitly provided in the group.unmapped."
	case ConceptMapGroupUnmappedModeOtherMap:
		return "Use the map identified by the canonical URL in the url element."
	}
	return "<unknown>"
}
