
package fhir430

import "encoding/json"
// Coverage is documented here http://hl7.org/fhir/StructureDefinition/Coverage
// Financial instrument which may be used to reimburse or pay for health care products and services. Includes both insurance and self-payment.
type Coverage struct {
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
	Type              *CodeableConcept             `json:"type,omitempty"`
	PolicyHolder      *Reference                   `json:"policyHolder,omitempty"`
	Subscriber        *Reference                   `json:"subscriber,omitempty"`
	SubscriberId      *string                      `json:"subscriberId,omitempty"`
	Beneficiary       Reference                    `json:"beneficiary"`
	Dependent         *string                      `json:"dependent,omitempty"`
	Relationship      *CodeableConcept             `json:"relationship,omitempty"`
	Period            *Period                      `json:"period,omitempty"`
	Payor             []Reference                  `json:"payor"`
	Class             []CoverageClass              `json:"class,omitempty"`
	Order             *int                         `json:"order,omitempty"`
	Network           *string                      `json:"network,omitempty"`
	CostToBeneficiary []CoverageCostToBeneficiary  `json:"costToBeneficiary,omitempty"`
	Subrogation       *bool                        `json:"subrogation,omitempty"`
	Contract          []Reference                  `json:"contract,omitempty"`
}

// A suite of underwriter specific classifiers.
// For example may be used to identify a class of coverage or employer group, Policy, Plan.
type CoverageClass struct {
	ID                *string         `json:"ID,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Value             string          `json:"value"`
	Name              *string         `json:"name,omitempty"`
}

// A suite of codes indicating the cost category and associated amount which have been detailed in the policy and may have been  included on the health card.
// For example by knowing the patient visit co-pay, the provider can collect the amount prior to undertaking treatment.
type CoverageCostToBeneficiary struct {
	ID                *string                              `json:"ID,omitempty"`
	Extension         []Extension                          `json:"extension,omitempty"`
	ModifierExtension []Extension                          `json:"modifierExtension,omitempty"`
	Type              *CodeableConcept                     `json:"type,omitempty"`
	ValueQuantity     Quantity                             `json:"valueQuantity"`
	ValueMoney        Money                                `json:"valueMoney"`
	Exception         []CoverageCostToBeneficiaryException `json:"exception,omitempty"`
}

// A suite of codes indicating exceptions or reductions to patient costs and their effective periods.
type CoverageCostToBeneficiaryException struct {
	ID                *string         `json:"ID,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Type              CodeableConcept `json:"type"`
	Period            *Period         `json:"period,omitempty"`
}

// This function returns resource reference information
func (r Coverage) ResourceRef() (string, *string) {
	return "Coverage", r.ID
}

// This function returns resource reference information
func (r Coverage) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherCoverage Coverage

// MarshalJSON marshals the given Coverage as JSON into a byte slice
func (r Coverage) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherCoverage
		ResourceType string `json:"resourceType"`
	}{
		OtherCoverage: OtherCoverage(r),
		ResourceType:  "Coverage",
	})
}

// UnmarshalCoverage unmarshals a Coverage.
func UnmarshalCoverage(b []byte) (Coverage, error) {
	var coverage Coverage
	if err := json.Unmarshal(b, &coverage); err != nil {
		return coverage, err
	}
	return coverage, nil
}
