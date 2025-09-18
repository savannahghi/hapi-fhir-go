
package fhir430

import "encoding/json"
// TestScript is documented here http://hl7.org/fhir/StructureDefinition/TestScript
// A structured set of tests against a FHIR server or client implementation to determine compliance against the FHIR specification.
type TestScript struct {
	ID                *string                 `json:"ID,omitempty"`
	Meta              *Meta                   `json:"meta,omitempty"`
	ImplicitRules     *string                 `json:"implicitRules,omitempty"`
	Language          *string                 `json:"language,omitempty"`
	Text              *Narrative              `json:"text,omitempty"`
	Contained         []json.RawMessage       `json:"contained,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Url               string                  `json:"url"`
	Identifier        *Identifier             `json:"identifier,omitempty"`
	Version           *string                 `json:"version,omitempty"`
	Name              string                  `json:"name"`
	Title             *string                 `json:"title,omitempty"`
	Status            PublicationStatus       `json:"status"`
	Experimental      *bool                   `json:"experimental,omitempty"`
	Date              *string                 `json:"date,omitempty"`
	Publisher         *string                 `json:"publisher,omitempty"`
	Contact           []ContactDetail         `json:"contact,omitempty"`
	Description       *string                 `json:"description,omitempty"`
	UseContext        []UsageContext          `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept       `json:"jurisdiction,omitempty"`
	Purpose           *string                 `json:"purpose,omitempty"`
	Origin            []TestScriptOrigin      `json:"origin,omitempty"`
	Destination       []TestScriptDestination `json:"destination,omitempty"`
	Metadata          *TestScriptMetadata     `json:"metadata,omitempty"`
	Fixture           []TestScriptFixture     `json:"fixture,omitempty"`
	Profile           []Reference             `json:"profile,omitempty"`
	Variable          []TestScriptVariable    `json:"variable,omitempty"`
	Setup             *TestScriptSetup        `json:"setup,omitempty"`
	Test              []TestScriptTest        `json:"test,omitempty"`
	Teardown          *TestScriptTeardown     `json:"teardown,omitempty"`
}

// An abstract server used in operations within this test script in the origin element.
// The purpose of this element is to define the profile of an origin element used elsewhere in the script.  Test engines could then use the origin-profile mapping to offer a filtered list of test systems that can serve as the sender for the interaction.
type TestScriptOrigin struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Index             int         `json:"index"`
	Profile           Coding      `json:"profile"`
}

// An abstract server used in operations within this test script in the destination element.
// The purpose of this element is to define the profile of a destination element used elsewhere in the script.  Test engines could then use the destination-profile mapping to offer a filtered list of test systems that can serve as the receiver for the interaction.
type TestScriptDestination struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Index             int         `json:"index"`
	Profile           Coding      `json:"profile"`
}

// The required capability must exist and are assumed to function correctly on the FHIR server being tested.
type TestScriptMetadata struct {
	ID                *string                        `json:"ID,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Link              []TestScriptMetadataLink       `json:"link,omitempty"`
	Capability        []TestScriptMetadataCapability `json:"capability"`
}

// A link to the FHIR specification that this test is covering.
type TestScriptMetadataLink struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Url               string      `json:"url"`
	Description       *string     `json:"description,omitempty"`
}

// Capabilities that must exist and are assumed to function correctly on the FHIR server being tested.
// When the metadata capabilities section is defined at TestScript.metadata or at TestScript.setup.metadata, and the server's conformance statement does not contain the elements defined in the minimal conformance statement, then all the tests in the TestScript are skipped.  When the metadata capabilities section is defined at TestScript.test.metadata and the server's conformance statement does not contain the elements defined in the minimal conformance statement, then only that test is skipped.  The "metadata.capabilities.required" and "metadata.capabilities.validated" elements only indicate whether the capabilities are the primary focus of the test script or not.  They do not impact the skipping logic.  Capabilities whose "metadata.capabilities.validated" flag is true are the primary focus of the test script.
type TestScriptMetadataCapability struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Required          bool        `json:"required"`
	Validated         bool        `json:"validated"`
	Description       *string     `json:"description,omitempty"`
	Origin            []int       `json:"origin,omitempty"`
	Destination       *int        `json:"destination,omitempty"`
	Link              []string    `json:"link,omitempty"`
	Capabilities      string      `json:"capabilities"`
}

// Fixture in the test script - by reference (uri). All fixtures are required for the test script to execute.
type TestScriptFixture struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Autocreate        bool        `json:"autocreate"`
	Autodelete        bool        `json:"autodelete"`
	Resource          *Reference  `json:"resource,omitempty"`
}

// Variable is set based either on element value in response body or on header field value in the response headers.
// Variables would be set based either on XPath/JSONPath expressions against fixtures (static and response), or headerField evaluations against response headers. If variable evaluates to nodelist or anything other than a primitive value, then test engine would report an error.  Variables would be used to perform clean replacements in "operation.params", "operation.requestHeader.value", and "operation.url" element values during operation calls and in "assert.value" during assertion evaluations. This limits the places that test engines would need to look for placeholders "${}".  Variables are scoped to the whole script. They are NOT evaluated at declaration. They are evaluated by test engine when used for substitutions in "operation.params", "operation.requestHeader.value", and "operation.url" element values during operation calls and in "assert.value" during assertion evaluations.  See example testscript-search.xml.
type TestScriptVariable struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Name              string      `json:"name"`
	DefaultValue      *string     `json:"defaultValue,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Expression        *string     `json:"expression,omitempty"`
	HeaderField       *string     `json:"headerField,omitempty"`
	Hint              *string     `json:"hint,omitempty"`
	Path              *string     `json:"path,omitempty"`
	SourceId          *string     `json:"sourceId,omitempty"`
}

// A series of required setup operations before tests are executed.
type TestScriptSetup struct {
	ID                *string                 `json:"ID,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Action            []TestScriptSetupAction `json:"action"`
}

// Action would contain either an operation or an assertion.
// An action should contain either an operation or an assertion but not both.  It can contain any number of variables.
type TestScriptSetupAction struct {
	ID                *string                         `json:"ID,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Operation         *TestScriptSetupActionOperation `json:"operation,omitempty"`
	Assert            *TestScriptSetupActionAssert    `json:"assert,omitempty"`
}

// The operation to perform.
type TestScriptSetupActionOperation struct {
	ID                *string                                       `json:"ID,omitempty"`
	Extension         []Extension                                   `json:"extension,omitempty"`
	ModifierExtension []Extension                                   `json:"modifierExtension,omitempty"`
	Type              *Coding                                       `json:"type,omitempty"`
	Resource          *string                                       `json:"resource,omitempty"`
	Label             *string                                       `json:"label,omitempty"`
	Description       *string                                       `json:"description,omitempty"`
	Accept            *string                                       `json:"accept,omitempty"`
	ContentType       *string                                       `json:"contentType,omitempty"`
	Destination       *int                                          `json:"destination,omitempty"`
	EncodeRequestUrl  bool                                          `json:"encodeRequestUrl"`
	Method            *TestScriptRequestMethodCode                  `json:"method,omitempty"`
	Origin            *int                                          `json:"origin,omitempty"`
	Params            *string                                       `json:"params,omitempty"`
	RequestHeader     []TestScriptSetupActionOperationRequestHeader `json:"requestHeader,omitempty"`
	RequestId         *string                                       `json:"requestId,omitempty"`
	ResponseId        *string                                       `json:"responseId,omitempty"`
	SourceId          *string                                       `json:"sourceId,omitempty"`
	TargetId          *string                                       `json:"targetId,omitempty"`
	Url               *string                                       `json:"url,omitempty"`
}

// Header elements would be used to set HTTP headers.
// This gives control to test-script writers to set headers explicitly based on test requirements.  It will allow for testing using:  - "If-Modified-Since" and "If-None-Match" headers.  See http://build.fhir.org/http.html#2.1.0.5.1 - "If-Match" header.  See http://build.fhir.org/http.html#2.1.0.11 - Conditional Create using "If-None-Exist".  See http://build.fhir.org/http.html#2.1.0.13.1 - Invalid "Content-Type" header for negative testing. - etc.
type TestScriptSetupActionOperationRequestHeader struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Field             string      `json:"field"`
	Value             string      `json:"value"`
}

// Evaluates the results of previous operations to determine if the server under test behaves appropriately.
// In order to evaluate an assertion, the request, response, and results of the most recently executed operation must always be maintained by the test engine.
type TestScriptSetupActionAssert struct {
	ID                        *string                      `json:"ID,omitempty"`
	Extension                 []Extension                  `json:"extension,omitempty"`
	ModifierExtension         []Extension                  `json:"modifierExtension,omitempty"`
	Label                     *string                      `json:"label,omitempty"`
	Description               *string                      `json:"description,omitempty"`
	Direction                 *AssertionDirectionType      `json:"direction,omitempty"`
	CompareToSourceId         *string                      `json:"compareToSourceId,omitempty"`
	CompareToSourceExpression *string                      `json:"compareToSourceExpression,omitempty"`
	CompareToSourcePath       *string                      `json:"compareToSourcePath,omitempty"`
	ContentType               *string                      `json:"contentType,omitempty"`
	Expression                *string                      `json:"expression,omitempty"`
	HeaderField               *string                      `json:"headerField,omitempty"`
	MinimumId                 *string                      `json:"minimumId,omitempty"`
	NavigationLinks           *bool                        `json:"navigationLinks,omitempty"`
	Operator                  *AssertionOperatorType       `json:"operator,omitempty"`
	Path                      *string                      `json:"path,omitempty"`
	RequestMethod             *TestScriptRequestMethodCode `json:"requestMethod,omitempty"`
	RequestURL                *string                      `json:"requestURL,omitempty"`
	Resource                  *string                      `json:"resource,omitempty"`
	Response                  *AssertionResponseTypes      `json:"response,omitempty"`
	ResponseCode              *string                      `json:"responseCode,omitempty"`
	SourceId                  *string                      `json:"sourceId,omitempty"`
	ValidateProfileId         *string                      `json:"validateProfileId,omitempty"`
	Value                     *string                      `json:"value,omitempty"`
	WarningOnly               bool                         `json:"warningOnly"`
}

// A test in this script.
type TestScriptTest struct {
	ID                *string                `json:"ID,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Name              *string                `json:"name,omitempty"`
	Description       *string                `json:"description,omitempty"`
	Action            []TestScriptTestAction `json:"action"`
}

// Action would contain either an operation or an assertion.
// An action should contain either an operation or an assertion but not both.  It can contain any number of variables.
type TestScriptTestAction struct {
	ID                *string                         `json:"ID,omitempty"`
	Extension         []Extension                     `json:"extension,omitempty"`
	ModifierExtension []Extension                     `json:"modifierExtension,omitempty"`
	Operation         *TestScriptSetupActionOperation `json:"operation,omitempty"`
	Assert            *TestScriptSetupActionAssert    `json:"assert,omitempty"`
}

// A series of operations required to clean up after all the tests are executed (successfully or otherwise).
type TestScriptTeardown struct {
	ID                *string                    `json:"ID,omitempty"`
	Extension         []Extension                `json:"extension,omitempty"`
	ModifierExtension []Extension                `json:"modifierExtension,omitempty"`
	Action            []TestScriptTeardownAction `json:"action"`
}

// The teardown action will only contain an operation.
// An action should contain either an operation or an assertion but not both.  It can contain any number of variables.
type TestScriptTeardownAction struct {
	ID                *string                        `json:"ID,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Operation         TestScriptSetupActionOperation `json:"operation,omitempty"`
}

// This function returns resource reference information
func (r TestScript) ResourceRef() (string, *string) {
	return "TestScript", r.ID
}

// This function returns resource reference information
func (r TestScript) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherTestScript TestScript

// MarshalJSON marshals the given TestScript as JSON into a byte slice
func (r TestScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherTestScript
		ResourceType string `json:"resourceType"`
	}{
		OtherTestScript: OtherTestScript(r),
		ResourceType:    "TestScript",
	})
}

// UnmarshalTestScript unmarshals a TestScript.
func UnmarshalTestScript(b []byte) (TestScript, error) {
	var testScript TestScript
	if err := json.Unmarshal(b, &testScript); err != nil {
		return testScript, err
	}
	return testScript, nil
}
