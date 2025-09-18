
package fhir430

import (
	"encoding/json"
	"fmt"
)
// CompartmentType is documented here http://hl7.org/fhir/ValueSet/compartment-type
type CompartmentType int

const (
	CompartmentTypePatient CompartmentType = iota
	CompartmentTypeEncounter
	CompartmentTypeRelatedPerson
	CompartmentTypePractitioner
	CompartmentTypeDevice
)

func (code CompartmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CompartmentType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal CompartmentType code `%s`", s)
	}
	switch s {
	case "Patient":
		*code = CompartmentTypePatient
	case "Encounter":
		*code = CompartmentTypeEncounter
	case "RelatedPerson":
		*code = CompartmentTypeRelatedPerson
	case "Practitioner":
		*code = CompartmentTypePractitioner
	case "Device":
		*code = CompartmentTypeDevice
	default:
		return fmt.Errorf("unknown CompartmentType code `%s`", s)
	}
	return nil
}
func (code CompartmentType) String() string {
	return code.Code()
}
func (code CompartmentType) Code() string {
	switch code {
	case CompartmentTypePatient:
		return "Patient"
	case CompartmentTypeEncounter:
		return "Encounter"
	case CompartmentTypeRelatedPerson:
		return "RelatedPerson"
	case CompartmentTypePractitioner:
		return "Practitioner"
	case CompartmentTypeDevice:
		return "Device"
	}
	return "<unknown>"
}
func (code CompartmentType) Display() string {
	switch code {
	case CompartmentTypePatient:
		return "Patient"
	case CompartmentTypeEncounter:
		return "Encounter"
	case CompartmentTypeRelatedPerson:
		return "RelatedPerson"
	case CompartmentTypePractitioner:
		return "Practitioner"
	case CompartmentTypeDevice:
		return "Device"
	}
	return "<unknown>"
}
func (code CompartmentType) Definition() string {
	switch code {
	case CompartmentTypePatient:
		return "The compartment definition is for the patient compartment."
	case CompartmentTypeEncounter:
		return "The compartment definition is for the encounter compartment."
	case CompartmentTypeRelatedPerson:
		return "The compartment definition is for the related-person compartment."
	case CompartmentTypePractitioner:
		return "The compartment definition is for the practitioner compartment."
	case CompartmentTypeDevice:
		return "The compartment definition is for the device compartment."
	}
	return "<unknown>"
}
