
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// SearchEntryMode is documented here http://hl7.org/fhir/ValueSet/search-entry-mode
type SearchEntryMode int

const (
	SearchEntryModeMatch SearchEntryMode = iota
	SearchEntryModeInclude
	SearchEntryModeOutcome
)

func (code SearchEntryMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *SearchEntryMode) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal SearchEntryMode code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "match":
		*code = SearchEntryModeMatch
	case "include":
		*code = SearchEntryModeInclude
	case "outcome":
		*code = SearchEntryModeOutcome
	default:
		return fmt.Errorf("unknown SearchEntryMode code `%s`", s)
	}
	return nil
}
func (code SearchEntryMode) String() string {
	return code.Code()
}
func (code SearchEntryMode) Code() string {
	switch code {
	case SearchEntryModeMatch:
		return "match"
	case SearchEntryModeInclude:
		return "include"
	case SearchEntryModeOutcome:
		return "outcome"
	}
	return "<unknown>"
}
func (code SearchEntryMode) Display() string {
	switch code {
	case SearchEntryModeMatch:
		return "Match"
	case SearchEntryModeInclude:
		return "Include"
	case SearchEntryModeOutcome:
		return "Outcome"
	}
	return "<unknown>"
}
func (code SearchEntryMode) Definition() string {
	switch code {
	case SearchEntryModeMatch:
		return "This resource matched the search specification."
	case SearchEntryModeInclude:
		return "This resource is returned because it is referred to from another resource in the search set."
	case SearchEntryModeOutcome:
		return "An OperationOutcome that provides additional information about the processing of a search."
	}
	return "<unknown>"
}
