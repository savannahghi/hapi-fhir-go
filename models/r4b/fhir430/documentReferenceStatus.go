
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// DocumentReferenceStatus is documented here http://hl7.org/fhir/ValueSet/document-reference-status
type DocumentReferenceStatus int

const (
	DocumentReferenceStatusCurrent DocumentReferenceStatus = iota
	DocumentReferenceStatusSuperseded
	DocumentReferenceStatusEnteredInError
)

func (code DocumentReferenceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DocumentReferenceStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal DocumentReferenceStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "current":
		*code = DocumentReferenceStatusCurrent
	case "superseded":
		*code = DocumentReferenceStatusSuperseded
	case "entered-in-error":
		*code = DocumentReferenceStatusEnteredInError
	default:
		return fmt.Errorf("unknown DocumentReferenceStatus code `%s`", s)
	}
	return nil
}
func (code DocumentReferenceStatus) String() string {
	return code.Code()
}
func (code DocumentReferenceStatus) Code() string {
	switch code {
	case DocumentReferenceStatusCurrent:
		return "current"
	case DocumentReferenceStatusSuperseded:
		return "superseded"
	case DocumentReferenceStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code DocumentReferenceStatus) Display() string {
	switch code {
	case DocumentReferenceStatusCurrent:
		return "Current"
	case DocumentReferenceStatusSuperseded:
		return "Superseded"
	case DocumentReferenceStatusEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code DocumentReferenceStatus) Definition() string {
	switch code {
	case DocumentReferenceStatusCurrent:
		return "This is the current reference for this document."
	case DocumentReferenceStatusSuperseded:
		return "This reference has been superseded by another reference."
	case DocumentReferenceStatusEnteredInError:
		return "This reference was created in error."
	}
	return "<unknown>"
}
