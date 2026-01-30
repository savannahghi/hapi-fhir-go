
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// SupplyRequestStatus is documented here http://hl7.org/fhir/ValueSet/supplyrequest-status
type SupplyRequestStatus int

const (
	SupplyRequestStatusDraft SupplyRequestStatus = iota
	SupplyRequestStatusActive
	SupplyRequestStatusSuspended
	SupplyRequestStatusCancelled
	SupplyRequestStatusCompleted
	SupplyRequestStatusEnteredInError
	SupplyRequestStatusUnknown
)

func (code SupplyRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SupplyRequestStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal SupplyRequestStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "draft":
		*code = SupplyRequestStatusDraft
	case "active":
		*code = SupplyRequestStatusActive
	case "suspended":
		*code = SupplyRequestStatusSuspended
	case "cancelled":
		*code = SupplyRequestStatusCancelled
	case "completed":
		*code = SupplyRequestStatusCompleted
	case "entered-in-error":
		*code = SupplyRequestStatusEnteredInError
	case "unknown":
		*code = SupplyRequestStatusUnknown
	default:
		return fmt.Errorf("unknown SupplyRequestStatus code `%s`", s)
	}
	return nil
}
func (code SupplyRequestStatus) String() string {
	return code.Code()
}
func (code SupplyRequestStatus) Code() string {
	switch code {
	case SupplyRequestStatusDraft:
		return "draft"
	case SupplyRequestStatusActive:
		return "active"
	case SupplyRequestStatusSuspended:
		return "suspended"
	case SupplyRequestStatusCancelled:
		return "cancelled"
	case SupplyRequestStatusCompleted:
		return "completed"
	case SupplyRequestStatusEnteredInError:
		return "entered-in-error"
	case SupplyRequestStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code SupplyRequestStatus) Display() string {
	switch code {
	case SupplyRequestStatusDraft:
		return "Draft"
	case SupplyRequestStatusActive:
		return "Active"
	case SupplyRequestStatusSuspended:
		return "Suspended"
	case SupplyRequestStatusCancelled:
		return "Cancelled"
	case SupplyRequestStatusCompleted:
		return "Completed"
	case SupplyRequestStatusEnteredInError:
		return "Entered in Error"
	case SupplyRequestStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code SupplyRequestStatus) Definition() string {
	switch code {
	case SupplyRequestStatusDraft:
		return "The request has been created but is not yet complete or ready for action."
	case SupplyRequestStatusActive:
		return "The request is ready to be acted upon."
	case SupplyRequestStatusSuspended:
		return "The authorization/request to act has been temporarily withdrawn but is expected to resume in the future."
	case SupplyRequestStatusCancelled:
		return "The authorization/request to act has been terminated prior to the full completion of the intended actions.  No further activity should occur."
	case SupplyRequestStatusCompleted:
		return "Activity against the request has been sufficiently completed to the satisfaction of the requester."
	case SupplyRequestStatusEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it.  (If real-world activity has occurred, the status should be \"cancelled\" rather than \"entered-in-error\".)."
	case SupplyRequestStatusUnknown:
		return "The authoring/source system does not know which of the status values currently applies for this observation. Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, but the authoring/source system does not know which."
	}
	return "<unknown>"
}
