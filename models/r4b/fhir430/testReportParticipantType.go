
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// TestReportParticipantType is documented here http://hl7.org/fhir/ValueSet/report-participant-type
type TestReportParticipantType int

const (
	TestReportParticipantTypeTestEngine TestReportParticipantType = iota
	TestReportParticipantTypeClient
	TestReportParticipantTypeServer
)

func (code TestReportParticipantType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *TestReportParticipantType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal TestReportParticipantType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "test-engine":
		*code = TestReportParticipantTypeTestEngine
	case "client":
		*code = TestReportParticipantTypeClient
	case "server":
		*code = TestReportParticipantTypeServer
	default:
		return fmt.Errorf("unknown TestReportParticipantType code `%s`", s)
	}
	return nil
}
func (code TestReportParticipantType) String() string {
	return code.Code()
}
func (code TestReportParticipantType) Code() string {
	switch code {
	case TestReportParticipantTypeTestEngine:
		return "test-engine"
	case TestReportParticipantTypeClient:
		return "client"
	case TestReportParticipantTypeServer:
		return "server"
	}
	return "<unknown>"
}
func (code TestReportParticipantType) Display() string {
	switch code {
	case TestReportParticipantTypeTestEngine:
		return "Test Engine"
	case TestReportParticipantTypeClient:
		return "Client"
	case TestReportParticipantTypeServer:
		return "Server"
	}
	return "<unknown>"
}
func (code TestReportParticipantType) Definition() string {
	switch code {
	case TestReportParticipantTypeTestEngine:
		return "The test execution engine."
	case TestReportParticipantTypeClient:
		return "A FHIR Client."
	case TestReportParticipantTypeServer:
		return "A FHIR Server."
	}
	return "<unknown>"
}
