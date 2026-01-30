package fhir430

import "encoding/json"

// SupplyDelivery is documented here http://hl7.org/fhir/StructureDefinition/SupplyDelivery
// Record of delivery of what is supplied.
type SupplyDelivery struct {
	ID                 *string                     `json:"id,omitempty"`
	Meta               *Meta                       `json:"meta,omitempty"`
	ImplicitRules      *string                     `json:"implicitRules,omitempty"`
	Language           *string                     `json:"language,omitempty"`
	Text               *Narrative                  `json:"text,omitempty"`
	Contained          []json.RawMessage           `json:"contained,omitempty"`
	Extension          []Extension                 `json:"extension,omitempty"`
	ModifierExtension  []Extension                 `json:"modifierExtension,omitempty"`
	Identifier         []Identifier                `json:"identifier,omitempty"`
	BasedOn            []Reference                 `json:"basedOn,omitempty"`
	PartOf             []Reference                 `json:"partOf,omitempty"`
	Status             *SupplyDeliveryStatus       `json:"status,omitempty"`
	Patient            *Reference                  `json:"patient,omitempty"`
	Type               *CodeableConcept            `json:"type,omitempty"`
	SuppliedItem       *SupplyDeliverySuppliedItem `json:"suppliedItem,omitempty"`
	OccurrenceDateTime *string                     `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod   *Period                     `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming   *Timing                     `json:"occurrenceTiming,omitempty"`
	Supplier           *Reference                  `json:"supplier,omitempty"`
	Destination        *Reference                  `json:"destination,omitempty"`
	Receiver           []Reference                 `json:"receiver,omitempty"`
}

// The item that is being delivered or has been supplied.
type SupplyDeliverySuppliedItem struct {
	ID                  *string          `json:"id,omitempty"`
	Extension           []Extension      `json:"extension,omitempty"`
	ModifierExtension   []Extension      `json:"modifierExtension,omitempty"`
	Quantity            *Quantity        `json:"quantity,omitempty"`
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `json:"itemReference,omitempty"`
}

// This function returns resource reference information
func (r SupplyDelivery) ResourceRef() (string, *string) {
	return "SupplyDelivery", r.ID
}

// This function returns resource reference information
func (r SupplyDelivery) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherSupplyDelivery SupplyDelivery

// MarshalJSON marshals the given SupplyDelivery as JSON into a byte slice
func (r SupplyDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSupplyDelivery
		ResourceType string `json:"resourceType"`
	}{
		OtherSupplyDelivery: OtherSupplyDelivery(r),
		ResourceType:        "SupplyDelivery",
	})
}

// UnmarshalSupplyDelivery unmarshals a SupplyDelivery.
func UnmarshalSupplyDelivery(b []byte) (SupplyDelivery, error) {
	var supplyDelivery SupplyDelivery
	if err := json.Unmarshal(b, &supplyDelivery); err != nil {
		return supplyDelivery, err
	}
	return supplyDelivery, nil
}
