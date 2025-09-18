
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ResponseType is documented here http://hl7.org/fhir/ValueSet/response-code
type ResponseType int

const (
	ResponseTypeOk ResponseType = iota
	ResponseTypeTransientError
	ResponseTypeFatalError
)

func (code ResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ResponseType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ResponseType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "ok":
		*code = ResponseTypeOk
	case "transient-error":
		*code = ResponseTypeTransientError
	case "fatal-error":
		*code = ResponseTypeFatalError
	default:
		return fmt.Errorf("unknown ResponseType code `%s`", s)
	}
	return nil
}
func (code ResponseType) String() string {
	return code.Code()
}
func (code ResponseType) Code() string {
	switch code {
	case ResponseTypeOk:
		return "ok"
	case ResponseTypeTransientError:
		return "transient-error"
	case ResponseTypeFatalError:
		return "fatal-error"
	}
	return "<unknown>"
}
func (code ResponseType) Display() string {
	switch code {
	case ResponseTypeOk:
		return "OK"
	case ResponseTypeTransientError:
		return "Transient Error"
	case ResponseTypeFatalError:
		return "Fatal Error"
	}
	return "<unknown>"
}
func (code ResponseType) Definition() string {
	switch code {
	case ResponseTypeOk:
		return "The message was accepted and processed without error."
	case ResponseTypeTransientError:
		return "Some internal unexpected error occurred - wait and try again. Note - this is usually used for things like database unavailable, which may be expected to resolve, though human intervention may be required."
	case ResponseTypeFatalError:
		return "The message was rejected because of a problem with the content. There is no point in re-sending without change. The response narrative SHALL describe the issue."
	}
	return "<unknown>"
}
