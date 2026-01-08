package fhir430

// Identifier is documented here http://hl7.org/fhir/StructureDefinition/Identifier
// Base StructureDefinition for Identifier Type: An identifier - identifies some entity uniquely and unambiguously. Typically this is used for business identifiers.
type Identifier struct {
	ID        *string          `json:"id,omitempty"`
	Extension []Extension      `json:"extension,omitempty"`
	Use       *IdentifierUse   `json:"use,omitempty"`
	Type      *CodeableConcept `json:"type,omitempty"`
	System    *string          `json:"system,omitempty"`
	Value     *string          `json:"value,omitempty"`
	Period    *Period          `json:"period,omitempty"`
	Assigner  *Reference       `json:"assigner,omitempty"`
}
