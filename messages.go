package mailtm

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
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
	Text           string      `json:"text"`
	Html           []string    `json:"html"`
	HasAttachments bool        `json:"hasAttachments"`
	Attachments    []string    `json:"attachments"`
	Size           int         `json:"size"`
	DownloadUrl    string      `json:"downloadUrl"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
}

func (c *MailClient) GetMessages(account *Account, page int) ([]Message, error) {
	var response []Message

	if account.Token == "" {
		return nil, errors.New("the messages haven't been fetched because auth token hasn't been found")
	}

	req, err := http.NewRequest("GET", string(c.service)+"/messages?page="+strconv.Itoa(page), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+account.Token)
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *MailClient) GetMessageByID(account *Account, id string) (*DetailedMessage, error) {
	var response DetailedMessage

	if account.Token == "" {
		return nil, errors.New("the message hasn't been fetched because auth token hasn't been found")
	}

	req, err := http.NewRequest("GET", string(c.service)+"/messages/"+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+account.Token)
	res, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *MailClient) DeleteMessageByID(account *Account, id string) error {
	if account.Token == "" {
		return errors.New("the message hasn't been deleted because auth token hasn't been found")
	}

	req, err := http.NewRequest("DELETE", string(c.service)+"/messages/"+id, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+account.Token)
	_, err = c.http.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *MailClient) SeenMessageByID(account *Account, id string) error {
	if account.Token == "" {
		return errors.New("the message hasn't been set to seen because auth token hasn't been found")
	}

	reqBody, err := json.Marshal(map[string]bool{
		"seen": true,
	})

	req, err := http.NewRequest("PATCH", string(c.service)+"/messages/"+id, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+account.Token)
	_, err = c.http.Do(req)
	if err != nil {
		return err
	}

	return nil
}
