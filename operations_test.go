package hapifhirgo

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

// TestInvokeOperation covers what separates this from CreateFHIRResource: the path and the
// body travel independently. Each case asserts the request that reached the server, because
// the whole point of the method is what it does *not* do to the payload on the way out.
func TestInvokeOperation(t *testing.T) {
	const submission = `{"resourceType":"Patient","id":"e174","active":true}`

	tests := []struct {
		name     string
		path     string
		payload  interface{}
		status   int
		reply    string
		wantPath string
		wantBody string
		wantErr  bool
		// wantCalled is false for cases that must fail before any request is sent.
		wantCalled bool
	}{
		{
			name:       "system level operation leaves the body untouched",
			path:       "$accept-from-exchange",
			payload:    json.RawMessage(submission),
			status:     http.StatusOK,
			reply:      `{"resourceType":"Parameters"}`,
			wantPath:   "/$accept-from-exchange",
			wantBody:   submission,
			wantCalled: true,
		},
		{
			name:       "instance level operation addresses the instance",
			path:       "Patient/123/$everything",
			payload:    json.RawMessage(`{"resourceType":"Parameters"}`),
			status:     http.StatusOK,
			reply:      `{"resourceType":"Bundle"}`,
			wantPath:   "/Patient/123/$everything",
			wantBody:   `{"resourceType":"Parameters"}`,
			wantCalled: true,
		},
		{
			name:       "a struct payload is marshalled as JSON",
			path:       "$accept-from-exchange",
			payload:    map[string]any{"resourceType": "Patient"},
			status:     http.StatusOK,
			reply:      `{"resourceType":"Parameters"}`,
			wantPath:   "/$accept-from-exchange",
			wantBody:   `{"resourceType":"Patient"}`,
			wantCalled: true,
		},
		{
			name:    "an empty path is refused before anything is sent",
			path:    "",
			payload: json.RawMessage(submission),
			wantErr: true,
		},
		{
			name:       "an operation outcome is reported as an error",
			path:       "$accept-from-exchange",
			payload:    json.RawMessage(submission),
			status:     http.StatusBadRequest,
			reply:      operationOutcome400,
			wantPath:   "/$accept-from-exchange",
			wantBody:   submission,
			wantErr:    true,
			wantCalled: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				gotPath   string
				gotBody   string
				gotMethod string
				called    bool
			)

			srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				called = true
				gotPath, gotMethod = r.URL.Path, r.Method

				body, _ := io.ReadAll(r.Body)
				gotBody = string(body)

				w.Header().Set("Content-Type", "application/fhir+json")
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.reply))
			})

			client, err := NewClient(srv.URL)
			if err != nil {
				t.Fatalf("NewClient() error = %v", err)
			}

			var result map[string]any

			err = client.InvokeOperation(context.Background(), tt.path, tt.payload, &result)
			if (err != nil) != tt.wantErr {
				t.Fatalf("InvokeOperation() error = %v, wantErr %v", err, tt.wantErr)
			}

			if called != tt.wantCalled {
				t.Fatalf("server called = %v, want %v", called, tt.wantCalled)
			}

			if !tt.wantCalled {
				return
			}

			if gotMethod != http.MethodPost {
				t.Errorf("method = %q, want POST", gotMethod)
			}

			if gotPath != tt.wantPath {
				t.Errorf("path = %q, want %q", gotPath, tt.wantPath)
			}

			if gotBody != tt.wantBody {
				t.Errorf("body = %q, want %q", gotBody, tt.wantBody)
			}
		})
	}
}

// TestInvokeOperation_DoesNotMutatePayload pins the difference from CreateFHIRResource,
// which writes resourceType and language into the map it is given. A caller here keeps the
// map it passed.
func TestInvokeOperation_DoesNotMutatePayload(t *testing.T) {
	srv := newTestServer(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/fhir+json")
		_, _ = w.Write([]byte(`{"resourceType":"Parameters"}`))
	})

	client, err := NewClient(srv.URL)
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}

	payload := map[string]any{"resourceType": "Patient"}

	var result map[string]any

	err = client.InvokeOperation(context.Background(), "$accept-from-exchange", payload, &result)
	if err != nil {
		t.Fatalf("InvokeOperation() error = %v", err)
	}

	if len(payload) != 1 || payload["resourceType"] != "Patient" {
		t.Errorf("payload was modified: %v", payload)
	}
}
