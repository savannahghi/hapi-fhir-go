
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// NoteType is documented here http://hl7.org/fhir/ValueSet/note-type
type NoteType int

const (
	NoteTypeDisplay NoteType = iota
	NoteTypePrint
	NoteTypePrintoper
)

func (code NoteType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *NoteType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal NoteType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "display":
		*code = NoteTypeDisplay
	case "print":
		*code = NoteTypePrint
	case "printoper":
		*code = NoteTypePrintoper
	default:
		return fmt.Errorf("unknown NoteType code `%s`", s)
	}
	return nil
}
func (code NoteType) String() string {
	return code.Code()
}
func (code NoteType) Code() string {
	switch code {
	case NoteTypeDisplay:
		return "display"
	case NoteTypePrint:
		return "print"
	case NoteTypePrintoper:
		return "printoper"
	}
	return "<unknown>"
}
func (code NoteType) Display() string {
	switch code {
	case NoteTypeDisplay:
		return "Display"
	case NoteTypePrint:
		return "Print (Form)"
	case NoteTypePrintoper:
		return "Print (Operator)"
	}
	return "<unknown>"
}
func (code NoteType) Definition() string {
	switch code {
	case NoteTypeDisplay:
		return "Display the note."
	case NoteTypePrint:
		return "Print the note on the form."
	case NoteTypePrintoper:
		return "Print the note for the operator."
	}
	return "<unknown>"
}
