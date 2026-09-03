// Package consent answers one question: does a patient have a valid FHIR R5
// Consent right now, for a given use?
//
// The package is pure: it has no I/O. FromR5 turns an R5 Consent into a
// version-neutral Directive, and Evaluate applies the validity rules to it
// for a Request under a Policy, returning a Result with a permit/deny decision
// and stable reason codes.
//
// Fetching the Consent from a HAPI server is done by Client.ValidateConsent in
// the root hapifhirgo package, which then calls into this one.
//
// See docs/consent-validation.md for the rules and for how a Consent must be
// built to pass them.
package consent
