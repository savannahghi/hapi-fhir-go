
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// AddressType is documented here http://hl7.org/fhir/ValueSet/address-type
type AddressType int

const (
	AddressTypePostal AddressType = iota
	AddressTypePhysical
	AddressTypeBoth
)

func (code AddressType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AddressType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal AddressType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "postal":
		*code = AddressTypePostal
	case "physical":
		*code = AddressTypePhysical
	case "both":
		*code = AddressTypeBoth
	default:
		return fmt.Errorf("unknown AddressType code `%s`", s)
	}
	return nil
}
func (code AddressType) String() string {
	return code.Code()
}
func (code AddressType) Code() string {
	switch code {
	case AddressTypePostal:
		return "postal"
	case AddressTypePhysical:
		return "physical"
	case AddressTypeBoth:
		return "both"
	}
	return "<unknown>"
}
func (code AddressType) Display() string {
	switch code {
	case AddressTypePostal:
		return "Postal"
	case AddressTypePhysical:
		return "Physical"
	case AddressTypeBoth:
		return "Postal & Physical"
	}
	return "<unknown>"
}
func (code AddressType) Definition() string {
	switch code {
	case AddressTypePostal:
		return "Mailing addresses - PO Boxes and care-of addresses."
	case AddressTypePhysical:
		return "A physical address that can be visited."
	case AddressTypeBoth:
		return "An address that is both physical and postal."
	}
	return "<unknown>"
}
