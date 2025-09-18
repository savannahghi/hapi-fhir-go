
package fhir430
// SubstanceAmount is documented here http://hl7.org/fhir/StructureDefinition/SubstanceAmount
// Base StructureDefinition for SubstanceAmount Type: Chemical substances are a single substance type whose primary defining element is the molecular structure. Chemical substances shall be defined on the basis of their complete covalent molecular structure; the presence of a salt (counter-ion) and/or solvates (water, alcohols) is also captured. Purity, grade, physical form or particle size are not taken into account in the definition of a chemical substance or in the assignment of a Substance ID.
type SubstanceAmount struct {
	ID                *string                        `json:"ID,omitempty"`
	Extension         []Extension                    `json:"extension,omitempty"`
	ModifierExtension []Extension                    `json:"modifierExtension,omitempty"`
	AmountQuantity    *Quantity                      `json:"amountQuantity,omitempty"`
	AmountRange       *Range                         `json:"amountRange,omitempty"`
	AmountString      *string                        `json:"amountString,omitempty"`
	AmountType        *CodeableConcept               `json:"amountType,omitempty"`
	AmountText        *string                        `json:"amountText,omitempty"`
	ReferenceRange    *SubstanceAmountReferenceRange `json:"referenceRange,omitempty"`
}

// Reference range of possible or expected values.
type SubstanceAmountReferenceRange struct {
	ID        *string     `json:"ID,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	LowLimit  *Quantity   `json:"lowLimit,omitempty"`
	HighLimit *Quantity   `json:"highLimit,omitempty"`
}
