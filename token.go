package hapifhirgo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	// defaultRefreshMargin is how long before expiry a cached token is minted again, so a
	// request never carries one that dies in flight.
	defaultRefreshMargin = 30 * time.Second

	defaultMintTimeout = 15 * time.Second
	maxTokenBody       = 1 << 20
)

// TokenMinter mints a fresh access token and says how long it lives. ClientCredentials is one;
// anything that talks to an identity provider can be another.
type TokenMinter func(ctx context.Context) (token string, expiresIn time.Duration, err error)

// CacheOption tunes CachedTokenProvider.
type CacheOption func(*cachedTokenProvider)

// WithRefreshMargin sets how long before expiry the cached token is minted again. The default
// is 30 seconds.
func WithRefreshMargin(d time.Duration) CacheOption {
	return func(p *cachedTokenProvider) {
		if d > 0 {
			p.margin = d
		}
	}
}

// CachedTokenProvider wraps a minter in a cache for WithTokenProvider: one token serves every
// request until it is about to expire, and callers arriving during a mint wait for its result
// rather than each asking the identity provider.
func CachedTokenProvider(mint TokenMinter, opts ...CacheOption) TokenProvider {
	p := &cachedTokenProvider{mint: mint, margin: defaultRefreshMargin, now: time.Now}
	for _, opt := range opts {
		opt(p)
	}

	return p.current
}

type cachedTokenProvider struct {
	mint   TokenMinter
	margin time.Duration
	now    func() time.Time

	mu     sync.Mutex
	token  string
	expiry time.Time
}

func (p *cachedTokenProvider) current(ctx context.Context) (string, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.token != "" && p.now().Add(p.margin).Before(p.expiry) {
		return p.token, nil
	}

	token, expiresIn, err := p.mint(ctx)
	if err != nil {
		return "", fmt.Errorf("hapifhirgo: mint token: %w", err)
	}

	if strings.TrimSpace(token) == "" {
		return "", errors.New("hapifhirgo: mint token: the identity provider answered without a token")
	}

	p.token = token
	p.expiry = p.now().Add(expiresIn)

	return p.token, nil
}

// ClientCredentials mints tokens from an OAuth2 token endpoint with the client credentials
// grant, which is how a service authenticates as itself. Keycloak's endpoint is
// {base}/realms/{realm}/protocol/openid-connect/token. Pass its Mint to CachedTokenProvider.
type ClientCredentials struct {
	TokenURL     string
	ClientID     string
	ClientSecret string
	Scopes       []string
	// HTTP is the client used to reach the endpoint. nil means a client with a 15 second
	// timeout, so a slow identity provider cannot stall the requests waiting on a token.
	HTTP *http.Client
}

type tokenResponse struct {
	AccessToken string      `json:"access_token"`
	ExpiresIn   json.Number `json:"expires_in"`
	Error       string      `json:"error"`
	Description string      `json:"error_description"`
}

// Mint asks the endpoint for a token. The error on a refusal names the endpoint's reason and
// never carries the secret.
func (c ClientCredentials) Mint(ctx context.Context) (string, time.Duration, error) {
	if c.TokenURL == "" || c.ClientID == "" {
		return "", 0, errors.New("hapifhirgo: client credentials need a token url and a client id")
	}

	form := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {c.ClientID},
		"client_secret": {c.ClientSecret},
	}
	if len(c.Scopes) > 0 {
		form.Set("scope", strings.Join(c.Scopes, " "))
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.TokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", 0, fmt.Errorf("hapifhirgo: token request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := c.HTTP
	if client == nil {
		client = &http.Client{Timeout: defaultMintTimeout}
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("hapifhirgo: token request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, maxTokenBody))
	if err != nil {
		return "", 0, fmt.Errorf("hapifhirgo: token response: %w", err)
	}

	var answer tokenResponse
	if err := json.Unmarshal(body, &answer); err != nil && resp.StatusCode < 400 {
		return "", 0, fmt.Errorf("hapifhirgo: token response (HTTP %d) is not JSON: %w", resp.StatusCode, err)
	}

	if resp.StatusCode >= 400 {
		reason := strings.TrimSpace(answer.Error + " " + answer.Description)
		if reason == "" {
			reason = "no reason given"
		}

		return "", 0, fmt.Errorf("hapifhirgo: token refused (HTTP %d): %s", resp.StatusCode, reason)
	}

	if answer.AccessToken == "" {
		return "", 0, fmt.Errorf("hapifhirgo: token response (HTTP %d) has no access_token", resp.StatusCode)
	}

	seconds, err := answer.ExpiresIn.Float64()
	if err != nil || seconds <= 0 {
		return "", 0, fmt.Errorf("hapifhirgo: token response has no usable expires_in (%q)", answer.ExpiresIn)
	}

	return answer.AccessToken, time.Duration(seconds * float64(time.Second)), nil
}
