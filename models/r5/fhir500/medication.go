package fhir500

import (
	"encoding/json"
)

// Medication is documented here http://hl7.org/fhir/StructureDefinition/Medication
type Medication struct {
	ID                           *string                `json:"id,omitempty"`
	Meta                         *Meta                  `json:"meta,omitempty"`
	ImplicitRules                *string                `json:"implicitRules,omitempty"`
	Language                     *string                `json:"language,omitempty"`
	Text                         *Narrative             `json:"text,omitempty"`
	Extension                    []Extension            `json:"extension,omitempty"`
	ModifierExtension            []Extension            `json:"modifierExtension,omitempty"`
	Identifier                   []Identifier           `json:"identifier,omitempty"`
	Code                         *CodeableConcept       `json:"code,omitempty"`
	Status                       *MedicationStatusCodes `json:"status,omitempty"`
	MarketingAuthorizationHolder *Reference             `json:"marketingAuthorizationHolder,omitempty"`
	DoseForm                     *CodeableConcept       `json:"doseForm,omitempty"`
	TotalVolume                  *Quantity              `json:"totalVolume,omitempty"`
	Ingredient                   []MedicationIngredient `json:"ingredient,omitempty"`
	Batch                        *MedicationBatch       `json:"batch,omitempty"`
	Definition                   *Reference             `json:"definition,omitempty"`
}

type MedicationIngredient struct {
	ID                      *string           `json:"id,omitempty"`
	Extension               []Extension       `json:"extension,omitempty"`
	ModifierExtension       []Extension       `json:"modifierExtension,omitempty"`
	Item                    CodeableReference `json:"item"`
	IsActive                *bool             `json:"isActive,omitempty"`
	StrengthRatio           *Ratio            `json:"strengthRatio,omitempty"`
	StrengthCodeableConcept *CodeableConcept  `json:"strengthCodeableConcept,omitempty"`
	StrengthQuantity        *Quantity         `json:"strengthQuantity,omitempty"`
}

type MedicationBatch struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	LotNumber         *string     `json:"lotNumber,omitempty"`
	ExpirationDate    *string     `json:"expirationDate,omitempty"`
}

type OtherMedication Medication

// MarshalJSON marshals the given Medication as JSON into a byte slice.
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

// MedicationStatusCodes is documented here http://hl7.org/fhir/ValueSet/medication-status
type MedicationStatusCodes string

const (
	MedicationStatusCodesActive         MedicationStatusCodes = "active"
	MedicationStatusCodesInactive       MedicationStatusCodes = "inactive"
	MedicationStatusCodesEnteredInError MedicationStatusCodes = "entered-in-error"
)

func (code MedicationStatusCodes) String() string {
	return code.Code()
}

func (code MedicationStatusCodes) Code() string {
	switch code {
	case MedicationStatusCodesActive:
		return "active"
	case MedicationStatusCodesInactive:
		return "inactive"
	case MedicationStatusCodesEnteredInError:
		return "entered-in-error"
	}

	return "<unknown>"
}
