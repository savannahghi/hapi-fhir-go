
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// FHIRDeviceStatus is documented here http://hl7.org/fhir/ValueSet/device-status
type FHIRDeviceStatus int

const (
	FHIRDeviceStatusActive FHIRDeviceStatus = iota
	FHIRDeviceStatusInactive
	FHIRDeviceStatusEnteredInError
	FHIRDeviceStatusUnknown
)

func (code FHIRDeviceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *FHIRDeviceStatus) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal FHIRDeviceStatus code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "active":
		*code = FHIRDeviceStatusActive
	case "inactive":
		*code = FHIRDeviceStatusInactive
	case "entered-in-error":
		*code = FHIRDeviceStatusEnteredInError
	case "unknown":
		*code = FHIRDeviceStatusUnknown
	default:
		return fmt.Errorf("unknown FHIRDeviceStatus code `%s`", s)
	}
	return nil
}
func (code FHIRDeviceStatus) String() string {
	return code.Code()
}
func (code FHIRDeviceStatus) Code() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "active"
	case FHIRDeviceStatusInactive:
		return "inactive"
	case FHIRDeviceStatusEnteredInError:
		return "entered-in-error"
	case FHIRDeviceStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code FHIRDeviceStatus) Display() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "Active"
	case FHIRDeviceStatusInactive:
		return "Inactive"
	case FHIRDeviceStatusEnteredInError:
		return "Entered in Error"
	case FHIRDeviceStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code FHIRDeviceStatus) Definition() string {
	switch code {
	case FHIRDeviceStatusActive:
		return "The device is available for use.  Note: For *implanted devices*  this means that the device is implanted in the patient."
	case FHIRDeviceStatusInactive:
		return "The device is no longer available for use (e.g. lost, expired, damaged).  Note: For *implanted devices*  this means that the device has been removed from the patient."
	case FHIRDeviceStatusEnteredInError:
		return "The device was entered in error and voided."
	case FHIRDeviceStatusUnknown:
		return "The status of the device has not been determined."
	}
	return "<unknown>"
}
