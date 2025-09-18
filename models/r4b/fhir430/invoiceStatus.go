
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// InvoiceStatus is documented here http://hl7.org/fhir/ValueSet/invoice-status
type InvoiceStatus int

const (
	InvoiceStatusDraft InvoiceStatus = iota
	InvoiceStatusIssued
	InvoiceStatusBalanced
	InvoiceStatusCancelled
	InvoiceStatusEnteredInError
)

func (code InvoiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *InvoiceStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal InvoiceStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "draft":
		*code = InvoiceStatusDraft
	case "issued":
		*code = InvoiceStatusIssued
	case "balanced":
		*code = InvoiceStatusBalanced
	case "cancelled":
		*code = InvoiceStatusCancelled
	case "entered-in-error":
		*code = InvoiceStatusEnteredInError
	default:
		return fmt.Errorf("unknown InvoiceStatus code `%s`", s)
	}
	return nil
}
func (code InvoiceStatus) String() string {
	return code.Code()
}
func (code InvoiceStatus) Code() string {
	switch code {
	case InvoiceStatusDraft:
		return "draft"
	case InvoiceStatusIssued:
		return "issued"
	case InvoiceStatusBalanced:
		return "balanced"
	case InvoiceStatusCancelled:
		return "cancelled"
	case InvoiceStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code InvoiceStatus) Display() string {
	switch code {
	case InvoiceStatusDraft:
		return "draft"
	case InvoiceStatusIssued:
		return "issued"
	case InvoiceStatusBalanced:
		return "balanced"
	case InvoiceStatusCancelled:
		return "cancelled"
	case InvoiceStatusEnteredInError:
		return "entered in error"
	}
	return "<unknown>"
}
func (code InvoiceStatus) Definition() string {
	switch code {
	case InvoiceStatusDraft:
		return "the invoice has been prepared but not yet finalized."
	case InvoiceStatusIssued:
		return "the invoice has been finalized and sent to the recipient."
	case InvoiceStatusBalanced:
		return "the invoice has been balaced / completely paid."
	case InvoiceStatusCancelled:
		return "the invoice was cancelled."
	case InvoiceStatusEnteredInError:
		return "the invoice was determined as entered in error before it was issued."
	}
	return "<unknown>"
}
