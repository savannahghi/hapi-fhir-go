
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// QuantityComparator is documented here http://hl7.org/fhir/ValueSet/quantity-comparator
type QuantityComparator int

const (
	QuantityComparatorLessThan QuantityComparator = iota
	QuantityComparatorLessOrEquals
	QuantityComparatorGreaterOrEquals
	QuantityComparatorGreaterThan
)

func (code QuantityComparator) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *QuantityComparator) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal QuantityComparator code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "<":
		*code = QuantityComparatorLessThan
	case "<=":
		*code = QuantityComparatorLessOrEquals
	case ">=":
		*code = QuantityComparatorGreaterOrEquals
	case ">":
		*code = QuantityComparatorGreaterThan
	default:
		return fmt.Errorf("unknown QuantityComparator code `%s`", s)
	}
	return nil
}
func (code QuantityComparator) String() string {
	return code.Code()
}
func (code QuantityComparator) Code() string {
	switch code {
	case QuantityComparatorLessThan:
		return "<"
	case QuantityComparatorLessOrEquals:
		return "<="
	case QuantityComparatorGreaterOrEquals:
		return ">="
	case QuantityComparatorGreaterThan:
		return ">"
	}
	return "<unknown>"
}
func (code QuantityComparator) Display() string {
	switch code {
	case QuantityComparatorLessThan:
		return "Less than"
	case QuantityComparatorLessOrEquals:
		return "Less or Equal to"
	case QuantityComparatorGreaterOrEquals:
		return "Greater or Equal to"
	case QuantityComparatorGreaterThan:
		return "Greater than"
	}
	return "<unknown>"
}
func (code QuantityComparator) Definition() string {
	switch code {
	case QuantityComparatorLessThan:
		return "The actual value is less than the given value."
	case QuantityComparatorLessOrEquals:
		return "The actual value is less than or equal to the given value."
	case QuantityComparatorGreaterOrEquals:
		return "The actual value is greater than or equal to the given value."
	case QuantityComparatorGreaterThan:
		return "The actual value is greater than the given value."
	}
	return "<unknown>"
}
