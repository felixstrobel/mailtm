package mailtm

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	DefaultBaseURL string = "https://api.mail.tm"
)

type Client struct {
	http    *http.Client
	baseUrl string
	token   string
}

type ClientOption func(*Client)

// WithBaseURL sets a custom base URL for the Client instance. Use this to override the default base url.
func WithBaseURL(url string) ClientOption {
	return func(client *Client) {
		client.baseUrl = url
	}
}

// WithHttpClient sets a custom HTTP client for the Client instance. Use this to override the default HTTP client configuration.
func WithHttpClient(httpClient *http.Client) ClientOption {
	return func(client *Client) {
		client.http = httpClient
	}
}

// New creates and initializes a new Client instance with optional configurations applied via ClientOption.
func New(opts ...ClientOption) *Client {
	httpClient := &http.Client{}

	client := &Client{
		http:    httpClient,
		baseUrl: DefaultBaseURL,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// authenticatedRequest sends an authenticated HTTP request with a bearer token and decodes the response into a result.
// Returns an error if authentication is missing, the request fails, or decoding the response fails.
func (c *Client) authenticatedRequest(req *http.Request, result interface{}) error {
	if len(c.token) == 0 {
		return errors.New("missing authentication")
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	return c.request(req, result)
}

// request performs an HTTP request, sets required headers, and decodes the response body into the provided result.
// Returns an error if the request fails or decoding fails.
func (c *Client) request(req *http.Request, result interface{}) error {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	//TODO: check status codes
	if result == nil {
		return nil
	}
	err = json.NewDecoder(res.Body).Decode(result)
	return err
}
