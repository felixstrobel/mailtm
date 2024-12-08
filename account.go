package mailtm

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type Account struct {
	ID        string    `json:"id"`
	Address   string    `json:"address"`
	Quota     int       `json:"quota"`
	Used      int       `json:"used"`
	Disabled  bool      `json:"isDisabled"`
	Deleted   bool      `json:"isDeleted"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c *Client) CreateAccount(ctx context.Context, address string, password string) (*Account, error) {
	reqBody := struct {
		address  string `json:"address"`
		password string `json:"password"`
	}{
		address:  address,
		password: password,
	}

	reqBodyBuffer := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBuffer).Encode(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseUrl+"/accounts", reqBodyBuffer)
	if err != nil {
		return nil, err
	}

	var result Account
	err = c.request(req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetAccount(ctx context.Context, id string) (*Account, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseUrl+"/accounts/"+id, nil)
	if err != nil {
		return nil, err
	}

	var result Account
	err = c.authenticatedRequest(req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteAccount(ctx context.Context, id string) error {
	req, err := http.NewRequestWithContext(ctx, "DELETE", c.baseUrl+"/accounts/"+id, nil)
	if err != nil {
		return err
	}

	// Status codes: 204 -> Account deleted
	err = c.authenticatedRequest(req, struct{}{}) //TODO: test if the use of an empty struct is allowed
	return err
}
