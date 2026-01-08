package fhir430

import "encoding/json"

// MedicinalProductInteraction is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductInteraction
// The interactions of the medicinal product with other medicinal products, or other forms of interactions.
type MedicinalProductInteraction struct {
	ID                *string                                  `json:"id,omitempty"`
	Meta              *Meta                                    `json:"meta,omitempty"`
	ImplicitRules     *string                                  `json:"implicitRules,omitempty"`
	Language          *string                                  `json:"language,omitempty"`
	Text              *Narrative                               `json:"text,omitempty"`
	Contained         []json.RawMessage                        `json:"contained,omitempty"`
	Extension         []Extension                              `json:"extension,omitempty"`
	ModifierExtension []Extension                              `json:"modifierExtension,omitempty"`
	Subject           []Reference                              `json:"subject,omitempty"`
	Description       *string                                  `json:"description,omitempty"`
	Interactant       []MedicinalProductInteractionInteractant `json:"interactant,omitempty"`
	Type              *CodeableConcept                         `json:"type,omitempty"`
	Effect            *CodeableConcept                         `json:"effect,omitempty"`
	Incidence         *CodeableConcept                         `json:"incidence,omitempty"`
	Management        *CodeableConcept                         `json:"management,omitempty"`
}

// The specific medication, food or laboratory test that interacts.
type MedicinalProductInteractionInteractant struct {
	ID                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	ItemReference       Reference       `json:"itemReference"`
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
}

// This function returns resource reference information
func (r MedicinalProductInteraction) ResourceRef() (string, *string) {
	return "MedicinalProductInteraction", r.ID
}

// This function returns resource reference information
func (r MedicinalProductInteraction) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedicinalProductInteraction MedicinalProductInteraction

// MarshalJSON marshals the given MedicinalProductInteraction as JSON into a byte slice
func (r MedicinalProductInteraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductInteraction
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductInteraction: OtherMedicinalProductInteraction(r),
		ResourceType:                     "MedicinalProductInteraction",
	})
}

// UnmarshalMedicinalProductInteraction unmarshals a MedicinalProductInteraction.
func UnmarshalMedicinalProductInteraction(b []byte) (MedicinalProductInteraction, error) {
	var medicinalProductInteraction MedicinalProductInteraction
	if err := json.Unmarshal(b, &medicinalProductInteraction); err != nil {
		return medicinalProductInteraction, err
	}
	return medicinalProductInteraction, nil
}
