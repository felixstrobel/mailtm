package mailtm

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type Account struct {
	ID         string    `json:"id"`
	Address    string    `json:"address"`
	Quota      int       `json:"quota"`
	Used       int       `json:"used"`
	IsDisabled bool      `json:"isDisabled"`
	IsDeleted  bool      `json:"isDeleted"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`

	Password string
	Token    string
}

func (c *MailClient) NewAccount() (*Account, error) {
	var password, err = RandomString(16)
	if err != nil {
		return nil, err
	}

	return c.NewAccountWithPassword(password)
}

func (c *MailClient) NewCustomAccount(username string, password string) (*Account, error) {
	var account Account

	domains, err := c.GetDomains()
	if err != nil {
		return nil, err
	}
	if len(domains) == 0 {
		return nil, errors.New("account hasn't been created due to receiving no domains from the server")
	}

	address := username + "@" + domains[0].TLD
	reqBody, err := json.Marshal(map[string]string{
		"address":  address,
		"password": password,
	})

	req, err := http.NewRequest("POST", string(c.service)+"/accounts", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &account)
	if err != nil {
		return nil, err
	}

	account.Address = address
	account.Password = password

	err = res.Body.Close()
	if err != nil {
		return nil, err
	}

	return &account, c.addAuthToken(&account)
}

func (c *MailClient) NewAccountWithPassword(password string) (*Account, error) {
	handle, err := RandomString(20)
	if err != nil {
		return nil, err
	}

	return c.NewCustomAccount(handle, password)
}

func (c *MailClient) RetrieveAccount(address string, password string) (*Account, error) {
	account := &Account{Address: address, Password: password}
	err := c.addAuthToken(account)
	if err != nil {
		return nil, err
	}

	err = c.UpdateAccountInformation(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (c *MailClient) UpdateAccountInformation(account *Account) error {
	var response Account

	req, err := http.NewRequest("GET", string(c.service)+"/me", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+account.Token)
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

	account.Quota = response.Quota
	account.Used = response.Used
	account.UpdatedAt = response.UpdatedAt

	err = res.Body.Close()
	if err != nil {
		return err
	}

	return nil
}

func (c *MailClient) DeleteAccount(account *Account) error {
	if account.Token == "" {
		return errors.New("the account hasn't been deleted because auth token hasn't been found")
	}

	req, err := http.NewRequest("DELETE", string(c.service)+"/accounts/"+account.ID+"?id="+account.ID, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+account.Token)
	res, err := c.http.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 204 {
		return errors.New("wasn't able to delete account")
	}

	return nil
}
