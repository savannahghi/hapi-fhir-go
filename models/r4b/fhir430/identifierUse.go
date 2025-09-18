
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// IdentifierUse is documented here http://hl7.org/fhir/ValueSet/identifier-use
type IdentifierUse int

const (
	IdentifierUseUsual IdentifierUse = iota
	IdentifierUseOfficial
	IdentifierUseTemp
	IdentifierUseSecondary
	IdentifierUseOld
)

func (code IdentifierUse) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *IdentifierUse) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal IdentifierUse code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "usual":
		*code = IdentifierUseUsual
	case "official":
		*code = IdentifierUseOfficial
	case "temp":
		*code = IdentifierUseTemp
	case "secondary":
		*code = IdentifierUseSecondary
	case "old":
		*code = IdentifierUseOld
	default:
		return fmt.Errorf("unknown IdentifierUse code `%s`", s)
	}
	return nil
}
func (code IdentifierUse) String() string {
	return code.Code()
}
func (code IdentifierUse) Code() string {
	switch code {
	case IdentifierUseUsual:
		return "usual"
	case IdentifierUseOfficial:
		return "official"
	case IdentifierUseTemp:
		return "temp"
	case IdentifierUseSecondary:
		return "secondary"
	case IdentifierUseOld:
		return "old"
	}
	return "<unknown>"
}
func (code IdentifierUse) Display() string {
	switch code {
	case IdentifierUseUsual:
		return "Usual"
	case IdentifierUseOfficial:
		return "Official"
	case IdentifierUseTemp:
		return "Temp"
	case IdentifierUseSecondary:
		return "Secondary"
	case IdentifierUseOld:
		return "Old"
	}
	return "<unknown>"
}
func (code IdentifierUse) Definition() string {
	switch code {
	case IdentifierUseUsual:
		return "The identifier recommended for display and use in real-world interactions."
	case IdentifierUseOfficial:
		return "The identifier considered to be most trusted for the identification of this item. Sometimes also known as \"primary\" and \"main\". The determination of \"official\" is subjective and implementation guides often provide additional guidelines for use."
	case IdentifierUseTemp:
		return "A temporary identifier."
	case IdentifierUseSecondary:
		return "An identifier that was assigned in secondary use - it serves to identify the object in a relative context, but cannot be consistently assigned to the same object again in a different context."
	case IdentifierUseOld:
		return "The identifier id no longer considered valid, but may be relevant for search purposes.  E.g. Changes to identifier schemes, account merges, etc."
	}
	return "<unknown>"
}
