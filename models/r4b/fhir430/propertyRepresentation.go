
package fhir430

import (
	"encoding/json"
	"fmt"
)
// PropertyRepresentation is documented here http://hl7.org/fhir/ValueSet/property-representation
type PropertyRepresentation int

const (
	PropertyRepresentationXmlAttr PropertyRepresentation = iota
	PropertyRepresentationXmlText
	PropertyRepresentationTypeAttr
	PropertyRepresentationCdaText
	PropertyRepresentationXhtml
)

func (code PropertyRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *PropertyRepresentation) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal PropertyRepresentation code `%s`", s)
	}
	switch s {
	case "xmlAttr":
		*code = PropertyRepresentationXmlAttr
	case "xmlText":
		*code = PropertyRepresentationXmlText
	case "typeAttr":
		*code = PropertyRepresentationTypeAttr
	case "cdaText":
		*code = PropertyRepresentationCdaText
	case "xhtml":
		*code = PropertyRepresentationXhtml
	default:
		return fmt.Errorf("unknown PropertyRepresentation code `%s`", s)
	}
	return nil
}
func (code PropertyRepresentation) String() string {
	return code.Code()
}
func (code PropertyRepresentation) Code() string {
	switch code {
	case PropertyRepresentationXmlAttr:
		return "xmlAttr"
	case PropertyRepresentationXmlText:
		return "xmlText"
	case PropertyRepresentationTypeAttr:
		return "typeAttr"
	case PropertyRepresentationCdaText:
		return "cdaText"
	case PropertyRepresentationXhtml:
		return "xhtml"
	}
	return "<unknown>"
}
func (code PropertyRepresentation) Display() string {
	switch code {
	case PropertyRepresentationXmlAttr:
		return "XML Attribute"
	case PropertyRepresentationXmlText:
		return "XML Text"
	case PropertyRepresentationTypeAttr:
		return "Type Attribute"
	case PropertyRepresentationCdaText:
		return "CDA Text Format"
	case PropertyRepresentationXhtml:
		return "XHTML"
	}
	return "<unknown>"
}
func (code PropertyRepresentation) Definition() string {
	switch code {
	case PropertyRepresentationXmlAttr:
		return "In XML, this property is represented as an attribute not an element."
	case PropertyRepresentationXmlText:
		return "This element is represented using the XML text attribute (primitives only)."
	case PropertyRepresentationTypeAttr:
		return "The type of this element is indicated using xsi:type."
	case PropertyRepresentationCdaText:
		return "Use CDA narrative instead of XHTML."
	case PropertyRepresentationXhtml:
		return "The property is represented using XHTML."
	}
	return "<unknown>"
}
