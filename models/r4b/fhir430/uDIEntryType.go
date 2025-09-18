
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// UDIEntryType is documented here http://hl7.org/fhir/ValueSet/udi-entry-type
type UDIEntryType int

const (
	UDIEntryTypeBarcode UDIEntryType = iota
	UDIEntryTypeRfid
	UDIEntryTypeManual
	UDIEntryTypeCard
	UDIEntryTypeSelfReported
	UDIEntryTypeUnknown
)

func (code UDIEntryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *UDIEntryType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal UDIEntryType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "barcode":
		*code = UDIEntryTypeBarcode
	case "rfid":
		*code = UDIEntryTypeRfid
	case "manual":
		*code = UDIEntryTypeManual
	case "card":
		*code = UDIEntryTypeCard
	case "self-reported":
		*code = UDIEntryTypeSelfReported
	case "unknown":
		*code = UDIEntryTypeUnknown
	default:
		return fmt.Errorf("unknown UDIEntryType code `%s`", s)
	}
	return nil
}
func (code UDIEntryType) String() string {
	return code.Code()
}
func (code UDIEntryType) Code() string {
	switch code {
	case UDIEntryTypeBarcode:
		return "barcode"
	case UDIEntryTypeRfid:
		return "rfid"
	case UDIEntryTypeManual:
		return "manual"
	case UDIEntryTypeCard:
		return "card"
	case UDIEntryTypeSelfReported:
		return "self-reported"
	case UDIEntryTypeUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code UDIEntryType) Display() string {
	switch code {
	case UDIEntryTypeBarcode:
		return "Barcode"
	case UDIEntryTypeRfid:
		return "RFID"
	case UDIEntryTypeManual:
		return "Manual"
	case UDIEntryTypeCard:
		return "Card"
	case UDIEntryTypeSelfReported:
		return "Self Reported"
	case UDIEntryTypeUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code UDIEntryType) Definition() string {
	switch code {
	case UDIEntryTypeBarcode:
		return "a barcodescanner captured the data from the device label."
	case UDIEntryTypeRfid:
		return "An RFID chip reader captured the data from the device label."
	case UDIEntryTypeManual:
		return "The data was read from the label by a person and manually entered. (e.g.  via a keyboard)."
	case UDIEntryTypeCard:
		return "The data originated from a patient's implant card and was read by an operator."
	case UDIEntryTypeSelfReported:
		return "The data originated from a patient source and was not directly scanned or read from a label or card."
	case UDIEntryTypeUnknown:
		return "The method of data capture has not been determined."
	}
	return "<unknown>"
}
