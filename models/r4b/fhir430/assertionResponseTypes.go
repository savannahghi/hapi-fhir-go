
package fhir430

import (
	"encoding/json"
	"fmt"
)
// AssertionResponseTypes is documented here http://hl7.org/fhir/ValueSet/assert-response-code-types
type AssertionResponseTypes int

const (
	AssertionResponseTypesOkay AssertionResponseTypes = iota
	AssertionResponseTypesCreated
	AssertionResponseTypesNoContent
	AssertionResponseTypesNotModified
	AssertionResponseTypesBad
	AssertionResponseTypesForbidden
	AssertionResponseTypesNotFound
	AssertionResponseTypesMethodNotAllowed
	AssertionResponseTypesConflict
	AssertionResponseTypesGone
	AssertionResponseTypesPreconditionFailed
	AssertionResponseTypesUnprocessable
)

func (code AssertionResponseTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AssertionResponseTypes) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal AssertionResponseTypes code `%s`", s)
	}
	switch s {
	case "okay":
		*code = AssertionResponseTypesOkay
	case "created":
		*code = AssertionResponseTypesCreated
	case "noContent":
		*code = AssertionResponseTypesNoContent
	case "notModified":
		*code = AssertionResponseTypesNotModified
	case "bad":
		*code = AssertionResponseTypesBad
	case "forbidden":
		*code = AssertionResponseTypesForbidden
	case "notFound":
		*code = AssertionResponseTypesNotFound
	case "methodNotAllowed":
		*code = AssertionResponseTypesMethodNotAllowed
	case "conflict":
		*code = AssertionResponseTypesConflict
	case "gone":
		*code = AssertionResponseTypesGone
	case "preconditionFailed":
		*code = AssertionResponseTypesPreconditionFailed
	case "unprocessable":
		*code = AssertionResponseTypesUnprocessable
	default:
		return fmt.Errorf("unknown AssertionResponseTypes code `%s`", s)
	}
	return nil
}
func (code AssertionResponseTypes) String() string {
	return code.Code()
}
func (code AssertionResponseTypes) Code() string {
	switch code {
	case AssertionResponseTypesOkay:
		return "okay"
	case AssertionResponseTypesCreated:
		return "created"
	case AssertionResponseTypesNoContent:
		return "noContent"
	case AssertionResponseTypesNotModified:
		return "notModified"
	case AssertionResponseTypesBad:
		return "bad"
	case AssertionResponseTypesForbidden:
		return "forbidden"
	case AssertionResponseTypesNotFound:
		return "notFound"
	case AssertionResponseTypesMethodNotAllowed:
		return "methodNotAllowed"
	case AssertionResponseTypesConflict:
		return "conflict"
	case AssertionResponseTypesGone:
		return "gone"
	case AssertionResponseTypesPreconditionFailed:
		return "preconditionFailed"
	case AssertionResponseTypesUnprocessable:
		return "unprocessable"
	}
	return "<unknown>"
}
func (code AssertionResponseTypes) Display() string {
	switch code {
	case AssertionResponseTypesOkay:
		return "okay"
	case AssertionResponseTypesCreated:
		return "created"
	case AssertionResponseTypesNoContent:
		return "noContent"
	case AssertionResponseTypesNotModified:
		return "notModified"
	case AssertionResponseTypesBad:
		return "bad"
	case AssertionResponseTypesForbidden:
		return "forbidden"
	case AssertionResponseTypesNotFound:
		return "notFound"
	case AssertionResponseTypesMethodNotAllowed:
		return "methodNotAllowed"
	case AssertionResponseTypesConflict:
		return "conflict"
	case AssertionResponseTypesGone:
		return "gone"
	case AssertionResponseTypesPreconditionFailed:
		return "preconditionFailed"
	case AssertionResponseTypesUnprocessable:
		return "unprocessable"
	}
	return "<unknown>"
}
func (code AssertionResponseTypes) Definition() string {
	switch code {
	case AssertionResponseTypesOkay:
		return "Response code is 200."
	case AssertionResponseTypesCreated:
		return "Response code is 201."
	case AssertionResponseTypesNoContent:
		return "Response code is 204."
	case AssertionResponseTypesNotModified:
		return "Response code is 304."
	case AssertionResponseTypesBad:
		return "Response code is 400."
	case AssertionResponseTypesForbidden:
		return "Response code is 403."
	case AssertionResponseTypesNotFound:
		return "Response code is 404."
	case AssertionResponseTypesMethodNotAllowed:
		return "Response code is 405."
	case AssertionResponseTypesConflict:
		return "Response code is 409."
	case AssertionResponseTypesGone:
		return "Response code is 410."
	case AssertionResponseTypesPreconditionFailed:
		return "Response code is 412."
	case AssertionResponseTypesUnprocessable:
		return "Response code is 422."
	}
	return "<unknown>"
}
