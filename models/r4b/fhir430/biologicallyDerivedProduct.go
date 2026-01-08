package fhir430

import "encoding/json"

// BiologicallyDerivedProduct is documented here http://hl7.org/fhir/StructureDefinition/BiologicallyDerivedProduct
/*
A material substance originating from a biological entity intended to be transplanted or infused
into another (possibly the same) biological entity.
*/
type BiologicallyDerivedProduct struct {
	ID                *string                                 `json:"id,omitempty"`
	Meta              *Meta                                   `json:"meta,omitempty"`
	ImplicitRules     *string                                 `json:"implicitRules,omitempty"`
	Language          *string                                 `json:"language,omitempty"`
	Text              *Narrative                              `json:"text,omitempty"`
	Contained         []json.RawMessage                       `json:"contained,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                            `json:"identifier,omitempty"`
	ProductCategory   *BiologicallyDerivedProductCategory     `json:"productCategory,omitempty"`
	ProductCode       *CodeableConcept                        `json:"productCode,omitempty"`
	Status            *BiologicallyDerivedProductStatus       `json:"status,omitempty"`
	Request           []Reference                             `json:"request,omitempty"`
	Quantity          *int                                    `json:"quantity,omitempty"`
	Parent            []Reference                             `json:"parent,omitempty"`
	Collection        *BiologicallyDerivedProductCollection   `json:"collection,omitempty"`
	Processing        []BiologicallyDerivedProductProcessing  `json:"processing,omitempty"`
	Manipulation      *BiologicallyDerivedProductManipulation `json:"manipulation,omitempty"`
	Storage           []BiologicallyDerivedProductStorage     `json:"storage,omitempty"`
}

// How this product was collected.
type BiologicallyDerivedProductCollection struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Collector         *Reference  `json:"collector,omitempty"`
	Source            *Reference  `json:"source,omitempty"`
	CollectedDateTime *string     `json:"collectedDateTime,omitempty"`
	CollectedPeriod   *Period     `json:"collectedPeriod,omitempty"`
}

// Any processing of the product during collection that does not change the fundamental nature of the product. For example adding anti-coagulants during the collection of Peripheral Blood Stem Cells.
type BiologicallyDerivedProductProcessing struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Description       *string          `json:"description,omitempty"`
	Procedure         *CodeableConcept `json:"procedure,omitempty"`
	Additive          *Reference       `json:"additive,omitempty"`
	TimeDateTime      *string          `json:"timeDateTime,omitempty"`
	TimePeriod        *Period          `json:"timePeriod,omitempty"`
}

// Any manipulation of product post-collection that is intended to alter the product.  For example a buffy-coat enrichment or CD8 reduction of Peripheral Blood Stem Cells to make it more suitable for infusion.
type BiologicallyDerivedProductManipulation struct {
	ID                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
	TimeDateTime      *string     `json:"timeDateTime,omitempty"`
	TimePeriod        *Period     `json:"timePeriod,omitempty"`
}

// Product storage.
type BiologicallyDerivedProductStorage struct {
	ID                *string                                 `json:"id,omitempty"`
	Extension         []Extension                             `json:"extension,omitempty"`
	ModifierExtension []Extension                             `json:"modifierExtension,omitempty"`
	Description       *string                                 `json:"description,omitempty"`
	Temperature       *json.Number                            `json:"temperature,omitempty"`
	Scale             *BiologicallyDerivedProductStorageScale `json:"scale,omitempty"`
	Duration          *Period                                 `json:"duration,omitempty"`
}

// This function returns resource reference information
func (r BiologicallyDerivedProduct) ResourceRef() (string, *string) {
	return "BiologicallyDerivedProduct", r.ID
}

// This function returns resource reference information
func (r BiologicallyDerivedProduct) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherBiologicallyDerivedProduct BiologicallyDerivedProduct

// MarshalJSON marshals the given BiologicallyDerivedProduct as JSON into a byte slice
func (r BiologicallyDerivedProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBiologicallyDerivedProduct
		ResourceType string `json:"resourceType"`
	}{
		OtherBiologicallyDerivedProduct: OtherBiologicallyDerivedProduct(r),
		ResourceType:                    "BiologicallyDerivedProduct",
	})
}

// UnmarshalBiologicallyDerivedProduct unmarshals a BiologicallyDerivedProduct.
func UnmarshalBiologicallyDerivedProduct(b []byte) (BiologicallyDerivedProduct, error) {
	var biologicallyDerivedProduct BiologicallyDerivedProduct
	if err := json.Unmarshal(b, &biologicallyDerivedProduct); err != nil {
		return biologicallyDerivedProduct, err
	}
	return biologicallyDerivedProduct, nil
}
