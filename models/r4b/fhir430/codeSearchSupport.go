
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// CodeSearchSupport is documented here http://hl7.org/fhir/ValueSet/code-search-support
type CodeSearchSupport int

const (
	CodeSearchSupportExplicit CodeSearchSupport = iota
	CodeSearchSupportAll
)

func (code CodeSearchSupport) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CodeSearchSupport) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal CodeSearchSupport code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "explicit":
		*code = CodeSearchSupportExplicit
	case "all":
		*code = CodeSearchSupportAll
	default:
		return fmt.Errorf("unknown CodeSearchSupport code `%s`", s)
	}
	return nil
}
func (code CodeSearchSupport) String() string {
	return code.Code()
}
func (code CodeSearchSupport) Code() string {
	switch code {
	case CodeSearchSupportExplicit:
		return "explicit"
	case CodeSearchSupportAll:
		return "all"
	}
	return "<unknown>"
}
func (code CodeSearchSupport) Display() string {
	switch code {
	case CodeSearchSupportExplicit:
		return "Explicit Codes"
	case CodeSearchSupportAll:
		return "Implicit Codes"
	}
	return "<unknown>"
}
func (code CodeSearchSupport) Definition() string {
	switch code {
	case CodeSearchSupportExplicit:
		return "The search for code on ValueSet only includes codes explicitly detailed on includes or expansions."
	case CodeSearchSupportAll:
		return "The search for code on ValueSet only includes all codes based on the expansion of the value set."
	}
	return "<unknown>"
}
