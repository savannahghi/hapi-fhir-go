package fhir430

import "encoding/json"

// DeviceUseStatement is documented here http://hl7.org/fhir/StructureDefinition/DeviceUseStatement
// A record of a device being used by a patient where the record is the result of a report from the patient or another clinician.
type DeviceUseStatement struct {
	ID                *string                  `json:"id,omitempty"`
	Meta              *Meta                    `json:"meta,omitempty"`
	ImplicitRules     *string                  `json:"implicitRules,omitempty"`
	Language          *string                  `json:"language,omitempty"`
	Text              *Narrative               `json:"text,omitempty"`
	Contained         []json.RawMessage        `json:"contained,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Identifier        []Identifier             `json:"identifier,omitempty"`
	BasedOn           []Reference              `json:"basedOn,omitempty"`
	Status            DeviceUseStatementStatus `json:"status"`
	Subject           Reference                `json:"subject"`
	DerivedFrom       []Reference              `json:"derivedFrom,omitempty"`
	TimingTiming      *Timing                  `json:"timingTiming,omitempty"`
	TimingPeriod      *Period                  `json:"timingPeriod,omitempty"`
	TimingDateTime    *string                  `json:"timingDateTime,omitempty"`
	RecordedOn        *string                  `json:"recordedOn,omitempty"`
	Source            *Reference               `json:"source,omitempty"`
	Device            Reference                `json:"device"`
	ReasonCode        []CodeableConcept        `json:"reasonCode,omitempty"`
	ReasonReference   []Reference              `json:"reasonReference,omitempty"`
	BodySite          *CodeableConcept         `json:"bodySite,omitempty"`
	Note              []Annotation             `json:"note,omitempty"`
}

// This function returns resource reference information
func (r DeviceUseStatement) ResourceRef() (string, *string) {
	return "DeviceUseStatement", r.ID
}

// This function returns resource reference information
func (r DeviceUseStatement) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherDeviceUseStatement DeviceUseStatement

// MarshalJSON marshals the given DeviceUseStatement as JSON into a byte slice
func (r DeviceUseStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceUseStatement
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceUseStatement: OtherDeviceUseStatement(r),
		ResourceType:            "DeviceUseStatement",
	})
}

// UnmarshalDeviceUseStatement unmarshals a DeviceUseStatement.
func UnmarshalDeviceUseStatement(b []byte) (DeviceUseStatement, error) {
	var deviceUseStatement DeviceUseStatement
	if err := json.Unmarshal(b, &deviceUseStatement); err != nil {
		return deviceUseStatement, err
	}
	return deviceUseStatement, nil
}
