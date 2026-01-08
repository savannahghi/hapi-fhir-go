package fhir430

import "encoding/json"

// Medication is documented here http://hl7.org/fhir/StructureDefinition/Medication
// This resource is primarily used for the identification and definition of a medication for the purposes of prescribing, dispensing, and administering a medication as well as for making statements about medication use.
type Medication struct {
	ID                *string                `json:"id,omitempty"`
	Meta              *Meta                  `json:"meta,omitempty"`
	ImplicitRules     *string                `json:"implicitRules,omitempty"`
	Language          *string                `json:"language,omitempty"`
	Text              *Narrative             `json:"text,omitempty"`
	Contained         []json.RawMessage      `json:"contained,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Identifier        []Identifier           `json:"identifier,omitempty"`
	Code              *CodeableConcept       `json:"code,omitempty"`
	Status            *string                `json:"status,omitempty"`
	Manufacturer      *Reference             `json:"manufacturer,omitempty"`
	Form              *CodeableConcept       `json:"form,omitempty"`
	Amount            *Ratio                 `json:"amount,omitempty"`
	Ingredient        []MedicationIngredient `json:"ingredient,omitempty"`
	Batch             *MedicationBatch       `json:"batch,omitempty"`
}

// Identifies a particular constituent of interest in the product.
// The ingredients need not be a complete list.  If an ingredient is not specified, this does not indicate whether an ingredient is present or absent.  If an ingredient is specified it does not mean that all ingredients are specified.  It is possible to specify both inactive and active ingredients.
type MedicationIngredient struct {
	ID                  *string         `json:"id,omitempty"`
	Extension           []Extension     `json:"extension,omitempty"`
	ModifierExtension   []Extension     `json:"modifierExtension,omitempty"`
	ItemCodeableConcept CodeableConcept `json:"itemCodeableConcept"`
	ItemReference       Reference       `json:"itemReference"`
	IsActive            *bool           `json:"isActive,omitempty"`
	Strength            *Ratio          `json:"strength,omitempty"`
}

// Information that only applies to packages (not products).
type MedicationBatch struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	LotNumber         *string     `json:"lotNumber,omitempty"`
	ExpirationDate    *string     `json:"expirationDate,omitempty"`
}

// This function returns resource reference information
func (r Medication) ResourceRef() (string, *string) {
	return "Medication", r.ID
}

// This function returns resource reference information
func (r Medication) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedication Medication

// MarshalJSON marshals the given Medication as JSON into a byte slice
func (r Medication) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedication
		ResourceType string `json:"resourceType"`
	}{
		OtherMedication: OtherMedication(r),
		ResourceType:    "Medication",
	})
}

// UnmarshalMedication unmarshals a Medication.
func UnmarshalMedication(b []byte) (Medication, error) {
	var medication Medication
	if err := json.Unmarshal(b, &medication); err != nil {
		return medication, err
	}
	return medication, nil
}
