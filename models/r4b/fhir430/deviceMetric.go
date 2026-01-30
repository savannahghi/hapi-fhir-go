package fhir430

import "encoding/json"

// DeviceMetric is documented here http://hl7.org/fhir/StructureDefinition/DeviceMetric
// Describes a measurement, calculation or setting capability of a medical device.
type DeviceMetric struct {
	ID                *string                        `json:"id,omitempty"`
	Meta              *Meta                          `json:"meta,omitempty"`
	ImplicitRules     *string                        `json:"implicitRules,omitempty"`
	Language          *string                        `json:"language,omitempty"`
	Text              *Narrative                     `json:"text,omitempty"`
	Contained         []json.RawMessage              `json:"contained,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                   `json:"identifier,omitempty"`
	Type              CodeableConcept                `json:"type"`
	Unit              *CodeableConcept               `json:"unit,omitempty"`
	Source            *Reference                     `json:"source,omitempty"`
	Parent            *Reference                     `json:"parent,omitempty"`
	OperationalStatus *DeviceMetricOperationalStatus `json:"operationalStatus,omitempty"`
	Color             *DeviceMetricColor             `json:"color,omitempty"`
	Category          DeviceMetricCategory           `json:"category"`
	MeasurementPeriod *Timing                        `json:"measurementPeriod,omitempty"`
	Calibration       []DeviceMetricCalibration      `json:"calibration,omitempty"`
}

// Describes the calibrations that have been performed or that are required to be performed.
type DeviceMetricCalibration struct {
	ID                *string                       `json:"id,omitempty"`
	Extension         []Extension                   `json:"extension,omitempty"`
	ModifierExtension []Extension                   `json:"modifierExtension,omitempty"`
	Type              *DeviceMetricCalibrationType  `json:"type,omitempty"`
	State             *DeviceMetricCalibrationState `json:"state,omitempty"`
	Time              *string                       `json:"time,omitempty"`
}

// This function returns resource reference information
func (r DeviceMetric) ResourceRef() (string, *string) {
	return "DeviceMetric", r.ID
}

// This function returns resource reference information
func (r DeviceMetric) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherDeviceMetric DeviceMetric

// MarshalJSON marshals the given DeviceMetric as JSON into a byte slice
func (r DeviceMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherDeviceMetric
		ResourceType string `json:"resourceType"`
	}{
		OtherDeviceMetric: OtherDeviceMetric(r),
		ResourceType:      "DeviceMetric",
	})
}

// UnmarshalDeviceMetric unmarshals a DeviceMetric.
func UnmarshalDeviceMetric(b []byte) (DeviceMetric, error) {
	var deviceMetric DeviceMetric
	if err := json.Unmarshal(b, &deviceMetric); err != nil {
		return deviceMetric, err
	}
	return deviceMetric, nil
}
