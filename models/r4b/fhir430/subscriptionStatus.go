
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// SubscriptionStatus is documented here http://hl7.org/fhir/ValueSet/subscription-status
type SubscriptionStatus int

const (
	SubscriptionStatusRequested SubscriptionStatus = iota
	SubscriptionStatusActive
	SubscriptionStatusError
	SubscriptionStatusOff
)

func (code SubscriptionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SubscriptionStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal SubscriptionStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "requested":
		*code = SubscriptionStatusRequested
	case "active":
		*code = SubscriptionStatusActive
	case "error":
		*code = SubscriptionStatusError
	case "off":
		*code = SubscriptionStatusOff
	default:
		return fmt.Errorf("unknown SubscriptionStatus code `%s`", s)
	}
	return nil
}
func (code SubscriptionStatus) String() string {
	return code.Code()
}
func (code SubscriptionStatus) Code() string {
	switch code {
	case SubscriptionStatusRequested:
		return "requested"
	case SubscriptionStatusActive:
		return "active"
	case SubscriptionStatusError:
		return "error"
	case SubscriptionStatusOff:
		return "off"
	}
	return "<unknown>"
}
func (code SubscriptionStatus) Display() string {
	switch code {
	case SubscriptionStatusRequested:
		return "Requested"
	case SubscriptionStatusActive:
		return "Active"
	case SubscriptionStatusError:
		return "Error"
	case SubscriptionStatusOff:
		return "Off"
	}
	return "<unknown>"
}
func (code SubscriptionStatus) Definition() string {
	switch code {
	case SubscriptionStatusRequested:
		return "The client has requested the subscription, and the server has not yet set it up."
	case SubscriptionStatusActive:
		return "The subscription is active."
	case SubscriptionStatusError:
		return "The server has an error executing the notification."
	case SubscriptionStatusOff:
		return "Too many errors have occurred or the subscription has expired."
	}
	return "<unknown>"
}
