
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// DeviceNameType is documented here http://hl7.org/fhir/ValueSet/device-nametype
type DeviceNameType int

const (
	DeviceNameTypeUdiLabelName DeviceNameType = iota
	DeviceNameTypeUserFriendlyName
	DeviceNameTypePatientReportedName
	DeviceNameTypeManufacturerName
	DeviceNameTypeModelName
	DeviceNameTypeOther
)

func (code DeviceNameType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DeviceNameType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal DeviceNameType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "udi-label-name":
		*code = DeviceNameTypeUdiLabelName
	case "user-friendly-name":
		*code = DeviceNameTypeUserFriendlyName
	case "patient-reported-name":
		*code = DeviceNameTypePatientReportedName
	case "manufacturer-name":
		*code = DeviceNameTypeManufacturerName
	case "model-name":
		*code = DeviceNameTypeModelName
	case "other":
		*code = DeviceNameTypeOther
	default:
		return fmt.Errorf("unknown DeviceNameType code `%s`", s)
	}
	return nil
}
func (code DeviceNameType) String() string {
	return code.Code()
}
func (code DeviceNameType) Code() string {
	switch code {
	case DeviceNameTypeUdiLabelName:
		return "udi-label-name"
	case DeviceNameTypeUserFriendlyName:
		return "user-friendly-name"
	case DeviceNameTypePatientReportedName:
		return "patient-reported-name"
	case DeviceNameTypeManufacturerName:
		return "manufacturer-name"
	case DeviceNameTypeModelName:
		return "model-name"
	case DeviceNameTypeOther:
		return "other"
	}
	return "<unknown>"
}
func (code DeviceNameType) Display() string {
	switch code {
	case DeviceNameTypeUdiLabelName:
		return "UDI Label name"
	case DeviceNameTypeUserFriendlyName:
		return "User Friendly name"
	case DeviceNameTypePatientReportedName:
		return "Patient Reported name"
	case DeviceNameTypeManufacturerName:
		return "Manufacturer name"
	case DeviceNameTypeModelName:
		return "Model name"
	case DeviceNameTypeOther:
		return "other"
	}
	return "<unknown>"
}
func (code DeviceNameType) Definition() string {
	switch code {
	case DeviceNameTypeUdiLabelName:
		return "UDI Label name."
	case DeviceNameTypeUserFriendlyName:
		return "User Friendly name."
	case DeviceNameTypePatientReportedName:
		return "Patient Reported name."
	case DeviceNameTypeManufacturerName:
		return "Manufacturer name."
	case DeviceNameTypeModelName:
		return "Model name."
	case DeviceNameTypeOther:
		return "other."
	}
	return "<unknown>"
}
