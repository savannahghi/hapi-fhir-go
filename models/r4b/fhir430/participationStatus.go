
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ParticipationStatus is documented here http://hl7.org/fhir/ValueSet/participationstatus
type ParticipationStatus int

const (
	ParticipationStatusAccepted ParticipationStatus = iota
	ParticipationStatusDeclined
	ParticipationStatusTentative
	ParticipationStatusNeedsAction
)

func (code ParticipationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ParticipationStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ParticipationStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "accepted":
		*code = ParticipationStatusAccepted
	case "declined":
		*code = ParticipationStatusDeclined
	case "tentative":
		*code = ParticipationStatusTentative
	case "needs-action":
		*code = ParticipationStatusNeedsAction
	default:
		return fmt.Errorf("unknown ParticipationStatus code `%s`", s)
	}
	return nil
}
func (code ParticipationStatus) String() string {
	return code.Code()
}
func (code ParticipationStatus) Code() string {
	switch code {
	case ParticipationStatusAccepted:
		return "accepted"
	case ParticipationStatusDeclined:
		return "declined"
	case ParticipationStatusTentative:
		return "tentative"
	case ParticipationStatusNeedsAction:
		return "needs-action"
	}
	return "<unknown>"
}
func (code ParticipationStatus) Display() string {
	switch code {
	case ParticipationStatusAccepted:
		return "Accepted"
	case ParticipationStatusDeclined:
		return "Declined"
	case ParticipationStatusTentative:
		return "Tentative"
	case ParticipationStatusNeedsAction:
		return "Needs Action"
	}
	return "<unknown>"
}
func (code ParticipationStatus) Definition() string {
	switch code {
	case ParticipationStatusAccepted:
		return "The participant has accepted the appointment."
	case ParticipationStatusDeclined:
		return "The participant has declined the appointment and will not participate in the appointment."
	case ParticipationStatusTentative:
		return "The participant has  tentatively accepted the appointment. This could be automatically created by a system and requires further processing before it can be accepted. There is no commitment that attendance will occur."
	case ParticipationStatusNeedsAction:
		return "The participant needs to indicate if they accept the appointment by changing this status to one of the other statuses."
	}
	return "<unknown>"
}
