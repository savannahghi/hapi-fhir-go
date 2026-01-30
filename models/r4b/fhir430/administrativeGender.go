
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// AdministrativeGender is documented here http://hl7.org/fhir/ValueSet/administrative-gender
type AdministrativeGender int

const (
	AdministrativeGenderMale AdministrativeGender = iota
	AdministrativeGenderFemale
	AdministrativeGenderOther
	AdministrativeGenderUnknown
)

func (code AdministrativeGender) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AdministrativeGender) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal AdministrativeGender code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "male":
		*code = AdministrativeGenderMale
	case "female":
		*code = AdministrativeGenderFemale
	case "other":
		*code = AdministrativeGenderOther
	case "unknown":
		*code = AdministrativeGenderUnknown
	default:
		return fmt.Errorf("unknown AdministrativeGender code `%s`", s)
	}
	return nil
}
func (code AdministrativeGender) String() string {
	return code.Code()
}
func (code AdministrativeGender) Code() string {
	switch code {
	case AdministrativeGenderMale:
		return "male"
	case AdministrativeGenderFemale:
		return "female"
	case AdministrativeGenderOther:
		return "other"
	case AdministrativeGenderUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code AdministrativeGender) Display() string {
	switch code {
	case AdministrativeGenderMale:
		return "Male"
	case AdministrativeGenderFemale:
		return "Female"
	case AdministrativeGenderOther:
		return "Other"
	case AdministrativeGenderUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code AdministrativeGender) Definition() string {
	switch code {
	case AdministrativeGenderMale:
		return "Male."
	case AdministrativeGenderFemale:
		return "Female."
	case AdministrativeGenderOther:
		return "Other."
	case AdministrativeGenderUnknown:
		return "Unknown."
	}
	return "<unknown>"
}
