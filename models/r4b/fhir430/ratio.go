package fhir430

// Ratio is documented here http://hl7.org/fhir/StructureDefinition/Ratio
// Base StructureDefinition for Ratio Type: A relationship of two Quantity values - expressed as a numerator and a denominator.
type Ratio struct {
	ID          *string     `json:"id,omitempty"`
	Extension   []Extension `json:"extension,omitempty"`
	Numerator   *Quantity   `json:"numerator,omitempty"`
	Denominator *Quantity   `json:"denominator,omitempty"`
}
