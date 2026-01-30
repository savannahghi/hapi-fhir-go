
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// StructureMapModelMode is documented here http://hl7.org/fhir/ValueSet/map-model-mode
type StructureMapModelMode int

const (
	StructureMapModelModeSource StructureMapModelMode = iota
	StructureMapModelModeQueried
	StructureMapModelModeTarget
	StructureMapModelModeProduced
)

func (code StructureMapModelMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *StructureMapModelMode) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal StructureMapModelMode code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "source":
		*code = StructureMapModelModeSource
	case "queried":
		*code = StructureMapModelModeQueried
	case "target":
		*code = StructureMapModelModeTarget
	case "produced":
		*code = StructureMapModelModeProduced
	default:
		return fmt.Errorf("unknown StructureMapModelMode code `%s`", s)
	}
	return nil
}
func (code StructureMapModelMode) String() string {
	return code.Code()
}
func (code StructureMapModelMode) Code() string {
	switch code {
	case StructureMapModelModeSource:
		return "source"
	case StructureMapModelModeQueried:
		return "queried"
	case StructureMapModelModeTarget:
		return "target"
	case StructureMapModelModeProduced:
		return "produced"
	}
	return "<unknown>"
}
func (code StructureMapModelMode) Display() string {
	switch code {
	case StructureMapModelModeSource:
		return "Source Structure Definition"
	case StructureMapModelModeQueried:
		return "Queried Structure Definition"
	case StructureMapModelModeTarget:
		return "Target Structure Definition"
	case StructureMapModelModeProduced:
		return "Produced Structure Definition"
	}
	return "<unknown>"
}
func (code StructureMapModelMode) Definition() string {
	switch code {
	case StructureMapModelModeSource:
		return "This structure describes an instance passed to the mapping engine that is used a source of data."
	case StructureMapModelModeQueried:
		return "This structure describes an instance that the mapping engine may ask for that is used a source of data."
	case StructureMapModelModeTarget:
		return "This structure describes an instance passed to the mapping engine that is used a target of data."
	case StructureMapModelModeProduced:
		return "This structure describes an instance that the mapping engine may ask to create that is used a target of data."
	}
	return "<unknown>"
}
