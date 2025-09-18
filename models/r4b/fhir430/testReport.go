
package fhir430

import "encoding/json"
// TestReport is documented here http://hl7.org/fhir/StructureDefinition/TestReport
// A summary of information based on the results of executing a TestScript.
type TestReport struct {
	ID                *string                 `json:"ID,omitempty"`
	Meta              *Meta                   `json:"meta,omitempty"`
	ImplicitRules     *string                 `json:"implicitRules,omitempty"`
	Language          *string                 `json:"language,omitempty"`
	Text              *Narrative              `json:"text,omitempty"`
	Contained         []json.RawMessage       `json:"contained,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Identifier        *Identifier             `json:"identifier,omitempty"`
	Name              *string                 `json:"name,omitempty"`
	Status            TestReportStatus        `json:"status"`
	TestScript        Reference               `json:"testScript"`
	Result            TestReportResult        `json:"result"`
	Score             *json.Number            `json:"score,omitempty"`
	Tester            *string                 `json:"tester,omitempty"`
	Issued            *string                 `json:"issued,omitempty"`
	Participant       []TestReportParticipant `json:"participant,omitempty"`
	Setup             *TestReportSetup        `json:"setup,omitempty"`
	Test              []TestReportTest        `json:"test,omitempty"`
	Teardown          *TestReportTeardown     `json:"teardown,omitempty"`
}

// A participant in the test execution, either the execution engine, a client, or a server.
type TestReportParticipant struct {
	ID                *string                   `json:"ID,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	Type              TestReportParticipantType `json:"type"`
	Uri               string                    `json:"uri"`
	Display           *string                   `json:"display,omitempty"`
}

// The results of the series of required setup operations before the tests were executed.
type TestReportSetup struct {
	ID                *string                 `json:"ID,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Action            []TestReportSetupAction `json:"action"`
}

// Action would contain either an operation or an assertion.
// An action should contain either an operation or an assertion but not both.  It can contain any number of variables.
type TestReportSetupAction struct {
	ID                *string                         `json:"ID,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Operation         *TestReportSetupActionOperation `json:"operation,omitempty"`
	Assert            *TestReportSetupActionAssert    `json:"assert,omitempty"`
}

// The operation performed.
type TestReportSetupActionOperation struct {
	ID                *string                `json:"ID,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Result            TestReportActionResult `json:"result"`
	Message           *string                `json:"message,omitempty"`
	Detail            *string                `json:"detail,omitempty"`
}

// The results of the assertion performed on the previous operations.
type TestReportSetupActionAssert struct {
	ID                *string                `json:"ID,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Result            TestReportActionResult `json:"result"`
	Message           *string                `json:"message,omitempty"`
	Detail            *string                `json:"detail,omitempty"`
}

// A test executed from the test script.
type TestReportTest struct {
	ID                *string                `json:"ID,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Name              *string                `json:"name,omitempty"`
	Description       *string                `json:"description,omitempty"`
	Action            []TestReportTestAction `json:"action"`
}

// Action would contain either an operation or an assertion.
// An action should contain either an operation or an assertion but not both.  It can contain any number of variables.
type TestReportTestAction struct {
	ID                *string                         `json:"ID,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Operation         *TestReportSetupActionOperation `json:"operation,omitempty"`
	Assert            *TestReportSetupActionAssert    `json:"assert,omitempty"`
}

// The results of the series of operations required to clean up after all the tests were executed (successfully or otherwise).
type TestReportTeardown struct {
	ID                *string                    `json:"ID,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Action            []TestReportTeardownAction `json:"action"`
}

// The teardown action will only contain an operation.
// An action should contain either an operation or an assertion but not both.  It can contain any number of variables.
type TestReportTeardownAction struct {
	ID                *string                        `json:"ID,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Operation         TestReportSetupActionOperation `json:"operation,omitempty"`
}

// This function returns resource reference information
func (r TestReport) ResourceRef() (string, *string) {
	return "TestReport", r.ID
}

// This function returns resource reference information
func (r TestReport) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherTestReport TestReport

// MarshalJSON marshals the given TestReport as JSON into a byte slice
func (r TestReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTestReport
		ResourceType string `json:"resourceType"`
	}{
		OtherTestReport: OtherTestReport(r),
		ResourceType:    "TestReport",
	})
}

// UnmarshalTestReport unmarshals a TestReport.
func UnmarshalTestReport(b []byte) (TestReport, error) {
	var testReport TestReport
	if err := json.Unmarshal(b, &testReport); err != nil {
		return testReport, err
	}
	return testReport, nil
}
