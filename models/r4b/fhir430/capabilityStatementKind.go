
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// CapabilityStatementKind is documented here http://hl7.org/fhir/ValueSet/capability-statement-kind
type CapabilityStatementKind int

const (
	CapabilityStatementKindInstance CapabilityStatementKind = iota
	CapabilityStatementKindCapability
	CapabilityStatementKindRequirements
)

func (code CapabilityStatementKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CapabilityStatementKind) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal CapabilityStatementKind code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "instance":
		*code = CapabilityStatementKindInstance
	case "capability":
		*code = CapabilityStatementKindCapability
	case "requirements":
		*code = CapabilityStatementKindRequirements
	default:
		return fmt.Errorf("unknown CapabilityStatementKind code `%s`", s)
	}
	return nil
}
func (code CapabilityStatementKind) String() string {
	return code.Code()
}
func (code CapabilityStatementKind) Code() string {
	switch code {
	case CapabilityStatementKindInstance:
		return "instance"
	case CapabilityStatementKindCapability:
		return "capability"
	case CapabilityStatementKindRequirements:
		return "requirements"
	}
	return "<unknown>"
}
func (code CapabilityStatementKind) Display() string {
	switch code {
	case CapabilityStatementKindInstance:
		return "Instance"
	case CapabilityStatementKindCapability:
		return "Capability"
	case CapabilityStatementKindRequirements:
		return "Requirements"
	}
	return "<unknown>"
}
func (code CapabilityStatementKind) Definition() string {
	switch code {
	case CapabilityStatementKindInstance:
		return "The CapabilityStatement instance represents the present capabilities of a specific system instance.  This is the kind returned by /metadata for a FHIR server end-point."
	case CapabilityStatementKindCapability:
		return "The CapabilityStatement instance represents the capabilities of a system or piece of software, independent of a particular installation."
	case CapabilityStatementKindRequirements:
		return "The CapabilityStatement instance represents a set of requirements for other systems to meet; e.g. as part of an implementation guide or 'request for proposal'."
	}
	return "<unknown>"
}
