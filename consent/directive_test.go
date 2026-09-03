package consent

import (
	"testing"
	"time"
)

func TestFromR5_ScreeningPermit(t *testing.T) {
	d, err := FromR5(loadR5(t, "r5_screening_permit.json"))
	if err != nil {
		t.Fatalf("FromR5: %v", err)
	}

	if d.ID != "scr-1" || d.Status != "active" || d.Subject != "Patient/pat-1" {
		t.Errorf("header fields wrong: %+v", d)
	}

	if !containsCode(d.Categories, DefaultScope) {
		t.Errorf("categories %v lack default scope", d.Categories)
	}

	if d.Date == nil || !d.Date.Equal(time.Date(2026, 9, 1, 7, 0, 0, 0, time.UTC)) {
		t.Errorf("date = %v, want 2026-09-01T07:00:00Z", d.Date)
	}

	if d.Period != nil {
		t.Errorf("period should be nil when Consent.period is absent, got %+v", d.Period)
	}

	if !d.HasPolicyBasis || !d.Verified || !d.HasVerification {
		t.Errorf("policy/verification flags wrong: %+v", d)
	}

	if d.BaseDecision == nil || *d.BaseDecision != Permit {
		t.Errorf("base decision = %v, want permit", d.BaseDecision)
	}

	if len(d.Provisions) != 1 {
		t.Fatalf("provisions = %d, want 1", len(d.Provisions))
	}

	// In R5 a provision is an exception to the base decision.
	p := d.Provisions[0]
	if p.Decision != Deny {
		t.Errorf("provision[0].Decision = %s, want deny (exception to permit)", p.Decision)
	}

	if len(p.Actors) != 1 || p.Actors[0] != "Organization/org-1" {
		t.Errorf("actors = %v", p.Actors)
	}

	if len(p.Data) != 1 || p.Data[0] != "Encounter/enc-1" {
		t.Errorf("data = %v", p.Data)
	}
}

func TestFromR5_Recommended(t *testing.T) {
	d, err := FromR5(loadR5(t, "r5_recommended.json"))
	if err != nil {
		t.Fatalf("FromR5: %v", err)
	}

	if d.Period == nil || d.Period.Start == nil || d.Period.End == nil {
		t.Fatalf("period not parsed: %+v", d.Period)
	}

	// "2027-09-03" as an inclusive end covers the whole day.
	wantEnd := time.Date(2027, 9, 4, 0, 0, 0, 0, time.UTC).Add(-time.Nanosecond)
	if !d.Period.End.Equal(wantEnd) {
		t.Errorf("period end = %s, want %s", d.Period.End, wantEnd)
	}

	if len(d.Provisions) != 2 {
		t.Fatalf("provisions = %d, want 2", len(d.Provisions))
	}

	if d.Provisions[0].Decision != Deny || d.Provisions[1].Decision != Deny {
		t.Errorf("top-level provisions must be deny exceptions to a permit base")
	}

	if len(d.Provisions[1].Children) != 1 || d.Provisions[1].Children[0].Decision != Permit {
		t.Errorf("nested provision must flip back to permit: %+v", d.Provisions[1].Children)
	}
}

func TestFromR5_BadDate(t *testing.T) {
	c := loadR5(t, "r5_screening_permit.json")
	bad := "not-a-date"
	c.Date = &bad

	if _, err := FromR5(c); err == nil {
		t.Error("expected error for unparsable date")
	}
}
