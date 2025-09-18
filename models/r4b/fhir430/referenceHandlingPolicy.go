
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ReferenceHandlingPolicy is documented here http://hl7.org/fhir/ValueSet/reference-handling-policy
type ReferenceHandlingPolicy int

const (
	ReferenceHandlingPolicyLiteral ReferenceHandlingPolicy = iota
	ReferenceHandlingPolicyLogical
	ReferenceHandlingPolicyResolves
	ReferenceHandlingPolicyEnforced
	ReferenceHandlingPolicyLocal
)

func (code ReferenceHandlingPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ReferenceHandlingPolicy) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ReferenceHandlingPolicy code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "literal":
		*code = ReferenceHandlingPolicyLiteral
	case "logical":
		*code = ReferenceHandlingPolicyLogical
	case "resolves":
		*code = ReferenceHandlingPolicyResolves
	case "enforced":
		*code = ReferenceHandlingPolicyEnforced
	case "local":
		*code = ReferenceHandlingPolicyLocal
	default:
		return fmt.Errorf("unknown ReferenceHandlingPolicy code `%s`", s)
	}
	return nil
}
func (code ReferenceHandlingPolicy) String() string {
	return code.Code()
}
func (code ReferenceHandlingPolicy) Code() string {
	switch code {
	case ReferenceHandlingPolicyLiteral:
		return "literal"
	case ReferenceHandlingPolicyLogical:
		return "logical"
	case ReferenceHandlingPolicyResolves:
		return "resolves"
	case ReferenceHandlingPolicyEnforced:
		return "enforced"
	case ReferenceHandlingPolicyLocal:
		return "local"
	}
	return "<unknown>"
}
func (code ReferenceHandlingPolicy) Display() string {
	switch code {
	case ReferenceHandlingPolicyLiteral:
		return "Literal References"
	case ReferenceHandlingPolicyLogical:
		return "Logical References"
	case ReferenceHandlingPolicyResolves:
		return "Resolves References"
	case ReferenceHandlingPolicyEnforced:
		return "Reference Integrity Enforced"
	case ReferenceHandlingPolicyLocal:
		return "Local References Only"
	}
	return "<unknown>"
}
func (code ReferenceHandlingPolicy) Definition() string {
	switch code {
	case ReferenceHandlingPolicyLiteral:
		return "The server supports and populates Literal references (i.e. using Reference.reference) where they are known (this code does not guarantee that all references are literal; see 'enforced')."
	case ReferenceHandlingPolicyLogical:
		return "The server allows logical references (i.e. using Reference.identifier)."
	case ReferenceHandlingPolicyResolves:
		return "The server will attempt to resolve logical references to literal references - i.e. converting Reference.identifier to Reference.reference (if resolution fails, the server may still accept resources; see logical)."
	case ReferenceHandlingPolicyEnforced:
		return "The server enforces that references have integrity - e.g. it ensures that references can always be resolved. This is typically the case for clinical record systems, but often not the case for middleware/proxy systems."
	case ReferenceHandlingPolicyLocal:
		return "The server does not support references that point to other servers."
	}
	return "<unknown>"
}
