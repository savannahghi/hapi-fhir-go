package consent

import (
	"fmt"
	"strings"
	"time"

	fhir500 "github.com/savannahghi/hapi-fhir-go/models/r5/fhir500"
)

// EvaluateR5 converts an R5 Consent and evaluates it.
func EvaluateR5(c fhir500.Consent, req Request, pol Policy) Result {
	d, err := FromR5(c)
	if err != nil {
		return malformed(d, req, err)
	}

	return Evaluate(d, req, pol)
}

func malformed(d Directive, req Request, err error) Result {
	r := newResult(d, req, evalTime(req))
	r.addReason(ReasonMalformedConsent, err.Error())

	return r
}

func evalTime(req Request) time.Time {
	if req.At.IsZero() {
		return time.Now().UTC()
	}

	return req.At
}

func newResult(d Directive, req Request, at time.Time) Result {
	r := Result{
		Decision:    Deny,
		PatientID:   req.PatientID,
		ConsentID:   d.ID,
		EvaluatedAt: at,
	}

	if d.Period != nil && d.Period.End != nil {
		end := *d.Period.End
		r.ExpiresAt = &end
	}

	return r
}

// Evaluate is pure: no I/O, and deterministic for a given Request.At.
//
// Checks run in order and stop at the first failure:
//
//  1. subject matches the patient
//  2. status is active
//  3. scope matches the policy
//  4. required categories are present
//  5. a policy basis exists (only when the policy demands it)
//  6. the consent was agreed on or before At
//  7. the consent period covers At
//  8. verified (only when the policy demands it)
//  9. the provision tree yields permit for the request
func Evaluate(d Directive, req Request, pol Policy) Result {
	pol = pol.withDefaults()
	at := evalTime(req)
	res := newResult(d, req, at)

	if !referenceMatches(d.Subject, "Patient/"+req.PatientID) {
		res.addReason(ReasonPatientMismatch,
			fmt.Sprintf("consent subject %q is not Patient/%s", d.Subject, req.PatientID))

		return res
	}

	if d.Status != "active" {
		res.addReason(ReasonStatusNotActive, fmt.Sprintf("consent status is %q, want active", d.Status))

		return res
	}

	if !containsCode(d.Categories, pol.Scope) {
		res.addReason(ReasonScopeMismatch,
			fmt.Sprintf("consent does not carry category %s|%s", pol.Scope.System, pol.Scope.Code))

		return res
	}

	for _, want := range pol.RequiredCategories {
		if !containsCode(d.Categories, want) {
			res.addReason(ReasonCategoryMissing,
				fmt.Sprintf("consent lacks required category %s|%s", want.System, want.Code))

			return res
		}
	}

	if pol.RequirePolicyBasis && !d.HasPolicyBasis {
		res.addReason(ReasonNoPolicy, "consent names no policy, policyRule, regulatoryBasis or policyBasis")

		return res
	}

	if d.Date != nil && at.Before(*d.Date) {
		res.addReason(ReasonNotYetEffective,
			fmt.Sprintf("consent was agreed on %s, after the evaluation time", d.Date.Format(time.RFC3339)))

		return res
	}

	if d.Period != nil {
		if d.Period.Start != nil && at.Before(*d.Period.Start) {
			res.addReason(ReasonNotYetEffective,
				fmt.Sprintf("consent period starts %s", d.Period.Start.Format(time.RFC3339)))

			return res
		}

		if d.Period.End != nil && at.After(*d.Period.End) {
			res.addReason(ReasonExpired,
				fmt.Sprintf("consent period ended %s", d.Period.End.Format(time.RFC3339)))

			return res
		}
	}

	if pol.RequireVerification && !d.Verified {
		res.addReason(ReasonUnverified, "consent has no verification with verified=true")

		return res
	}

	base := pol.DefaultDecision
	if d.BaseDecision != nil {
		base = *d.BaseDecision
	}

	decision, path := resolve(d.Provisions, "provision", req, at)
	if decision == nil {
		decision = &base
		path = "base decision"
	}

	res.Decision = *decision
	if *decision == Deny {
		res.addReason(ReasonProvisionDenies, fmt.Sprintf("%s denies the request", path))

		return res
	}

	res.Valid = true

	return res
}

// resolve walks the provision tree. Each matching provision contributes the
// decision of its deepest matching descendant; across siblings, deny wins.
// It returns nil when nothing matches.
func resolve(provs []Provision, prefix string, req Request, at time.Time) (*Decision, string) {
	var (
		result *Decision
		path   string
	)

	for i := range provs {
		p := &provs[i]
		if !p.matches(req, at) {
			continue
		}

		here := fmt.Sprintf("%s[%d]", prefix, i)
		dec := p.Decision

		if child, childPath := resolve(p.Children, here+".provision", req, at); child != nil {
			dec = *child
			here = childPath
		}

		if result == nil || dec == Deny {
			d := dec
			result = &d
			path = here
		}
	}

	return result, path
}

// matches reports whether the provision applies to the request. Every
// constraint the provision populates must be satisfied by a value the request
// supplies; a request that says nothing about actors cannot match a provision
// that names actors.
func (p *Provision) matches(req Request, at time.Time) bool {
	if !p.Period.Covers(at) {
		return false
	}

	if len(p.Actors) > 0 && !anyReferenceMatches(p.Actors, req.Actor) {
		return false
	}

	if len(p.Purpose) > 0 && !intersects(p.Purpose, req.Purpose) {
		return false
	}

	if len(p.Action) > 0 && !intersects(p.Action, req.Action) {
		return false
	}

	if len(p.Class) > 0 && !intersects(p.Class, req.Class) {
		return false
	}

	if len(p.Data) > 0 && !anyDataMatches(p.Data, req.Data) {
		return false
	}

	return true
}

func anyReferenceMatches(refs []string, want string) bool {
	if want == "" {
		return false
	}

	for _, r := range refs {
		if referenceMatches(r, want) || referenceMatches(want, r) {
			return true
		}
	}

	return false
}

func anyDataMatches(refs, wants []string) bool {
	for _, w := range wants {
		if anyReferenceMatches(refs, w) {
			return true
		}
	}

	return false
}

func intersects(have, want []Code) bool {
	for _, h := range have {
		for _, w := range want {
			if h.Matches(w) {
				return true
			}
		}
	}

	return false
}

func containsCode(have []Code, want Code) bool {
	for _, h := range have {
		if h.Matches(want) {
			return true
		}
	}

	return false
}

// String renders a Decision for messages.
func (d Decision) String() string { return strings.ToLower(string(d)) }
