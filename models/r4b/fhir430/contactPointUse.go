
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ContactPointUse is documented here http://hl7.org/fhir/ValueSet/contact-point-use
type ContactPointUse int

const (
	ContactPointUseHome ContactPointUse = iota
	ContactPointUseWork
	ContactPointUseTemp
	ContactPointUseOld
	ContactPointUseMobile
)

func (code ContactPointUse) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ContactPointUse) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ContactPointUse code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "home":
		*code = ContactPointUseHome
	case "work":
		*code = ContactPointUseWork
	case "temp":
		*code = ContactPointUseTemp
	case "old":
		*code = ContactPointUseOld
	case "mobile":
		*code = ContactPointUseMobile
	default:
		return fmt.Errorf("unknown ContactPointUse code `%s`", s)
	}
	return nil
}
func (code ContactPointUse) String() string {
	return code.Code()
}
func (code ContactPointUse) Code() string {
	switch code {
	case ContactPointUseHome:
		return "home"
	case ContactPointUseWork:
		return "work"
	case ContactPointUseTemp:
		return "temp"
	case ContactPointUseOld:
		return "old"
	case ContactPointUseMobile:
		return "mobile"
	}
	return "<unknown>"
}
func (code ContactPointUse) Display() string {
	switch code {
	case ContactPointUseHome:
		return "Home"
	case ContactPointUseWork:
		return "Work"
	case ContactPointUseTemp:
		return "Temp"
	case ContactPointUseOld:
		return "Old"
	case ContactPointUseMobile:
		return "Mobile"
	}
	return "<unknown>"
}
func (code ContactPointUse) Definition() string {
	switch code {
	case ContactPointUseHome:
		return "A communication contact point at a home; attempted contacts for business purposes might intrude privacy and chances are one will contact family or other household members instead of the person one wishes to call. Typically used with urgent cases, or if no other contacts are available."
	case ContactPointUseWork:
		return "An office contact point. First choice for business related contacts during business hours."
	case ContactPointUseTemp:
		return "A temporary contact point. The period can provide more detailed information."
	case ContactPointUseOld:
		return "This contact point is no longer in use (or was never correct, but retained for records)."
	case ContactPointUseMobile:
		return "A telecommunication device that moves and stays with its owner. May have characteristics of all other use codes, suitable for urgent matters, not the first choice for routine business."
	}
	return "<unknown>"
}
