
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// LocationStatus is documented here http://hl7.org/fhir/ValueSet/location-status
type LocationStatus int

const (
	LocationStatusActive LocationStatus = iota
	LocationStatusSuspended
	LocationStatusInactive
)

func (code LocationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *LocationStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal LocationStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "active":
		*code = LocationStatusActive
	case "suspended":
		*code = LocationStatusSuspended
	case "inactive":
		*code = LocationStatusInactive
	default:
		return fmt.Errorf("unknown LocationStatus code `%s`", s)
	}
	return nil
}
func (code LocationStatus) String() string {
	return code.Code()
}
func (code LocationStatus) Code() string {
	switch code {
	case LocationStatusActive:
		return "active"
	case LocationStatusSuspended:
		return "suspended"
	case LocationStatusInactive:
		return "inactive"
	}
	return "<unknown>"
}
func (code LocationStatus) Display() string {
	switch code {
	case LocationStatusActive:
		return "Active"
	case LocationStatusSuspended:
		return "Suspended"
	case LocationStatusInactive:
		return "Inactive"
	}
	return "<unknown>"
}
func (code LocationStatus) Definition() string {
	switch code {
	case LocationStatusActive:
		return "The location is operational."
	case LocationStatusSuspended:
		return "The location is temporarily closed."
	case LocationStatusInactive:
		return "The location is no longer used."
	}
	return "<unknown>"
}
