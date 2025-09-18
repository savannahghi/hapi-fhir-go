
package fhir430

import "encoding/json"
// Invoice is documented here http://hl7.org/fhir/StructureDefinition/Invoice
// Invoice containing collected ChargeItems from an Account with calculated individual and total price for Billing purpose.
type Invoice struct {
	ID                  *string                         `json:"ID,omitempty"`
	Meta                *Meta                           `json:"meta,omitempty"`
	ImplicitRules       *string                         `json:"implicitRules,omitempty"`
	Language            *string                         `json:"language,omitempty"`
	Text                *Narrative                      `json:"text,omitempty"`
	Contained           []json.RawMessage               `json:"contained,omitempty"`
	Extension           []Extension                     `json:"extension,omitempty"`
	ModifierExtension   []Extension                     `json:"modifierExtension,omitempty"`
	Identifier          []Identifier                    `json:"identifier,omitempty"`
	Status              InvoiceStatus                   `json:"status"`
	CancelledReason     *string                         `json:"cancelledReason,omitempty"`
	Type                *CodeableConcept                `json:"type,omitempty"`
	Subject             *Reference                      `json:"subject,omitempty"`
	Recipient           *Reference                      `json:"recipient,omitempty"`
	Date                *string                         `json:"date,omitempty"`
	Participant         []InvoiceParticipant            `json:"participant,omitempty"`
	Issuer              *Reference                      `json:"issuer,omitempty"`
	Account             *Reference                      `json:"account,omitempty"`
	LineItem            []InvoiceLineItem               `json:"lineItem,omitempty"`
	TotalPriceComponent []InvoiceLineItemPriceComponent `json:"totalPriceComponent,omitempty"`
	TotalNet            *Money                          `json:"totalNet,omitempty"`
	TotalGross          *Money                          `json:"totalGross,omitempty"`
	PaymentTerms        *string                         `json:"paymentTerms,omitempty"`
	Note                []Annotation                    `json:"note,omitempty"`
}

// Indicates who or what performed or participated in the charged service.
type InvoiceParticipant struct {
	ID                *string          `json:"ID,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Actor             Reference        `json:"actor"`
}

// Each line item represents one charge for goods and services rendered. Details such as date, code and amount are found in the referenced ChargeItem resource.
type InvoiceLineItem struct {
	ID                        *string                         `json:"ID,omitempty"`
	Extension                 []Extension                     `json:"extension,omitempty"`
	ModifierExtension         []Extension                     `json:"modifierExtension,omitempty"`
	Sequence                  *int                            `json:"sequence,omitempty"`
	ChargeItemReference       Reference                       `json:"chargeItemReference"`
	ChargeItemCodeableConcept CodeableConcept                 `json:"chargeItemCodeableConcept"`
	PriceComponent            []InvoiceLineItemPriceComponent `json:"priceComponent,omitempty"`
}

// The price for a ChargeItem may be calculated as a base price with surcharges/deductions that apply in certain conditions. A ChargeItemDefinition resource that defines the prices, factors and conditions that apply to a billing code is currently under development. The priceComponent element can be used to offer transparency to the recipient of the Invoice as to how the prices have been calculated.
type InvoiceLineItemPriceComponent struct {
	ID                *string                   `json:"ID,omitempty"`
	Extension         []Extension               `json:"extension,omitempty"`
	ModifierExtension []Extension               `json:"modifierExtension,omitempty"`
	Type              InvoicePriceComponentType `json:"type"`
	Code              *CodeableConcept          `json:"code,omitempty"`
	Factor            *json.Number              `json:"factor,omitempty"`
	Amount            *Money                    `json:"amount,omitempty"`
}

// This function returns resource reference information
func (r Invoice) ResourceRef() (string, *string) {
	return "Invoice", r.ID
}

// This function returns resource reference information
func (r Invoice) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherInvoice Invoice

// MarshalJSON marshals the given Invoice as JSON into a byte slice
func (r Invoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherInvoice
		ResourceType string `json:"resourceType"`
	}{
		OtherInvoice: OtherInvoice(r),
		ResourceType: "Invoice",
	})
}

// UnmarshalInvoice unmarshals a Invoice.
func UnmarshalInvoice(b []byte) (Invoice, error) {
	var invoice Invoice
	if err := json.Unmarshal(b, &invoice); err != nil {
		return invoice, err
	}
	return invoice, nil
}
