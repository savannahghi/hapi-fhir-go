package fhir430

// ProdCharacteristic is documented here http://hl7.org/fhir/StructureDefinition/ProdCharacteristic
// Base StructureDefinition for ProdCharacteristic Type: The marketing status describes the date when a medicinal product is actually put on the market or the date as of which it is no longer available.
type ProdCharacteristic struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Height            *Quantity        `json:"height,omitempty"`
	Width             *Quantity        `json:"width,omitempty"`
	Depth             *Quantity        `json:"depth,omitempty"`
	Weight            *Quantity        `json:"weight,omitempty"`
	NominalVolume     *Quantity        `json:"nominalVolume,omitempty"`
	ExternalDiameter  *Quantity        `json:"externalDiameter,omitempty"`
	Shape             *string          `json:"shape,omitempty"`
	Color             []string         `json:"color,omitempty"`
	Imprint           []string         `json:"imprint,omitempty"`
	Image             []Attachment     `json:"image,omitempty"`
	Scoring           *CodeableConcept `json:"scoring,omitempty"`
}
