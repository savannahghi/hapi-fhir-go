
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// StructureMapContextType is documented here http://hl7.org/fhir/ValueSet/map-context-type
type StructureMapContextType int

const (
	StructureMapContextTypeType StructureMapContextType = iota
	StructureMapContextTypeVariable
)

func (code StructureMapContextType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *StructureMapContextType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal StructureMapContextType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "type":
		*code = StructureMapContextTypeType
	case "variable":
		*code = StructureMapContextTypeVariable
	default:
		return fmt.Errorf("unknown StructureMapContextType code `%s`", s)
	}
	return nil
}
func (code StructureMapContextType) String() string {
	return code.Code()
}
func (code StructureMapContextType) Code() string {
	switch code {
	case StructureMapContextTypeType:
		return "type"
	case StructureMapContextTypeVariable:
		return "variable"
	}
	return "<unknown>"
}
func (code StructureMapContextType) Display() string {
	switch code {
	case StructureMapContextTypeType:
		return "Type"
	case StructureMapContextTypeVariable:
		return "Variable"
	}
	return "<unknown>"
}
func (code StructureMapContextType) Definition() string {
	switch code {
	case StructureMapContextTypeType:
		return "The context specifies a type."
	case StructureMapContextTypeVariable:
		return "The context specifies a variable."
	}
	return "<unknown>"
}
