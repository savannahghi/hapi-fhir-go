package consent

import (
	"fmt"
	"strings"
	"time"

	fhir500 "github.com/savannahghi/hapi-fhir-go/models/r5/fhir500"
)

// Directive is the flattened view of a Consent that Evaluate works on.
// Build one with FromR5, or construct it directly for tests.
type Directive struct {
	ID string
	// Status is the raw FHIR status code, e.g. "active".
	Status string
	// Subject is the reference the consent applies to, e.g. "Patient/123".
	Subject string
	// Categories holds Consent.category codings.
	Categories []Code
	// Date is when the consent was agreed (Consent.date).
	Date *time.Time
	// Period is Consent.period. End is inclusive and already widened for
	// partial dates.
	Period *Interval
	// HasPolicyBasis is true when the consent names the policy it rests on.
	HasPolicyBasis bool
	// Verified is true when any verification entry is verified.
	Verified bool
	// HasVerification is true when the consent carries any verification entry.
	HasVerification bool
	// BaseDecision is the consent's own default; nil means it states none.
	BaseDecision *Decision
	// Provisions are exceptions to the base decision. Each already carries
	// its effective Decision.
	Provisions []Provision
}

// Interval is a closed time range. A nil bound is open.
type Interval struct {
	Start *time.Time
	End   *time.Time
}

// Covers reports whether at lies inside the interval, bounds inclusive.
func (i *Interval) Covers(at time.Time) bool {
	if i == nil {
		return true
	}

	if i.Start != nil && at.Before(*i.Start) {
		return false
	}

	if i.End != nil && at.After(*i.End) {
		return false
	}

	return true
}

// Provision is one rule in the consent's exception tree.
type Provision struct {
	// Decision is what applies when this provision matches.
	Decision Decision
	Period   *Interval
	// Actors are references such as "Organization/7".
	Actors  []string
	Purpose []Code
	Action  []Code
	Class   []Code
	// Data are references to specific records, e.g. "Encounter/9".
	Data     []string
	Children []Provision
}

// FromR5 converts an R5 Consent into a Directive.
//
// R5 rules: Consent.decision is the base decision. Every Consent.provision is
// an exception to it, and every nested provision is an exception to its
// parent, so effective decisions alternate with depth.
func FromR5(c fhir500.Consent) (Directive, error) {
	d := Directive{
		HasPolicyBasis: len(c.RegulatoryBasis) > 0 || c.PolicyBasis != nil || len(c.PolicyText) > 0,
	}

	if c.ID != nil {
		d.ID = *c.ID
	}

	if c.Status != nil {
		d.Status = string(*c.Status)
	}

	if c.Subject != nil && c.Subject.Reference != nil {
		d.Subject = *c.Subject.Reference
	}

	for _, cat := range c.Category {
		d.Categories = append(d.Categories, codesFromR5Concept(cat)...)
	}

	if c.Date != nil && *c.Date != "" {
		t, err := ParseTime(*c.Date)
		if err != nil {
			return d, fmt.Errorf("date: %w", err)
		}

		d.Date = &t
	}

	if c.Period != nil {
		iv, err := intervalFromStrings(string(c.Period.Start), string(c.Period.End))
		if err != nil {
			return d, fmt.Errorf("period: %w", err)
		}

		d.Period = iv
	}

	for _, v := range c.Verification {
		if v == nil {
			continue
		}

		d.HasVerification = true
		d.Verified = d.Verified || v.Verified
	}

	if c.Decision != nil {
		dec := decisionFromCode(string(*c.Decision))
		d.BaseDecision = &dec
	}

	parent := Permit
	if d.BaseDecision != nil {
		parent = *d.BaseDecision
	}

	for i, p := range c.Provision {
		if p == nil {
			continue
		}

		pv, err := provisionFromR5(*p, parent)
		if err != nil {
			return d, fmt.Errorf("provision[%d]: %w", i, err)
		}

		d.Provisions = append(d.Provisions, pv)
	}

	return d, nil
}

func provisionFromR5(p fhir500.ConsentProvision, parent Decision) (Provision, error) {
	out := Provision{Decision: parent.Opposite()}

	if p.Period != nil {
		iv, err := intervalFromStrings(string(p.Period.Start), string(p.Period.End))
		if err != nil {
			return out, fmt.Errorf("period: %w", err)
		}

		out.Period = iv
	}

	for _, a := range p.Actor {
		if a != nil && a.Reference != nil && a.Reference.Reference != nil {
			out.Actors = append(out.Actors, *a.Reference.Reference)
		}
	}

	for _, a := range p.Action {
		out.Action = append(out.Action, codesFromR5Concept(a)...)
	}

	for _, c := range p.Purpose {
		if c != nil {
			out.Purpose = appendCode(out.Purpose, c.System, c.Code)
		}
	}

	for _, c := range p.ResourceType {
		if c != nil {
			out.Class = appendCode(out.Class, c.System, c.Code)
		}
	}

	for _, dt := range p.Data {
		if dt.Reference != nil && dt.Reference.Reference != nil {
			out.Data = append(out.Data, *dt.Reference.Reference)
		}
	}

	for i, child := range p.Provision {
		if child == nil {
			continue
		}

		cp, err := provisionFromR5(*child, out.Decision)
		if err != nil {
			return out, fmt.Errorf("provision[%d]: %w", i, err)
		}

		out.Children = append(out.Children, cp)
	}

	return out, nil
}

func decisionFromCode(code string) Decision {
	if strings.EqualFold(code, string(Deny)) {
		return Deny
	}

	return Permit
}

func codesFromR5Concept(cc *fhir500.CodeableConcept) []Code {
	if cc == nil {
		return nil
	}

	var out []Code
	for _, c := range cc.Coding {
		if c != nil {
			out = appendCode(out, c.System, c.Code)
		}
	}

	return out
}

func appendCode(out []Code, system, code *string) []Code {
	if code == nil || *code == "" {
		return out
	}

	c := Code{Code: *code}
	if system != nil {
		c.System = *system
	}

	return append(out, c)
}

// intervalFromStrings builds an Interval from FHIR start/end strings. Both
// empty yields nil (no constraint).
func intervalFromStrings(start, end string) (*Interval, error) {
	if start == "" && end == "" {
		return nil, nil //nolint:nilnil // nil means "no period constraint"
	}

	iv := &Interval{}

	if start != "" {
		t, err := ParseTime(start)
		if err != nil {
			return nil, fmt.Errorf("start: %w", err)
		}

		iv.Start = &t
	}

	if end != "" {
		t, err := ParsePeriodEnd(end)
		if err != nil {
			return nil, fmt.Errorf("end: %w", err)
		}

		iv.End = &t
	}

	return iv, nil
}
