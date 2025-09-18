
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// RequestPriority is documented here http://hl7.org/fhir/ValueSet/request-priority
type RequestPriority int

const (
	RequestPriorityRoutine RequestPriority = iota
	RequestPriorityUrgent
	RequestPriorityAsap
	RequestPriorityStat
)

func (code RequestPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *RequestPriority) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal RequestPriority code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "routine":
		*code = RequestPriorityRoutine
	case "urgent":
		*code = RequestPriorityUrgent
	case "asap":
		*code = RequestPriorityAsap
	case "stat":
		*code = RequestPriorityStat
	default:
		return fmt.Errorf("unknown RequestPriority code `%s`", s)
	}
	return nil
}
func (code RequestPriority) String() string {
	return code.Code()
}
func (code RequestPriority) Code() string {
	switch code {
	case RequestPriorityRoutine:
		return "routine"
	case RequestPriorityUrgent:
		return "urgent"
	case RequestPriorityAsap:
		return "asap"
	case RequestPriorityStat:
		return "stat"
	}
	return "<unknown>"
}
func (code RequestPriority) Display() string {
	switch code {
	case RequestPriorityRoutine:
		return "Routine"
	case RequestPriorityUrgent:
		return "Urgent"
	case RequestPriorityAsap:
		return "ASAP"
	case RequestPriorityStat:
		return "STAT"
	}
	return "<unknown>"
}
func (code RequestPriority) Definition() string {
	switch code {
	case RequestPriorityRoutine:
		return "The request has normal priority."
	case RequestPriorityUrgent:
		return "The request should be actioned promptly - higher priority than routine."
	case RequestPriorityAsap:
		return "The request should be actioned as soon as possible - higher priority than urgent."
	case RequestPriorityStat:
		return "The request should be actioned immediately - highest possible priority.  E.g. an emergency."
	}
	return "<unknown>"
}
