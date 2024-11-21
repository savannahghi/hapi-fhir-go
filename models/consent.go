package models

import (
	"github.com/savannahghi/scalarutils"
)

// Consent models a fhir consent resource.
type FHIRConsent struct {
	ID         *string                `json:"id,omitempty"`
	Status     *ConsentStatusEnum     `json:"status"`
	Scope      *FHIRCodeableConcept   `json:"scope"`
	Category   []*FHIRCodeableConcept `json:"category"`
	PolicyRule *FHIRCodeableConcept   `json:"policyRule,omitempty"`
	Provision  *FHIRConsentProvision  `json:"provision,omitempty"`
	Patient    *FHIRReference         `json:"patient,omitempty"`
	Meta       *FHIRMeta              `json:"meta,omitempty"`
	Extension  []Extension            `json:"extension,omitempty"`
	DateTime   *scalarutils.DateTime  `json:"dateTime,omitempty"`
}

// FHIRConsentProvision models a fhir consent provision.
type FHIRConsentProvision struct {
	ID   *string                    `json:"id,omitempty"`
	Type *ConsentProvisionTypeEnum  `json:"type,omitempty"`
	Data []FHIRConsentProvisionData `json:"data,omitempty"`
}

// FHIRConsentProvisionData models a consent provision data.
type FHIRConsentProvisionData struct {
	ID                *string                `json:"id,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Meaning           ConsentDataMeaningEnum `json:"meaning,omitempty"`
	Reference         *FHIRReference         `json:"reference,omitempty"`
}
