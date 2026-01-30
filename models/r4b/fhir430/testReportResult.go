
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// TestReportResult is documented here http://hl7.org/fhir/ValueSet/report-result-codes
type TestReportResult int

const (
	TestReportResultPass TestReportResult = iota
	TestReportResultFail
	TestReportResultPending
)

func (code TestReportResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *TestReportResult) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal TestReportResult code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "pass":
		*code = TestReportResultPass
	case "fail":
		*code = TestReportResultFail
	case "pending":
		*code = TestReportResultPending
	default:
		return fmt.Errorf("unknown TestReportResult code `%s`", s)
	}
	return nil
}
func (code TestReportResult) String() string {
	return code.Code()
}
func (code TestReportResult) Code() string {
	switch code {
	case TestReportResultPass:
		return "pass"
	case TestReportResultFail:
		return "fail"
	case TestReportResultPending:
		return "pending"
	}
	return "<unknown>"
}
func (code TestReportResult) Display() string {
	switch code {
	case TestReportResultPass:
		return "Pass"
	case TestReportResultFail:
		return "Fail"
	case TestReportResultPending:
		return "Pending"
	}
	return "<unknown>"
}
func (code TestReportResult) Definition() string {
	switch code {
	case TestReportResultPass:
		return "All test operations successfully passed all asserts."
	case TestReportResultFail:
		return "One or more test operations failed one or more asserts."
	case TestReportResultPending:
		return "One or more test operations is pending execution completion."
	}
	return "<unknown>"
}
