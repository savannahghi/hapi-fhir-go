package hapifhirgo

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"
)

const encounterJSON = `{"resourceType":"Encounter","id":"x","status":"completed"}`

// capturedRequest holds the parts of an outbound request a test wants to assert on.
type capturedRequest struct {
	method  string
	headers http.Header
	body    []byte
}

// captureAndWrite records the inbound request, then writes body back to the caller.
func captureAndWrite(t *testing.T, got *capturedRequest, body string) http.HandlerFunc {
	t.Helper()

	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("read request body: %v", err)
		}

		got.method = r.Method
		got.headers = r.Header.Clone()
		got.body = payload

		w.Header().Set("Content-Type", fhirContentType)
		_, _ = io.WriteString(w, body)
	}
}

// TestFHIRPathPatch_SendsJSONPatchContentType pins the media type of the patch
// body. FHIRPathPatch builds an RFC 6902 array, so labelling it as a FHIR
// resource makes HAPI reject it with "Content does not appear to be FHIR JSON,
// first non-whitespace character was: '['". This has regressed twice, so assert
// the wire format directly.
func TestFHIRPathPatch_SendsJSONPatchContentType(t *testing.T) {
	var got capturedRequest

	srv := newTestServer(t, captureAndWrite(t, &got, encounterJSON))

	c, err := NewClient(srv.URL)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	resource := map[string]interface{}{}
	if err := c.FHIRPathPatch(
		context.Background(),
		"Encounter",
		"x",
		map[string]interface{}{"status": "completed"},
		&resource,
	); err != nil {
		t.Fatalf("FHIRPathPatch: %v", err)
	}

	if got.method != http.MethodPatch {
		t.Errorf("method = %s, want %s", got.method, http.MethodPatch)
	}

	if ct := got.headers.Get("Content-Type"); ct != jsonPatchContentType {
		t.Errorf("Content-Type = %q, want %q", ct, jsonPatchContentType)
	}

	// the response is still a FHIR resource, so Accept must not change
	if accept := got.headers.Get("Accept"); accept != fhirContentType {
		t.Errorf("Accept = %q, want %q", accept, fhirContentType)
	}

	var patches []map[string]interface{}
	if err := json.Unmarshal(got.body, &patches); err != nil {
		t.Fatalf("patch body is not a JSON array: %v (body: %s)", err, got.body)
	}

	if len(patches) != 1 {
		t.Fatalf("sent %d patches, want 1", len(patches))
	}

	if patches[0]["op"] != "replace" || patches[0]["path"] != "/status" || patches[0]["value"] != "completed" {
		t.Errorf("patch = %v, want replace /status completed", patches[0])
	}

	if resource["status"] != "completed" {
		t.Errorf("patched resource not decoded: %v", resource)
	}
}

// TestFHIRPathPatch_KeepsClientHeaders guards the shape of the previous fix for
// this bug (v1.15.2), which special-cased PATCH by skipping setHeaders entirely
// and so dropped Accept, Cache-Control and WithDefaultHeaders from every patch.
func TestFHIRPathPatch_KeepsClientHeaders(t *testing.T) {
	var got capturedRequest

	srv := newTestServer(t, captureAndWrite(t, &got, encounterJSON))

	c, err := NewClient(
		srv.URL,
		WithBearerToken("token"),
		WithDefaultHeaders(map[string]string{"X-Tenant": "tenant-1"}),
	)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	resource := map[string]interface{}{}
	if err := c.FHIRPathPatch(
		context.Background(),
		"Encounter",
		"x",
		map[string]interface{}{"status": "completed"},
		&resource,
	); err != nil {
		t.Fatalf("FHIRPathPatch: %v", err)
	}

	for header, want := range map[string]string{
		"Authorization": "Bearer token",
		"X-Tenant":      "tenant-1",
		"Cache-Control": "no-cache",
		"Content-Type":  jsonPatchContentType,
	} {
		if value := got.headers.Get(header); value != want {
			t.Errorf("%s = %q, want %q", header, value, want)
		}
	}
}

// TestFHIRPathPatch_SkipsEmptyValues covers the payload filter. A nil value used
// to panic: reflect.ValueOf(nil) is an invalid Value and IsZero panics on it.
func TestFHIRPathPatch_SkipsEmptyValues(t *testing.T) {
	var got capturedRequest

	srv := newTestServer(t, captureAndWrite(t, &got, encounterJSON))

	c, err := NewClient(srv.URL)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	resource := map[string]interface{}{}
	if err := c.FHIRPathPatch(
		context.Background(),
		"Encounter",
		"x",
		map[string]interface{}{
			"status":  "completed",
			"empty":   "",
			"missing": nil,
		},
		&resource,
	); err != nil {
		t.Fatalf("FHIRPathPatch: %v", err)
	}

	var patches []map[string]interface{}
	if err := json.Unmarshal(got.body, &patches); err != nil {
		t.Fatalf("patch body is not a JSON array: %v", err)
	}

	if len(patches) != 1 {
		t.Fatalf("sent %d patches, want only the non-empty one: %v", len(patches), patches)
	}

	if patches[0]["path"] != "/status" {
		t.Errorf("patched %v, want /status", patches[0]["path"])
	}
}

// TestFHIRPathPatch_SurfacesOperationOutcome checks that a rejected patch still
// returns the server's OperationOutcome rather than a decode error.
func TestFHIRPathPatch_SurfacesOperationOutcome(t *testing.T) {
	srv := newTestServer(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", fhirContentType)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = io.WriteString(w, operationOutcome400)
	})

	c, err := NewClient(srv.URL)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	resource := map[string]interface{}{}

	err = c.FHIRPathPatch(
		context.Background(),
		"Encounter",
		"x",
		map[string]interface{}{"status": "completed"},
		&resource,
	)
	if err == nil {
		t.Fatal("expected an error for a 400 response")
	}

	var apiErr APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error is not an APIError: %v", err)
	}

	if apiErr.StatusCode != http.StatusBadRequest {
		t.Errorf("StatusCode = %d, want %d", apiErr.StatusCode, http.StatusBadRequest)
	}
}

// TestPutFHIRResource_SendsFHIRContentType is the counterpart assertion: only
// FHIRPathPatch overrides the media type, everything else still sends resources
// as application/fhir+json.
func TestPutFHIRResource_SendsFHIRContentType(t *testing.T) {
	var got capturedRequest

	srv := newTestServer(t, captureAndWrite(t, &got, patientJSON))

	c, err := NewClient(srv.URL)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	resource := map[string]interface{}{}
	if err := c.PutFHIRResource(
		context.Background(),
		"Patient",
		"x",
		map[string]interface{}{"resourceType": "Patient", "id": "x"},
		&resource,
		false,
	); err != nil {
		t.Fatalf("PutFHIRResource: %v", err)
	}

	if ct := got.headers.Get("Content-Type"); ct != fhirContentType {
		t.Errorf("Content-Type = %q, want %q", ct, fhirContentType)
	}
}
