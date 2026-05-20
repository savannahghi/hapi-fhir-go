package hapifhirgo

import (
	"errors"
	"net"
	"net/http"
	"os"
	"time"
)

const (
	defaultTimeout = 1 * time.Minute

	defaultMaxIdleConns        = 400
	defaultMaxIdleConnsPerHost = 150
	defaultIdleConnTimeout     = 90 * time.Second
	defaultTLSHandshakeTimeout = 5 * time.Second
	defaultDialTimeout         = 5 * time.Second
	defaultDialKeepAlive       = 30 * time.Second
)

type Client struct {
	baseURL string

	HTTP *http.Client

	authCreds *authCredential

	// CREnabledHAPIFHIRBaseURL is the base url of a HAPI FHIR server with Clinical Reasoning Module enabled
	CREnabledHAPIFHIRBaseURL string

	// defaultHeaders are merged into every outbound request by setHeaders.
	// Set via WithDefaultHeaders.
	defaultHeaders map[string]string

	// omitCacheControl, when true, suppresses the library's default
	// Cache-Control: no-cache header. Set via WithoutCacheControlHeader.
	omitCacheControl bool
}

type authCredential struct {
	username string
	password string
}

// ClientOption allows customization of the client.
type ClientOption func(c *Client)

func WithTimeout(t time.Duration) func(c *Client) {
	return func(c *Client) {
		c.HTTP.Timeout = t
	}
}

func WithBasicAuth(username, password string) ClientOption {
	return func(c *Client) {
		if c.authCreds == nil {
			c.authCreds = &authCredential{}
		}

		c.authCreds.username = username
		c.authCreds.password = password
	}
}

// WithCREnabledHAPIFHIR is an option that allows the clients to define servers with clinical reasoning module enabled
func WithCREnabledHAPIFHIR(CREnabledHAPIFHIRBaseURL string) ClientOption {
	return func(c *Client) {
		c.CREnabledHAPIFHIRBaseURL = CREnabledHAPIFHIRBaseURL
	}
}

// WithHTTPClient replaces the underlying *http.Client. Use this when you need
// full control — e.g. swapping in a client with its own transport, timeout
// and observability wrappers. Other options that mutate the client (timeout,
// retry) still apply on top.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		if httpClient != nil {
			c.HTTP = httpClient
		}
	}
}

// WithTransport replaces the http.RoundTripper used by the client. Most callers
// who only need to tune the connection pool should use this instead of
// WithHTTPClient. The library's Client.Timeout and other client-level settings
// remain unchanged.
func WithTransport(rt http.RoundTripper) ClientOption {
	return func(c *Client) {
		if rt == nil {
			return
		}
		if c.HTTP == nil {
			c.HTTP = &http.Client{Timeout: defaultTimeout}
		}
		c.HTTP.Transport = rt
	}
}

// WithDefaultHeaders adds headers that will be set on every outbound request.
// Existing headers (Content-Type, Accept, Authorization) are not overwritten.
// Call this multiple times to merge; later calls take precedence on key
// collisions within the merged set.
func WithDefaultHeaders(headers map[string]string) ClientOption {
	return func(c *Client) {
		if c.defaultHeaders == nil {
			c.defaultHeaders = make(map[string]string, len(headers))
		}
		for k, v := range headers {
			c.defaultHeaders[k] = v
		}
	}
}

// WithoutCacheControlHeader suppresses the default `Cache-Control: no-cache`
// header that the library otherwise sets on every request. Useful when the
// upstream FHIR server's caching is beneficial for static reads (Questionnaire,
// ValueSet, CapabilityStatement) and you want HTTP-level caching to apply.
func WithoutCacheControlHeader() ClientOption {
	return func(c *Client) {
		c.omitCacheControl = true
	}
}

// WithRetry installs a retry policy applied via a RoundTripper wrapper around
// the client's existing Transport. Retries are limited to idempotent methods
// (GET/HEAD/PUT/DELETE/OPTIONS by default) and only on retryable status codes
// (408, 429, 500, 502, 503, 504 by default) or transport errors.
//
// POST is intentionally not retried by default because FHIR create operations
// are not idempotent without `If-None-Exist`. Callers who need POST retry must
// add http.MethodPost to RetryableMethods explicitly.
//
// If the request body cannot be replayed (Body is set but GetBody is nil) the
// request is sent once with no retry attempts, regardless of method.
func WithRetry(policy RetryPolicy) ClientOption {
	return func(c *Client) {
		if c.HTTP == nil {
			c.HTTP = &http.Client{Timeout: defaultTimeout}
		}
		policy.applyDefaults()
		inner := c.HTTP.Transport
		if inner == nil {
			inner = http.DefaultTransport
		}
		c.HTTP.Transport = &retryRoundTripper{
			inner:  inner,
			policy: policy,
		}
	}
}

// NewClientFromEnvVars creates a new client where the needed fields are
// retrieved from the environment variables.
func NewClientFromEnvVars() (*Client, error) {
	return NewClient(os.Getenv("HAPI_FHIR_BASE_URL"))
}

// NewClient creates a new HAPI FHIR api client plus optional configurations.
//
// The default client is configured with a tuned http.Transport — see
// defaultTransport() — and a 1 minute Client.Timeout. Override either via
// WithTransport / WithHTTPClient / WithTimeout.
func NewClient(baseURL string, options ...ClientOption) (*Client, error) {
	if baseURL == "" {
		return nil, errors.New("baseURL is empty")
	}

	client := &Client{
		HTTP: &http.Client{
			Timeout:   defaultTimeout,
			Transport: defaultTransport(),
		},
		baseURL: baseURL,
	}

	for _, opt := range options {
		opt(client)
	}

	return client, nil
}

// defaultTransport builds the library's default *http.Transport. Sized for a
// medium-fleet deployment talking to one or two upstream FHIR servers. The Go
// standard http.DefaultTransport's MaxIdleConnsPerHost of 2 causes severe
// connection-pool starvation under any real concurrency, so we override here.
func defaultTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   defaultDialTimeout,
			KeepAlive: defaultDialKeepAlive,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          defaultMaxIdleConns,
		MaxIdleConnsPerHost:   defaultMaxIdleConnsPerHost,
		MaxConnsPerHost:       0, // unlimited concurrent in-flight; let app-level cap admit
		IdleConnTimeout:       defaultIdleConnTimeout,
		TLSHandshakeTimeout:   defaultTLSHandshakeTimeout,
		ExpectContinueTimeout: 1 * time.Second,
	}
}
