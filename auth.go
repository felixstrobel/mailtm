package mailtm

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type AuthenticationResponse struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func (c *Client) Authenticate(ctx context.Context, address string, password string) error {
	reqBody := struct {
		Address  string `json:"address"`
		Password string `json:"password"`
	}{
		Address:  address,
		Password: password,
	}

	reqBodyBuffer := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBuffer).Encode(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.baseUrl+"/token", reqBodyBuffer)
	if err != nil {
		return err
	}

	var result AuthenticationResponse
	err = c.request(req, &result)
	if err != nil {
		return err
	}

	return nil
}
