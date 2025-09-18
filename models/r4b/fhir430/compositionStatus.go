
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// CompositionStatus is documented here http://hl7.org/fhir/ValueSet/composition-status
type CompositionStatus int

const (
	CompositionStatusPreliminary CompositionStatus = iota
	CompositionStatusFinal
	CompositionStatusAmended
	CompositionStatusEnteredInError
)

func (code CompositionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CompositionStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal CompositionStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "preliminary":
		*code = CompositionStatusPreliminary
	case "final":
		*code = CompositionStatusFinal
	case "amended":
		*code = CompositionStatusAmended
	case "entered-in-error":
		*code = CompositionStatusEnteredInError
	default:
		return fmt.Errorf("unknown CompositionStatus code `%s`", s)
	}
	return nil
}
func (code CompositionStatus) String() string {
	return code.Code()
}
func (code CompositionStatus) Code() string {
	switch code {
	case CompositionStatusPreliminary:
		return "preliminary"
	case CompositionStatusFinal:
		return "final"
	case CompositionStatusAmended:
		return "amended"
	case CompositionStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code CompositionStatus) Display() string {
	switch code {
	case CompositionStatusPreliminary:
		return "Preliminary"
	case CompositionStatusFinal:
		return "Final"
	case CompositionStatusAmended:
		return "Amended"
	case CompositionStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code CompositionStatus) Definition() string {
	switch code {
	case CompositionStatusPreliminary:
		return "This is a preliminary composition or document (also known as initial or interim). The content may be incomplete or unverified."
	case CompositionStatusFinal:
		return "This version of the composition is complete and verified by an appropriate person and no further work is planned. Any subsequent updates would be on a new version of the composition."
	case CompositionStatusAmended:
		return "The composition content or the referenced resources have been modified (edited or added to) subsequent to being released as \"final\" and the composition is complete and verified by an authorized person."
	case CompositionStatusEnteredInError:
		return "The composition or document was originally created/issued in error, and this is an amendment that marks that the entire series should not be considered as valid."
	}
	return "<unknown>"
}
