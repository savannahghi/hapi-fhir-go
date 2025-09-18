
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ContributorType is documented here http://hl7.org/fhir/ValueSet/contributor-type
type ContributorType int

const (
	ContributorTypeAuthor ContributorType = iota
	ContributorTypeEditor
	ContributorTypeReviewer
	ContributorTypeEndorser
)

func (code ContributorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ContributorType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ContributorType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "author":
		*code = ContributorTypeAuthor
	case "editor":
		*code = ContributorTypeEditor
	case "reviewer":
		*code = ContributorTypeReviewer
	case "endorser":
		*code = ContributorTypeEndorser
	default:
		return fmt.Errorf("unknown ContributorType code `%s`", s)
	}
	return nil
}
func (code ContributorType) String() string {
	return code.Code()
}
func (code ContributorType) Code() string {
	switch code {
	case ContributorTypeAuthor:
		return "author"
	case ContributorTypeEditor:
		return "editor"
	case ContributorTypeReviewer:
		return "reviewer"
	case ContributorTypeEndorser:
		return "endorser"
	}
	return "<unknown>"
}
func (code ContributorType) Display() string {
	switch code {
	case ContributorTypeAuthor:
		return "Author"
	case ContributorTypeEditor:
		return "Editor"
	case ContributorTypeReviewer:
		return "Reviewer"
	case ContributorTypeEndorser:
		return "Endorser"
	}
	return "<unknown>"
}
func (code ContributorType) Definition() string {
	switch code {
	case ContributorTypeAuthor:
		return "An author of the content of the module."
	case ContributorTypeEditor:
		return "An editor of the content of the module."
	case ContributorTypeReviewer:
		return "A reviewer of the content of the module."
	case ContributorTypeEndorser:
		return "An endorser of the content of the module."
	}
	return "<unknown>"
}
