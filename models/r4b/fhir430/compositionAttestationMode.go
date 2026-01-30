
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// CompositionAttestationMode is documented here http://hl7.org/fhir/ValueSet/composition-attestation-mode
type CompositionAttestationMode int

const (
	CompositionAttestationModePersonal CompositionAttestationMode = iota
	CompositionAttestationModeProfessional
	CompositionAttestationModeLegal
	CompositionAttestationModeOfficial
)

func (code CompositionAttestationMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CompositionAttestationMode) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal CompositionAttestationMode code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "personal":
		*code = CompositionAttestationModePersonal
	case "professional":
		*code = CompositionAttestationModeProfessional
	case "legal":
		*code = CompositionAttestationModeLegal
	case "official":
		*code = CompositionAttestationModeOfficial
	default:
		return fmt.Errorf("unknown CompositionAttestationMode code `%s`", s)
	}
	return nil
}
func (code CompositionAttestationMode) String() string {
	return code.Code()
}
func (code CompositionAttestationMode) Code() string {
	switch code {
	case CompositionAttestationModePersonal:
		return "personal"
	case CompositionAttestationModeProfessional:
		return "professional"
	case CompositionAttestationModeLegal:
		return "legal"
	case CompositionAttestationModeOfficial:
		return "official"
	}
	return "<unknown>"
}
func (code CompositionAttestationMode) Display() string {
	switch code {
	case CompositionAttestationModePersonal:
		return "Personal"
	case CompositionAttestationModeProfessional:
		return "Professional"
	case CompositionAttestationModeLegal:
		return "Legal"
	case CompositionAttestationModeOfficial:
		return "Official"
	}
	return "<unknown>"
}
func (code CompositionAttestationMode) Definition() string {
	switch code {
	case CompositionAttestationModePersonal:
		return "The person authenticated the content in their personal capacity."
	case CompositionAttestationModeProfessional:
		return "The person authenticated the content in their professional capacity."
	case CompositionAttestationModeLegal:
		return "The person authenticated the content and accepted legal responsibility for its content."
	case CompositionAttestationModeOfficial:
		return "The organization authenticated the content as consistent with their policies and procedures."
	}
	return "<unknown>"
}
