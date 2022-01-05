package mailtm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TokenResponse struct {
	Token string `json:"token"`
}

func (c *MailClient) GetAuthToken() (string, error) {
	var tokenResponse TokenResponse

	reqBody, err := json.Marshal(map[string]string{
		"address":  c.Account.Address,
		"password": c.Account.Password,
	})

	req, err := http.NewRequest("POST", c.Service.Url+"/token", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	fmt.Printf("[DEBUG] Requested token [%d]:\t\t%s:%s\n", res.StatusCode, c.Account.Address, c.Account.Password)

	c.Token = tokenResponse.Token

	return tokenResponse.Token, nil
}

func (c *MailClient) GetAuthTokenCredentials(address string, password string) (string, error) {
	var tokenResponse TokenResponse

	reqBody, err := json.Marshal(map[string]string{
		"address":  address,
		"password": password,
	})

	req, err := http.NewRequest("POST", c.Service.Url+"/token", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		return "", err
	}

	fmt.Printf("[DEBUG] Requested token [%d]:\t\t%s:%s\n", res.StatusCode, address, password)

	c.Token = tokenResponse.Token
	c.Account.Address = address
	c.Account.Password = password

	return tokenResponse.Token, nil
}
