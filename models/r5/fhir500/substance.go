package fhir500

import (
	"encoding/json"
)

type Substance struct {
	ID                *string               `json:"id,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `json:"identifier,omitempty"`
	Instance          bool                  `json:"instance"`
	Status            *FHIRSubstanceStatus  `json:"status,omitempty"`
	Category          []CodeableConcept     `json:"category,omitempty"`
	Code              CodeableReference     `json:"code"`
	Description       *string               `json:"description,omitempty"`
	Expiry            *string               `json:"expiry,omitempty"`
	Quantity          *Quantity             `json:"quantity,omitempty"`
	Ingredient        []SubstanceIngredient `json:"ingredient,omitempty"`
}

type SubstanceIngredient struct {
	ID                       *string         `json:"id,omitempty"`
	Extension                []Extension     `json:"extension,omitempty"`
	ModifierExtension        []Extension     `json:"modifierExtension,omitempty"`
	Quantity                 *Ratio          `json:"quantity,omitempty"`
	SubstanceCodeableConcept CodeableConcept `json:"substanceCodeableConcept"`
	SubstanceReference       Reference       `json:"substanceReference"`
}

type OtherSubstance Substance

// MarshalJSON marshals the given Substance as JSON into a byte slice.
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

type FHIRSubstanceStatus string

const (
	FHIRSubstanceStatusActive         FHIRSubstanceStatus = "active"
	FHIRSubstanceStatusInactive       FHIRSubstanceStatus = "inactive"
	FHIRSubstanceStatusEnteredInError FHIRSubstanceStatus = "entered-in-error"
)

func (code FHIRSubstanceStatus) String() string {
	return code.Code()
}

func (code FHIRSubstanceStatus) Code() string {
	switch code {
	case FHIRSubstanceStatusActive:
		return "active"
	case FHIRSubstanceStatusInactive:
		return "inactive"
	case FHIRSubstanceStatusEnteredInError:
		return "entered-in-error"
	}

	return "<unknown>"
}
