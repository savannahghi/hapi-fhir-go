
package fhir430

import (
	"encoding/json"
	"fmt"
	"strings"
)
// ExtensionContextType is documented here http://hl7.org/fhir/ValueSet/extension-context-type
type ExtensionContextType int

const (
	ExtensionContextTypeFhirpath ExtensionContextType = iota
	ExtensionContextTypeElement
	ExtensionContextTypeExtension
)

func (code ExtensionContextType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ExtensionContextType) UnmarshalJSON(input []byte) error {
	var s string
	if err := json.Unmarshal(input, &s); err != nil {
		return fmt.Errorf("failed to Unmarshal ExtensionContextType code `%s`", s)
	}
	s = strings.ToLower(s)
	switch s {
	case "fhirpath":
		*code = ExtensionContextTypeFhirpath
	case "element":
		*code = ExtensionContextTypeElement
	case "extension":
		*code = ExtensionContextTypeExtension
	default:
		return fmt.Errorf("unknown ExtensionContextType code `%s`", s)
	}
	return nil
}
func (code ExtensionContextType) String() string {
	return code.Code()
}
func (code ExtensionContextType) Code() string {
	switch code {
	case ExtensionContextTypeFhirpath:
		return "fhirpath"
	case ExtensionContextTypeElement:
		return "element"
	case ExtensionContextTypeExtension:
		return "extension"
	}
	return "<unknown>"
}
func (code ExtensionContextType) Display() string {
	switch code {
	case ExtensionContextTypeFhirpath:
		return "FHIRPath"
	case ExtensionContextTypeElement:
		return "Element ID"
	case ExtensionContextTypeExtension:
		return "Extension URL"
	}
	return "<unknown>"
}
func (code ExtensionContextType) Definition() string {
	switch code {
	case ExtensionContextTypeFhirpath:
		return "The context is all elements that match the FHIRPath query found in the expression."
	case ExtensionContextTypeElement:
		return "The context is any element that has an ElementDefinition.id that matches that found in the expression. This includes ElementDefinition Ids that have slicing identifiers. The full path for the element is [url]#[elementid]. If there is no #, the Element id is one defined in the base specification."
	case ExtensionContextTypeExtension:
		return "The context is a particular extension from a particular StructureDefinition, and the expression is just a uri that identifies the extension."
	}
	return "<unknown>"
}
