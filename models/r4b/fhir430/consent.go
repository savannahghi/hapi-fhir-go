
package fhir430

import "encoding/json"
// Consent is documented here http://hl7.org/fhir/StructureDefinition/Consent
// A record of a healthcare consumer’s  choices, which permits or denies identified recipient(s) or recipient role(s) to perform one or more actions within a given policy context, for specific purposes and periods of time.
type Consent struct {
	ID                *string               `json:"ID,omitempty"`
	Meta              *Meta                 `json:"meta,omitempty"`
	ImplicitRules     *string               `json:"implicitRules,omitempty"`
	Language          *string               `json:"language,omitempty"`
	Text              *Narrative            `json:"text,omitempty"`
	Contained         []json.RawMessage     `json:"contained,omitempty"`
	Extension         []Extension           `json:"extension,omitempty"`
	ModifierExtension []Extension           `json:"modifierExtension,omitempty"`
	Identifier        []Identifier          `json:"identifier,omitempty"`
	Status            ConsentState          `json:"status"`
	Scope             CodeableConcept       `json:"scope"`
	Category          []CodeableConcept     `json:"category"`
	Patient           *Reference            `json:"patient,omitempty"`
	DateTime          *string               `json:"dateTime,omitempty"`
	Performer         []Reference           `json:"performer,omitempty"`
	Organization      []Reference           `json:"organization,omitempty"`
	SourceAttachment  *Attachment           `json:"sourceAttachment,omitempty"`
	SourceReference   *Reference            `json:"sourceReference,omitempty"`
	Policy            []ConsentPolicy       `json:"policy,omitempty"`
	PolicyRule        *CodeableConcept      `json:"policyRule,omitempty"`
	Verification      []ConsentVerification `json:"verification,omitempty"`
	Provision         *ConsentProvision     `json:"provision,omitempty"`
}

// The references to the policies that are included in this consent scope. Policies may be organizational, but are often defined jurisdictionally, or in law.
type ConsentPolicy struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Authority         *string     `json:"authority,omitempty"`
	Uri               *string     `json:"uri,omitempty"`
}

// Whether a treatment instruction (e.g. artificial respiration yes or no) was verified with the patient, his/her family or another authorized person.
type ConsentVerification struct {
	ID                *string     `json:"ID,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Verified          bool        `json:"verified"`
	VerifiedWith      *Reference  `json:"verifiedWith,omitempty"`
	VerificationDate  *string     `json:"verificationDate,omitempty"`
}

// An exception to the base policy of this consent. An exception can be an addition or removal of access permissions.
type ConsentProvision struct {
	ID                *string                 `json:"ID,omitempty"`
	Extension         []Extension             `json:"extension,omitempty"`
	ModifierExtension []Extension             `json:"modifierExtension,omitempty"`
	Type              *ConsentProvisionType   `json:"type,omitempty"`
	Period            *Period                 `json:"period,omitempty"`
	Actor             []ConsentProvisionActor `json:"actor,omitempty"`
	Action            []CodeableConcept       `json:"action,omitempty"`
	SecurityLabel     []Coding                `json:"securityLabel,omitempty"`
	Purpose           []Coding                `json:"purpose,omitempty"`
	Class             []Coding                `json:"class,omitempty"`
	Code              []CodeableConcept       `json:"code,omitempty"`
	DataPeriod        *Period                 `json:"dataPeriod,omitempty"`
	Data              []ConsentProvisionData  `json:"data,omitempty"`
	Provision         []ConsentProvision      `json:"provision,omitempty"`
}

// Who or what is controlled by this rule. Use group to identify a set of actors by some property they share (e.g. 'admitting officers').
// There is no specific actor associated with the exception
type ConsentProvisionActor struct {
	ID                *string         `json:"ID,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Role              CodeableConcept `json:"role"`
	Reference         Reference       `json:"reference"`
}

// The resources controlled by this rule if specific resources are referenced.
// all data
type ConsentProvisionData struct {
	ID                *string            `json:"ID,omitempty"`
	Extension         []Extension        `json:"extension,omitempty"`
	ModifierExtension []Extension        `json:"modifierExtension,omitempty"`
	Meaning           ConsentDataMeaning `json:"meaning"`
	Reference         Reference          `json:"reference"`
}

// This function returns resource reference information
func (r Consent) ResourceRef() (string, *string) {
	return "Consent", r.ID
}

// This function returns resource reference information
func (r Consent) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherConsent Consent

// MarshalJSON marshals the given Consent as JSON into a byte slice
func (r Consent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherConsent
		ResourceType string `json:"resourceType"`
	}{
		OtherConsent: OtherConsent(r),
		ResourceType: "Consent",
	})
}

// UnmarshalConsent unmarshals a Consent.
func UnmarshalConsent(b []byte) (Consent, error) {
	var consent Consent
	if err := json.Unmarshal(b, &consent); err != nil {
		return consent, err
	}
	return consent, nil
}
