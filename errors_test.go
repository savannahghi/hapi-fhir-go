package hapifhirgo

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestAPIError(t *testing.T) {
	tests := []struct {
		name            string
		status          int
		contentType     string
		body            string
		wantIssues      int
		wantDiagnostics string
		wantOutcome     bool
		wantError       string
	}{
		{
			name:            "sad case: an operation outcome is decoded and its issues parsed",
			status:          http.StatusNotFound,
			contentType:     "application/fhir+json",
			body:            `{"resourceType":"OperationOutcome","issue":[{"severity":"error","code":"not-found","diagnostics":"Resource Patient/x is not known"}]}`,
			wantIssues:      1,
			wantDiagnostics: "Resource Patient/x is not known",
			wantOutcome:     true,
			wantError:       "FHIR error (HTTP 404)",
		},
		{
			name:            "sad case: the details text stands in for missing diagnostics, and warnings are left out",
			status:          http.StatusUnprocessableEntity,
			contentType:     "application/fhir+json",
			body:            `{"resourceType":"OperationOutcome","issue":[{"severity":"warning","diagnostics":"unknown extension"},{"severity":"error","details":{"text":"Observation.code is required"}},{"severity":"fatal","diagnostics":"no subject"}]}`,
			wantIssues:      3,
			wantDiagnostics: "Observation.code is required; no subject",
			wantOutcome:     true,
			wantError:       "FHIR error (HTTP 422)",
		},
		{
			name:        "sad case: a gateway page keeps its status and body with no outcome",
			status:      http.StatusBadGateway,
			contentType: "text/html",
			body:        `<html><body>502 Bad Gateway</body></html>`,
			wantIssues:  0,
			wantOutcome: false,
			wantError:   "FHIR error (HTTP 502)",
		},
		{
			name:        "sad case: an empty answer still carries its status",
			status:      http.StatusServiceUnavailable,
			contentType: "application/fhir+json",
			body:        "",
			wantIssues:  0,
			wantOutcome: false,
			wantError:   "FHIR error (HTTP 503)",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			srv := newTestServer(t, func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", tc.contentType)
				w.WriteHeader(tc.status)
				_, _ = w.Write([]byte(tc.body))
			})

			c, err := NewClient(srv.URL)
			if err != nil {
				t.Fatalf("build client: %v", err)
			}

			var out map[string]interface{}

			err = c.GetFHIRResource(context.Background(), "Patient", "x", &out)

			var apiErr APIError
			if !errors.As(err, &apiErr) {
				t.Fatalf("err = %T %v, want an APIError", err, err)
			}

			if apiErr.StatusCode != tc.status {
				t.Errorf("status = %d, want %d", apiErr.StatusCode, tc.status)
			}

			if len(apiErr.Issues) != tc.wantIssues {
				t.Errorf("issues = %d, want %d", len(apiErr.Issues), tc.wantIssues)
			}

			if got := apiErr.Diagnostics(); got != tc.wantDiagnostics {
				t.Errorf("diagnostics = %q, want %q", got, tc.wantDiagnostics)
			}

			if (apiErr.OperationOutcome != nil) != tc.wantOutcome {
				t.Errorf("outcome present = %v, want %v", apiErr.OperationOutcome != nil, tc.wantOutcome)
			}

			if string(apiErr.Body) != tc.body {
				t.Errorf("body = %q, want %q", apiErr.Body, tc.body)
			}

			if !strings.HasPrefix(err.Error(), tc.wantError) {
				t.Errorf("Error() = %q, want prefix %q", err.Error(), tc.wantError)
			}
		})
	}
}
