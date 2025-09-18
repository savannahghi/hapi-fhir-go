
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// NamingSystemType is documented here http://hl7.org/fhir/ValueSet/namingsystem-type
type NamingSystemType int

const (
	NamingSystemTypeCodesystem NamingSystemType = iota
	NamingSystemTypeIdentifier
	NamingSystemTypeRoot
)

func (code NamingSystemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *NamingSystemType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal NamingSystemType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "codesystem":
		*code = NamingSystemTypeCodesystem
	case "identifier":
		*code = NamingSystemTypeIdentifier
	case "root":
		*code = NamingSystemTypeRoot
	default:
		return fmt.Errorf("unknown NamingSystemType code `%s`", s)
	}
	return nil
}
func (code NamingSystemType) String() string {
	return code.Code()
}
func (code NamingSystemType) Code() string {
	switch code {
	case NamingSystemTypeCodesystem:
		return "codesystem"
	case NamingSystemTypeIdentifier:
		return "identifier"
	case NamingSystemTypeRoot:
		return "root"
	}
	return "<unknown>"
}
func (code NamingSystemType) Display() string {
	switch code {
	case NamingSystemTypeCodesystem:
		return "Code System"
	case NamingSystemTypeIdentifier:
		return "Identifier"
	case NamingSystemTypeRoot:
		return "Root"
	}
	return "<unknown>"
}
func (code NamingSystemType) Definition() string {
	switch code {
	case NamingSystemTypeCodesystem:
		return "The naming system is used to define concepts and symbols to represent those concepts; e.g. UCUM, LOINC, NDC code, local lab codes, etc."
	case NamingSystemTypeIdentifier:
		return "The naming system is used to manage identifiers (e.g. license numbers, order numbers, etc.)."
	case NamingSystemTypeRoot:
		return "The naming system is used as the root for other identifiers and naming systems."
	}
	return "<unknown>"
}
