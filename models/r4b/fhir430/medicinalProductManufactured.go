
package fhir430

import "encoding/json"
// MedicinalProductManufactured is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductManufactured
// The manufactured item as contained in the packaged medicinal product.
type MedicinalProductManufactured struct {
	ID                      *string             `json:"ID,omitempty"`
	Meta                    *Meta               `json:"meta,omitempty"`
	ImplicitRules           *string             `json:"implicitRules,omitempty"`
	Language                *string             `json:"language,omitempty"`
	Text                    *Narrative          `json:"text,omitempty"`
	Contained               []json.RawMessage   `json:"contained,omitempty"`
	Extension               []Extension         `json:"extension,omitempty"`
	ModifierExtension       []Extension         `json:"modifierExtension,omitempty"`
	ManufacturedDoseForm    CodeableConcept     `json:"manufacturedDoseForm"`
	UnitOfPresentation      *CodeableConcept    `json:"unitOfPresentation,omitempty"`
	Quantity                Quantity            `json:"quantity"`
	Manufacturer            []Reference         `json:"manufacturer,omitempty"`
	Ingredient              []Reference         `json:"ingredient,omitempty"`
	PhysicalCharacteristics *ProdCharacteristic `json:"physicalCharacteristics,omitempty"`
	OtherCharacteristics    []CodeableConcept   `json:"otherCharacteristics,omitempty"`
}

// This function returns resource reference information
func (r MedicinalProductManufactured) ResourceRef() (string, *string) {
	return "MedicinalProductManufactured", r.ID
}

// This function returns resource reference information
func (r MedicinalProductManufactured) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedicinalProductManufactured MedicinalProductManufactured

// MarshalJSON marshals the given MedicinalProductManufactured as JSON into a byte slice
func (r MedicinalProductManufactured) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductManufactured
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductManufactured: OtherMedicinalProductManufactured(r),
		ResourceType:                      "MedicinalProductManufactured",
	})
}

// UnmarshalMedicinalProductManufactured unmarshals a MedicinalProductManufactured.
func UnmarshalMedicinalProductManufactured(b []byte) (MedicinalProductManufactured, error) {
	var medicinalProductManufactured MedicinalProductManufactured
	if err := json.Unmarshal(b, &medicinalProductManufactured); err != nil {
		return medicinalProductManufactured, err
	}
	return medicinalProductManufactured, nil
}
