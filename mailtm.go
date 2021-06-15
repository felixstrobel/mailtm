package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}
type Domain struct {
	Path      string    `json:"@id"`
	Type      string    `json:"@type"`
	Id        string    `json:"id"`
	Name      string    `json:"domain"`
	IsActive  bool      `json:"isActive"`
	IsPrivate bool      `json:"isPrivate"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
type DomainResponse struct {
	Context     string   `json:"@context"`
	Path        string   `json:"@id"`
	Type        string   `json:"@type"`
	Domains     []Domain `json:"hydra:member"`
	DomainCount int      `json:"hydra:totalItems"`
}
type TokenResponse struct {
	Token string `json:"token"`
}
type Message struct {
	Path           string    `json:"@id"`
	Type           string    `json:"@type"`
	Id             string    `json:"id"`
	AccountId      string    `json:"accountId"`
	MessageId      string    `json:"msgid"`
	From           User      `json:"from"`
	To             []User    `json:"to"`
	Subject        string    `json:"subject"`
	Intro          string    `json:"intro"`
	Seen           bool      `json:"seen"`
	IsDeleted      bool      `json:"isDeleted"`
	HasAttachments bool      `json:"hasAttachments"`
	Size           int       `json:"size"`
	DownloadUrl    string    `json:"downloadUrl"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
type MessageResponse struct {
	Context      string    `json:"@context"`
	Path         string    `json:"@id"`
	Type         string    `json:"@type"`
	Messages     []Message `json:"hydra:member"`
	MessageCount int       `json:"hydra:totalItems"`
}
type AccountInfo struct {
	Context    string    `json:"@context"`
	Path       string    `json:"@id"`
	Type       string    `json:"@type"`
	Id         string    `json:"id"`
	Email      string    `json:"address"`
	Quota      int       `json:"quota"`
	UsedCount  int       `json:"used"`
	IsDisabled bool      `json:"isDisabled"`
	IsDeleted  bool      `json:"isDeleted"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
type MailClient struct {
	URL         string
	Email       string
	Password    string
	BearerToken string
	Information AccountInfo
}

func NewMailClient() *MailClient {
	return &MailClient{
		URL:      "https://api.mail.tm",
		Email:    "",
		Password: "",
		Information: AccountInfo{
			Context:    "",
			Path:       "",
			Type:       "",
			Id:         "",
			Email:      "",
			Quota:      0,
			UsedCount:  0,
			IsDisabled: false,
			IsDeleted:  false,
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		},
	}
}

func (c *MailClient) GetAvailableDomains() []Domain {
	res, err := http.Get(c.URL + "/domains")
	if err != nil {
		log.Fatal("Server is not reachable:", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	var domainResponse DomainResponse

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Server response body is not readable:", err)
	}

	err = json.Unmarshal(body, &domainResponse)
	if err != nil {
		log.Fatal("Server response is not parsable:", err)
	}

	return domainResponse.Domains
}
func (c *MailClient) Register(username string, domain string, password string) {
	c.Email = username + "@" + domain
	c.Password = password

	res, err := http.Post(
		c.URL+"/accounts",
		"application/json",
		bytes.NewBuffer([]byte("{\"address\":\""+username+"@"+domain+"\",\"password\":\""+password+"\"}")),
	)
	if err != nil {
		log.Fatal("Registering failed:", err)
	}

	var accountInfo AccountInfo

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Server response body is not readable:", err)
	}

	err = json.Unmarshal(body, &accountInfo)
	if err != nil {
		log.Fatal("Server response is not parsable:", err)
	}

	c.Information = accountInfo
}
func (c *MailClient) Login() {
	res, err := http.Post(
		c.URL+"/token",
		"application/json",
		bytes.NewBuffer([]byte("{\"address\":\""+c.Email+"\",\"password\":\""+c.Password+"\"}")),
	)
	if err != nil {
		log.Fatal("Server is not reachable:", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	var tokenResponse TokenResponse

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Server response body is not readable:", err)
	}

	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Fatal("Server response is not parsable:", err)
	}

	c.BearerToken = tokenResponse.Token
}
func (c *MailClient) Delete() {
	var client *http.Client = &http.Client{}

	req, err := http.NewRequest("DELETE", c.URL+"/accounts/"+c.Information.Id, nil)
	if err != nil {
		log.Fatal("Server is not reachable:", err)
	}

	req.Header.Add("Authorization", "Bearer "+c.BearerToken)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Reqeuest failed: ", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
}
func (c *MailClient) GetMessages(page int) []Message {
	var client *http.Client = &http.Client{}

	req, err := http.NewRequest("GET", c.URL+"/messages?page="+strconv.Itoa(page), nil)
	if err != nil {
		log.Fatal("Server is not reachable:", err)
	}

	req.Header.Add("Authorization", "Bearer "+c.BearerToken)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Reqeuest failed: ", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	var messageResponse MessageResponse

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Server response body is not readable:", err)
	}
	fmt.Printf("%s\n", body)

	err = json.Unmarshal(body, &messageResponse)
	if err != nil {
		log.Fatal("Server response is not parsable:", err)
	}

	return messageResponse.Messages
}
func (c *MailClient) GetMessage(id string) Message {
	var client *http.Client = &http.Client{}

	req, err := http.NewRequest("GET", c.URL+"/messages/"+id, nil)
	if err != nil {
		log.Fatal("Server is not reachable:", err)
	}

	req.Header.Add("Authorization", "Bearer "+c.BearerToken)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Reqeuest failed: ", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	var message Message

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Server response body is not readable:", err)
	}

	err = json.Unmarshal(body, &message)
	if err != nil {
		log.Fatal("Server response is not parsable:", err)
	}

	return message
}
func (c *MailClient) DeleteMessage(id string) {
	var client *http.Client = &http.Client{}

	req, err := http.NewRequest("DELETE", c.URL+"/messages/"+id, nil)
	if err != nil {
		log.Fatal("Server is not reachable:", err)
	}

	req.Header.Add("Authorization", "Bearer "+c.BearerToken)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Reqeuest failed: ", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
}
func (c *MailClient) MarkMessageAsSeen(id string) {
	var client *http.Client = &http.Client{}

	req, err := http.NewRequest("PATCH", c.URL+"/messages/"+id, bytes.NewBufferString("true"))
	if err != nil {
		log.Fatal("Server is not reachable:", err)
	}

	req.Header.Add("Content-Type", "merge-patch+json")
	req.Header.Add("Authorization", "Bearer "+c.BearerToken)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Reqeuest failed: ", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
}
func (c *MailClient) GetMessageSource(id string) string {
	var client *http.Client = &http.Client{}

	req, err := http.NewRequest("GET", c.URL+"/sources/"+id, nil)
	if err != nil {
		log.Fatal("Server is not reachable:", err)
	}

	req.Header.Add("Authorization", "Bearer "+c.BearerToken)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Reqeuest failed: ", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Server response body is not readable:", err)
	}

	return string(body)
}

func main() {}
