
package fhir430

import "encoding/json"
// DetectedIssue is documented here http://hl7.org/fhir/StructureDefinition/DetectedIssue
// Indicates an actual or potential clinical issue with or between one or more active or proposed clinical actions for a patient; e.g. Drug-drug interaction, Ineffective treatment frequency, Procedure-condition conflict, etc.
type DetectedIssue struct {
	ID                 *string                   `json:"ID,omitempty"`
	Meta               *Meta                     `json:"meta,omitempty"`
	ImplicitRules      *string                   `json:"implicitRules,omitempty"`
	Language           *string                   `json:"language,omitempty"`
	Text               *Narrative                `json:"text,omitempty"`
	Contained          []json.RawMessage         `json:"contained,omitempty"`
	Extension          []Extension               `json:"extension,omitempty"`
	ModifierExtension  []Extension               `json:"modifierExtension,omitempty"`
	Identifier         []Identifier              `json:"identifier,omitempty"`
	Status             ObservationStatus         `json:"status"`
	Code               *CodeableConcept          `json:"code,omitempty"`
	Severity           *DetectedIssueSeverity    `json:"severity,omitempty"`
	Patient            *Reference                `json:"patient,omitempty"`
	IdentifiedDateTime *string                   `json:"identifiedDateTime,omitempty"`
	IdentifiedPeriod   *Period                   `json:"identifiedPeriod,omitempty"`
	Author             *Reference                `json:"author,omitempty"`
	Implicated         []Reference               `json:"implicated,omitempty"`
	Evidence           []DetectedIssueEvidence   `json:"evidence,omitempty"`
	Detail             *string                   `json:"detail,omitempty"`
	Reference          *string                   `json:"reference,omitempty"`
	Mitigation         []DetectedIssueMitigation `json:"mitigation,omitempty"`
}

// Supporting evidence or manifestations that provide the basis for identifying the detected issue such as a GuidanceResponse or MeasureReport.
type DetectedIssueEvidence struct {
	ID                *string           `json:"ID,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Code              []CodeableConcept `json:"code,omitempty"`
	Detail            []Reference       `json:"detail,omitempty"`
}

// Indicates an action that has been taken or is committed to reduce or eliminate the likelihood of the risk identified by the detected issue from manifesting.  Can also reflect an observation of known mitigating factors that may reduce/eliminate the need for any action.
type DetectedIssueMitigation struct {
	ID                *string         `json:"ID,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Action            CodeableConcept `json:"action"`
	Date              *string         `json:"date,omitempty"`
	Author            *Reference      `json:"author,omitempty"`
}

// This function returns resource reference information
func (r DetectedIssue) ResourceRef() (string, *string) {
	return "DetectedIssue", r.ID
}

// This function returns resource reference information
func (r DetectedIssue) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherDetectedIssue DetectedIssue

// MarshalJSON marshals the given DetectedIssue as JSON into a byte slice
func (r DetectedIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDetectedIssue
		ResourceType string `json:"resourceType"`
	}{
		OtherDetectedIssue: OtherDetectedIssue(r),
		ResourceType:       "DetectedIssue",
	})
}

// UnmarshalDetectedIssue unmarshals a DetectedIssue.
func UnmarshalDetectedIssue(b []byte) (DetectedIssue, error) {
	var detectedIssue DetectedIssue
	if err := json.Unmarshal(b, &detectedIssue); err != nil {
		return detectedIssue, err
	}
	return detectedIssue, nil
}
