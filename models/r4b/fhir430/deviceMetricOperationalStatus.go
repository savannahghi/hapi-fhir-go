
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// DeviceMetricOperationalStatus is documented here http://hl7.org/fhir/ValueSet/metric-operational-status
type DeviceMetricOperationalStatus int

const (
	DeviceMetricOperationalStatusOn DeviceMetricOperationalStatus = iota
	DeviceMetricOperationalStatusOff
	DeviceMetricOperationalStatusStandby
	DeviceMetricOperationalStatusEnteredInError
)

func (code DeviceMetricOperationalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DeviceMetricOperationalStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal DeviceMetricOperationalStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "on":
		*code = DeviceMetricOperationalStatusOn
	case "off":
		*code = DeviceMetricOperationalStatusOff
	case "standby":
		*code = DeviceMetricOperationalStatusStandby
	case "entered-in-error":
		*code = DeviceMetricOperationalStatusEnteredInError
	default:
		return fmt.Errorf("unknown DeviceMetricOperationalStatus code `%s`", s)
	}
	return nil
}
func (code DeviceMetricOperationalStatus) String() string {
	return code.Code()
}
func (code DeviceMetricOperationalStatus) Code() string {
	switch code {
	case DeviceMetricOperationalStatusOn:
		return "on"
	case DeviceMetricOperationalStatusOff:
		return "off"
	case DeviceMetricOperationalStatusStandby:
		return "standby"
	case DeviceMetricOperationalStatusEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code DeviceMetricOperationalStatus) Display() string {
	switch code {
	case DeviceMetricOperationalStatusOn:
		return "On"
	case DeviceMetricOperationalStatusOff:
		return "Off"
	case DeviceMetricOperationalStatusStandby:
		return "Standby"
	case DeviceMetricOperationalStatusEnteredInError:
		return "Entered In Error"
	}
	return "<unknown>"
}
func (code DeviceMetricOperationalStatus) Definition() string {
	switch code {
	case DeviceMetricOperationalStatusOn:
		return "The DeviceMetric is operating and will generate DeviceObservations."
	case DeviceMetricOperationalStatusOff:
		return "The DeviceMetric is not operating."
	case DeviceMetricOperationalStatusStandby:
		return "The DeviceMetric is operating, but will not generate any DeviceObservations."
	case DeviceMetricOperationalStatusEnteredInError:
		return "The DeviceMetric was entered in error."
	}
	return "<unknown>"
}
