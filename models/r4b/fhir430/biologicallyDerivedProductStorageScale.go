
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// BiologicallyDerivedProductStorageScale is documented here http://hl7.org/fhir/ValueSet/product-storage-scale
type BiologicallyDerivedProductStorageScale int

const (
	BiologicallyDerivedProductStorageScaleFarenheit BiologicallyDerivedProductStorageScale = iota
	BiologicallyDerivedProductStorageScaleCelsius
	BiologicallyDerivedProductStorageScaleKelvin
)

func (code BiologicallyDerivedProductStorageScale) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *BiologicallyDerivedProductStorageScale) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal BiologicallyDerivedProductStorageScale code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "farenheit":
		*code = BiologicallyDerivedProductStorageScaleFarenheit
	case "celsius":
		*code = BiologicallyDerivedProductStorageScaleCelsius
	case "kelvin":
		*code = BiologicallyDerivedProductStorageScaleKelvin
	default:
		return fmt.Errorf("unknown BiologicallyDerivedProductStorageScale code `%s`", s)
	}
	return nil
}
func (code BiologicallyDerivedProductStorageScale) String() string {
	return code.Code()
}
func (code BiologicallyDerivedProductStorageScale) Code() string {
	switch code {
	case BiologicallyDerivedProductStorageScaleFarenheit:
		return "farenheit"
	case BiologicallyDerivedProductStorageScaleCelsius:
		return "celsius"
	case BiologicallyDerivedProductStorageScaleKelvin:
		return "kelvin"
	}
	return "<unknown>"
}
func (code BiologicallyDerivedProductStorageScale) Display() string {
	switch code {
	case BiologicallyDerivedProductStorageScaleFarenheit:
		return "Fahrenheit"
	case BiologicallyDerivedProductStorageScaleCelsius:
		return "Celsius"
	case BiologicallyDerivedProductStorageScaleKelvin:
		return "Kelvin"
	}
	return "<unknown>"
}
func (code BiologicallyDerivedProductStorageScale) Definition() string {
	switch code {
	case BiologicallyDerivedProductStorageScaleFarenheit:
		return "Fahrenheit temperature scale."
	case BiologicallyDerivedProductStorageScaleCelsius:
		return "Celsius or centigrade temperature scale."
	case BiologicallyDerivedProductStorageScaleKelvin:
		return "Kelvin absolute thermodynamic temperature scale."
	}
	return "<unknown>"
}
