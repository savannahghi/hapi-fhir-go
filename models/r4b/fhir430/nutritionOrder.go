package fhir430

import "encoding/json"

// NutritionOrder is documented here http://hl7.org/fhir/StructureDefinition/NutritionOrder
// A request to supply a diet, formula feeding (enteral) or oral nutritional supplement to a patient/resident.
type NutritionOrder struct {
	ID                     *string                       `json:"id,omitempty"`
	Meta                   *Meta                         `json:"meta,omitempty"`
	ImplicitRules          *string                       `json:"implicitRules,omitempty"`
	Language               *string                       `json:"language,omitempty"`
	Text                   *Narrative                    `json:"text,omitempty"`
	Contained              []json.RawMessage             `json:"contained,omitempty"`
	Extension              []Extension                   `json:"extension,omitempty"`
	ModifierExtension      []Extension                   `json:"modifierExtension,omitempty"`
	Identifier             []Identifier                  `json:"identifier,omitempty"`
	InstantiatesCanonical  []string                      `json:"instantiatesCanonical,omitempty"`
	InstantiatesUri        []string                      `json:"instantiatesUri,omitempty"`
	Instantiates           []string                      `json:"instantiates,omitempty"`
	Status                 RequestStatus                 `json:"status"`
	Intent                 RequestIntent                 `json:"intent"`
	Patient                Reference                     `json:"patient"`
	Encounter              *Reference                    `json:"encounter,omitempty"`
	DateTime               string                        `json:"dateTime"`
	Orderer                *Reference                    `json:"orderer,omitempty"`
	AllergyIntolerance     []Reference                   `json:"allergyIntolerance,omitempty"`
	FoodPreferenceModifier []CodeableConcept             `json:"foodPreferenceModifier,omitempty"`
	ExcludeFoodModifier    []CodeableConcept             `json:"excludeFoodModifier,omitempty"`
	OralDiet               *NutritionOrderOralDiet       `json:"oralDiet,omitempty"`
	Supplement             []NutritionOrderSupplement    `json:"supplement,omitempty"`
	EnteralFormula         *NutritionOrderEnteralFormula `json:"enteralFormula,omitempty"`
	Note                   []Annotation                  `json:"note,omitempty"`
}

// Diet given orally in contrast to enteral (tube) feeding.
type NutritionOrderOralDiet struct {
	ID                   *string                          `json:"id,omitempty"`
	Extension            []Extension                      `json:"extension,omitempty"`
	ModifierExtension    []Extension                      `json:"modifierExtension,omitempty"`
	Type                 []CodeableConcept                `json:"type,omitempty"`
	Schedule             []Timing                         `json:"schedule,omitempty"`
	Nutrient             []NutritionOrderOralDietNutrient `json:"nutrient,omitempty"`
	Texture              []NutritionOrderOralDietTexture  `json:"texture,omitempty"`
	FluidConsistencyType []CodeableConcept                `json:"fluidConsistencyType,omitempty"`
	Instruction          *string                          `json:"instruction,omitempty"`
}

// Class that defines the quantity and type of nutrient modifications (for example carbohydrate, fiber or sodium) required for the oral diet.
type NutritionOrderOralDietNutrient struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Modifier          *CodeableConcept `json:"modifier,omitempty"`
	Amount            *Quantity        `json:"amount,omitempty"`
}

// Class that describes any texture modifications required for the patient to safely consume various types of solid foods.
type NutritionOrderOralDietTexture struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Modifier          *CodeableConcept `json:"modifier,omitempty"`
	FoodType          *CodeableConcept `json:"foodType,omitempty"`
}

// Oral nutritional products given in order to add further nutritional value to the patient's diet.
type NutritionOrderSupplement struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `json:"type,omitempty"`
	ProductName       *string          `json:"productName,omitempty"`
	Schedule          []Timing         `json:"schedule,omitempty"`
	Quantity          *Quantity        `json:"quantity,omitempty"`
	Instruction       *string          `json:"instruction,omitempty"`
}

// Feeding provided through the gastrointestinal tract via a tube, catheter, or stoma that delivers nutrition distal to the oral cavity.
type NutritionOrderEnteralFormula struct {
	ID                        *string                                      `json:"id,omitempty"`
	Extension                 []Extension                                  `json:"extension,omitempty"`
	ModifierExtension         []Extension                                  `json:"modifierExtension,omitempty"`
	BaseFormulaType           *CodeableConcept                             `json:"baseFormulaType,omitempty"`
	BaseFormulaProductName    *string                                      `json:"baseFormulaProductName,omitempty"`
	AdditiveType              *CodeableConcept                             `json:"additiveType,omitempty"`
	AdditiveProductName       *string                                      `json:"additiveProductName,omitempty"`
	CaloricDensity            *Quantity                                    `json:"caloricDensity,omitempty"`
	RouteofAdministration     *CodeableConcept                             `json:"routeofAdministration,omitempty"`
	Administration            []NutritionOrderEnteralFormulaAdministration `json:"administration,omitempty"`
	MaxVolumeToDeliver        *Quantity                                    `json:"maxVolumeToDeliver,omitempty"`
	AdministrationInstruction *string                                      `json:"administrationInstruction,omitempty"`
}

// Formula administration instructions as structured data.  This repeating structure allows for changing the administration rate or volume over time for both bolus and continuous feeding.  An example of this would be an instruction to increase the rate of continuous feeding every 2 hours.
// See implementation notes below for further discussion on how to order continuous vs bolus enteral feeding using this resource.
type NutritionOrderEnteralFormulaAdministration struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Schedule          *Timing     `json:"schedule,omitempty"`
	Quantity          *Quantity   `json:"quantity,omitempty"`
	RateQuantity      *Quantity   `json:"rateQuantity,omitempty"`
	RateRatio         *Ratio      `json:"rateRatio,omitempty"`
}

// This function returns resource reference information
func (r NutritionOrder) ResourceRef() (string, *string) {
	return "NutritionOrder", r.ID
}

// This function returns resource reference information
func (r NutritionOrder) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherNutritionOrder NutritionOrder

// MarshalJSON marshals the given NutritionOrder as JSON into a byte slice
func (r NutritionOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherNutritionOrder
		ResourceType string `json:"resourceType"`
	}{
		OtherNutritionOrder: OtherNutritionOrder(r),
		ResourceType:        "NutritionOrder",
	})
}

// UnmarshalNutritionOrder unmarshals a NutritionOrder.
func UnmarshalNutritionOrder(b []byte) (NutritionOrder, error) {
	var nutritionOrder NutritionOrder
	if err := json.Unmarshal(b, &nutritionOrder); err != nil {
		return nutritionOrder, err
	}
	return nutritionOrder, nil
}
