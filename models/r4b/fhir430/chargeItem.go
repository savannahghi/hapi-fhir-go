
package fhir430

import "encoding/json"
// ChargeItem is documented here http://hl7.org/fhir/StructureDefinition/ChargeItem
// The resource ChargeItem describes the provision of healthcare provider products for a certain patient, therefore referring not only to the product, but containing in addition details of the provision, like date, time, amounts and participating organizations and persons. Main Usage of the ChargeItem is to enable the billing process and internal cost allocation.
type ChargeItem struct {
	ID                     *string               `json:"ID,omitempty"`
	Meta                   *Meta                 `json:"meta,omitempty"`
	ImplicitRules          *string               `json:"implicitRules,omitempty"`
	Language               *string               `json:"language,omitempty"`
	Text                   *Narrative            `json:"text,omitempty"`
	Contained              []json.RawMessage     `json:"contained,omitempty"`
	Extension              []Extension           `json:"extension,omitempty"`
	ModifierExtension      []Extension           `json:"modifierExtension,omitempty"`
	Identifier             []Identifier          `json:"identifier,omitempty"`
	DefinitionUri          []string              `json:"definitionUri,omitempty"`
	DefinitionCanonical    []string              `json:"definitionCanonical,omitempty"`
	Status                 ChargeItemStatus      `json:"status"`
	PartOf                 []Reference           `json:"partOf,omitempty"`
	Code                   CodeableConcept       `json:"code"`
	Subject                Reference             `json:"subject"`
	Context                *Reference            `json:"context,omitempty"`
	OccurrenceDateTime     *string               `json:"occurrenceDateTime,omitempty"`
	OccurrencePeriod       *Period               `json:"occurrencePeriod,omitempty"`
	OccurrenceTiming       *Timing               `json:"occurrenceTiming,omitempty"`
	Performer              []ChargeItemPerformer `json:"performer,omitempty"`
	PerformingOrganization *Reference            `json:"performingOrganization,omitempty"`
	RequestingOrganization *Reference            `json:"requestingOrganization,omitempty"`
	CostCenter             *Reference            `json:"costCenter,omitempty"`
	Quantity               *Quantity             `json:"quantity,omitempty"`
	Bodysite               []CodeableConcept     `json:"bodysite,omitempty"`
	FactorOverride         *json.Number          `json:"factorOverride,omitempty"`
	PriceOverride          *Money                `json:"priceOverride,omitempty"`
	OverrideReason         *string               `json:"overrideReason,omitempty"`
	Enterer                *Reference            `json:"enterer,omitempty"`
	EnteredDate            *string               `json:"enteredDate,omitempty"`
	Reason                 []CodeableConcept     `json:"reason,omitempty"`
	Service                []Reference           `json:"service,omitempty"`
	ProductReference       *Reference            `json:"productReference,omitempty"`
	ProductCodeableConcept *CodeableConcept      `json:"productCodeableConcept,omitempty"`
	Account                []Reference           `json:"account,omitempty"`
	Note                   []Annotation          `json:"note,omitempty"`
	SupportingInformation  []Reference           `json:"supportingInformation,omitempty"`
}

// Indicates who or what performed or participated in the charged service.
type ChargeItemPerformer struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Function          *CodeableConcept `json:"function,omitempty"`
	Actor             Reference        `json:"actor"`
}

// This function returns resource reference information
func (r ChargeItem) ResourceRef() (string, *string) {
	return "ChargeItem", r.ID
}

// This function returns resource reference information
func (r ChargeItem) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherChargeItem ChargeItem

// MarshalJSON marshals the given ChargeItem as JSON into a byte slice
func (r ChargeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherChargeItem
		ResourceType string `json:"resourceType"`
	}{
		OtherChargeItem: OtherChargeItem(r),
		ResourceType:    "ChargeItem",
	})
}

// UnmarshalChargeItem unmarshals a ChargeItem.
func UnmarshalChargeItem(b []byte) (ChargeItem, error) {
	var chargeItem ChargeItem
	if err := json.Unmarshal(b, &chargeItem); err != nil {
		return chargeItem, err
	}
	return chargeItem, nil
}
