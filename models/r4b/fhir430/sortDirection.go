
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// SortDirection is documented here http://hl7.org/fhir/ValueSet/sort-direction
type SortDirection int

const (
	SortDirectionAscending SortDirection = iota
	SortDirectionDescending
)

func (code SortDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SortDirection) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal SortDirection code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "ascending":
		*code = SortDirectionAscending
	case "descending":
		*code = SortDirectionDescending
	default:
		return fmt.Errorf("unknown SortDirection code `%s`", s)
	}
	return nil
}
func (code SortDirection) String() string {
	return code.Code()
}
func (code SortDirection) Code() string {
	switch code {
	case SortDirectionAscending:
		return "ascending"
	case SortDirectionDescending:
		return "descending"
	}
	return "<unknown>"
}
func (code SortDirection) Display() string {
	switch code {
	case SortDirectionAscending:
		return "Ascending"
	case SortDirectionDescending:
		return "Descending"
	}
	return "<unknown>"
}
func (code SortDirection) Definition() string {
	switch code {
	case SortDirectionAscending:
		return "Sort by the value ascending, so that lower values appear first."
	case SortDirectionDescending:
		return "Sort by the value descending, so that lower values appear last."
	}
	return "<unknown>"
}
