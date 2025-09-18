
package fhir430

import "encoding/json"
// SupplyRequest is documented here http://hl7.org/fhir/StructureDefinition/SupplyRequest
// A record of a request for a medication, substance or device used in the healthcare setting.
type SupplyRequest struct {
	ID                  *string                  `json:"ID,omitempty"`
	Meta                *Meta                    `json:"meta,omitempty"`
	ImplicitRules       *string                  `json:"implicitRules,omitempty"`
	Language            *string                  `json:"language,omitempty"`
	Text                *Narrative               `json:"text,omitempty"`
	Contained           []json.RawMessage        `json:"contained,omitempty"`
	Extension           []Extension              `json:"extension,omitempty"`
	ModifierExtension   []Extension              `json:"modifierExtension,omitempty"`
	Identifier          []Identifier             `json:"identifier,omitempty"`
	Status              *SupplyRequestStatus     `json:"status,omitempty"`
	Category            *CodeableConcept         `json:"category,omitempty"`
	Priority            *RequestPriority         `json:"priority,omitempty"`
	ItemCodeableConcept CodeableConcept          `json:"itemCodeableConcept"`
	ItemReference       Reference                `json:"itemReference"`
	Quantity            Quantity                 `json:"quantity"`
	Parameter           []SupplyRequestParameter `json:"parameter,omitempty"`
	OccurrenceDateTime  *string                  `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod    *Period                  `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming    *Timing                  `json:"occurrenceTiming,omitempty"`
	AuthoredOn          *string                  `json:"authoredOn,omitempty"`
	Requester           *Reference               `json:"requester,omitempty"`
	Supplier            []Reference              `json:"supplier,omitempty"`
	ReasonCode          []CodeableConcept        `json:"reasonCode,omitempty"`
	ReasonReference     []Reference              `json:"reasonReference,omitempty"`
	DeliverFrom         *Reference               `json:"deliverFrom,omitempty"`
	DeliverTo           *Reference               `json:"deliverTo,omitempty"`
}

// Specific parameters for the ordered item.  For example, the size of the indicated item.
type SupplyRequestParameter struct {
	ID                   *string          `json:"ID,omitempty"`
	Extension            []Extension      `json:"extension,omitempty"`
	ModifierExtension    []Extension      `json:"modifierExtension,omitempty"`
	Code                 *CodeableConcept `json:"code,omitempty"`
	ValueCodeableConcept *CodeableConcept `json:"valueCodeableConcept,omitempty"`
	ValueQuantity        *Quantity        `json:"valueQuantity,omitempty"`
	ValueRange           *Range           `json:"valueRange,omitempty"`
	ValueBoolean         *bool            `json:"valueBoolean,omitempty"`
}

// This function returns resource reference information
func (r SupplyRequest) ResourceRef() (string, *string) {
	return "SupplyRequest", r.ID
}

// This function returns resource reference information
func (r SupplyRequest) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherSupplyRequest SupplyRequest

// MarshalJSON marshals the given SupplyRequest as JSON into a byte slice
func (r SupplyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSupplyRequest
		ResourceType string `json:"resourceType"`
	}{
		OtherSupplyRequest: OtherSupplyRequest(r),
		ResourceType:       "SupplyRequest",
	})
}

// UnmarshalSupplyRequest unmarshals a SupplyRequest.
func UnmarshalSupplyRequest(b []byte) (SupplyRequest, error) {
	var supplyRequest SupplyRequest
	if err := json.Unmarshal(b, &supplyRequest); err != nil {
		return supplyRequest, err
	}
	return supplyRequest, nil
}
