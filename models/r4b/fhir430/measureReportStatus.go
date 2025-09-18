
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// MeasureReportStatus is documented here http://hl7.org/fhir/ValueSet/measure-report-status
type MeasureReportStatus int

const (
	MeasureReportStatusComplete MeasureReportStatus = iota
	MeasureReportStatusPending
	MeasureReportStatusError
)

func (code MeasureReportStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MeasureReportStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal MeasureReportStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "complete":
		*code = MeasureReportStatusComplete
	case "pending":
		*code = MeasureReportStatusPending
	case "error":
		*code = MeasureReportStatusError
	default:
		return fmt.Errorf("unknown MeasureReportStatus code `%s`", s)
	}
	return nil
}
func (code MeasureReportStatus) String() string {
	return code.Code()
}
func (code MeasureReportStatus) Code() string {
	switch code {
	case MeasureReportStatusComplete:
		return "complete"
	case MeasureReportStatusPending:
		return "pending"
	case MeasureReportStatusError:
		return "error"
	}
	return "<unknown>"
}
func (code MeasureReportStatus) Display() string {
	switch code {
	case MeasureReportStatusComplete:
		return "Complete"
	case MeasureReportStatusPending:
		return "Pending"
	case MeasureReportStatusError:
		return "Error"
	}
	return "<unknown>"
}
func (code MeasureReportStatus) Definition() string {
	switch code {
	case MeasureReportStatusComplete:
		return "The report is complete and ready for use."
	case MeasureReportStatusPending:
		return "The report is currently being generated."
	case MeasureReportStatusError:
		return "An error occurred attempting to generate the report."
	}
	return "<unknown>"
}
