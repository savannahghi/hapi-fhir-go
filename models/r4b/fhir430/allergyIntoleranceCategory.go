
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// AllergyIntoleranceCategory is documented here http://hl7.org/fhir/ValueSet/allergy-intolerance-category
type AllergyIntoleranceCategory int

const (
	AllergyIntoleranceCategoryFood AllergyIntoleranceCategory = iota
	AllergyIntoleranceCategoryMedication
	AllergyIntoleranceCategoryEnvironment
	AllergyIntoleranceCategoryBiologic
)

func (code AllergyIntoleranceCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *AllergyIntoleranceCategory) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal AllergyIntoleranceCategory code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "food":
		*code = AllergyIntoleranceCategoryFood
	case "medication":
		*code = AllergyIntoleranceCategoryMedication
	case "environment":
		*code = AllergyIntoleranceCategoryEnvironment
	case "biologic":
		*code = AllergyIntoleranceCategoryBiologic
	default:
		return fmt.Errorf("unknown AllergyIntoleranceCategory code `%s`", s)
	}
	return nil
}
func (code AllergyIntoleranceCategory) String() string {
	return code.Code()
}
func (code AllergyIntoleranceCategory) Code() string {
	switch code {
	case AllergyIntoleranceCategoryFood:
		return "food"
	case AllergyIntoleranceCategoryMedication:
		return "medication"
	case AllergyIntoleranceCategoryEnvironment:
		return "environment"
	case AllergyIntoleranceCategoryBiologic:
		return "biologic"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceCategory) Display() string {
	switch code {
	case AllergyIntoleranceCategoryFood:
		return "Food"
	case AllergyIntoleranceCategoryMedication:
		return "Medication"
	case AllergyIntoleranceCategoryEnvironment:
		return "Environment"
	case AllergyIntoleranceCategoryBiologic:
		return "Biologic"
	}
	return "<unknown>"
}
func (code AllergyIntoleranceCategory) Definition() string {
	switch code {
	case AllergyIntoleranceCategoryFood:
		return "Any substance consumed to provide nutritional support for the body."
	case AllergyIntoleranceCategoryMedication:
		return "Substances administered to achieve a physiological effect."
	case AllergyIntoleranceCategoryEnvironment:
		return "Any substances that are encountered in the environment, including any substance not already classified as food, medication, or biologic."
	case AllergyIntoleranceCategoryBiologic:
		return "A preparation that is synthesized from living organisms or their products, especially a human or animal protein, such as a hormone or antitoxin, that is used as a diagnostic, preventive, or therapeutic agent. Examples of biologic medications include: vaccines; allergenic extracts, which are used for both diagnosis and treatment (for example, allergy shots); gene therapies; cellular therapies.  There are other biologic products, such as tissues, which are not typically associated with allergies."
	}
	return "<unknown>"
}
