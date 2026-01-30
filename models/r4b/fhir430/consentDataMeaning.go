
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ConsentDataMeaning is documented here http://hl7.org/fhir/ValueSet/consent-data-meaning
type ConsentDataMeaning int

const (
	ConsentDataMeaningInstance ConsentDataMeaning = iota
	ConsentDataMeaningRelated
	ConsentDataMeaningDependents
	ConsentDataMeaningAuthoredby
)

func (code ConsentDataMeaning) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ConsentDataMeaning) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ConsentDataMeaning code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "instance":
		*code = ConsentDataMeaningInstance
	case "related":
		*code = ConsentDataMeaningRelated
	case "dependents":
		*code = ConsentDataMeaningDependents
	case "authoredby":
		*code = ConsentDataMeaningAuthoredby
	default:
		return fmt.Errorf("unknown ConsentDataMeaning code `%s`", s)
	}
	return nil
}
func (code ConsentDataMeaning) String() string {
	return code.Code()
}
func (code ConsentDataMeaning) Code() string {
	switch code {
	case ConsentDataMeaningInstance:
		return "instance"
	case ConsentDataMeaningRelated:
		return "related"
	case ConsentDataMeaningDependents:
		return "dependents"
	case ConsentDataMeaningAuthoredby:
		return "authoredby"
	}
	return "<unknown>"
}
func (code ConsentDataMeaning) Display() string {
	switch code {
	case ConsentDataMeaningInstance:
		return "Instance"
	case ConsentDataMeaningRelated:
		return "Related"
	case ConsentDataMeaningDependents:
		return "Dependents"
	case ConsentDataMeaningAuthoredby:
		return "AuthoredBy"
	}
	return "<unknown>"
}
func (code ConsentDataMeaning) Definition() string {
	switch code {
	case ConsentDataMeaningInstance:
		return "The consent applies directly to the instance of the resource."
	case ConsentDataMeaningRelated:
		return "The consent applies directly to the instance of the resource and instances it refers to."
	case ConsentDataMeaningDependents:
		return "The consent applies directly to the instance of the resource and instances that refer to it."
	case ConsentDataMeaningAuthoredby:
		return "The consent applies to instances of resources that are authored by."
	}
	return "<unknown>"
}
