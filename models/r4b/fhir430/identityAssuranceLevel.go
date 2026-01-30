
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// IdentityAssuranceLevel is documented here http://hl7.org/fhir/ValueSet/identity-assuranceLevel
type IdentityAssuranceLevel int

const (
	IdentityAssuranceLevelLevel1 IdentityAssuranceLevel = iota
	IdentityAssuranceLevelLevel2
	IdentityAssuranceLevelLevel3
	IdentityAssuranceLevelLevel4
)

func (code IdentityAssuranceLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *IdentityAssuranceLevel) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal IdentityAssuranceLevel code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "level1":
		*code = IdentityAssuranceLevelLevel1
	case "level2":
		*code = IdentityAssuranceLevelLevel2
	case "level3":
		*code = IdentityAssuranceLevelLevel3
	case "level4":
		*code = IdentityAssuranceLevelLevel4
	default:
		return fmt.Errorf("unknown IdentityAssuranceLevel code `%s`", s)
	}
	return nil
}
func (code IdentityAssuranceLevel) String() string {
	return code.Code()
}
func (code IdentityAssuranceLevel) Code() string {
	switch code {
	case IdentityAssuranceLevelLevel1:
		return "level1"
	case IdentityAssuranceLevelLevel2:
		return "level2"
	case IdentityAssuranceLevelLevel3:
		return "level3"
	case IdentityAssuranceLevelLevel4:
		return "level4"
	}
	return "<unknown>"
}
func (code IdentityAssuranceLevel) Display() string {
	switch code {
	case IdentityAssuranceLevelLevel1:
		return "Level 1"
	case IdentityAssuranceLevelLevel2:
		return "Level 2"
	case IdentityAssuranceLevelLevel3:
		return "Level 3"
	case IdentityAssuranceLevelLevel4:
		return "Level 4"
	}
	return "<unknown>"
}
func (code IdentityAssuranceLevel) Definition() string {
	switch code {
	case IdentityAssuranceLevelLevel1:
		return "Little or no confidence in the asserted identity's accuracy."
	case IdentityAssuranceLevelLevel2:
		return "Some confidence in the asserted identity's accuracy."
	case IdentityAssuranceLevelLevel3:
		return "High confidence in the asserted identity's accuracy."
	case IdentityAssuranceLevelLevel4:
		return "Very high confidence in the asserted identity's accuracy."
	}
	return "<unknown>"
}
