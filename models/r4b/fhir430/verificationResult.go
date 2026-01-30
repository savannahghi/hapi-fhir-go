package fhir430

import "encoding/json"

// VerificationResult is documented here http://hl7.org/fhir/StructureDefinition/VerificationResult
// Describes validation requirements, source(s), status and dates for one or more elements.
type VerificationResult struct {
	ID                *string                           `json:"id,omitempty"`
	Meta              *Meta                             `json:"meta,omitempty"`
	ImplicitRules     *string                           `json:"implicitRules,omitempty"`
	Language          *string                           `json:"language,omitempty"`
	Text              *Narrative                        `json:"text,omitempty"`
	Contained         []json.RawMessage                 `json:"contained,omitempty"`
	Extension         []Extension                       `json:"extension,omitempty"`
	ModifierExtension []Extension                       `json:"modifierExtension,omitempty"`
	Target            []Reference                       `json:"target,omitempty"`
	TargetLocation    []string                          `json:"targetLocation,omitempty"`
	Need              *CodeableConcept                  `json:"need,omitempty"`
	Status            string                            `json:"status"`
	StatusDate        *string                           `json:"statusDate,omitempty"`
	ValidationType    *CodeableConcept                  `json:"validationType,omitempty"`
	ValidationProcess []CodeableConcept                 `json:"validationProcess,omitempty"`
	Frequency         *Timing                           `json:"frequency,omitempty"`
	LastPerformed     *string                           `json:"lastPerformed,omitempty"`
	NextScheduled     *string                           `json:"nextScheduled,omitempty"`
	FailureAction     *CodeableConcept                  `json:"failureAction,omitempty"`
	PrimarySource     []VerificationResultPrimarySource `json:"primarySource,omitempty"`
	Attestation       *VerificationResultAttestation    `json:"attestation,omitempty"`
	Validator         []VerificationResultValidator     `json:"validator,omitempty"`
}

// Information about the primary source(s) involved in validation.
type VerificationResultPrimarySource struct {
	ID                  *string           `json:"id,omitempty"`
	Extension           []Extension       `json:"extension,omitempty"`
	ModifierExtension   []Extension       `json:"modifierExtension,omitempty"`
	Who                 *Reference        `json:"who,omitempty"`
	Type                []CodeableConcept `json:"type,omitempty"`
	CommunicationMethod []CodeableConcept `json:"communicationMethod,omitempty"`
	ValidationStatus    *CodeableConcept  `json:"validationStatus,omitempty"`
	ValidationDate      *string           `json:"validationDate,omitempty"`
	CanPushUpdates      *CodeableConcept  `json:"canPushUpdates,omitempty"`
	PushTypeAvailable   []CodeableConcept `json:"pushTypeAvailable,omitempty"`
}

// Information about the entity attesting to information.
type VerificationResultAttestation struct {
	ID                        *string          `json:"id,omitempty"`
	Extension                 []Extension      `json:"extension,omitempty"`
	ModifierExtension         []Extension      `json:"modifierExtension,omitempty"`
	Who                       *Reference       `json:"who,omitempty"`
	OnBehalfOf                *Reference       `json:"onBehalfOf,omitempty"`
	CommunicationMethod       *CodeableConcept `json:"communicationMethod,omitempty"`
	Date                      *string          `json:"date,omitempty"`
	SourceIdentityCertificate *string          `json:"sourceIdentityCertificate,omitempty"`
	ProxyIdentityCertificate  *string          `json:"proxyIdentityCertificate,omitempty"`
	ProxySignature            *Signature       `json:"proxySignature,omitempty"`
	SourceSignature           *Signature       `json:"sourceSignature,omitempty"`
}

// Information about the entity validating information.
type VerificationResultValidator struct {
	ID                   *string     `json:"id,omitempty"`
	Extension            []Extension `json:"extension,omitempty"`
	ModifierExtension    []Extension `json:"modifierExtension,omitempty"`
	Organization         Reference   `json:"organization"`
	IdentityCertificate  *string     `json:"identityCertificate,omitempty"`
	AttestationSignature *Signature  `json:"attestationSignature,omitempty"`
}

// This function returns resource reference information
func (r VerificationResult) ResourceRef() (string, *string) {
	return "VerificationResult", r.ID
}

// This function returns resource reference information
func (r VerificationResult) ContainedResources() []json.RawMessage {
	return r.Contained
}

type OtherVerificationResult VerificationResult

// MarshalJSON marshals the given VerificationResult as JSON into a byte slice
func (r VerificationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherVerificationResult
		ResourceType string `json:"resourceType"`
	}{
		OtherVerificationResult: OtherVerificationResult(r),
		ResourceType:            "VerificationResult",
	})
}

// UnmarshalVerificationResult unmarshals a VerificationResult.
func UnmarshalVerificationResult(b []byte) (VerificationResult, error) {
	var verificationResult VerificationResult
	if err := json.Unmarshal(b, &verificationResult); err != nil {
		return verificationResult, err
	}
	return verificationResult, nil
}
