package fhir430

import "encoding/json"

// Account is documented here http://hl7.org/fhir/StructureDefinition/Account
// A financial tool for tracking value accrued for a particular purpose.  In the healthcare field, used to track charges for a patient, cost centers, etc.
type Account struct {
	ID                *string            `json:"ID,omitempty"`
	Meta              *Meta              `json:"meta,omitempty"`
	ImplicitRules     *string            `json:"implicitRules,omitempty"`
	Language          *string            `json:"language,omitempty"`
	Text              *Narrative         `json:"text,omitempty"`
	Contained         []json.RawMessage  `json:"contained,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Identifier        []Identifier       `json:"identifier,omitempty"`
	Status            AccountStatus      `json:"status"`
	Type              *CodeableConcept   `json:"type,omitempty"`
	Name              *string            `json:"name,omitempty"`
	Subject           []Reference        `json:"subject,omitempty"`
	ServicePeriod     *Period            `json:"servicePeriod,omitempty"`
	Coverage          []AccountCoverage  `json:"coverage,omitempty"`
	Owner             *Reference         `json:"owner,omitempty"`
	Description       *string            `json:"description,omitempty"`
	Guarantor         []AccountGuarantor `json:"guarantor,omitempty"`
	PartOf            *Reference         `json:"partOf,omitempty"`
}

// The party(s) that are responsible for covering the payment of this account, and what order should they be applied to the account.
/*
Typically. this may be some form of insurance, internal charges, or self-pay.

Local or jurisdictional business rules may determine which coverage covers which types of billable items charged to the account, and in which order.
Where the order is important, a local/jurisdictional extension may be defined to specify the order for the type of charge.
*/
type AccountCoverage struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Coverage          Reference   `json:"coverage"`
	Priority          *int        `json:"priority,omitempty"`
}

// The parties responsible for balancing the account if other payment options fall short.
type AccountGuarantor struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Party             Reference   `json:"party"`
	OnHold            *bool       `json:"onHold,omitempty"`
	Period            *Period     `json:"period,omitempty"`
}

// This function returns resource reference information
func (r Account) ResourceRef() (string, *string) {
	return "Account", r.ID
}

// This function returns resource reference information
func (r Account) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherAccount Account

// MarshalJSON marshals the given Account as JSON into a byte slice
func (r Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherAccount
		ResourceType string `json:"resourceType"`
	}{
		OtherAccount: OtherAccount(r),
		ResourceType: "Account",
	})
}

// UnmarshalAccount unmarshals a Account.
func UnmarshalAccount(b []byte) (Account, error) {
	var account Account
	if err := json.Unmarshal(b, &account); err != nil {
		return account, err
	}
	return account, nil
}
