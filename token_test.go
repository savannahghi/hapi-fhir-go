package hapifhirgo

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// minter stands in for an identity provider: each call mints the next numbered token, or
// nothing at all when blank.
type minter struct {
	calls    atomic.Int32
	lifetime time.Duration
	blank    bool
	err      error
}

func (m *minter) mint(context.Context) (string, time.Duration, error) {
	n := m.calls.Add(1)

	switch {
	case m.err != nil:
		return "", 0, m.err
	case m.blank:
		return "", m.lifetime, nil
	default:
		return "t-" + strconv.Itoa(int(n)), m.lifetime, nil
	}
}

func TestCachedTokenProvider(t *testing.T) {
	tests := []struct {
		name      string
		minter    *minter
		opts      []CacheOption
		calls     int
		wantToken string
		wantMints int32
		wantErr   string
	}{
		{
			name:      "happy case: one mint serves every call while the token is fresh",
			minter:    &minter{lifetime: time.Hour},
			calls:     3,
			wantToken: "t-1",
			wantMints: 1,
		},
		{
			name:      "happy case: a token inside the refresh margin is minted again",
			minter:    &minter{lifetime: 10 * time.Second},
			calls:     2,
			wantToken: "t-2",
			wantMints: 2,
		},
		{
			name:      "happy case: a wider margin mints earlier",
			minter:    &minter{lifetime: 2 * time.Minute},
			opts:      []CacheOption{WithRefreshMargin(5 * time.Minute)},
			calls:     2,
			wantToken: "t-2",
			wantMints: 2,
		},
		{
			name:    "sad case: a refused mint names the step",
			minter:  &minter{err: errors.New("401 invalid_client")},
			calls:   1,
			wantErr: "mint token: 401 invalid_client",
		},
		{
			name:    "sad case: an empty token is refused",
			minter:  &minter{blank: true, lifetime: time.Hour},
			calls:   1,
			wantErr: "answered without a token",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			provider := CachedTokenProvider(tc.minter.mint, tc.opts...)

			var (
				token string
				err   error
			)

			for range tc.calls {
				token, err = provider(context.Background())
			}

			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("err = %v, want it to contain %q", err, tc.wantErr)
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if token != tc.wantToken {
				t.Errorf("token = %q, want %q", token, tc.wantToken)
			}

			if got := tc.minter.calls.Load(); got != tc.wantMints {
				t.Errorf("mints = %d, want %d", got, tc.wantMints)
			}
		})
	}
}

func TestCachedTokenProvider_ConcurrentCallers(t *testing.T) {
	tests := []struct {
		name      string
		callers   int
		wantMints int32
	}{
		{name: "happy case: one caller mints once", callers: 1, wantMints: 1},
		{name: "happy case: fifty callers at once share one mint", callers: 50, wantMints: 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := &minter{lifetime: time.Hour}
			provider := CachedTokenProvider(m.mint)

			var wg sync.WaitGroup

			tokens := make(chan string, tc.callers)

			for range tc.callers {
				wg.Go(func() {
					token, err := provider(context.Background())
					if err != nil {
						tokens <- "error: " + err.Error()

						return
					}

					tokens <- token
				})
			}

			wg.Wait()
			close(tokens)

			for token := range tokens {
				if token != "t-1" {
					t.Errorf("token = %q, want t-1", token)
				}
			}

			if got := m.calls.Load(); got != tc.wantMints {
				t.Errorf("mints = %d, want %d", got, tc.wantMints)
			}
		})
	}
}

func TestClientCredentials_Mint(t *testing.T) {
	tests := []struct {
		name        string
		credentials ClientCredentials
		status      int
		reply       string
		wantForm    map[string]string
		wantToken   string
		wantTTL     time.Duration
		wantErr     string
	}{
		{
			name:        "happy case: the grant goes as a form and the token comes back with its lifetime",
			credentials: ClientCredentials{ClientID: "svc", ClientSecret: "s3cret", Scopes: []string{"fhir", "read"}},
			status:      http.StatusOK,
			reply:       `{"access_token":"abc","expires_in":300,"token_type":"Bearer"}`,
			wantForm:    map[string]string{"grant_type": "client_credentials", "client_id": "svc", "client_secret": "s3cret", "scope": "fhir read"},
			wantToken:   "abc",
			wantTTL:     300 * time.Second,
		},
		{
			name:        "happy case: a lifetime sent as a string is read",
			credentials: ClientCredentials{ClientID: "svc", ClientSecret: "s3cret"},
			status:      http.StatusOK,
			reply:       `{"access_token":"abc","expires_in":"60"}`,
			wantToken:   "abc",
			wantTTL:     time.Minute,
		},
		{
			name:        "sad case: a refused client carries the endpoint's reason and not the secret",
			credentials: ClientCredentials{ClientID: "svc", ClientSecret: "s3cret"},
			status:      http.StatusUnauthorized,
			reply:       `{"error":"invalid_client","error_description":"Invalid client credentials"}`,
			wantErr:     "token refused (HTTP 401): invalid_client Invalid client credentials",
		},
		{
			name:        "sad case: a refusal with no body still names the status",
			credentials: ClientCredentials{ClientID: "svc", ClientSecret: "s3cret"},
			status:      http.StatusBadGateway,
			reply:       `<html>bad gateway</html>`,
			wantErr:     "token refused (HTTP 502): no reason given",
		},
		{
			name:        "sad case: an answer without a token is refused",
			credentials: ClientCredentials{ClientID: "svc", ClientSecret: "s3cret"},
			status:      http.StatusOK,
			reply:       `{"token_type":"Bearer"}`,
			wantErr:     "has no access_token",
		},
		{
			name:        "sad case: an answer without a lifetime is refused",
			credentials: ClientCredentials{ClientID: "svc", ClientSecret: "s3cret"},
			status:      http.StatusOK,
			reply:       `{"access_token":"abc"}`,
			wantErr:     "no usable expires_in",
		},
		{
			name:        "sad case: a client id is required",
			credentials: ClientCredentials{ClientSecret: "s3cret"},
			wantErr:     "need a token url and a client id",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var gotForm map[string]string

			srv := newTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				if err := r.ParseForm(); err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)

					return
				}

				gotForm = map[string]string{}
				for key := range r.PostForm {
					gotForm[key] = r.PostForm.Get(key)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(tc.status)
				_, _ = w.Write([]byte(tc.reply))
			})

			credentials := tc.credentials
			if credentials.ClientID != "" {
				credentials.TokenURL = srv.URL + "/realms/test/protocol/openid-connect/token"
			}

			token, ttl, err := credentials.Mint(context.Background())

			if tc.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
					t.Fatalf("err = %v, want it to contain %q", err, tc.wantErr)
				}

				if strings.Contains(err.Error(), "s3cret") {
					t.Fatal("the secret leaked into the error")
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if token != tc.wantToken || ttl != tc.wantTTL {
				t.Errorf("got (%q, %s), want (%q, %s)", token, ttl, tc.wantToken, tc.wantTTL)
			}

			for key, want := range tc.wantForm {
				if gotForm[key] != want {
					t.Errorf("form %s = %q, want %q", key, gotForm[key], want)
				}
			}
		})
	}
}
