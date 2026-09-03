package consent

import (
	"testing"
	"time"
)

func TestEvaluate_Checks(t *testing.T) {
	baseReq := Request{PatientID: "pat-1", At: sept3}

	tests := []struct {
		name       string
		mutate     func(d *Directive)
		req        Request
		pol        Policy
		wantValid  bool
		wantReason string
	}{
		{
			name:      "happy path",
			mutate:    func(*Directive) {},
			req:       baseReq,
			wantValid: true,
		},
		{
			name:       "patient mismatch",
			mutate:     func(*Directive) {},
			req:        Request{PatientID: "someone-else", At: sept3},
			wantReason: ReasonPatientMismatch,
		},
		{
			name:      "absolute subject url still matches",
			mutate:    func(d *Directive) { d.Subject = "https://fhir.example.test/fhir/Patient/pat-1" },
			req:       baseReq,
			wantValid: true,
		},
		{
			name:      "versioned subject reference still matches",
			mutate:    func(d *Directive) { d.Subject = "Patient/pat-1/_history/3" },
			req:       baseReq,
			wantValid: true,
		},
		{
			name:       "status not active",
			mutate:     func(d *Directive) { d.Status = "draft" },
			req:        baseReq,
			wantReason: ReasonStatusNotActive,
		},
		{
			name:       "scope mismatch",
			mutate:     func(d *Directive) { d.Categories = []Code{{System: "http://loinc.org", Code: "59284-0"}} },
			req:        baseReq,
			wantReason: ReasonScopeMismatch,
		},
		{
			name:       "required category missing",
			mutate:     func(*Directive) {},
			req:        baseReq,
			pol:        Policy{RequiredCategories: []Code{{System: "http://loinc.org", Code: "59284-0"}}},
			wantReason: ReasonCategoryMissing,
		},
		{
			name:       "policy basis required but absent",
			mutate:     func(d *Directive) { d.HasPolicyBasis = false },
			req:        baseReq,
			pol:        Policy{RequirePolicyBasis: true},
			wantReason: ReasonNoPolicy,
		},
		{
			name:      "policy basis absent but not required",
			mutate:    func(d *Directive) { d.HasPolicyBasis = false },
			req:       baseReq,
			wantValid: true,
		},
		{
			name:       "agreed after evaluation time",
			mutate:     func(d *Directive) { d.Date = tp(sept3.Add(time.Hour)) },
			req:        baseReq,
			wantReason: ReasonNotYetEffective,
		},
		{
			name:       "period not started",
			mutate:     func(d *Directive) { d.Period.Start = tp(sept3.Add(time.Minute)) },
			req:        baseReq,
			wantReason: ReasonNotYetEffective,
		},
		{
			name:       "period expired",
			mutate:     func(d *Directive) { d.Period.End = tp(sept3.Add(-time.Minute)) },
			req:        baseReq,
			wantReason: ReasonExpired,
		},
		{
			name:      "period start inclusive",
			mutate:    func(d *Directive) { d.Period.Start = tp(sept3) },
			req:       baseReq,
			wantValid: true,
		},
		{
			name:      "period end inclusive",
			mutate:    func(d *Directive) { d.Period.End = tp(sept3) },
			req:       baseReq,
			wantValid: true,
		},
		{
			name:      "open-ended period",
			mutate:    func(d *Directive) { d.Period.End = nil },
			req:       baseReq,
			wantValid: true,
		},
		{
			name:      "no period at all",
			mutate:    func(d *Directive) { d.Period = nil },
			req:       baseReq,
			wantValid: true,
		},
		{
			name:       "verification required but unverified",
			mutate:     func(d *Directive) { d.Verified = false },
			req:        baseReq,
			pol:        Policy{RequireVerification: true},
			wantReason: ReasonUnverified,
		},
		{
			name:      "unverified but verification not required",
			mutate:    func(d *Directive) { d.Verified = false },
			req:       baseReq,
			wantValid: true,
		},
		{
			name:       "base decision deny",
			mutate:     func(d *Directive) { d.BaseDecision = dp(Deny) },
			req:        baseReq,
			wantReason: ReasonProvisionDenies,
		},
		{
			name:      "no base decision falls back to policy default permit",
			mutate:    func(d *Directive) { d.BaseDecision = nil },
			req:       baseReq,
			wantValid: true,
		},
		{
			name:       "no base decision with policy default deny",
			mutate:     func(d *Directive) { d.BaseDecision = nil },
			req:        baseReq,
			pol:        Policy{DefaultDecision: Deny},
			wantReason: ReasonProvisionDenies,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			d := activeDirective()
			tc.mutate(&d)

			res := Evaluate(d, tc.req, tc.pol)

			if res.Valid != tc.wantValid {
				t.Fatalf("Valid = %v, want %v; reasons %+v", res.Valid, tc.wantValid, res.Reasons)
			}

			if tc.wantValid && res.Decision != Permit {
				t.Errorf("Decision = %s, want permit", res.Decision)
			}

			if !tc.wantValid && !hasReason(res, tc.wantReason) {
				t.Errorf("reasons %+v lack %s", res.Reasons, tc.wantReason)
			}

			if res.ConsentID != "c-1" || res.PatientID != tc.req.PatientID || !res.EvaluatedAt.Equal(sept3) {
				t.Errorf("result header wrong: %+v", res)
			}
		})
	}
}

func TestEvaluate_ExpiresAt(t *testing.T) {
	d := activeDirective()
	res := Evaluate(d, Request{PatientID: "pat-1", At: sept3}, Policy{})

	if res.ExpiresAt == nil || !res.ExpiresAt.Equal(*d.Period.End) {
		t.Errorf("ExpiresAt = %v, want %v", res.ExpiresAt, d.Period.End)
	}

	d.Period.End = nil
	res = Evaluate(d, Request{PatientID: "pat-1", At: sept3}, Policy{})

	if res.ExpiresAt != nil {
		t.Errorf("ExpiresAt = %v, want nil for open-ended period", res.ExpiresAt)
	}
}

func TestEvaluate_DefaultsToNow(t *testing.T) {
	d := activeDirective()
	d.Date = tp(time.Now().Add(-time.Hour))
	d.Period = nil

	res := Evaluate(d, Request{PatientID: "pat-1"}, Policy{})
	if !res.Valid {
		t.Fatalf("expected valid, got %+v", res.Reasons)
	}

	if time.Since(res.EvaluatedAt) > time.Minute {
		t.Errorf("EvaluatedAt should default to now, got %s", res.EvaluatedAt)
	}
}

func TestEvaluate_Provisions(t *testing.T) {
	hmarkt := Code{System: "http://terminology.hl7.org/CodeSystem/v3-ActReason", Code: "HMARKT"}
	treat := Code{System: "http://terminology.hl7.org/CodeSystem/v3-ActReason", Code: "TREAT"}
	disclose := Code{System: "http://terminology.hl7.org/CodeSystem/consentaction", Code: "disclose"}
	access := Code{System: "http://terminology.hl7.org/CodeSystem/consentaction", Code: "access"}

	denyMarketingDisclosure := Provision{Decision: Deny, Purpose: []Code{hmarkt}, Action: []Code{disclose}}
	denyOrgExceptTreatment := Provision{
		Decision: Deny,
		Actors:   []string{"Organization/org-blocked"},
		Children: []Provision{{Decision: Permit, Purpose: []Code{treat}}},
	}

	tests := []struct {
		name       string
		provisions []Provision
		req        Request
		wantValid  bool
		wantPath   string
	}{
		{
			name:       "plain request ignores constrained exceptions",
			provisions: []Provision{denyMarketingDisclosure, denyOrgExceptTreatment},
			req:        Request{PatientID: "pat-1", At: sept3},
			wantValid:  true,
		},
		{
			name:       "matching purpose and action hits the deny exception",
			provisions: []Provision{denyMarketingDisclosure},
			req:        Request{PatientID: "pat-1", At: sept3, Purpose: []Code{hmarkt}, Action: []Code{disclose}},
			wantValid:  false,
			wantPath:   "provision[0]",
		},
		{
			name:       "matching purpose but different action does not match",
			provisions: []Provision{denyMarketingDisclosure},
			req:        Request{PatientID: "pat-1", At: sept3, Purpose: []Code{hmarkt}, Action: []Code{access}},
			wantValid:  true,
		},
		{
			name:       "bare code without system matches",
			provisions: []Provision{denyMarketingDisclosure},
			req:        Request{PatientID: "pat-1", At: sept3, Purpose: []Code{{Code: "HMARKT"}}, Action: []Code{{Code: "disclose"}}},
			wantValid:  false,
		},
		{
			name:       "blocked actor is denied",
			provisions: []Provision{denyOrgExceptTreatment},
			req:        Request{PatientID: "pat-1", At: sept3, Actor: "Organization/org-blocked"},
			wantValid:  false,
			wantPath:   "provision[0]",
		},
		{
			name:       "blocked actor is permitted for treatment (nested permit under deny)",
			provisions: []Provision{denyOrgExceptTreatment},
			req:        Request{PatientID: "pat-1", At: sept3, Actor: "Organization/org-blocked", Purpose: []Code{treat}},
			wantValid:  true,
		},
		{
			name:       "other actor is unaffected",
			provisions: []Provision{denyOrgExceptTreatment},
			req:        Request{PatientID: "pat-1", At: sept3, Actor: "Organization/org-other"},
			wantValid:  true,
		},
		{
			name: "deny wins among matching siblings",
			provisions: []Provision{
				{Decision: Permit, Actors: []string{"Practitioner/p1"}},
				{Decision: Deny, Actors: []string{"Practitioner/p1"}},
			},
			req:       Request{PatientID: "pat-1", At: sept3, Actor: "Practitioner/p1"},
			wantValid: false,
			wantPath:  "provision[1]",
		},
		{
			name: "expired exception is ignored",
			provisions: []Provision{{
				Decision: Deny,
				Actors:   []string{"Practitioner/p1"},
				Period:   &Interval{End: tp(sept3.Add(-time.Hour))},
			}},
			req:       Request{PatientID: "pat-1", At: sept3, Actor: "Practitioner/p1"},
			wantValid: true,
		},
		{
			name:       "unconstrained exception applies to everything",
			provisions: []Provision{{Decision: Deny}},
			req:        Request{PatientID: "pat-1", At: sept3},
			wantValid:  false,
		},
		{
			name:       "data constraint matches specific record",
			provisions: []Provision{{Decision: Deny, Data: []string{"Encounter/enc-9"}}},
			req:        Request{PatientID: "pat-1", At: sept3, Data: []string{"Encounter/enc-9"}},
			wantValid:  false,
		},
		{
			name:       "data constraint does not match other record",
			provisions: []Provision{{Decision: Deny, Data: []string{"Encounter/enc-9"}}},
			req:        Request{PatientID: "pat-1", At: sept3, Data: []string{"Encounter/enc-1"}},
			wantValid:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			d := activeDirective()
			d.Provisions = tc.provisions

			res := Evaluate(d, tc.req, Policy{})

			if res.Valid != tc.wantValid {
				t.Fatalf("Valid = %v, want %v; reasons %+v", res.Valid, tc.wantValid, res.Reasons)
			}

			if !tc.wantValid {
				if !hasReason(res, ReasonProvisionDenies) {
					t.Errorf("expected PROVISION_DENIES, got %+v", res.Reasons)
				}

				if tc.wantPath != "" && !containsPath(res, tc.wantPath) {
					t.Errorf("reason message should name %s: %+v", tc.wantPath, res.Reasons)
				}
			}
		})
	}
}

func containsPath(r Result, path string) bool {
	for _, re := range r.Reasons {
		if re.Code == ReasonProvisionDenies && len(re.Message) >= len(path) && re.Message[:len(path)] == path {
			return true
		}
	}

	return false
}

// TestEvaluate_Fixtures runs the adapter + evaluator over the JSON fixtures,
// including the shape written by the sil_advantage builder.
func TestEvaluate_Fixtures(t *testing.T) {
	at := time.Date(2026, 9, 3, 9, 0, 0, 0, time.UTC)

	t.Run("R5 screening permit, plain request", func(t *testing.T) {
		res := EvaluateR5(loadR5(t, "r5_screening_permit.json"), Request{PatientID: "pat-1", At: at}, Policy{})
		if !res.Valid {
			t.Fatalf("expected valid, got %+v", res.Reasons)
		}

		if res.ExpiresAt != nil {
			t.Errorf("builder writes no period, ExpiresAt should be nil")
		}
	})

	t.Run("R5 screening permit, actor named: the builder's provision reads as a deny exception", func(t *testing.T) {
		res := EvaluateR5(loadR5(t, "r5_screening_permit.json"),
			Request{PatientID: "pat-1", At: at, Actor: "Organization/org-1", Data: []string{"Encounter/enc-1"}}, Policy{})
		if res.Valid {
			t.Fatal("R5 provisions are exceptions; an actor-scoped provision under decision=permit denies that actor")
		}

		if !hasReason(res, ReasonProvisionDenies) {
			t.Errorf("want PROVISION_DENIES, got %+v", res.Reasons)
		}
	})

	t.Run("R5 screening deny is recorded inactive", func(t *testing.T) {
		res := EvaluateR5(loadR5(t, "r5_screening_deny.json"), Request{PatientID: "pat-1", At: at}, Policy{})
		if res.Valid || !hasReason(res, ReasonStatusNotActive) {
			t.Errorf("want STATUS_NOT_ACTIVE, got %+v", res)
		}
	})

	t.Run("R5 recommended shape", func(t *testing.T) {
		c := loadR5(t, "r5_recommended.json")
		pol := Policy{RequireVerification: true, RequirePolicyBasis: true}

		if res := EvaluateR5(c, Request{PatientID: "pat-1", At: at}, pol); !res.Valid {
			t.Errorf("plain: %+v", res.Reasons)
		}

		res := EvaluateR5(c, Request{PatientID: "pat-1", At: at,
			Purpose: []Code{{Code: "HMARKT"}}, Action: []Code{{Code: "disclose"}}}, pol)
		if res.Valid {
			t.Error("marketing disclosure should be denied")
		}

		res = EvaluateR5(c, Request{PatientID: "pat-1", At: at, Actor: "Organization/org-blocked"}, pol)
		if res.Valid {
			t.Error("blocked org should be denied")
		}

		res = EvaluateR5(c, Request{PatientID: "pat-1", At: at, Actor: "Organization/org-blocked",
			Purpose: []Code{{Code: "ETREAT"}}}, pol)
		if !res.Valid {
			t.Errorf("blocked org should be permitted for emergency treatment: %+v", res.Reasons)
		}

		res = EvaluateR5(c, Request{PatientID: "pat-1", At: time.Date(2028, 1, 1, 0, 0, 0, 0, time.UTC)}, pol)
		if res.Valid || !hasReason(res, ReasonExpired) {
			t.Errorf("want EXPIRED in 2028, got %+v", res)
		}
	})

	t.Run("malformed date surfaces as MALFORMED_CONSENT", func(t *testing.T) {
		c := loadR5(t, "r5_screening_permit.json")
		bad := "soon"
		c.Date = &bad

		res := EvaluateR5(c, Request{PatientID: "pat-1", At: at}, Policy{})
		if res.Valid || !hasReason(res, ReasonMalformedConsent) {
			t.Errorf("want MALFORMED_CONSENT, got %+v", res)
		}
	})
}
