
package fhir430

import "encoding/json"
// ChargeItemDefinition is documented here http://hl7.org/fhir/StructureDefinition/ChargeItemDefinition
// The ChargeItemDefinition resource provides the properties that apply to the (billing) codes necessary to calculate costs and prices. The properties may differ largely depending on type and realm, therefore this resource gives only a rough structure and requires profiling for each type of billing code system.
type ChargeItemDefinition struct {
	ID                *string                             `json:"ID,omitempty"`
	Meta              *Meta                               `json:"meta,omitempty"`
	ImplicitRules     *string                             `json:"implicitRules,omitempty"`
	Language          *string                             `json:"language,omitempty"`
	Text              *Narrative                          `json:"text,omitempty"`
	Contained         []json.RawMessage                   `json:"contained,omitempty"`
	Extension         []Extension                         `json:"extension,omitempty"`
	ModifierExtension []Extension                         `json:"modifierExtension,omitempty"`
	Url               string                              `json:"url"`
	Identifier        []Identifier                        `json:"identifier,omitempty"`
	Version           *string                             `json:"version,omitempty"`
	Title             *string                             `json:"title,omitempty"`
	DerivedFromUri    []string                            `json:"derivedFromUri,omitempty"`
	PartOf            []string                            `json:"partOf,omitempty"`
	Replaces          []string                            `json:"replaces,omitempty"`
	Status            PublicationStatus                   `json:"status"`
	Experimental      *bool                               `json:"experimental,omitempty"`
	Date              *string                             `json:"date,omitempty"`
	Publisher         *string                             `json:"publisher,omitempty"`
	Contact           []ContactDetail                     `json:"contact,omitempty"`
	Description       *string                             `json:"description,omitempty"`
	UseContext        []UsageContext                      `json:"useContext,omitempty"`
	Jurisdiction      []CodeableConcept                   `json:"jurisdiction,omitempty"`
	ApprovalDate      *string                             `json:"approvalDate,omitempty"`
	LastReviewDate    *string                             `json:"lastReviewDate,omitempty"`
	EffectivePeriod   *Period                             `json:"effectivePeriod,omitempty"`
	Code              *CodeableConcept                    `json:"code,omitempty"`
	Instance          []Reference                         `json:"instance,omitempty"`
	Applicability     []ChargeItemDefinitionApplicability `json:"applicability,omitempty"`
	PropertyGroup     []ChargeItemDefinitionPropertyGroup `json:"propertyGroup,omitempty"`
}

// Expressions that describe applicability criteria for the billing code.
// The applicability conditions can be used to ascertain whether a billing item is allowed in a specific context. E.g. some billing codes may only be applicable in out-patient settings, only to male/female patients or only to children.
type ChargeItemDefinitionApplicability struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Description       *string     `json:"description,omitempty"`
	Language          *string     `json:"language,omitempty"`
	Expression        *string     `json:"expression,omitempty"`
}

// Group of properties which are applicable under the same conditions. If no applicability rules are established for the group, then all properties always apply.
type ChargeItemDefinitionPropertyGroup struct {
	ID                *string                                           `json:"ID,omitempty"`
	Extension         []Extension                                       `json:"extension,omitempty"`
	ModifierExtension []Extension                                       `json:"modifierExtension,omitempty"`
	Applicability     []ChargeItemDefinitionApplicability               `json:"applicability,omitempty"`
	PriceComponent    []ChargeItemDefinitionPropertyGroupPriceComponent `json:"priceComponent,omitempty"`
}

// The price for a ChargeItem may be calculated as a base price with surcharges/deductions that apply in certain conditions. A ChargeItemDefinition resource that defines the prices, factors and conditions that apply to a billing code is currently under development. The priceComponent element can be used to offer transparency to the recipient of the Invoice of how the prices have been calculated.
type ChargeItemDefinitionPropertyGroupPriceComponent struct {
	ID                *string                   `json:"ID,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	Type              InvoicePriceComponentType `json:"type"`
	Code              *CodeableConcept          `json:"code,omitempty"`
	Factor            *json.Number              `json:"factor,omitempty"`
	Amount            *Money                    `json:"amount,omitempty"`
}

// This function returns resource reference information
func (r ChargeItemDefinition) ResourceRef() (string, *string) {
	return "ChargeItemDefinition", r.ID
}

// This function returns resource reference information
func (r ChargeItemDefinition) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherChargeItemDefinition ChargeItemDefinition

// MarshalJSON marshals the given ChargeItemDefinition as JSON into a byte slice
func (r ChargeItemDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherChargeItemDefinition
		ResourceType string `json:"resourceType"`
	}{
		OtherChargeItemDefinition: OtherChargeItemDefinition(r),
		ResourceType:              "ChargeItemDefinition",
	})
}

// UnmarshalChargeItemDefinition unmarshals a ChargeItemDefinition.
func UnmarshalChargeItemDefinition(b []byte) (ChargeItemDefinition, error) {
	var chargeItemDefinition ChargeItemDefinition
	if err := json.Unmarshal(b, &chargeItemDefinition); err != nil {
		return chargeItemDefinition, err
	}
	return chargeItemDefinition, nil
}
