package fhir430

import "encoding/json"

// MedicinalProductIngredient is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductIngredient
// An ingredient of a manufactured item or pharmaceutical product.
type MedicinalProductIngredient struct {
	ID                  *string                                        `json:"id,omitempty"`
	Meta                *Meta                                          `json:"meta,omitempty"`
	ImplicitRules       *string                                        `json:"implicitRules,omitempty"`
	Language            *string                                        `json:"language,omitempty"`
	Text                *Narrative                                     `json:"text,omitempty"`
	Contained           []json.RawMessage                              `json:"contained,omitempty"`
	Extension           []Extension                                    `json:"extension,omitempty"`
	ModifierExtension   []Extension                                    `json:"modifierExtension,omitempty"`
	Identifier          *Identifier                                    `json:"identifier,omitempty"`
	Role                CodeableConcept                                `json:"role"`
	AllergenicIndicator *bool                                          `json:"allergenicIndicator,omitempty"`
	Manufacturer        []Reference                                    `json:"manufacturer,omitempty"`
	SpecifiedSubstance  []MedicinalProductIngredientSpecifiedSubstance `json:"specifiedSubstance,omitempty"`
	Substance           *MedicinalProductIngredientSubstance           `json:"substance,omitempty"`
}

// A specified substance that comprises this ingredient.
type MedicinalProductIngredientSpecifiedSubstance struct {
	ID                *string                                                `json:"id,omitempty"`
	Extension         []Extension                                            `json:"extension,omitempty"`
	ModifierExtension []Extension                                            `json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                        `json:"code"`
	Group             CodeableConcept                                        `json:"group"`
	Confidentiality   *CodeableConcept                                       `json:"confidentiality,omitempty"`
	Strength          []MedicinalProductIngredientSpecifiedSubstanceStrength `json:"strength,omitempty"`
}

// Quantity of the substance or specified substance present in the manufactured item or pharmaceutical product.
type MedicinalProductIngredientSpecifiedSubstanceStrength struct {
	ID                    *string                                                                 `json:"id,omitempty"`
	Extension             []Extension                                                             `json:"extension,omitempty"`
	ModifierExtension     []Extension                                                             `json:"modifierExtension,omitempty"`
	Presentation          Ratio                                                                   `json:"presentation"`
	PresentationLowLimit  *Ratio                                                                  `json:"presentationLowLimit,omitempty"`
	Concentration         *Ratio                                                                  `json:"concentration,omitempty"`
	ConcentrationLowLimit *Ratio                                                                  `json:"concentrationLowLimit,omitempty"`
	MeasurementPoint      *string                                                                 `json:"measurementPoint,omitempty"`
	Country               []CodeableConcept                                                       `json:"country,omitempty"`
	ReferenceStrength     []MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength `json:"referenceStrength,omitempty"`
}

// Strength expressed in terms of a reference substance.
type MedicinalProductIngredientSpecifiedSubstanceStrengthReferenceStrength struct {
	ID                *string           `json:"id,omitempty"`
	Extension         []Extension       `json:"extension,omitempty"`
	ModifierExtension []Extension       `json:"modifierExtension,omitempty"`
	Substance         *CodeableConcept  `json:"substance,omitempty"`
	Strength          Ratio             `json:"strength"`
	StrengthLowLimit  *Ratio            `json:"strengthLowLimit,omitempty"`
	MeasurementPoint  *string           `json:"measurementPoint,omitempty"`
	Country           []CodeableConcept `json:"country,omitempty"`
}

// The ingredient substance.
type MedicinalProductIngredientSubstance struct {
	ID                *string                                                `json:"id,omitempty"`
	Extension         []Extension                                            `json:"extension,omitempty"`
	ModifierExtension []Extension                                            `json:"modifierExtension,omitempty"`
	Code              CodeableConcept                                        `json:"code"`
	Strength          []MedicinalProductIngredientSpecifiedSubstanceStrength `json:"strength,omitempty"`
}

// This function returns resource reference information
func (r MedicinalProductIngredient) ResourceRef() (string, *string) {
	return "MedicinalProductIngredient", r.ID
}

// This function returns resource reference information
func (r MedicinalProductIngredient) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedicinalProductIngredient MedicinalProductIngredient

// MarshalJSON marshals the given MedicinalProductIngredient as JSON into a byte slice
func (r MedicinalProductIngredient) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductIngredient
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductIngredient: OtherMedicinalProductIngredient(r),
		ResourceType:                    "MedicinalProductIngredient",
	})
}

// UnmarshalMedicinalProductIngredient unmarshals a MedicinalProductIngredient.
func UnmarshalMedicinalProductIngredient(b []byte) (MedicinalProductIngredient, error) {
	var medicinalProductIngredient MedicinalProductIngredient
	if err := json.Unmarshal(b, &medicinalProductIngredient); err != nil {
		return medicinalProductIngredient, err
	}
	return medicinalProductIngredient, nil
}
