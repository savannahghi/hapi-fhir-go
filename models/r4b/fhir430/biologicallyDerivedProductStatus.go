
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// BiologicallyDerivedProductStatus is documented here http://hl7.org/fhir/ValueSet/product-status
type BiologicallyDerivedProductStatus int

const (
	BiologicallyDerivedProductStatusAvailable BiologicallyDerivedProductStatus = iota
	BiologicallyDerivedProductStatusUnavailable
)

func (code BiologicallyDerivedProductStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *BiologicallyDerivedProductStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal BiologicallyDerivedProductStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "available":
		*code = BiologicallyDerivedProductStatusAvailable
	case "unavailable":
		*code = BiologicallyDerivedProductStatusUnavailable
	default:
		return fmt.Errorf("unknown BiologicallyDerivedProductStatus code `%s`", s)
	}
	return nil
}
func (code BiologicallyDerivedProductStatus) String() string {
	return code.Code()
}
func (code BiologicallyDerivedProductStatus) Code() string {
	switch code {
	case BiologicallyDerivedProductStatusAvailable:
		return "available"
	case BiologicallyDerivedProductStatusUnavailable:
		return "unavailable"
	}
	return "<unknown>"
}
func (code BiologicallyDerivedProductStatus) Display() string {
	switch code {
	case BiologicallyDerivedProductStatusAvailable:
		return "Available"
	case BiologicallyDerivedProductStatusUnavailable:
		return "Unavailable"
	}
	return "<unknown>"
}
func (code BiologicallyDerivedProductStatus) Definition() string {
	switch code {
	case BiologicallyDerivedProductStatusAvailable:
		return "Product is currently available for use."
	case BiologicallyDerivedProductStatusUnavailable:
		return "Product is not currently available for use."
	}
	return "<unknown>"
}
