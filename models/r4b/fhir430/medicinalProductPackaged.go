package fhir430

import "encoding/json"

// MedicinalProductPackaged is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductPackaged
// A medicinal product in a container or package.
type MedicinalProductPackaged struct {
	ID                     *string                                   `json:"id,omitempty"`
	Meta                   *Meta                                     `json:"meta,omitempty"`
	ImplicitRules          *string                                   `json:"implicitRules,omitempty"`
	Language               *string                                   `json:"language,omitempty"`
	Text                   *Narrative                                `json:"text,omitempty"`
	Contained              []json.RawMessage                         `json:"contained,omitempty"`
	Extension              []Extension                               `json:"extension,omitempty"`
	ModifierExtension      []Extension                               `json:"modifierExtension,omitempty"`
	Identifier             []Identifier                              `json:"identifier,omitempty"`
	Subject                []Reference                               `json:"subject,omitempty"`
	Description            *string                                   `json:"description,omitempty"`
	LegalStatusOfSupply    *CodeableConcept                          `json:"legalStatusOfSupply,omitempty"`
	MarketingStatus        []MarketingStatus                         `json:"marketingStatus,omitempty"`
	MarketingAuthorization *Reference                                `json:"marketingAuthorization,omitempty"`
	Manufacturer           []Reference                               `json:"manufacturer,omitempty"`
	BatchIdentifier        []MedicinalProductPackagedBatchIdentifier `json:"batchIdentifier,omitempty"`
	PackageItem            []MedicinalProductPackagedPackageItem     `json:"packageItem"`
}

// Batch numbering.
type MedicinalProductPackagedBatchIdentifier struct {
	ID                 *string     `json:"id,omitempty"`
	Extension          []Extension `json:"extension,omitempty"`
	ModifierExtension  []Extension `json:"modifierExtension,omitempty"`
	OuterPackaging     Identifier  `json:"outerPackaging"`
	ImmediatePackaging *Identifier `json:"immediatePackaging,omitempty"`
}

// A packaging item, as a contained for medicine, possibly with other packaging items within.
type MedicinalProductPackagedPackageItem struct {
	ID                      *string                               `json:"id,omitempty"`
	Extension               []Extension                           `json:"extension,omitempty"`
	ModifierExtension       []Extension                           `json:"modifierExtension,omitempty"`
	Identifier              []Identifier                          `json:"identifier,omitempty"`
	Type                    CodeableConcept                       `json:"type"`
	Quantity                Quantity                              `json:"quantity"`
	Material                []CodeableConcept                     `json:"material,omitempty"`
	AlternateMaterial       []CodeableConcept                     `json:"alternateMaterial,omitempty"`
	Device                  []Reference                           `json:"device,omitempty"`
	ManufacturedItem        []Reference                           `json:"manufacturedItem,omitempty"`
	PackageItem             []MedicinalProductPackagedPackageItem `json:"packageItem,omitempty"`
	PhysicalCharacteristics *ProdCharacteristic                   `json:"physicalCharacteristics,omitempty"`
	OtherCharacteristics    []CodeableConcept                     `json:"otherCharacteristics,omitempty"`
	ShelfLifeStorage        []ProductShelfLife                    `json:"shelfLifeStorage,omitempty"`
	Manufacturer            []Reference                           `json:"manufacturer,omitempty"`
}

// This function returns resource reference information
func (r MedicinalProductPackaged) ResourceRef() (string, *string) {
	return "MedicinalProductPackaged", r.ID
}

// This function returns resource reference information
func (r MedicinalProductPackaged) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherMedicinalProductPackaged MedicinalProductPackaged

// MarshalJSON marshals the given MedicinalProductPackaged as JSON into a byte slice
func (r MedicinalProductPackaged) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductPackaged
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductPackaged: OtherMedicinalProductPackaged(r),
		ResourceType:                  "MedicinalProductPackaged",
	})
}

// UnmarshalMedicinalProductPackaged unmarshals a MedicinalProductPackaged.
func UnmarshalMedicinalProductPackaged(b []byte) (MedicinalProductPackaged, error) {
	var medicinalProductPackaged MedicinalProductPackaged
	if err := json.Unmarshal(b, &medicinalProductPackaged); err != nil {
		return medicinalProductPackaged, err
	}
	return medicinalProductPackaged, nil
}
