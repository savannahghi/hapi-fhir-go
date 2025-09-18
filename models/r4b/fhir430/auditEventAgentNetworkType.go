
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// AuditEventAgentNetworkType is documented here http://hl7.org/fhir/ValueSet/network-type
type AuditEventAgentNetworkType int

const (
	AuditEventAgentNetworkType1 AuditEventAgentNetworkType = iota
	AuditEventAgentNetworkType2
	AuditEventAgentNetworkType3
	AuditEventAgentNetworkType4
	AuditEventAgentNetworkType5
)

func (code AuditEventAgentNetworkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AuditEventAgentNetworkType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal AuditEventAgentNetworkType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "1":
		*code = AuditEventAgentNetworkType1
	case "2":
		*code = AuditEventAgentNetworkType2
	case "3":
		*code = AuditEventAgentNetworkType3
	case "4":
		*code = AuditEventAgentNetworkType4
	case "5":
		*code = AuditEventAgentNetworkType5
	default:
		return fmt.Errorf("unknown AuditEventAgentNetworkType code `%s`", s)
	}
	return nil
}
func (code AuditEventAgentNetworkType) String() string {
	return code.Code()
}
func (code AuditEventAgentNetworkType) Code() string {
	switch code {
	case AuditEventAgentNetworkType1:
		return "1"
	case AuditEventAgentNetworkType2:
		return "2"
	case AuditEventAgentNetworkType3:
		return "3"
	case AuditEventAgentNetworkType4:
		return "4"
	case AuditEventAgentNetworkType5:
		return "5"
	}
	return "<unknown>"
}
func (code AuditEventAgentNetworkType) Display() string {
	switch code {
	case AuditEventAgentNetworkType1:
		return "Machine Name"
	case AuditEventAgentNetworkType2:
		return "IP Address"
	case AuditEventAgentNetworkType3:
		return "Telephone Number"
	case AuditEventAgentNetworkType4:
		return "Email address"
	case AuditEventAgentNetworkType5:
		return "URI"
	}
	return "<unknown>"
}
func (code AuditEventAgentNetworkType) Definition() string {
	switch code {
	case AuditEventAgentNetworkType1:
		return "The machine name, including DNS name."
	case AuditEventAgentNetworkType2:
		return "The assigned Internet Protocol (IP) address."
	case AuditEventAgentNetworkType3:
		return "The assigned telephone number."
	case AuditEventAgentNetworkType4:
		return "The assigned email address."
	case AuditEventAgentNetworkType5:
		return "URI (User directory, HTTP-PUT, ftp, etc.)."
	}
	return "<unknown>"
}
