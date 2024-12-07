package mailtm

import (
	"context"
	"net/http"
	"time"
)

type Domain struct {
	ID        string    `json:"id"`
	Domain    string    `json:"domain"` // Refers to the top-level domain
	Active    bool      `json:"isActive"`
	Private   bool      `json:"isPrivate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c *Client) GetDomains(ctx context.Context) ([]*Domain, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseUrl+"/domains", nil)
	if err != nil {
		return nil, err
	}

	var result []*Domain
	err = c.request(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetDomain(ctx context.Context, id string) (*Domain, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseUrl+"/domains/"+id, nil)
	if err != nil {
		return nil, err
	}

	var result *Domain
	err = c.request(req, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
