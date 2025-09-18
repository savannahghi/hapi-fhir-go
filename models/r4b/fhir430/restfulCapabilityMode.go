
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// RestfulCapabilityMode is documented here http://hl7.org/fhir/ValueSet/restful-capability-mode
type RestfulCapabilityMode int

const (
	RestfulCapabilityModeClient RestfulCapabilityMode = iota
	RestfulCapabilityModeServer
)

func (code RestfulCapabilityMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *RestfulCapabilityMode) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal RestfulCapabilityMode code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "client":
		*code = RestfulCapabilityModeClient
	case "server":
		*code = RestfulCapabilityModeServer
	default:
		return fmt.Errorf("unknown RestfulCapabilityMode code `%s`", s)
	}
	return nil
}
func (code RestfulCapabilityMode) String() string {
	return code.Code()
}
func (code RestfulCapabilityMode) Code() string {
	switch code {
	case RestfulCapabilityModeClient:
		return "client"
	case RestfulCapabilityModeServer:
		return "server"
	}
	return "<unknown>"
}
func (code RestfulCapabilityMode) Display() string {
	switch code {
	case RestfulCapabilityModeClient:
		return "Client"
	case RestfulCapabilityModeServer:
		return "Server"
	}
	return "<unknown>"
}
func (code RestfulCapabilityMode) Definition() string {
	switch code {
	case RestfulCapabilityModeClient:
		return "The application acts as a client for this resource."
	case RestfulCapabilityModeServer:
		return "The application acts as a server for this resource."
	}
	return "<unknown>"
}
