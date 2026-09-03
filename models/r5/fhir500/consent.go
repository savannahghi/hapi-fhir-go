package fhir500

import (
	"github.com/savannahghi/scalarutils"
)

// Consent is documented here http://hl7.org/fhir/R5/consent.html
//
// A record of a healthcare consumer's choices or choices made on their behalf
// by a third party, which permits or denies identified recipient(s) or
// recipient role(s) to perform one or more actions within a given policy
// context, for specific purposes and periods of time.
//
// R5 structure: the base rule is Consent.decision (permit | deny); every
// Consent.provision is an exception to that base rule, and every nested
// provision.provision is an exception to its parent.
type Consent struct {
	ID                *string            `json:"id,omitempty"`
	Meta              *Meta              `json:"meta,omitempty"`
	ImplicitRules     *string            `json:"implicitRules,omitempty"`
	Language          *string            `json:"language,omitempty"`
	Text              *Narrative         `json:"text,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Identifier        []*Identifier      `json:"identifier,omitempty"`
	Status            *ConsentStatusEnum `json:"status"`
	Category          []*CodeableConcept `json:"category,omitempty"`
	// Subject is the Patient | Practitioner | Group the consent applies to.
	Subject *Reference `json:"subject,omitempty"`
	// Date is a FHIR `date` (YYYY, YYYY-MM or YYYY-MM-DD). It is kept as a
	// string so that partial dates round-trip unchanged.
	Date *string `json:"date,omitempty"`
	// Period is the effective period for this Consent and all provisions
	// unless a provision specifies its own.
	Period           *Period                `json:"period,omitempty"`
	Grantor          []*Reference           `json:"grantor,omitempty"`
	Grantee          []*Reference           `json:"grantee,omitempty"`
	Manager          []*Reference           `json:"manager,omitempty"`
	Controller       []*Reference           `json:"controller,omitempty"`
	SourceAttachment []*Attachment          `json:"sourceAttachment,omitempty"`
	SourceReference  []*Reference           `json:"sourceReference,omitempty"`
	RegulatoryBasis  []*CodeableConcept     `json:"regulatoryBasis,omitempty"`
	PolicyBasis      *ConsentPolicyBasis    `json:"policyBasis,omitempty"`
	PolicyText       []*Reference           `json:"policyText,omitempty"`
	Verification     []*ConsentVerification `json:"verification,omitempty"`
	// Decision is the base rule: permit or deny.
	Decision  *ConsentProvisionTypeEnum `json:"decision,omitempty"`
	Provision []*ConsentProvision       `json:"provision,omitempty"`
}

// ConsentPolicyBasis is a reference to the computable policy this consent is based on.
type ConsentPolicyBasis struct {
	ID        *string     `json:"id,omitempty"`
	Extension []Extension `json:"extension,omitempty"`
	Reference *Reference  `json:"reference,omitempty"`
	URL       *string     `json:"url,omitempty"`
}

// ConsentVerification records whether the consent was verified with the subject or a proxy.
type ConsentVerification struct {
	ID                *string                `json:"id,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Verified          bool                   `json:"verified"`
	VerificationType  *CodeableConcept       `json:"verificationType,omitempty"`
	VerifiedBy        *Reference             `json:"verifiedBy,omitempty"`
	VerifiedWith      *Reference             `json:"verifiedWith,omitempty"`
	VerificationDate  []scalarutils.DateTime `json:"verificationDate,omitempty"`
}

// ConsentProvision is an exception to the base policy of this consent.
type ConsentProvision struct {
	ID                *string                  `json:"id,omitempty"`
	Extension         []Extension              `json:"extension,omitempty"`
	ModifierExtension []Extension              `json:"modifierExtension,omitempty"`
	Period            *Period                  `json:"period,omitempty"`
	Actor             []*ConsentProvisionActor `json:"actor,omitempty"`
	Action            []*CodeableConcept       `json:"action,omitempty"`
	SecurityLabel     []*Coding                `json:"securityLabel,omitempty"`
	Purpose           []*Coding                `json:"purpose,omitempty"`
	DocumentType      []*Coding                `json:"documentType,omitempty"`
	ResourceType      []*Coding                `json:"resourceType,omitempty"`
	Code              []*CodeableConcept       `json:"code,omitempty"`
	DataPeriod        *Period                  `json:"dataPeriod,omitempty"`
	Data              []ConsentProvisionData   `json:"data,omitempty"`
	Expression        *Expression              `json:"expression,omitempty"`
	Provision         []*ConsentProvision      `json:"provision,omitempty"`
}

// ConsentProvisionActor is who or what is controlled by this provision.
type ConsentProvisionActor struct {
	ID                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `json:"role,omitempty"`
	Reference         *Reference       `json:"reference,omitempty"`
}

// ConsentProvisionData identifies the resources controlled by this provision.
type ConsentProvisionData struct {
	ID                *string                `json:"id,omitempty"`
	Extension         []Extension            `json:"extension,omitempty"`
	ModifierExtension []Extension            `json:"modifierExtension,omitempty"`
	Meaning           ConsentDataMeaningEnum `json:"meaning,omitempty"`
	Reference         *Reference             `json:"reference,omitempty"`
}
