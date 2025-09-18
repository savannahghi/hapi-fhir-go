
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ParticipantRequired is documented here http://hl7.org/fhir/ValueSet/participantrequired
type ParticipantRequired int

const (
	ParticipantRequiredRequired ParticipantRequired = iota
	ParticipantRequiredOptional
	ParticipantRequiredInformationOnly
)

func (code ParticipantRequired) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ParticipantRequired) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ParticipantRequired code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "required":
		*code = ParticipantRequiredRequired
	case "optional":
		*code = ParticipantRequiredOptional
	case "information-only":
		*code = ParticipantRequiredInformationOnly
	default:
		return fmt.Errorf("unknown ParticipantRequired code `%s`", s)
	}
	return nil
}
func (code ParticipantRequired) String() string {
	return code.Code()
}
func (code ParticipantRequired) Code() string {
	switch code {
	case ParticipantRequiredRequired:
		return "required"
	case ParticipantRequiredOptional:
		return "optional"
	case ParticipantRequiredInformationOnly:
		return "information-only"
	}
	return "<unknown>"
}
func (code ParticipantRequired) Display() string {
	switch code {
	case ParticipantRequiredRequired:
		return "Required"
	case ParticipantRequiredOptional:
		return "Optional"
	case ParticipantRequiredInformationOnly:
		return "Information Only"
	}
	return "<unknown>"
}
func (code ParticipantRequired) Definition() string {
	switch code {
	case ParticipantRequiredRequired:
		return "The participant is required to attend the appointment."
	case ParticipantRequiredOptional:
		return "The participant may optionally attend the appointment."
	case ParticipantRequiredInformationOnly:
		return "The participant is excluded from the appointment, and might not be informed of the appointment taking place. (Appointment is about them, not for them - such as 2 doctors discussing results about a patient's test)."
	}
	return "<unknown>"
}
