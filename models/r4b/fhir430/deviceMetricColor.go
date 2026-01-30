
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// DeviceMetricColor is documented here http://hl7.org/fhir/ValueSet/metric-color
type DeviceMetricColor int

const (
	DeviceMetricColorBlack DeviceMetricColor = iota
	DeviceMetricColorRed
	DeviceMetricColorGreen
	DeviceMetricColorYellow
	DeviceMetricColorBlue
	DeviceMetricColorMagenta
	DeviceMetricColorCyan
	DeviceMetricColorWhite
)

func (code DeviceMetricColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DeviceMetricColor) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal DeviceMetricColor code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "black":
		*code = DeviceMetricColorBlack
	case "red":
		*code = DeviceMetricColorRed
	case "green":
		*code = DeviceMetricColorGreen
	case "yellow":
		*code = DeviceMetricColorYellow
	case "blue":
		*code = DeviceMetricColorBlue
	case "magenta":
		*code = DeviceMetricColorMagenta
	case "cyan":
		*code = DeviceMetricColorCyan
	case "white":
		*code = DeviceMetricColorWhite
	default:
		return fmt.Errorf("unknown DeviceMetricColor code `%s`", s)
	}
	return nil
}
func (code DeviceMetricColor) String() string {
	return code.Code()
}
func (code DeviceMetricColor) Code() string {
	switch code {
	case DeviceMetricColorBlack:
		return "black"
	case DeviceMetricColorRed:
		return "red"
	case DeviceMetricColorGreen:
		return "green"
	case DeviceMetricColorYellow:
		return "yellow"
	case DeviceMetricColorBlue:
		return "blue"
	case DeviceMetricColorMagenta:
		return "magenta"
	case DeviceMetricColorCyan:
		return "cyan"
	case DeviceMetricColorWhite:
		return "white"
	}
	return "<unknown>"
}
func (code DeviceMetricColor) Display() string {
	switch code {
	case DeviceMetricColorBlack:
		return "Color Black"
	case DeviceMetricColorRed:
		return "Color Red"
	case DeviceMetricColorGreen:
		return "Color Green"
	case DeviceMetricColorYellow:
		return "Color Yellow"
	case DeviceMetricColorBlue:
		return "Color Blue"
	case DeviceMetricColorMagenta:
		return "Color Magenta"
	case DeviceMetricColorCyan:
		return "Color Cyan"
	case DeviceMetricColorWhite:
		return "Color White"
	}
	return "<unknown>"
}
func (code DeviceMetricColor) Definition() string {
	switch code {
	case DeviceMetricColorBlack:
		return "Color for representation - black."
	case DeviceMetricColorRed:
		return "Color for representation - red."
	case DeviceMetricColorGreen:
		return "Color for representation - green."
	case DeviceMetricColorYellow:
		return "Color for representation - yellow."
	case DeviceMetricColorBlue:
		return "Color for representation - blue."
	case DeviceMetricColorMagenta:
		return "Color for representation - magenta."
	case DeviceMetricColorCyan:
		return "Color for representation - cyan."
	case DeviceMetricColorWhite:
		return "Color for representation - white."
	}
	return "<unknown>"
}
