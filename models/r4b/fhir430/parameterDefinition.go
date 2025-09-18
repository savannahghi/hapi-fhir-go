
package fhir430
// ParameterDefinition is documented here http://hl7.org/fhir/StructureDefinition/ParameterDefinition
// Base StructureDefinition for ParameterDefinition Type: The parameters to the module. This collection specifies both the input and output parameters. Input parameters are provided by the caller as part of the $evaluate operation. Output parameters are included in the GuidanceResponse.
type ParameterDefinition struct {
	ID            *string               `json:"ID,omitempty"`
	Extension     []Extension           `json:"extension,omitempty"`
	Name          *string               `json:"name,omitempty"`
	Use           OperationParameterUse `json:"use"`
	Min           *int                  `json:"min,omitempty"`
	Max           *string               `json:"max,omitempty"`
	Documentation *string               `json:"documentation,omitempty"`
	Type          string                `json:"type"`
	Profile       *string               `json:"profile,omitempty"`
}
