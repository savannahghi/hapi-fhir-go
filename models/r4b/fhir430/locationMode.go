
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// LocationMode is documented here http://hl7.org/fhir/ValueSet/location-mode
type LocationMode int

const (
	LocationModeInstance LocationMode = iota
	LocationModeKind
)

func (code LocationMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *LocationMode) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal LocationMode code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "instance":
		*code = LocationModeInstance
	case "kind":
		*code = LocationModeKind
	default:
		return fmt.Errorf("unknown LocationMode code `%s`", s)
	}
	return nil
}
func (code LocationMode) String() string {
	return code.Code()
}
func (code LocationMode) Code() string {
	switch code {
	case LocationModeInstance:
		return "instance"
	case LocationModeKind:
		return "kind"
	}
	return "<unknown>"
}
func (code LocationMode) Display() string {
	switch code {
	case LocationModeInstance:
		return "Instance"
	case LocationModeKind:
		return "Kind"
	}
	return "<unknown>"
}
func (code LocationMode) Definition() string {
	switch code {
	case LocationModeInstance:
		return "The Location resource represents a specific instance of a location (e.g. Operating Theatre 1A)."
	case LocationModeKind:
		return "The Location represents a class of locations (e.g. Any Operating Theatre) although this class of locations could be constrained within a specific boundary (such as organization, or parent location, address etc.)."
	}
	return "<unknown>"
}
