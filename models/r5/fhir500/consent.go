package fhir500

import (
	"github.com/savannahghi/scalarutils"
)

// Consent models a fhir consent resource.
type Consent struct {
	ID         *string               `json:"id,omitempty"`
	Status     *ConsentStatusEnum    `json:"status"`
	Scope      *CodeableConcept      `json:"scope"`
	Category   []*CodeableConcept    `json:"category"`
	PolicyRule *CodeableConcept      `json:"policyRule,omitempty"`
	Provision  *ConsentProvision     `json:"provision,omitempty"`
	Patient    *Reference            `json:"patient,omitempty"`
	Meta       *Meta                 `json:"meta,omitempty"`
	Extension  []Extension           `json:"extension,omitempty"`
	DateTime   *scalarutils.DateTime `json:"dateTime,omitempty"`
}

// ConsentProvision models a fhir consent provision.
type ConsentProvision struct {
	ID   *string                   `json:"id,omitempty"`
	Type *ConsentProvisionTypeEnum `json:"type,omitempty"`
	Data []ConsentProvisionData    `json:"data,omitempty"`
}

// ConsentProvisionData models a consent provision data.
type ConsentProvisionData struct {
	ID                *string                `json:"id,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Meaning           ConsentDataMeaningEnum `json:"meaning,omitempty"`
	Reference         *Reference             `json:"reference,omitempty"`
}
