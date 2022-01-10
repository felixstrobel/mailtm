package mailtm

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type Domain struct {
	ID        string    `json:"id"`
	Path      string    `json:"domain"`
	IsActive  bool      `json:"isActive"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c *MailClient) GetDomains() ([]Domain, error) {
	var domainResponse []Domain

	req, err := http.NewRequest("GET", c.Service.Url+"/domains", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &domainResponse)
	if err != nil {
		return nil, err
	}

	if len(domainResponse) != 0 {
		c.Domain = domainResponse[0]
	} else {
		return nil, errors.New("no domains found on the server")
	}

	return domainResponse, nil
}

func (c *MailClient) GetDomainByID(id string) (*Domain, error) {
	var domainResponse Domain

	req, err := http.NewRequest("GET", c.Service.Url+"/domains/"+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &domainResponse)
	if err != nil {
		return nil, err
	}

	return &domainResponse, nil
}
