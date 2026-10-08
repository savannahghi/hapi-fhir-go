package hapifhirgo

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"testing"
)

const observationJSON = `{"resourceType":"Observation","id":"o1","meta":{"versionId":"1"},"status":"final","code":{"text":"sbp"}}`

type observationModel struct {
	ResourceType string      `json:"resourceType"`
	ID           string      `json:"id"`
	Status       string      `json:"status"`
	Value        json.Number `json:"valueDecimal"`
}

func TestMetadata(t *testing.T) {
	tests := []struct {
		name        string
		status      int
		reply       string
		params      map[string]any
		wantQuery   string
		wantVersion string
		wantStatus  int
	}{
		{
			name:        "happy case: the capability statement is read without a token",
			status:      http.StatusOK,
			reply:       `{"resourceType":"CapabilityStatement","fhirVersion":"5.0.0"}`,
			params:      map[string]any{"_summary": "true"},
			wantQuery:   "_summary=true",
			wantVersion: "5.0.0",
		},
		{
			name:       "sad case: a store that is down answers its status",
			status:     http.StatusServiceUnavailable,
			reply:      "",
			wantStatus: http.StatusServiceUnavailable,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var got *http.Request

			srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				got = r.Clone(context.Background())
				w.Header().Set("Content-Type", "application/fhir+json")
				w.WriteHeader(tc.status)
				_, _ = w.Write([]byte(tc.reply))
			})

			var minted atomic.Int32

			c, err := NewClient(srv.URL, WithTokenProvider(func(context.Context) (string, error) {
				minted.Add(1)

				return "token", nil
			}))
			if err != nil {
				t.Fatalf("build client: %v", err)
			}

			var capability struct {
				FHIRVersion string `json:"fhirVersion"`
			}

			err = c.Metadata(context.Background(), tc.params, &capability)

			if minted.Load() != 0 {
				t.Fatal("metadata asked for a token")
			}

			if tc.wantStatus != 0 {
				var apiErr APIError
				if !asAPIError(err, &apiErr) || apiErr.StatusCode != tc.wantStatus {
					t.Fatalf("err = %v, want an APIError with status %d", err, tc.wantStatus)
				}

				return
			}

			if err != nil {
				t.Fatalf("Metadata: %v", err)
			}

			if got.URL.Path != "/metadata" || got.URL.RawQuery != tc.wantQuery {
				t.Errorf("request = %s?%s, want /metadata?%s", got.URL.Path, got.URL.RawQuery, tc.wantQuery)
			}

			if got.Header.Get("Authorization") != "" {
				t.Error("metadata carried credentials")
			}

			if capability.FHIRVersion != tc.wantVersion {
				t.Errorf("fhirVersion = %q, want %q", capability.FHIRVersion, tc.wantVersion)
			}
		})
	}
}

func TestPutResource(t *testing.T) {
	tests := []struct {
		name         string
		resourceType string
		id           string
		resource     interface{}
		wantBody     string
		wantInBody   []string
		wantErr      string
		wantBad      bool
		wantCalled   bool
	}{
		{
			name:         "happy case: a resource carrying its type is sent as it is",
			resourceType: "Observation",
			id:           "o1",
			resource:     observationModel{ResourceType: "Observation", ID: "o1", Status: "final", Value: "128.50"},
			wantBody:     `{"resourceType":"Observation","id":"o1","status":"final","valueDecimal":128.50}`,
			wantCalled:   true,
		},
		{
			name:         "happy case: a resource without a type is stamped and its numbers kept",
			resourceType: "Observation",
			id:           "o1",
			resource:     map[string]any{"id": "o1", "valueQuantity": map[string]any{"value": json.Number("128.50")}},
			wantInBody:   []string{`"resourceType":"Observation"`, `"value":128.50`},
			wantCalled:   true,
		},
		{
			name:         "sad case: a resource of another type is refused before anything is sent",
			resourceType: "Observation",
			id:           "o1",
			resource:     map[string]any{"resourceType": "Patient"},
			wantErr:      "the resource is a Patient, not a Observation",
			wantBad:      true,
		},
		{
			name:         "sad case: something that is not an object is refused",
			resourceType: "Observation",
			id:           "o1",
			resource:     []string{"not", "a", "resource"},
			wantErr:      "not a JSON object",
			wantBad:      true,
		},
		{
			name:         "sad case: an id is required",
			resourceType: "Observation",
			resource:     map[string]any{"id": "o1"},
			wantErr:      "an id are required",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var (
				called atomic.Bool
				method string
				path   string
				body   string
			)

			srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				called.Store(true)
				method, path = r.Method, r.URL.Path

				raw, _ := io.ReadAll(r.Body)
				body = string(raw)

				w.Header().Set("Content-Type", "application/fhir+json")
				w.WriteHeader(http.StatusCreated)
				_, _ = w.Write([]byte(observationJSON))
			})

			c, err := NewClient(srv.URL)
			if err != nil {
				t.Fatalf("build client: %v", err)
			}

			var stored struct {
				Meta struct {
					VersionID string `json:"versionId"`
				} `json:"meta"`
			}

			err = c.PutResource(context.Background(), tc.resourceType, tc.id, tc.resource, &stored)

			if called.Load() != tc.wantCalled {
				t.Fatalf("called = %v, want %v", called.Load(), tc.wantCalled)
			}

			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("err = %v, want it to contain %q", err, tc.wantErr)
				}

				if errors.Is(err, ErrBadResource) != tc.wantBad {
					t.Fatalf("errors.Is(err, ErrBadResource) = %v, want %v", !tc.wantBad, tc.wantBad)
				}

				return
			}

			if err != nil {
				t.Fatalf("PutResource: %v", err)
			}

			if method != http.MethodPut || path != "/Observation/o1" {
				t.Errorf("request = %s %s, want PUT /Observation/o1", method, path)
			}

			if tc.wantBody != "" && body != tc.wantBody {
				t.Errorf("body = %s, want %s", body, tc.wantBody)
			}

			for _, want := range tc.wantInBody {
				if !strings.Contains(body, want) {
					t.Errorf("body %s lacks %s", body, want)
				}
			}

			if stored.Meta.VersionID != "1" {
				t.Errorf("stored version = %q, want 1", stored.Meta.VersionID)
			}
		})
	}
}

func TestCreateResource(t *testing.T) {
	tests := []struct {
		name         string
		resourceType string
		resource     interface{}
		wantBody     string
		wantErr      string
		wantCalls    int32
	}{
		{
			name:         "happy case: the resource is posted to its type with no validation call first",
			resourceType: "Observation",
			resource:     map[string]any{"id": "o1"},
			wantBody:     `{"id":"o1","resourceType":"Observation"}`,
			wantCalls:    1,
		},
		{
			name:         "sad case: a resource of another type is refused before anything is sent",
			resourceType: "Observation",
			resource:     observationModel{ResourceType: "Patient"},
			wantErr:      "the resource is a Patient, not a Observation",
		},
		{
			name:     "sad case: a resource type is required",
			resource: map[string]any{"id": "o1"},
			wantErr:  "a resource type is required",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var (
				calls atomic.Int32
				paths []string
				body  string
			)

			srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				calls.Add(1)
				paths = append(paths, r.Method+" "+r.URL.Path)

				raw, _ := io.ReadAll(r.Body)
				body = string(raw)

				w.Header().Set("Content-Type", "application/fhir+json")
				w.WriteHeader(http.StatusCreated)
				_, _ = w.Write([]byte(observationJSON))
			})

			c, err := NewClient(srv.URL)
			if err != nil {
				t.Fatalf("build client: %v", err)
			}

			err = c.CreateResource(context.Background(), tc.resourceType, tc.resource, nil)

			if calls.Load() != tc.wantCalls {
				t.Fatalf("calls = %d (%v), want %d", calls.Load(), paths, tc.wantCalls)
			}

			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("err = %v, want it to contain %q", err, tc.wantErr)
				}

				return
			}

			if err != nil {
				t.Fatalf("CreateResource: %v", err)
			}

			if paths[0] != "POST /Observation" {
				t.Errorf("request = %s, want POST /Observation", paths[0])
			}

			if body != tc.wantBody {
				t.Errorf("body = %s, want %s", body, tc.wantBody)
			}
		})
	}
}

// asAPIError is errors.As for the value type, kept local so the test reads as the caller
// would write it.
func asAPIError(err error, target *APIError) bool {
	for err != nil {
		if apiErr, ok := err.(APIError); ok {
			*target = apiErr

			return true
		}

		unwrapper, ok := err.(interface{ Unwrap() error })
		if !ok {
			return false
		}

		err = unwrapper.Unwrap()
	}

	return false
}
