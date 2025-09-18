
package fhir430

import "encoding/json"
// PaymentNotice is documented here http://hl7.org/fhir/StructureDefinition/PaymentNotice
// This resource provides the status of the payment for goods and services rendered, and the request and response resource references.
type PaymentNotice struct {
	ID                *string                      `json:"ID,omitempty"`
	Meta              *Meta                        `json:"meta,omitempty"`
	ImplicitRules     *string                      `json:"implicitRules,omitempty"`
	Language          *string                      `json:"language,omitempty"`
	Text              *Narrative                   `json:"text,omitempty"`
	Contained         []json.RawMessage            `json:"contained,omitempty"`
	Extension         []Extension                  `json:"extension,omitempty"`
	ModifierExtension []Extension                  `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                 `json:"identifier,omitempty"`
	Status            FinancialResourceStatusCodes `json:"status"`
	Request           *Reference                   `json:"request,omitempty"`
	Response          *Reference                   `json:"response,omitempty"`
	Created           string                       `json:"created"`
	Provider          *Reference                   `json:"provider,omitempty"`
	Payment           Reference                    `json:"payment"`
	PaymentDate       *string                      `json:"paymentDate,omitempty"`
	Payee             *Reference                   `json:"payee,omitempty"`
	Recipient         Reference                    `json:"recipient"`
	Amount            Money                        `json:"amount"`
	PaymentStatus     *CodeableConcept             `json:"paymentStatus,omitempty"`
}

// This function returns resource reference information
func (r PaymentNotice) ResourceRef() (string, *string) {
	return "PaymentNotice", r.ID
}

// This function returns resource reference information
func (r PaymentNotice) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherPaymentNotice PaymentNotice

// MarshalJSON marshals the given PaymentNotice as JSON into a byte slice
func (r PaymentNotice) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPaymentNotice
		ResourceType string `json:"resourceType"`
	}{
		OtherPaymentNotice: OtherPaymentNotice(r),
		ResourceType:       "PaymentNotice",
	})
}

// UnmarshalPaymentNotice unmarshals a PaymentNotice.
func UnmarshalPaymentNotice(b []byte) (PaymentNotice, error) {
	var paymentNotice PaymentNotice
	if err := json.Unmarshal(b, &paymentNotice); err != nil {
		return paymentNotice, err
	}
	return paymentNotice, nil
}
