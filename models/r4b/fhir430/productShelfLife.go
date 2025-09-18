
package fhir430
// ProductShelfLife is documented here http://hl7.org/fhir/StructureDefinition/ProductShelfLife
// Base StructureDefinition for ProductShelfLife Type: The shelf-life and storage information for a medicinal product item or container can be described using this class.
type ProductShelfLife struct {
	ID                           *string           `json:"ID,omitempty"`
	Extension                    []Extension       `json:"extension,omitempty"`
	ModifierExtension            []Extension       `json:"modifierExtension,omitempty"`
	Identifier                   *Identifier       `json:"identifier,omitempty"`
	Type                         CodeableConcept   `json:"type"`
	Period                       Quantity          `json:"period"`
	SpecialPrecautionsForStorage []CodeableConcept `json:"specialPrecautionsForStorage,omitempty"`
}
