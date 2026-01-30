
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// TestReportActionResult is documented here http://hl7.org/fhir/ValueSet/report-action-result-codes
type TestReportActionResult int

const (
	TestReportActionResultPass TestReportActionResult = iota
	TestReportActionResultSkip
	TestReportActionResultFail
	TestReportActionResultWarning
	TestReportActionResultError
)

func (code TestReportActionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *TestReportActionResult) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal TestReportActionResult code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "pass":
		*code = TestReportActionResultPass
	case "skip":
		*code = TestReportActionResultSkip
	case "fail":
		*code = TestReportActionResultFail
	case "warning":
		*code = TestReportActionResultWarning
	case "error":
		*code = TestReportActionResultError
	default:
		return fmt.Errorf("unknown TestReportActionResult code `%s`", s)
	}
	return nil
}
func (code TestReportActionResult) String() string {
	return code.Code()
}
func (code TestReportActionResult) Code() string {
	switch code {
	case TestReportActionResultPass:
		return "pass"
	case TestReportActionResultSkip:
		return "skip"
	case TestReportActionResultFail:
		return "fail"
	case TestReportActionResultWarning:
		return "warning"
	case TestReportActionResultError:
		return "error"
	}
	return "<unknown>"
}
func (code TestReportActionResult) Display() string {
	switch code {
	case TestReportActionResultPass:
		return "Pass"
	case TestReportActionResultSkip:
		return "Skip"
	case TestReportActionResultFail:
		return "Fail"
	case TestReportActionResultWarning:
		return "Warning"
	case TestReportActionResultError:
		return "Error"
	}
	return "<unknown>"
}
func (code TestReportActionResult) Definition() string {
	switch code {
	case TestReportActionResultPass:
		return "The action was successful."
	case TestReportActionResultSkip:
		return "The action was skipped."
	case TestReportActionResultFail:
		return "The action failed."
	case TestReportActionResultWarning:
		return "The action passed but with warnings."
	case TestReportActionResultError:
		return "The action encountered a fatal error and the engine was unable to process."
	}
	return "<unknown>"
}
