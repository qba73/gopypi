package pypi

import (
	"net/http"
	"time"
)

type Client struct {
	BaseURL    string
	httpClient *http.Client

	Package *PackageService
}

// NewClient creates a new PyPI client. It also allow
// to override default http client.
func NewClient(opts ...func(*Client) error) (*Client, error) {
	client := Client{
		BaseURL: "https://pypi.org",
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}

	for _, opt := range opts {
		if err := opt(&client); err != nil {
			return nil, err
		}
	}

	client.Package = &PackageService{client: &client}
	return &client, nil
}

// WithHTTPClient allows user of the PyPI API client
// to override the default HTTP client.
func WithHTTPClient(client *http.Client) func(*Client) error {
	return func(c *Client) error {
		c.httpClient = client
		return nil
	}
}
