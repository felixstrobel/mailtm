package mailtm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Account struct {
	ID         string `json:"id"`
	Address    string `json:"address"`
	Password   string
	Token      string
	Quota      int       `json:"quota"`
	Used       int       `json:"used"`
	IsDisabled bool      `json:"isDisabled"`
	IsDeleted  bool      `json:"isDeleted"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (c *MailClient) CreateAccount() (*Account, error) {
	var account Account
	var password, err = RandomString(16)
	if err != nil {
		return nil, err
	}
	mailName, err := RandomString(20)
	if err != nil {
		return nil, err
	}

	address := mailName + "@" + c.Domain.Path

	reqBody, err := json.Marshal(map[string]string{
		"address":  address,
		"password": password,
	})

	req, err := http.NewRequest("POST", c.Service.Url+"/accounts", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &account)
	if err != nil {
		return nil, err
	}

	account.Address = address
	account.Password = password
	c.Account = account

	fmt.Printf("[DEBUG] Created account [%d]:\t\t%s:%s\n", res.StatusCode, address, password)

	return &account, nil
}

func (c *MailClient) GetAccountByID(id string) (*Account, error) {
	var account Account

	if c.Token == "" {
		return nil, errors.New("please fetch the auth-token first using GetAuthToken()")
	}

	req, err := http.NewRequest("GET", c.Service.Url+"/accounts/"+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+ c.Token)
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

func (c *MailClient) DeleteAccountByID(id string) error {
	if c.Token == "" {
		return errors.New("please fetch the auth-token first using GetAuthToken()")
	}

	req, err := http.NewRequest("DELETE", c.Service.Url+"/accounts/"+id, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+ c.Token)
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	fmt.Printf("[DEBUG] Delete account [%d]:\t\t%s\n", res.StatusCode, id)

	return nil
}

func (c *MailClient) GetCurrentAccountInformation() (*Account, error) {
	var account Account

	req, err := http.NewRequest("GET", c.Service.Url+"/me", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Set("Accept", "application/json")
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &account)
	if err != nil {
		return nil, err
	}

	fmt.Printf("[DEBUG] Get account [%d]:\t\t%s\n", res.StatusCode, account.Address)

	return &account, nil
}
