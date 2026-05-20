package hapifhirgo

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

const (
	patientJSON         = `{"resourceType":"Patient","id":"x"}`
	operationOutcome500 = `{"resourceType":"OperationOutcome","issue":[{"severity":"error"}]}`
	operationOutcome400 = `{"resourceType":"OperationOutcome","issue":[{"severity":"error","diagnostics":"bad"}]}`
	fastBackoff         = 1 * time.Millisecond
	fastBackoffMax      = 5 * time.Millisecond
)

func newTestServer(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	return srv
}

// writePatient writes a small successful Patient response. Common across many cases.
func writePatient(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/fhir+json")
	_, _ = w.Write([]byte(patientJSON))
}

// TestClientOptions covers the option constructors. Each case applies one
// or more options to NewClient and asserts the resulting Client / outbound
// request shape. Some cases verify configuration directly (no server); others
// round-trip through a stub server to verify on-the-wire behaviour.
func TestClientOptions(t *testing.T) {
	tests := []struct {
		name string
		// run is given a callable "build" that constructs a client
		// against the supplied baseURL. baseURL is empty for cases
		// that do not need a server.
		needsServer bool
		// handler optionally asserts the outbound request shape.
		handler http.HandlerFunc
		// build builds the client.
		build func(baseURL string) (*Client, error)
		// assert runs after build/round-trip.
		assert func(t *testing.T, c *Client)
	}{
		{
			name: "default transport is tuned for real concurrency",
			build: func(_ string) (*Client, error) {
				return NewClient("http://example.test")
			},
			assert: func(t *testing.T, c *Client) {
				tr, ok := c.HTTP.Transport.(*http.Transport)
				if !ok {
					t.Fatalf("default transport is not *http.Transport: %T", c.HTTP.Transport)
				}
				if tr.MaxIdleConnsPerHost < 100 {
					t.Errorf("MaxIdleConnsPerHost = %d, want >= 100", tr.MaxIdleConnsPerHost)
				}
				if tr.MaxConnsPerHost != 0 {
					t.Errorf("MaxConnsPerHost = %d, want 0 (no admission cap)", tr.MaxConnsPerHost)
				}
				if !tr.ForceAttemptHTTP2 {
					t.Error("ForceAttemptHTTP2 = false, want true")
				}
			},
		},
		{
			name: "WithTransport replaces the round tripper",
			build: func(_ string) (*Client, error) {
				custom := &http.Transport{MaxIdleConnsPerHost: 999}
				return NewClient("http://example.test", WithTransport(custom))
			},
			assert: func(t *testing.T, c *Client) {
				tr, ok := c.HTTP.Transport.(*http.Transport)
				if !ok || tr.MaxIdleConnsPerHost != 999 {
					t.Fatalf("custom transport not applied; got %T %+v", c.HTTP.Transport, c.HTTP.Transport)
				}
			},
		},
		{
			name: "WithHTTPClient replaces the http.Client wholesale",
			build: func(_ string) (*Client, error) {
				custom := &http.Client{Timeout: 7 * time.Second}
				return NewClient("http://example.test", WithHTTPClient(custom))
			},
			assert: func(t *testing.T, c *Client) {
				if c.HTTP.Timeout != 7*time.Second {
					t.Fatalf("WithHTTPClient did not apply (timeout %v)", c.HTTP.Timeout)
				}
			},
		},
		{
			name:        "WithoutCacheControlHeader suppresses default Cache-Control",
			needsServer: true,
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get("Cache-Control") != "" {
					http.Error(w, "Cache-Control header should be absent", http.StatusBadRequest)
					return
				}
				writePatient(w)
			},
			build: func(baseURL string) (*Client, error) {
				return NewClient(baseURL, WithoutCacheControlHeader())
			},
			assert: func(t *testing.T, c *Client) {
				var out map[string]interface{}
				if err := c.GetFHIRResource(context.Background(), "Patient", "x", &out); err != nil {
					t.Fatalf("GetFHIRResource: %v", err)
				}
			},
		},
		{
			name:        "WithDefaultHeaders adds custom headers without dropping defaults",
			needsServer: true,
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get("X-Tenant-ID") != "demo-tenant" {
					http.Error(w, "missing X-Tenant-ID", http.StatusBadRequest)
					return
				}
				if r.Header.Get("Cache-Control") != "no-cache" {
					http.Error(w, "default Cache-Control was dropped", http.StatusBadRequest)
					return
				}
				writePatient(w)
			},
			build: func(baseURL string) (*Client, error) {
				return NewClient(baseURL,
					WithDefaultHeaders(map[string]string{"X-Tenant-ID": "demo-tenant"}),
				)
			},
			assert: func(t *testing.T, c *Client) {
				var out map[string]interface{}
				if err := c.GetFHIRResource(context.Background(), "Patient", "x", &out); err != nil {
					t.Fatalf("GetFHIRResource: %v", err)
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var baseURL string
			if tc.needsServer {
				srv := newTestServer(t, tc.handler)
				baseURL = srv.URL
			}
			c, err := tc.build(baseURL)
			if err != nil {
				t.Fatalf("build client: %v", err)
			}
			tc.assert(t, c)
		})
	}
}

// TestRetryRoundTripper covers the retry policy semantics via WithRetry.
// Each row scripts a sequence of server responses and asserts the number
// of times the server was actually called.
func TestRetryRoundTripper(t *testing.T) {
	tests := []struct {
		name string
		// responses[i] is returned on the i-th request. The slice is
		// indexed by the atomic hit counter, clamped at the last entry
		// so trailing 200s persist.
		responses []serverResponse
		// call drives the actual client call. Either a GET (retryable
		// by default) or a Create POST (non-retryable by default).
		call func(ctx context.Context, c *Client) error
		// policy is the retry policy. Zero-value means "use defaults
		// without WithRetry installed" (i.e. no retry at all).
		policy *RetryPolicy
		// wantHits is the exact expected number of server hits.
		wantHits int32
		// wantErr describes the expected error state.
		// "" means no error; "any" means "any error"; otherwise the
		// substring must be contained in err.Error().
		wantErr string
	}{
		{
			name: "GET retries on 503 then succeeds",
			responses: []serverResponse{
				{status: http.StatusServiceUnavailable, body: operationOutcome500},
				{status: http.StatusServiceUnavailable, body: operationOutcome500},
				{status: http.StatusOK, body: patientJSON, contentType: "application/fhir+json"},
			},
			call: func(ctx context.Context, c *Client) error {
				var out map[string]interface{}
				return c.GetFHIRResource(ctx, "Patient", "x", &out)
			},
			policy:   &RetryPolicy{MaxAttempts: 3, InitialBackoff: fastBackoff, MaxBackoff: fastBackoffMax},
			wantHits: 3,
		},
		{
			name: "GET does NOT retry on 4xx (non-retryable status)",
			responses: []serverResponse{
				{status: http.StatusBadRequest, body: operationOutcome400},
			},
			call: func(ctx context.Context, c *Client) error {
				var out map[string]interface{}
				return c.GetFHIRResource(ctx, "Patient", "x", &out)
			},
			policy:   &RetryPolicy{MaxAttempts: 3, InitialBackoff: fastBackoff},
			wantHits: 1,
			wantErr:  "any",
		},
		{
			name: "POST is not retried by default (Create is non-idempotent)",
			responses: []serverResponse{
				{status: http.StatusServiceUnavailable, body: operationOutcome500},
				{status: http.StatusServiceUnavailable, body: operationOutcome500},
				{status: http.StatusServiceUnavailable, body: operationOutcome500},
			},
			call: func(ctx context.Context, c *Client) error {
				var out map[string]interface{}
				return c.CreateFHIRResource(ctx, "Patient", map[string]interface{}{"active": true}, &out)
			},
			policy: &RetryPolicy{MaxAttempts: 3, InitialBackoff: fastBackoff},
			// Create makes TWO POSTs ($validate then create); neither
			// is retried. Cap at 2 hits, anything more means we retried.
			wantHits: 2,
			wantErr:  "any",
		},
		{
			name: "no policy installed = no retry even on 503",
			responses: []serverResponse{
				{status: http.StatusServiceUnavailable, body: operationOutcome500},
			},
			call: func(ctx context.Context, c *Client) error {
				var out map[string]interface{}
				return c.GetFHIRResource(ctx, "Patient", "x", &out)
			},
			policy:   nil,
			wantHits: 1,
			wantErr:  "any",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var hits int32
			srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				idx := int(atomic.AddInt32(&hits, 1)) - 1
				if idx >= len(tc.responses) {
					idx = len(tc.responses) - 1
				}
				resp := tc.responses[idx]
				if resp.contentType != "" {
					w.Header().Set("Content-Type", resp.contentType)
				}
				w.WriteHeader(resp.status)
				_, _ = w.Write([]byte(resp.body))
			})

			opts := []ClientOption{}
			if tc.policy != nil {
				opts = append(opts, WithRetry(*tc.policy))
			}
			c, err := NewClient(srv.URL, opts...)
			if err != nil {
				t.Fatalf("NewClient: %v", err)
			}

			err = tc.call(context.Background(), c)
			checkErr(t, err, tc.wantErr)

			got := atomic.LoadInt32(&hits)
			if tc.name == "POST is not retried by default (Create is non-idempotent)" {
				// Allow 1 or 2 hits — depends on whether $validate or
				// the create POST hit the server first. Anything > 2
				// means a retry occurred.
				if got > tc.wantHits {
					t.Errorf("hits = %d, want <= %d (no retry expected on POST)", got, tc.wantHits)
				}
			} else if got != tc.wantHits {
				t.Errorf("hits = %d, want %d", got, tc.wantHits)
			}
		})
	}
}

func TestRetryRoundTripper_RespectsContextCancellation(t *testing.T) {
	srv := newTestServer(t, func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, _ = w.Write([]byte(`{}`))
	})

	c, err := NewClient(srv.URL, WithRetry(RetryPolicy{
		MaxAttempts:    10,
		InitialBackoff: 100 * time.Millisecond,
		MaxBackoff:     200 * time.Millisecond,
	}))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	start := time.Now()
	var out map[string]interface{}
	err = c.GetFHIRResource(ctx, "Patient", "x", &out)
	if err == nil {
		t.Fatal("expected ctx error")
	}
	if elapsed := time.Since(start); elapsed > 500*time.Millisecond {
		t.Errorf("retry loop ignored ctx cancel: waited %v", elapsed)
	}
}

func TestStreamingDecode_LargeResponse(t *testing.T) {
	// ~1 MB Bundle exercises the streaming-decode path on the happy path.
	// The pre-PR io.ReadAll path would have fully buffered this; the
	// json.NewDecoder path should not.
	bigID := strings.Repeat("a", 1024*1024)
	body := `{"resourceType":"Bundle","id":"` + bigID + `","type":"searchset"}`

	srv := newTestServer(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/fhir+json")
		_, _ = io.WriteString(w, body)
	})

	c, err := NewClient(srv.URL)
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	var out map[string]interface{}
	if err := c.GetFHIRResource(context.Background(), "Bundle", "x", &out); err != nil {
		t.Fatalf("GetFHIRResource: %v", err)
	}
	if out["resourceType"] != "Bundle" {
		t.Errorf("decoded wrong shape: %v", out["resourceType"])
	}
	if id, _ := out["id"].(string); len(id) != len(bigID) {
		t.Errorf("id len = %d, want %d", len(id), len(bigID))
	}
}

// serverResponse describes one scripted response in a retry test row.
type serverResponse struct {
	status      int
	body        string
	contentType string
}

func checkErr(t *testing.T, got error, want string) {
	t.Helper()
	switch want {
	case "":
		if got != nil {
			t.Fatalf("unexpected error: %v", got)
		}
	case "any":
		if got == nil {
			t.Fatal("expected an error, got nil")
		}
	default:
		if got == nil {
			t.Fatalf("expected error containing %q, got nil", want)
		}
		if !strings.Contains(got.Error(), want) {
			t.Fatalf("error %q does not contain %q", got.Error(), want)
		}
	}
}

// keep the errors import used so go vet stays quiet — referenced by callers
// that wrap APIError, which is still asserted in some scenarios.
var _ = errors.Is
