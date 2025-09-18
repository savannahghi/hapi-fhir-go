
package fhir430

import "encoding/json"
// SampledData is documented here http://hl7.org/fhir/StructureDefinition/SampledData
// Base StructureDefinition for SampledData Type: A series of measurements taken by a device, with upper and lower limits. There may be more than one dimension in the data.
type SampledData struct {
	ID         *string      `json:"ID,omitempty"`
	Extension  []Extension  `json:"extension,omitempty"`
	Origin     Quantity     `json:"origin"`
	Period     json.Number  `json:"period"`
	Factor     *json.Number `json:"factor,omitempty"`
	LowerLimit *json.Number `json:"lowerLimit,omitempty"`
	UpperLimit *json.Number `json:"upperLimit,omitempty"`
	Dimensions int          `json:"dimensions"`
	Data       *string      `json:"data,omitempty"`
}
