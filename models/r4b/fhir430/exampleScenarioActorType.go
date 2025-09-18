
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ExampleScenarioActorType is documented here http://hl7.org/fhir/ValueSet/examplescenario-actor-type
type ExampleScenarioActorType int

const (
	ExampleScenarioActorTypePerson ExampleScenarioActorType = iota
	ExampleScenarioActorTypeEntity
)

func (code ExampleScenarioActorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ExampleScenarioActorType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ExampleScenarioActorType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "person":
		*code = ExampleScenarioActorTypePerson
	case "entity":
		*code = ExampleScenarioActorTypeEntity
	default:
		return fmt.Errorf("unknown ExampleScenarioActorType code `%s`", s)
	}
	return nil
}
func (code ExampleScenarioActorType) String() string {
	return code.Code()
}
func (code ExampleScenarioActorType) Code() string {
	switch code {
	case ExampleScenarioActorTypePerson:
		return "person"
	case ExampleScenarioActorTypeEntity:
		return "entity"
	}
	return "<unknown>"
}
func (code ExampleScenarioActorType) Display() string {
	switch code {
	case ExampleScenarioActorTypePerson:
		return "Person"
	case ExampleScenarioActorTypeEntity:
		return "System"
	}
	return "<unknown>"
}
func (code ExampleScenarioActorType) Definition() string {
	switch code {
	case ExampleScenarioActorTypePerson:
		return "A person."
	case ExampleScenarioActorTypeEntity:
		return "A system."
	}
	return "<unknown>"
}
