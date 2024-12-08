package mailtm

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Service string

const (
	DefaultBaseURL string = "https://api.mail.tm"
)

type Client struct {
	http    *http.Client
	baseUrl string
	token   string
}

type ClientOption func(*Client)

func WithBaseURL(url string) ClientOption {
	return func(client *Client) {
		client.baseUrl = url
	}
}

func WithHttpClient(httpClient *http.Client) ClientOption {
	return func(client *Client) {
		client.http = httpClient
	}
}

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

func (c *Client) authenticatedRequest(req *http.Request, result interface{}) error {
	if len(c.token) == 0 {
		return errors.New("missing authentication")
	}

	req.Header.Set("Authorization", "Bearer "+c.token)

	return c.request(req, result)
}

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
