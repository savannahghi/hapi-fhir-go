
package fhir430
// Signature is documented here http://hl7.org/fhir/StructureDefinition/Signature
// Base StructureDefinition for Signature Type: A signature along with supporting context. The signature may be a digital signature that is cryptographic in nature, or some other signature acceptable to the domain. This other signature may be as simple as a graphical image representing a hand-written signature, or a signature ceremony Different signature approaches have different utilities.
type Signature struct {
	ID           *string     `json:"ID,omitempty"`
	Extension    []Extension `json:"extension,omitempty"`
	Type         []Coding    `json:"type"`
	When         string      `json:"when"`
	Who          Reference   `json:"who"`
	OnBehalfOf   *Reference  `json:"onBehalfOf,omitempty"`
	TargetFormat *string     `json:"targetFormat,omitempty"`
	SigFormat    *string     `json:"sigFormat,omitempty"`
	Data         *string     `json:"data,omitempty"`
}
