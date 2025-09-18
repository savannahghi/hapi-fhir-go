
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// DeviceMetricCalibrationState is documented here http://hl7.org/fhir/ValueSet/metric-calibration-state
type DeviceMetricCalibrationState int

const (
	DeviceMetricCalibrationStateNotCalibrated DeviceMetricCalibrationState = iota
	DeviceMetricCalibrationStateCalibrationRequired
	DeviceMetricCalibrationStateCalibrated
	DeviceMetricCalibrationStateUnspecified
)

func (code DeviceMetricCalibrationState) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DeviceMetricCalibrationState) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal DeviceMetricCalibrationState code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "not-calibrated":
		*code = DeviceMetricCalibrationStateNotCalibrated
	case "calibration-required":
		*code = DeviceMetricCalibrationStateCalibrationRequired
	case "calibrated":
		*code = DeviceMetricCalibrationStateCalibrated
	case "unspecified":
		*code = DeviceMetricCalibrationStateUnspecified
	default:
		return fmt.Errorf("unknown DeviceMetricCalibrationState code `%s`", s)
	}
	return nil
}
func (code DeviceMetricCalibrationState) String() string {
	return code.Code()
}
func (code DeviceMetricCalibrationState) Code() string {
	switch code {
	case DeviceMetricCalibrationStateNotCalibrated:
		return "not-calibrated"
	case DeviceMetricCalibrationStateCalibrationRequired:
		return "calibration-required"
	case DeviceMetricCalibrationStateCalibrated:
		return "calibrated"
	case DeviceMetricCalibrationStateUnspecified:
		return "unspecified"
	}
	return "<unknown>"
}
func (code DeviceMetricCalibrationState) Display() string {
	switch code {
	case DeviceMetricCalibrationStateNotCalibrated:
		return "Not Calibrated"
	case DeviceMetricCalibrationStateCalibrationRequired:
		return "Calibration Required"
	case DeviceMetricCalibrationStateCalibrated:
		return "Calibrated"
	case DeviceMetricCalibrationStateUnspecified:
		return "Unspecified"
	}
	return "<unknown>"
}
func (code DeviceMetricCalibrationState) Definition() string {
	switch code {
	case DeviceMetricCalibrationStateNotCalibrated:
		return "The metric has not been calibrated."
	case DeviceMetricCalibrationStateCalibrationRequired:
		return "The metric needs to be calibrated."
	case DeviceMetricCalibrationStateCalibrated:
		return "The metric has been calibrated."
	case DeviceMetricCalibrationStateUnspecified:
		return "The state of calibration of this metric is unspecified."
	}
	return "<unknown>"
}
