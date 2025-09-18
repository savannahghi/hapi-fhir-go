
package fhir430
// UsageContext is documented here http://hl7.org/fhir/StructureDefinition/UsageContext
// Base StructureDefinition for UsageContext Type: Specifies clinical/business/etc. metadata that can be used to retrieve, index and/or categorize an artifact. This metadata can either be specific to the applicable population (e.g., age category, DRG) or the specific context of care (e.g., venue, care setting, provider of care).
type UsageContext struct {
	ID                   *string         `json:"ID,omitempty"`
	Extension            []Extension     `json:"extension,omitempty"`
	Code                 Coding          `json:"code"`
	ValueCodeableConcept CodeableConcept `json:"valueCodeableConcept"`
	ValueQuantity        Quantity        `json:"valueQuantity"`
	ValueRange           Range           `json:"valueRange"`
	ValueReference       Reference       `json:"valueReference"`
}
