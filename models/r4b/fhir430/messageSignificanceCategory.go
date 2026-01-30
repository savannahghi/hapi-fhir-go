
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// MessageSignificanceCategory is documented here http://hl7.org/fhir/ValueSet/message-significance-category
type MessageSignificanceCategory int

const (
	MessageSignificanceCategoryConsequence MessageSignificanceCategory = iota
	MessageSignificanceCategoryCurrency
	MessageSignificanceCategoryNotification
)

func (code MessageSignificanceCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MessageSignificanceCategory) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal MessageSignificanceCategory code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "consequence":
		*code = MessageSignificanceCategoryConsequence
	case "currency":
		*code = MessageSignificanceCategoryCurrency
	case "notification":
		*code = MessageSignificanceCategoryNotification
	default:
		return fmt.Errorf("unknown MessageSignificanceCategory code `%s`", s)
	}
	return nil
}
func (code MessageSignificanceCategory) String() string {
	return code.Code()
}
func (code MessageSignificanceCategory) Code() string {
	switch code {
	case MessageSignificanceCategoryConsequence:
		return "consequence"
	case MessageSignificanceCategoryCurrency:
		return "currency"
	case MessageSignificanceCategoryNotification:
		return "notification"
	}
	return "<unknown>"
}
func (code MessageSignificanceCategory) Display() string {
	switch code {
	case MessageSignificanceCategoryConsequence:
		return "Consequence"
	case MessageSignificanceCategoryCurrency:
		return "Currency"
	case MessageSignificanceCategoryNotification:
		return "Notification"
	}
	return "<unknown>"
}
func (code MessageSignificanceCategory) Definition() string {
	switch code {
	case MessageSignificanceCategoryConsequence:
		return "The message represents/requests a change that should not be processed more than once; e.g., making a booking for an appointment."
	case MessageSignificanceCategoryCurrency:
		return "The message represents a response to query for current information. Retrospective processing is wrong and/or wasteful."
	case MessageSignificanceCategoryNotification:
		return "The content is not necessarily intended to be current, and it can be reprocessed, though there may be version issues created by processing old notifications."
	}
	return "<unknown>"
}
