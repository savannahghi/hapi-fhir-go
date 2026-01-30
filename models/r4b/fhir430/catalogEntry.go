package fhir430

import "encoding/json"

// CatalogEntry is documented here http://hl7.org/fhir/StructureDefinition/CatalogEntry
// Catalog entries are wrappers that contextualize items included in a catalog.
type CatalogEntry struct {
	ID                       *string                    `json:"id,omitempty"`
	Meta                     *Meta                      `json:"meta,omitempty"`
	ImplicitRules            *string                    `json:"implicitRules,omitempty"`
	Language                 *string                    `json:"language,omitempty"`
	Text                     *Narrative                 `json:"text,omitempty"`
	Contained                []json.RawMessage          `json:"contained,omitempty"`
	Extension                []Extension                `json:"extension,omitempty"`
	ModifierExtension        []Extension                `json:"modifierExtension,omitempty"`
	Identifier               []Identifier               `json:"identifier,omitempty"`
	Type                     *CodeableConcept           `json:"type,omitempty"`
	Orderable                bool                       `json:"orderable"`
	ReferencedItem           Reference                  `json:"referencedItem"`
	AdditionalIdentifier     []Identifier               `json:"additionalIdentifier,omitempty"`
	Classification           []CodeableConcept          `json:"classification,omitempty"`
	Status                   *PublicationStatus         `json:"status,omitempty"`
	ValidityPeriod           *Period                    `json:"validityPeriod,omitempty"`
	ValidTo                  *string                    `json:"validTo,omitempty"`
	LastUpdated              *string                    `json:"lastUpdated,omitempty"`
	AdditionalCharacteristic []CodeableConcept          `json:"additionalCharacteristic,omitempty"`
	AdditionalClassification []CodeableConcept          `json:"additionalClassification,omitempty"`
	RelatedEntry             []CatalogEntryRelatedEntry `json:"relatedEntry,omitempty"`
}

// Used for example, to point to a substance, or to a device used to administer a medication.
type CatalogEntryRelatedEntry struct {
	ID                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Relationtype      CatalogEntryRelationType `json:"relationtype"`
	Item              Reference                `json:"item"`
}

// This function returns resource reference information
func (r CatalogEntry) ResourceRef() (string, *string) {
	return "CatalogEntry", r.ID
}

// This function returns resource reference information
func (r CatalogEntry) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCatalogEntry CatalogEntry

// MarshalJSON marshals the given CatalogEntry as JSON into a byte slice
func (r CatalogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCatalogEntry
		ResourceType string `json:"resourceType"`
	}{
		OtherCatalogEntry: OtherCatalogEntry(r),
		ResourceType:      "CatalogEntry",
	})
}

// UnmarshalCatalogEntry unmarshals a CatalogEntry.
func UnmarshalCatalogEntry(b []byte) (CatalogEntry, error) {
	var catalogEntry CatalogEntry
	if err := json.Unmarshal(b, &catalogEntry); err != nil {
		return catalogEntry, err
	}
	return catalogEntry, nil
}
