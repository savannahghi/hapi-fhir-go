
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// FinancialResourceStatusCodes is documented here http://hl7.org/fhir/ValueSet/fm-status
type FinancialResourceStatusCodes int

const (
	FinancialResourceStatusCodesActive FinancialResourceStatusCodes = iota
	FinancialResourceStatusCodesCancelled
	FinancialResourceStatusCodesDraft
	FinancialResourceStatusCodesEnteredInError
)

func (code FinancialResourceStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *FinancialResourceStatusCodes) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal FinancialResourceStatusCodes code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "active":
		*code = FinancialResourceStatusCodesActive
	case "cancelled":
		*code = FinancialResourceStatusCodesCancelled
	case "draft":
		*code = FinancialResourceStatusCodesDraft
	case "entered-in-error":
		*code = FinancialResourceStatusCodesEnteredInError
	default:
		return fmt.Errorf("unknown FinancialResourceStatusCodes code `%s`", s)
	}
	return nil
}
func (code FinancialResourceStatusCodes) String() string {
	return code.Code()
}
func (code FinancialResourceStatusCodes) Code() string {
	switch code {
	case FinancialResourceStatusCodesActive:
		return "active"
	case FinancialResourceStatusCodesCancelled:
		return "cancelled"
	case FinancialResourceStatusCodesDraft:
		return "draft"
	case FinancialResourceStatusCodesEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code FinancialResourceStatusCodes) Display() string {
	switch code {
	case FinancialResourceStatusCodesActive:
		return "Active"
	case FinancialResourceStatusCodesCancelled:
		return "Cancelled"
	case FinancialResourceStatusCodesDraft:
		return "Draft"
	case FinancialResourceStatusCodesEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code FinancialResourceStatusCodes) Definition() string {
	switch code {
	case FinancialResourceStatusCodesActive:
		return "The instance is currently in-force."
	case FinancialResourceStatusCodesCancelled:
		return "The instance is withdrawn, rescinded or reversed."
	case FinancialResourceStatusCodesDraft:
		return "A new instance the contents of which is not complete."
	case FinancialResourceStatusCodesEnteredInError:
		return "The instance was entered in error."
	}
	return "<unknown>"
}
