
package fhir430

import (
	"encoding/json"
	"fmt"
)
// HTTPVerb is documented here http://hl7.org/fhir/ValueSet/http-verb
type HTTPVerb int

const (
	HTTPVerbGET HTTPVerb = iota
	HTTPVerbHEAD
	HTTPVerbPOST
	HTTPVerbPUT
	HTTPVerbDELETE
	HTTPVerbPATCH
)

func (code HTTPVerb) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *HTTPVerb) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal HTTPVerb code `%s`", s)
	}
	switch s {
	case "GET":
		*code = HTTPVerbGET
	case "HEAD":
		*code = HTTPVerbHEAD
	case "POST":
		*code = HTTPVerbPOST
	case "PUT":
		*code = HTTPVerbPUT
	case "DELETE":
		*code = HTTPVerbDELETE
	case "PATCH":
		*code = HTTPVerbPATCH
	default:
		return fmt.Errorf("unknown HTTPVerb code `%s`", s)
	}
	return nil
}
func (code HTTPVerb) String() string {
	return code.Code()
}
func (code HTTPVerb) Code() string {
	switch code {
	case HTTPVerbGET:
		return "GET"
	case HTTPVerbHEAD:
		return "HEAD"
	case HTTPVerbPOST:
		return "POST"
	case HTTPVerbPUT:
		return "PUT"
	case HTTPVerbDELETE:
		return "DELETE"
	case HTTPVerbPATCH:
		return "PATCH"
	}
	return "<unknown>"
}
func (code HTTPVerb) Display() string {
	switch code {
	case HTTPVerbGET:
		return "GET"
	case HTTPVerbHEAD:
		return "HEAD"
	case HTTPVerbPOST:
		return "POST"
	case HTTPVerbPUT:
		return "PUT"
	case HTTPVerbDELETE:
		return "DELETE"
	case HTTPVerbPATCH:
		return "PATCH"
	}
	return "<unknown>"
}
func (code HTTPVerb) Definition() string {
	switch code {
	case HTTPVerbGET:
		return "HTTP GET Command."
	case HTTPVerbHEAD:
		return "HTTP HEAD Command."
	case HTTPVerbPOST:
		return "HTTP POST Command."
	case HTTPVerbPUT:
		return "HTTP PUT Command."
	case HTTPVerbDELETE:
		return "HTTP DELETE Command."
	case HTTPVerbPATCH:
		return "HTTP PATCH Command."
	}
	return "<unknown>"
}
