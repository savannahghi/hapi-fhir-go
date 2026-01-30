package fhir430

import "encoding/json"

// Substance is documented here http://hl7.org/fhir/StructureDefinition/Substance
// A homogeneous material with a definite composition.
type Substance struct {
	ID                *string               `json:"id,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Contained         []json.RawMessage     `json:"contained,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `json:"identifier,omitempty"`
	Status            *FHIRSubstanceStatus  `json:"status,omitempty"`
	Category          []CodeableConcept     `json:"category,omitempty"`
	Code              CodeableConcept       `json:"code"`
	Description       *string               `json:"description,omitempty"`
	Instance          []SubstanceInstance   `json:"instance,omitempty"`
	Ingredient        []SubstanceIngredient `json:"ingredient,omitempty"`
}

// Substance may be used to describe a kind of substance, or a specific package/container of the substance: an instance.
// If this element is not present, then the substance resource describes a kind of substance
type SubstanceInstance struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Identifier        *Identifier `json:"identifier,omitempty"`
	Expiry            *string     `json:"expiry,omitempty"`
	Quantity          *Quantity   `json:"quantity,omitempty"`
}

// A substance can be composed of other substances.
type SubstanceIngredient struct {
	ID                       *string         `json:"id,omitempty"`
	Extension                []Extension     `json:"extension,omitempty"`
	ModifierExtension        []Extension     `json:"modifierExtension,omitempty"`
	Quantity                 *Ratio          `json:"quantity,omitempty"`
	SubstanceCodeableConcept CodeableConcept `json:"substanceCodeableConcept"`
	SubstanceReference       Reference       `json:"substanceReference"`
}

// This function returns resource reference information
func (r Substance) ResourceRef() (string, *string) {
	return "Substance", r.ID
}

// This function returns resource reference information
func (r Substance) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherSubstance Substance

// MarshalJSON marshals the given Substance as JSON into a byte slice
func (r Substance) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstance
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstance: OtherSubstance(r),
		ResourceType:   "Substance",
	})
}

// UnmarshalSubstance unmarshals a Substance.
func UnmarshalSubstance(b []byte) (Substance, error) {
	var substance Substance
	if err := json.Unmarshal(b, &substance); err != nil {
		return substance, err
	}
	return substance, nil
}
