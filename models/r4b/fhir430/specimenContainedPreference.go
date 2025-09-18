
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// SpecimenContainedPreference is documented here http://hl7.org/fhir/ValueSet/specimen-contained-preference
type SpecimenContainedPreference int

const (
	SpecimenContainedPreferencePreferred SpecimenContainedPreference = iota
	SpecimenContainedPreferenceAlternate
)

func (code SpecimenContainedPreference) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SpecimenContainedPreference) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal SpecimenContainedPreference code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "preferred":
		*code = SpecimenContainedPreferencePreferred
	case "alternate":
		*code = SpecimenContainedPreferenceAlternate
	default:
		return fmt.Errorf("unknown SpecimenContainedPreference code `%s`", s)
	}
	return nil
}
func (code SpecimenContainedPreference) String() string {
	return code.Code()
}
func (code SpecimenContainedPreference) Code() string {
	switch code {
	case SpecimenContainedPreferencePreferred:
		return "preferred"
	case SpecimenContainedPreferenceAlternate:
		return "alternate"
	}
	return "<unknown>"
}
func (code SpecimenContainedPreference) Display() string {
	switch code {
	case SpecimenContainedPreferencePreferred:
		return "Preferred"
	case SpecimenContainedPreferenceAlternate:
		return "Alternate"
	}
	return "<unknown>"
}
func (code SpecimenContainedPreference) Definition() string {
	switch code {
	case SpecimenContainedPreferencePreferred:
		return "This type of contained specimen is preferred to collect this kind of specimen."
	case SpecimenContainedPreferenceAlternate:
		return "This type of conditioned specimen is an alternate."
	}
	return "<unknown>"
}
