package mailtm

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Addressee struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Message struct {
	ID             string      `json:"id"`
	AccountID      string      `json:"accountId"`
	MessageID      string      `json:"msgid"`
	From           Addressee   `json:"from"`
	To             []Addressee `json:"to"`
	Subject        string      `json:"subject"`
	Intro          string      `json:"intro"`
	Seen           bool        `json:"seen"`
	IsDeleted      bool        `json:"isDeleted"`
	HasAttachments bool        `json:"hasAttachments"`
	Size           int         `json:"size"`
	DownloadUrl    string      `json:"downloadUrl"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
}

type DetailedMessage struct {
	ID             string      `json:"id"`
	AccountID      string      `json:"accountId"`
	MessageID      string      `json:"msgid"`
	From           Addressee   `json:"from"`
	To             []Addressee `json:"to"`
	CC             []Addressee `json:"cc"`
	BCC            []Addressee `json:"bcc"`
	Subject        string      `json:"subject"`
	Seen           bool        `json:"seen"`
	Flagged        bool        `json:"flagged"`
	IsDeleted      bool        `json:"isDeleted"`
	Verifications  []string    `json:"verifications"`
	Retention      bool        `json:"retention"`
	RetentionDate  time.Time   `json:"retentionDate"`
	Test           string      `json:"test"`
	Html           []string    `json:"html"`
	HasAttachments bool        `json:"hasAttachments"`
	Attachments    []string    `json:"attachments"`
	Size           int         `json:"size"`
	DownloadUrl    string      `json:"downloadUrl"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
}

func (c *MailClient) GetMessages(page int) ([]Message, error) {
	var messagesResponse []Message

	if c.Token == "" {
		return nil, errors.New("please fetch the auth-token first using GetAuthToken()")
	}

	req, err := http.NewRequest("GET", c.Service.Url+"/messages?page="+strconv.Itoa(page), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &messagesResponse)
	if err != nil {
		return nil, err
	}

	return messagesResponse, nil
}

func (c *MailClient) GetMessageByID(id string) (*DetailedMessage, error) {
	var message DetailedMessage

	if c.Token == "" {
		return nil, errors.New("please fetch the auth-token first using GetAuthToken()")
	}

	req, err := http.NewRequest("GET", c.Service.Url+"/messages/"+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (c *MailClient) DeleteMessageByID(id string) error {
	if c.Token == "" {
		return errors.New("please fetch the auth-token first using GetAuthToken()")
	}

	req, err := http.NewRequest("DELETE", c.Service.Url+"/messages/"+id, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	_, err = c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *MailClient) SeenMessageByID(id string) error {
	if c.Token == "" {
		return errors.New("please fetch the auth-token first using GetAuthToken()")
	}

	reqBody, err := json.Marshal(map[string]bool{
		"seen":  true,
	})

	req, err := http.NewRequest("PATCH", c.Service.Url+"/messages/"+id, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	_, err = c.HttpClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}