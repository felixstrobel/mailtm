package mailtm

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func (c *MailClient) addAuthToken(account *Account) error {
	var response Response

	reqBody, err := json.Marshal(map[string]string{
		"address":  account.Address,
		"password": account.Password,
	})

	req, err := http.NewRequest("POST", string(c.Service)+"/token", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := c.http.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return err
	}

	account.Token = response.Token

	err = res.Body.Close()
	if err != nil {
		return err
	}

	return nil
}
