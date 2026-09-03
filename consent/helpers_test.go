package consent

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	fhir500 "github.com/savannahghi/hapi-fhir-go/models/r5/fhir500"
)

// sept3 is the reference evaluation instant used across tests.
var sept3 = time.Date(2026, 9, 3, 9, 0, 0, 0, time.UTC)

func readFixture(t *testing.T, name string) []byte {
	t.Helper()

	b, err := os.ReadFile(filepath.Join("testdata", name))
	if err != nil {
		t.Fatalf("read fixture %s: %v", name, err)
	}

	return b
}

func loadR5(t *testing.T, name string) fhir500.Consent {
	t.Helper()

	var c fhir500.Consent
	if err := json.Unmarshal(readFixture(t, name), &c); err != nil {
		t.Fatalf("decode %s as R5: %v", name, err)
	}

	return c
}

func tp(t time.Time) *time.Time { return &t }

func dp(d Decision) *Decision { return &d }

func hasReason(r Result, code string) bool {
	for _, re := range r.Reasons {
		if re.Code == code {
			return true
		}
	}

	return false
}

// activeDirective is a minimal valid directive for Patient/pat-1 that tests
// mutate to trigger one failure at a time.
func activeDirective() Directive {
	start := sept3.Add(-24 * time.Hour)
	end := sept3.Add(365 * 24 * time.Hour)

	return Directive{
		ID:              "c-1",
		Status:          "active",
		Subject:         "Patient/pat-1",
		Categories:      []Code{DefaultScope},
		Date:            tp(start),
		Period:          &Interval{Start: tp(start), End: tp(end)},
		HasPolicyBasis:  true,
		HasVerification: true,
		Verified:        true,
		BaseDecision:    dp(Permit),
	}
}
