
package fhir430

import (
	"encoding/json"
	"fmt"
)
// AuditEventAction is documented here http://hl7.org/fhir/ValueSet/audit-event-action
type AuditEventAction int

const (
	AuditEventActionC AuditEventAction = iota
	AuditEventActionR
	AuditEventActionU
	AuditEventActionD
	AuditEventActionE
)

func (code AuditEventAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AuditEventAction) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal AuditEventAction code `%s`", s)
	}
	switch s {
	case "C":
		*code = AuditEventActionC
	case "R":
		*code = AuditEventActionR
	case "U":
		*code = AuditEventActionU
	case "D":
		*code = AuditEventActionD
	case "E":
		*code = AuditEventActionE
	default:
		return fmt.Errorf("unknown AuditEventAction code `%s`", s)
	}
	return nil
}
func (code AuditEventAction) String() string {
	return code.Code()
}
func (code AuditEventAction) Code() string {
	switch code {
	case AuditEventActionC:
		return "C"
	case AuditEventActionR:
		return "R"
	case AuditEventActionU:
		return "U"
	case AuditEventActionD:
		return "D"
	case AuditEventActionE:
		return "E"
	}
	return "<unknown>"
}
func (code AuditEventAction) Display() string {
	switch code {
	case AuditEventActionC:
		return "Create"
	case AuditEventActionR:
		return "Read/View/Print"
	case AuditEventActionU:
		return "Update"
	case AuditEventActionD:
		return "Delete"
	case AuditEventActionE:
		return "Execute"
	}
	return "<unknown>"
}
func (code AuditEventAction) Definition() string {
	switch code {
	case AuditEventActionC:
		return "Create a new database object, such as placing an order."
	case AuditEventActionR:
		return "Display or print data, such as a doctor census."
	case AuditEventActionU:
		return "Update data, such as revise patient information."
	case AuditEventActionD:
		return "Delete items, such as a doctor master file record."
	case AuditEventActionE:
		return "Perform a system or application function such as log-on, program execution or use of an object's method, or perform a query/search operation."
	}
	return "<unknown>"
}
