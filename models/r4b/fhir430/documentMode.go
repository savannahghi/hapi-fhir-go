
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// DocumentMode is documented here http://hl7.org/fhir/ValueSet/document-mode
type DocumentMode int

const (
	DocumentModeProducer DocumentMode = iota
	DocumentModeConsumer
)

func (code DocumentMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DocumentMode) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal DocumentMode code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "producer":
		*code = DocumentModeProducer
	case "consumer":
		*code = DocumentModeConsumer
	default:
		return fmt.Errorf("unknown DocumentMode code `%s`", s)
	}
	return nil
}
func (code DocumentMode) String() string {
	return code.Code()
}
func (code DocumentMode) Code() string {
	switch code {
	case DocumentModeProducer:
		return "producer"
	case DocumentModeConsumer:
		return "consumer"
	}
	return "<unknown>"
}
func (code DocumentMode) Display() string {
	switch code {
	case DocumentModeProducer:
		return "Producer"
	case DocumentModeConsumer:
		return "Consumer"
	}
	return "<unknown>"
}
func (code DocumentMode) Definition() string {
	switch code {
	case DocumentModeProducer:
		return "The application produces documents of the specified type."
	case DocumentModeConsumer:
		return "The application consumes documents of the specified type."
	}
	return "<unknown>"
}
