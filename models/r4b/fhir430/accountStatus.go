package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)

// AccountStatus is documented here http://hl7.org/fhir/ValueSet/account-status
type AccountStatus int

const (
	AccountStatusActive AccountStatus = iota
	AccountStatusInactive
	AccountStatusEnteredInError
	AccountStatusOnHold
	AccountStatusUnknown
)

func (code AccountStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AccountStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal AccountStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "active":
		*code = AccountStatusActive
	case "inactive":
		*code = AccountStatusInactive
	case "entered-in-error":
		*code = AccountStatusEnteredInError
	case "on-hold":
		*code = AccountStatusOnHold
	case "unknown":
		*code = AccountStatusUnknown
	default:
		return fmt.Errorf("unknown AccountStatus code `%s`", s)
	}
	return nil
}
func (code AccountStatus) String() string {
	return code.Code()
}
func (code AccountStatus) Code() string {
	switch code {
	case AccountStatusActive:
		return "active"
	case AccountStatusInactive:
		return "inactive"
	case AccountStatusEnteredInError:
		return "entered-in-error"
	case AccountStatusOnHold:
		return "on-hold"
	case AccountStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code AccountStatus) Display() string {
	switch code {
	case AccountStatusActive:
		return "Active"
	case AccountStatusInactive:
		return "Inactive"
	case AccountStatusEnteredInError:
		return "Entered in error"
	case AccountStatusOnHold:
		return "On Hold"
	case AccountStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code AccountStatus) Definition() string {
	switch code {
	case AccountStatusActive:
		return "This account is active and may be used."
	case AccountStatusInactive:
		return "This account is inactive and should not be used to track financial information."
	case AccountStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	case AccountStatusOnHold:
		return "This account is on hold."
	case AccountStatusUnknown:
		return "The account status is unknown."
	}
	return "<unknown>"
}
