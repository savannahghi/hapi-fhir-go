package hapifhirgo

import (
	"context"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/savannahghi/hapi-fhir-go/consent"
)

const consentNotFoundOutcome = `{"resourceType":"OperationOutcome","issue":[{"severity":"error","code":"not-found"}]}`

var consentAt = time.Date(2026, 9, 3, 9, 0, 0, 0, time.UTC)

func consentFixture(t *testing.T, name string) []byte {
	t.Helper()

	b, err := os.ReadFile(filepath.Join("consent", "testdata", name))
	if err != nil {
		t.Fatalf("read fixture %s: %v", name, err)
	}

	return b
}

func writeFHIR(w http.ResponseWriter, status int, body []byte) {
	w.Header().Set("Content-Type", "application/fhir+json")
	w.WriteHeader(status)
	_, _ = w.Write(body)
}

func consentClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()

	srv := newTestServer(t, handler)

	c, err := NewClient(srv.URL)
	if err != nil {
		t.Fatal(err)
	}

	return c
}

func hasReasonCode(r consent.Result, code string) bool {
	for _, re := range r.Reasons {
		if re.Code == code {
			return true
		}
	}

	return false
}

func TestValidateConsent_RequiredFields(t *testing.T) {
	c := consentClient(t, func(w http.ResponseWriter, _ *http.Request) {
		t.Error("no request expected")
		writeFHIR(w, http.StatusInternalServerError, nil)
	})

	_, err := c.ValidateConsent(context.Background(), consent.Request{ConsentID: "x"}, consent.Policy{})
	if !errors.Is(err, consent.ErrPatientIDRequired) {
		t.Errorf("err = %v, want ErrPatientIDRequired", err)
	}

	_, err = c.ValidateConsent(context.Background(), consent.Request{PatientID: "pat-1"}, consent.Policy{})
	if !errors.Is(err, consent.ErrConsentIDRequired) {
		t.Errorf("err = %v, want ErrConsentIDRequired", err)
	}
}

func TestValidateConsent_ByID(t *testing.T) {
	permit := consentFixture(t, "r5_screening_permit.json")

	tests := []struct {
		name       string
		handler    http.HandlerFunc
		req        consent.Request
		wantErr    bool
		wantValid  bool
		wantReason string
		wantID     string
	}{
		{
			name: "found and valid",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet || !strings.HasSuffix(r.URL.Path, "/Consent/scr-1") {
					t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
				}

				writeFHIR(w, http.StatusOK, permit)
			},
			req:       consent.Request{PatientID: "pat-1", ConsentID: "scr-1", At: consentAt},
			wantValid: true,
			wantID:    "scr-1",
		},
		{
			name:       "wrong patient",
			handler:    func(w http.ResponseWriter, _ *http.Request) { writeFHIR(w, http.StatusOK, permit) },
			req:        consent.Request{PatientID: "pat-2", ConsentID: "scr-1", At: consentAt},
			wantReason: consent.ReasonPatientMismatch,
		},
		{
			name: "404 is a decision, not an error",
			handler: func(w http.ResponseWriter, _ *http.Request) {
				writeFHIR(w, http.StatusNotFound, []byte(consentNotFoundOutcome))
			},
			req:        consent.Request{PatientID: "pat-1", ConsentID: "nope", At: consentAt},
			wantReason: consent.ReasonConsentNotFound,
			wantID:     "nope",
		},
		{
			name: "500 is an error",
			handler: func(w http.ResponseWriter, _ *http.Request) {
				writeFHIR(w, http.StatusInternalServerError, []byte(operationOutcome500))
			},
			req:     consent.Request{PatientID: "pat-1", ConsentID: "scr-1", At: consentAt},
			wantErr: true,
		},
		{
			name: "wrong resource type is malformed",
			handler: func(w http.ResponseWriter, _ *http.Request) {
				writeFHIR(w, http.StatusOK, []byte(`{"resourceType":"Patient","id":"scr-1"}`))
			},
			req:        consent.Request{PatientID: "pat-1", ConsentID: "scr-1", At: consentAt},
			wantReason: consent.ReasonMalformedConsent,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := consentClient(t, tc.handler)

			res, err := c.ValidateConsent(context.Background(), tc.req, consent.Policy{})
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}

				return
			}

			if err != nil {
				t.Fatal(err)
			}

			if res.Valid != tc.wantValid {
				t.Fatalf("Valid = %v, want %v: %+v", res.Valid, tc.wantValid, res.Reasons)
			}

			if tc.wantReason != "" && !hasReasonCode(res, tc.wantReason) {
				t.Errorf("reasons %+v lack %s", res.Reasons, tc.wantReason)
			}

			if tc.wantID != "" && res.ConsentID != tc.wantID {
				t.Errorf("ConsentID = %q, want %q", res.ConsentID, tc.wantID)
			}
		})
	}
}
