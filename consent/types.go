package consent

import (
	"errors"
	"strings"
	"time"
)

// Decision is the outcome of evaluating a consent: permit or deny.
type Decision string

const (
	// Permit allows the requested use.
	Permit Decision = "permit"
	// Deny refuses the requested use.
	Deny Decision = "deny"
)

// Opposite flips permit to deny and deny to permit. R5 provisions are
// exceptions to their parent's decision, so each level flips.
func (d Decision) Opposite() Decision {
	if d == Deny {
		return Permit
	}

	return Deny
}

// Code is a (system, code) pair. An empty System on either side of a
// comparison matches any system, so callers may pass bare codes.
type Code struct {
	System string
	Code   string
}

// Matches reports whether two codes are equal, treating an empty system on
// either side as a wildcard.
func (c Code) Matches(o Code) bool {
	if c.Code == "" || o.Code == "" || c.Code != o.Code {
		return false
	}

	if c.System == "" || o.System == "" {
		return true
	}

	return c.System == o.System
}

// IsZero reports whether the code carries no value.
func (c Code) IsZero() bool { return c.System == "" && c.Code == "" }

// Request describes the access being checked.
type Request struct {
	// PatientID is the bare logical id of the patient, e.g. "123". Required.
	PatientID string
	// ConsentID is the logical id of the Consent to check. Required by
	// Client.ValidateConsent, which reads exactly that resource and confirms
	// it belongs to PatientID.
	ConsentID string
	// At is the evaluation instant. Zero means time.Now().
	At time.Time

	// The fields below are the context matched against provisions. A
	// provision constraint is only considered matched when the request
	// specifies a value for it: a request with no Actor never matches a
	// provision that names actors. Leave them empty to ask the plain
	// question "does this patient have a valid base consent?".

	// Actor is the party performing the action, e.g. "Practitioner/45" or
	// "Organization/7".
	Actor string
	// Purpose codes, usually from http://terminology.hl7.org/CodeSystem/v3-ActReason
	// (TREAT, HPAYMT, HOPERAT, HRESCH, HMARKT ...).
	Purpose []Code
	// Action codes from http://terminology.hl7.org/CodeSystem/consentaction
	// (access, collect, use, disclose, correct).
	Action []Code
	// Class codes: the kind of data, e.g. a FHIR resource type such as
	// {System: "http://hl7.org/fhir/fhir-types", Code: "Observation"}.
	Class []Code
	// Data references the specific records being accessed, e.g. "Encounter/9".
	Data []string
}

// Policy tightens the validity rules. The zero value is usable.
type Policy struct {
	// Scope is the consent scope the patient must have consented to. It is
	// matched against Consent.category codings.
	// Default: consentscope | patient-privacy.
	Scope Code
	// RequiredCategories must all be present in Consent.category.
	RequiredCategories []Code
	// RequireVerification demands at least one verification.verified = true.
	RequireVerification bool
	// RequirePolicyBasis demands at least one of regulatoryBasis, policyBasis
	// or policyText.
	RequirePolicyBasis bool
	// DefaultDecision applies when the consent states no base decision.
	// Default: Permit.
	DefaultDecision Decision
}

// DefaultScope is the scope a Policy checks for when none is set.
var DefaultScope = Code{
	System: "http://terminology.hl7.org/CodeSystem/consentscope",
	Code:   "patient-privacy",
}

func (p Policy) withDefaults() Policy {
	if p.Scope.IsZero() {
		p.Scope = DefaultScope
	}

	if p.DefaultDecision == "" {
		p.DefaultDecision = Permit
	}

	return p
}

// ErrPatientIDRequired is returned when Request.PatientID is empty.
var ErrPatientIDRequired = errors.New("consent: Request.PatientID is required")

// ErrConsentIDRequired is returned when Request.ConsentID is empty.
var ErrConsentIDRequired = errors.New("consent: Request.ConsentID is required")

// Reason codes. They are stable and safe to switch on.
const (
	ReasonConsentNotFound  = "CONSENT_NOT_FOUND"
	ReasonPatientMismatch  = "PATIENT_MISMATCH"
	ReasonStatusNotActive  = "STATUS_NOT_ACTIVE"
	ReasonScopeMismatch    = "SCOPE_MISMATCH"
	ReasonCategoryMissing  = "CATEGORY_MISSING"
	ReasonNoPolicy         = "NO_POLICY"
	ReasonNotYetEffective  = "NOT_YET_EFFECTIVE"
	ReasonExpired          = "EXPIRED"
	ReasonUnverified       = "UNVERIFIED"
	ReasonProvisionDenies  = "PROVISION_DENIES"
	ReasonMalformedConsent = "MALFORMED_CONSENT"
)

// Reason explains a Result. Codes are listed above.
type Reason struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Result is the outcome of validating a consent.
type Result struct {
	Valid       bool       `json:"valid"`
	Decision    Decision   `json:"decision"`
	PatientID   string     `json:"patientId"`
	ConsentID   string     `json:"consentId,omitempty"`
	EvaluatedAt time.Time  `json:"evaluatedAt"`
	ExpiresAt   *time.Time `json:"expiresAt,omitempty"`
	// Reasons explain a denial, or carry warnings on a valid result.
	Reasons []Reason `json:"reasons,omitempty"`
}

func (r *Result) addReason(code, msg string) {
	r.Reasons = append(r.Reasons, Reason{Code: code, Message: msg})
}

// referenceMatches reports whether a FHIR reference points at want (a
// relative reference such as "Patient/123"). Absolute URLs and versioned
// references are accepted.
func referenceMatches(got, want string) bool {
	if got == "" || want == "" {
		return false
	}

	if i := strings.Index(got, "/_history/"); i >= 0 {
		got = got[:i]
	}

	return got == want || strings.HasSuffix(got, "/"+want)
}
