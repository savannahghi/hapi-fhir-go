package models

import "encoding/json"

// Ingredient is documented here http://hl7.org/fhir/StructureDefinition/Ingredient
type Ingredient struct {
	ID                  *string                  `json:"id,omitempty"`
	Meta                *Meta                    `json:"meta,omitempty"`
	ImplicitRules       *string                  `json:"implicitRules,omitempty"`
	Language            *string                  `json:"language,omitempty"`
	Text                *Narrative               `json:"text,omitempty"`
	Extension           []Extension              `json:"extension,omitempty"`
	ModifierExtension   []Extension              `json:"modifierExtension,omitempty"`
	Identifier          *Identifier              `json:"identifier,omitempty"`
	Status              PublicationStatus        `json:"status"`
	For                 []Reference              `json:"for,omitempty"`
	Role                CodeableConcept          `json:"role"`
	Function            []CodeableConcept        `json:"function,omitempty"`
	Group               *CodeableConcept         `json:"group,omitempty"`
	AllergenicIndicator *bool                    `json:"allergenicIndicator,omitempty"`
	Comment             *string                  `json:"comment,omitempty"`
	Manufacturer        []IngredientManufacturer `json:"manufacturer,omitempty"`
	Substance           IngredientSubstance      `json:"substance"`
}

type IngredientManufacturer struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Manufacturer      Reference   `json:"manufacturer"`
}

type IngredientSubstance struct {
	ID                *string                       `json:"id,omitempty"`
	Extension         []Extension                   `json:"extension,omitempty"`
	ModifierExtension []Extension                   `json:"modifierExtension,omitempty"`
	Code              CodeableReference             `json:"code"`
	Strength          []IngredientSubstanceStrength `json:"strength,omitempty"`
}

//nolint:lll
type IngredientSubstanceStrength struct {
	ID                           *string                                        `json:"id,omitempty"`
	Extension                    []Extension                                    `json:"extension,omitempty"`
	ModifierExtension            []Extension                                    `json:"modifierExtension,omitempty"`
	PresentationRatio            *Ratio                                         `json:"presentationRatio,omitempty"`
	PresentationRatioRange       *RatioRange                                    `json:"presentationRatioRange,omitempty"`
	PresentationCodeableConcept  *CodeableConcept                               `json:"presentationCodeableConcept,omitempty"`
	PresentationQuantity         *Quantity                                      `json:"presentationQuantity,omitempty"`
	TextPresentation             *string                                        `json:"textPresentation,omitempty"`
	ConcentrationRatio           *Ratio                                         `json:"concentrationRatio,omitempty"`
	ConcentrationRatioRange      *RatioRange                                    `json:"concentrationRatioRange,omitempty"`
	ConcentrationCodeableConcept *CodeableConcept                               `json:"concentrationCodeableConcept,omitempty"`
	ConcentrationQuantity        *Quantity                                      `json:"concentrationQuantity,omitempty"`
	TextConcentration            *string                                        `json:"textConcentration,omitempty"`
	Basis                        *CodeableConcept                               `json:"basis,omitempty"`
	MeasurementPoint             *string                                        `json:"measurementPoint,omitempty"`
	Country                      []CodeableConcept                              `json:"country,omitempty"`
	ReferenceStrength            []IngredientSubstanceStrengthReferenceStrength `json:"referenceStrength,omitempty"`
}

type IngredientSubstanceStrengthReferenceStrength struct {
	ID                 *string           `json:"id,omitempty"`
	Extension          []Extension       `json:"extension,omitempty"`
	ModifierExtension  []Extension       `json:"modifierExtension,omitempty"`
	Substance          CodeableReference `json:"substance"`
	StrengthRatio      Ratio             `json:"strengthRatio"`
	StrengthRatioRange RatioRange        `json:"strengthRatioRange"`
	StrengthQuantity   Quantity          `json:"strengthQuantity"`
	MeasurementPoint   *string           `json:"measurementPoint,omitempty"`
	Country            []CodeableConcept `json:"country,omitempty"`
}

type OtherIngredient Ingredient

// MarshalJSON marshals the given Ingredient as JSON into a byte slice.
func (r Ingredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherIngredient
		ResourceType string `json:"resourceType"`
	}{
		OtherIngredient: OtherIngredient(r),
		ResourceType:    "Ingredient",
	})
}

// UnmarshalIngredient unmarshals a Ingredient.
func UnmarshalIngredient(b []byte) (Ingredient, error) {
	var ingredient Ingredient
	if err := json.Unmarshal(b, &ingredient); err != nil {
		return ingredient, err
	}

	return ingredient, nil
}
