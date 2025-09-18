
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// StructureMapSourceListMode is documented here http://hl7.org/fhir/ValueSet/map-source-list-mode
type StructureMapSourceListMode int

const (
	StructureMapSourceListModeFirst StructureMapSourceListMode = iota
	StructureMapSourceListModeNot_first
	StructureMapSourceListModeLast
	StructureMapSourceListModeNot_last
	StructureMapSourceListModeOnly_one
)

func (code StructureMapSourceListMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *StructureMapSourceListMode) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal StructureMapSourceListMode code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "first":
		*code = StructureMapSourceListModeFirst
	case "not_first":
		*code = StructureMapSourceListModeNot_first
	case "last":
		*code = StructureMapSourceListModeLast
	case "not_last":
		*code = StructureMapSourceListModeNot_last
	case "only_one":
		*code = StructureMapSourceListModeOnly_one
	default:
		return fmt.Errorf("unknown StructureMapSourceListMode code `%s`", s)
	}
	return nil
}
func (code StructureMapSourceListMode) String() string {
	return code.Code()
}
func (code StructureMapSourceListMode) Code() string {
	switch code {
	case StructureMapSourceListModeFirst:
		return "first"
	case StructureMapSourceListModeNot_first:
		return "not_first"
	case StructureMapSourceListModeLast:
		return "last"
	case StructureMapSourceListModeNot_last:
		return "not_last"
	case StructureMapSourceListModeOnly_one:
		return "only_one"
	}
	return "<unknown>"
}
func (code StructureMapSourceListMode) Display() string {
	switch code {
	case StructureMapSourceListModeFirst:
		return "First"
	case StructureMapSourceListModeNot_first:
		return "All but the first"
	case StructureMapSourceListModeLast:
		return "Last"
	case StructureMapSourceListModeNot_last:
		return "All but the last"
	case StructureMapSourceListModeOnly_one:
		return "Enforce only one"
	}
	return "<unknown>"
}
func (code StructureMapSourceListMode) Definition() string {
	switch code {
	case StructureMapSourceListModeFirst:
		return "Only process this rule for the first in the list."
	case StructureMapSourceListModeNot_first:
		return "Process this rule for all but the first."
	case StructureMapSourceListModeLast:
		return "Only process this rule for the last in the list."
	case StructureMapSourceListModeNot_last:
		return "Process this rule for all but the last."
	case StructureMapSourceListModeOnly_one:
		return "Only process this rule is there is only item."
	}
	return "<unknown>"
}
