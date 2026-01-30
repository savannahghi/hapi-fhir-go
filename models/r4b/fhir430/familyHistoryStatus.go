
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// FamilyHistoryStatus is documented here http://hl7.org/fhir/ValueSet/history-status
type FamilyHistoryStatus int

const (
	FamilyHistoryStatusPartial FamilyHistoryStatus = iota
	FamilyHistoryStatusCompleted
	FamilyHistoryStatusEnteredInError
	FamilyHistoryStatusHealthUnknown
)

func (code FamilyHistoryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *FamilyHistoryStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal FamilyHistoryStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "partial":
		*code = FamilyHistoryStatusPartial
	case "completed":
		*code = FamilyHistoryStatusCompleted
	case "entered-in-error":
		*code = FamilyHistoryStatusEnteredInError
	case "health-unknown":
		*code = FamilyHistoryStatusHealthUnknown
	default:
		return fmt.Errorf("unknown FamilyHistoryStatus code `%s`", s)
	}
	return nil
}
func (code FamilyHistoryStatus) String() string {
	return code.Code()
}
func (code FamilyHistoryStatus) Code() string {
	switch code {
	case FamilyHistoryStatusPartial:
		return "partial"
	case FamilyHistoryStatusCompleted:
		return "completed"
	case FamilyHistoryStatusEnteredInError:
		return "entered-in-error"
	case FamilyHistoryStatusHealthUnknown:
		return "health-unknown"
	}
	return "<unknown>"
}
func (code FamilyHistoryStatus) Display() string {
	switch code {
	case FamilyHistoryStatusPartial:
		return "Partial"
	case FamilyHistoryStatusCompleted:
		return "Completed"
	case FamilyHistoryStatusEnteredInError:
		return "Entered in Error"
	case FamilyHistoryStatusHealthUnknown:
		return "Health Unknown"
	}
	return "<unknown>"
}
func (code FamilyHistoryStatus) Definition() string {
	switch code {
	case FamilyHistoryStatusPartial:
		return "Some health information is known and captured, but not complete - see notes for details."
	case FamilyHistoryStatusCompleted:
		return "All available related health information is captured as of the date (and possibly time) when the family member history was taken."
	case FamilyHistoryStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	case FamilyHistoryStatusHealthUnknown:
		return "Health information for this family member is unavailable/unknown."
	}
	return "<unknown>"
}
