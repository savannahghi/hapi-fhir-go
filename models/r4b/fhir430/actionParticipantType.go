
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ActionParticipantType is documented here http://hl7.org/fhir/ValueSet/action-participant-type
type ActionParticipantType int

const (
	ActionParticipantTypePatient ActionParticipantType = iota
	ActionParticipantTypePractitioner
	ActionParticipantTypeRelatedPerson
	ActionParticipantTypeDevice
)

func (code ActionParticipantType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ActionParticipantType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ActionParticipantType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "patient":
		*code = ActionParticipantTypePatient
	case "practitioner":
		*code = ActionParticipantTypePractitioner
	case "related-person":
		*code = ActionParticipantTypeRelatedPerson
	case "device":
		*code = ActionParticipantTypeDevice
	default:
		return fmt.Errorf("unknown ActionParticipantType code `%s`", s)
	}
	return nil
}
func (code ActionParticipantType) String() string {
	return code.Code()
}
func (code ActionParticipantType) Code() string {
	switch code {
	case ActionParticipantTypePatient:
		return "patient"
	case ActionParticipantTypePractitioner:
		return "practitioner"
	case ActionParticipantTypeRelatedPerson:
		return "related-person"
	case ActionParticipantTypeDevice:
		return "device"
	}
	return "<unknown>"
}
func (code ActionParticipantType) Display() string {
	switch code {
	case ActionParticipantTypePatient:
		return "Patient"
	case ActionParticipantTypePractitioner:
		return "Practitioner"
	case ActionParticipantTypeRelatedPerson:
		return "Related Person"
	case ActionParticipantTypeDevice:
		return "Device"
	}
	return "<unknown>"
}
func (code ActionParticipantType) Definition() string {
	switch code {
	case ActionParticipantTypePatient:
		return "The participant is the patient under evaluation."
	case ActionParticipantTypePractitioner:
		return "The participant is a practitioner involved in the patient's care."
	case ActionParticipantTypeRelatedPerson:
		return "The participant is a person related to the patient."
	case ActionParticipantTypeDevice:
		return "The participant is a system or device used in the care of the patient."
	}
	return "<unknown>"
}
