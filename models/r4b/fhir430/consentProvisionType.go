
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ConsentProvisionType is documented here http://hl7.org/fhir/ValueSet/consent-provision-type
type ConsentProvisionType int

const (
	ConsentProvisionTypeDeny ConsentProvisionType = iota
	ConsentProvisionTypePermit
)

func (code ConsentProvisionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ConsentProvisionType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ConsentProvisionType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "deny":
		*code = ConsentProvisionTypeDeny
	case "permit":
		*code = ConsentProvisionTypePermit
	default:
		return fmt.Errorf("unknown ConsentProvisionType code `%s`", s)
	}
	return nil
}
func (code ConsentProvisionType) String() string {
	return code.Code()
}
func (code ConsentProvisionType) Code() string {
	switch code {
	case ConsentProvisionTypeDeny:
		return "deny"
	case ConsentProvisionTypePermit:
		return "permit"
	}
	return "<unknown>"
}
func (code ConsentProvisionType) Display() string {
	switch code {
	case ConsentProvisionTypeDeny:
		return "Opt Out"
	case ConsentProvisionTypePermit:
		return "Opt In"
	}
	return "<unknown>"
}
func (code ConsentProvisionType) Definition() string {
	switch code {
	case ConsentProvisionTypeDeny:
		return "Consent is denied for actions meeting these rules."
	case ConsentProvisionTypePermit:
		return "Consent is provided for actions meeting these rules."
	}
	return "<unknown>"
}
