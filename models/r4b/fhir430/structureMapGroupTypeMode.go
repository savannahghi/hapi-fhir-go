
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// StructureMapGroupTypeMode is documented here http://hl7.org/fhir/ValueSet/map-group-type-mode
type StructureMapGroupTypeMode int

const (
	StructureMapGroupTypeModeNone StructureMapGroupTypeMode = iota
	StructureMapGroupTypeModeTypes
	StructureMapGroupTypeModeTypeAndTypes
)

func (code StructureMapGroupTypeMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *StructureMapGroupTypeMode) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal StructureMapGroupTypeMode code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "none":
		*code = StructureMapGroupTypeModeNone
	case "types":
		*code = StructureMapGroupTypeModeTypes
	case "type-and-types":
		*code = StructureMapGroupTypeModeTypeAndTypes
	default:
		return fmt.Errorf("unknown StructureMapGroupTypeMode code `%s`", s)
	}
	return nil
}
func (code StructureMapGroupTypeMode) String() string {
	return code.Code()
}
func (code StructureMapGroupTypeMode) Code() string {
	switch code {
	case StructureMapGroupTypeModeNone:
		return "none"
	case StructureMapGroupTypeModeTypes:
		return "types"
	case StructureMapGroupTypeModeTypeAndTypes:
		return "type-and-types"
	}
	return "<unknown>"
}
func (code StructureMapGroupTypeMode) Display() string {
	switch code {
	case StructureMapGroupTypeModeNone:
		return "Not a Default"
	case StructureMapGroupTypeModeTypes:
		return "Default for Type Combination"
	case StructureMapGroupTypeModeTypeAndTypes:
		return "Default for type + combination"
	}
	return "<unknown>"
}
func (code StructureMapGroupTypeMode) Definition() string {
	switch code {
	case StructureMapGroupTypeModeNone:
		return "This group is not a default group for the types."
	case StructureMapGroupTypeModeTypes:
		return "This group is a default mapping group for the specified types and for the primary source type."
	case StructureMapGroupTypeModeTypeAndTypes:
		return "This group is a default mapping group for the specified types."
	}
	return "<unknown>"
}
