
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// BindingStrength is documented here http://hl7.org/fhir/ValueSet/binding-strength
type BindingStrength int

const (
	BindingStrengthRequired BindingStrength = iota
	BindingStrengthExtensible
	BindingStrengthPreferred
	BindingStrengthExample
)

func (code BindingStrength) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *BindingStrength) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal BindingStrength code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "required":
		*code = BindingStrengthRequired
	case "extensible":
		*code = BindingStrengthExtensible
	case "preferred":
		*code = BindingStrengthPreferred
	case "example":
		*code = BindingStrengthExample
	default:
		return fmt.Errorf("unknown BindingStrength code `%s`", s)
	}
	return nil
}
func (code BindingStrength) String() string {
	return code.Code()
}
func (code BindingStrength) Code() string {
	switch code {
	case BindingStrengthRequired:
		return "required"
	case BindingStrengthExtensible:
		return "extensible"
	case BindingStrengthPreferred:
		return "preferred"
	case BindingStrengthExample:
		return "example"
	}
	return "<unknown>"
}
func (code BindingStrength) Display() string {
	switch code {
	case BindingStrengthRequired:
		return "Required"
	case BindingStrengthExtensible:
		return "Extensible"
	case BindingStrengthPreferred:
		return "Preferred"
	case BindingStrengthExample:
		return "Example"
	}
	return "<unknown>"
}
func (code BindingStrength) Definition() string {
	switch code {
	case BindingStrengthRequired:
		return "To be conformant, the concept in this element SHALL be from the specified value set."
	case BindingStrengthExtensible:
		return "To be conformant, the concept in this element SHALL be from the specified value set if any of the codes within the value set can apply to the concept being communicated.  If the value set does not cover the concept (based on human review), alternate codings (or, data type allowing, text) may be included instead."
	case BindingStrengthPreferred:
		return "Instances are encouraged to draw from the specified codes for interoperability purposes but are not required to do so to be considered conformant."
	case BindingStrengthExample:
		return "Instances are not expected or even encouraged to draw from the specified value set.  The value set merely provides examples of the types of concepts intended to be included."
	}
	return "<unknown>"
}
