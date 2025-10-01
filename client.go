package hapifhirgo

import (
	"errors"
	"net/http"
	"os"
	"time"
)

const (
	defaultTimeout = 1 * time.Minute
)

type Client struct {
	baseURL string

	HTTP *http.Client

	authCreds *authCredential

	SDCEnabledServerURL string
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

func WithSDCEnabledServerURL(SDCEnabledServerURL string) ClientOption {
	return func(c *Client) {
		c.SDCEnabledServerURL = SDCEnabledServerURL
	}
}

// NewClientFromEnvVars creates a new client where the needed fields are
// retrieved from the environment variables.
func NewClientFromEnvVars() (*Client, error) {
	return NewClient(os.Getenv("HAPI_FHIR_BASE_URL"))
}

// NewClient creates a new HAPI FHIR api client plus optional configurations
func NewClient(baseURL string, options ...ClientOption) (*Client, error) {
	if baseURL == "" {
		return nil, errors.New("baseURL is empty")
	}

	client := &Client{
		HTTP: &http.Client{
			Timeout: defaultTimeout,
		},
		baseURL: baseURL,
	}

	for _, opt := range options {
		opt(client)
	}

	return client, nil
}
