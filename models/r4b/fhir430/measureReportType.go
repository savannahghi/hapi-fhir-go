
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// MeasureReportType is documented here http://hl7.org/fhir/ValueSet/measure-report-type
type MeasureReportType int

const (
	MeasureReportTypeIndividual MeasureReportType = iota
	MeasureReportTypeSubjectList
	MeasureReportTypeSummary
	MeasureReportTypeDataCollection
)

func (code MeasureReportType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *MeasureReportType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal MeasureReportType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "individual":
		*code = MeasureReportTypeIndividual
	case "subject-list":
		*code = MeasureReportTypeSubjectList
	case "summary":
		*code = MeasureReportTypeSummary
	case "data-collection":
		*code = MeasureReportTypeDataCollection
	default:
		return fmt.Errorf("unknown MeasureReportType code `%s`", s)
	}
	return nil
}
func (code MeasureReportType) String() string {
	return code.Code()
}
func (code MeasureReportType) Code() string {
	switch code {
	case MeasureReportTypeIndividual:
		return "individual"
	case MeasureReportTypeSubjectList:
		return "subject-list"
	case MeasureReportTypeSummary:
		return "summary"
	case MeasureReportTypeDataCollection:
		return "data-collection"
	}
	return "<unknown>"
}
func (code MeasureReportType) Display() string {
	switch code {
	case MeasureReportTypeIndividual:
		return "Individual"
	case MeasureReportTypeSubjectList:
		return "Subject List"
	case MeasureReportTypeSummary:
		return "Summary"
	case MeasureReportTypeDataCollection:
		return "Data Collection"
	}
	return "<unknown>"
}
func (code MeasureReportType) Definition() string {
	switch code {
	case MeasureReportTypeIndividual:
		return "An individual report that provides information on the performance for a given measure with respect to a single subject."
	case MeasureReportTypeSubjectList:
		return "A subject list report that includes a listing of subjects that satisfied each population criteria in the measure."
	case MeasureReportTypeSummary:
		return "A summary report that returns the number of members in each population criteria for the measure."
	case MeasureReportTypeDataCollection:
		return "A data collection report that contains data-of-interest for the measure."
	}
	return "<unknown>"
}
