
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// SupplyDeliveryStatus is documented here http://hl7.org/fhir/ValueSet/supplydelivery-status
type SupplyDeliveryStatus int

const (
	SupplyDeliveryStatusInProgress SupplyDeliveryStatus = iota
	SupplyDeliveryStatusCompleted
	SupplyDeliveryStatusAbandoned
	SupplyDeliveryStatusEnteredInError
)

func (code SupplyDeliveryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SupplyDeliveryStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal SupplyDeliveryStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "in-progress":
		*code = SupplyDeliveryStatusInProgress
	case "completed":
		*code = SupplyDeliveryStatusCompleted
	case "abandoned":
		*code = SupplyDeliveryStatusAbandoned
	case "entered-in-error":
		*code = SupplyDeliveryStatusEnteredInError
	default:
		return fmt.Errorf("unknown SupplyDeliveryStatus code `%s`", s)
	}
	return nil
}
func (code SupplyDeliveryStatus) String() string {
	return code.Code()
}
func (code SupplyDeliveryStatus) Code() string {
	switch code {
	case SupplyDeliveryStatusInProgress:
		return "in-progress"
	case SupplyDeliveryStatusCompleted:
		return "completed"
	case SupplyDeliveryStatusAbandoned:
		return "abandoned"
	case SupplyDeliveryStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code SupplyDeliveryStatus) Display() string {
	switch code {
	case SupplyDeliveryStatusInProgress:
		return "In Progress"
	case SupplyDeliveryStatusCompleted:
		return "Delivered"
	case SupplyDeliveryStatusAbandoned:
		return "Abandoned"
	case SupplyDeliveryStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}
func (code SupplyDeliveryStatus) Definition() string {
	switch code {
	case SupplyDeliveryStatusInProgress:
		return "Supply has been requested, but not delivered."
	case SupplyDeliveryStatusCompleted:
		return "Supply has been delivered (\"completed\")."
	case SupplyDeliveryStatusAbandoned:
		return "Delivery was not completed."
	case SupplyDeliveryStatusEnteredInError:
		return "This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be \"abandoned\" rather than \"entered-in-error\".)."
	}
	return "<unknown>"
}
