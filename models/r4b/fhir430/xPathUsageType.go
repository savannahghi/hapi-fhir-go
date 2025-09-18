package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)

// XPathUsageType is documented here http://hl7.org/fhir/ValueSet/search-xpath-usage
type XPathUsageType int

const (
	XPathUsageTypeNormal XPathUsageType = iota
	XPathUsageTypePhonetic
	XPathUsageTypeNearby
	XPathUsageTypeDistance
	XPathUsageTypeOther
)

func (code XPathUsageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *XPathUsageType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal XPathUsageType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "normal":
		*code = XPathUsageTypeNormal
	case "phonetic":
		*code = XPathUsageTypePhonetic
	case "nearby":
		*code = XPathUsageTypeNearby
	case "distance":
		*code = XPathUsageTypeDistance
	case "other":
		*code = XPathUsageTypeOther
	default:
		return fmt.Errorf("unknown XPathUsageType code `%s`", s)
	}
	return nil
}
func (code XPathUsageType) String() string {
	return code.Code()
}
func (code XPathUsageType) Code() string {
	switch code {
	case XPathUsageTypeNormal:
		return "normal"
	case XPathUsageTypePhonetic:
		return "phonetic"
	case XPathUsageTypeNearby:
		return "nearby"
	case XPathUsageTypeDistance:
		return "distance"
	case XPathUsageTypeOther:
		return "other"
	}
	return "<unknown>"
}
func (code XPathUsageType) Display() string {
	switch code {
	case XPathUsageTypeNormal:
		return "Normal"
	case XPathUsageTypePhonetic:
		return "Phonetic"
	case XPathUsageTypeNearby:
		return "Nearby"
	case XPathUsageTypeDistance:
		return "Distance"
	case XPathUsageTypeOther:
		return "Other"
	}
	return "<unknown>"
}
func (code XPathUsageType) Definition() string {
	switch code {
	case XPathUsageTypeNormal:
		return "The search parameter is derived directly from the selected nodes based on the type definitions."
	case XPathUsageTypePhonetic:
		return "The search parameter is derived by a phonetic transform from the selected nodes."
	case XPathUsageTypeNearby:
		return "The search parameter is based on a spatial transform of the selected nodes."
	case XPathUsageTypeDistance:
		return "The search parameter is based on a spatial transform of the selected nodes, using physical distance from the middle."
	case XPathUsageTypeOther:
		return "The interpretation of the xpath statement is unknown (and can't be automated)."
	}
	return "<unknown>"
}
