package mailtm

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type Recipient struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Sender struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Attachment struct {
	ID          string `json:"id"`
	Filename    string `json:"filename"`
	ContentType string `json:"contentType"`
	Disposition string `json:"disposition"`
	Encoding    string `json:"transferEncoding"`
	Related     bool   `json:"related"`
	Size        int    `json:"size"`
	DownloadUrl string `json:"downloadUrl"`
}

type SimpleMessage struct {
	ID             string      `json:"id"`
	MessageID      string      `json:"msgid"`
	Sender         Sender      `json:"from"`
	Recipients     []Recipient `json:"to"`
	Subject        string      `json:"subject"`
	Intro          string      `json:"intro"`
	Seen           bool        `json:"seen"`
	Deleted        bool        `json:"isDeleted"`
	HasAttachments bool        `json:"hasAttachments"`
	Size           int         `json:"size"`
	DownloadUrl    string      `json:"downloadUrl"`
	SourceUrl      string      `json:"sourceUrl"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
	AccountID      string      `json:"accountId"`
}

type FullMessage struct {
	ID             string       `json:"id"`
	MessageID      string       `json:"msgid"`
	Sender         Sender       `json:"from"`
	Recipients     []Recipient  `json:"to"`
	CC             []Recipient  `json:"cc"`
	BCC            []Recipient  `json:"bcc"`
	Subject        string       `json:"subject"`
	Intro          string       `json:"intro"`
	Seen           bool         `json:"seen"`
	Flagged        bool         `json:"flagged"`
	Deleted        bool         `json:"isDeleted"`
	Verifications  []string     `json:"verifications"`
	Retention      bool         `json:"retention"`
	RetentionDate  time.Time    `json:"retentionDate"`
	Text           string       `json:"text"`
	Html           []string     `json:"html"`
	HasAttachments bool         `json:"hasAttachments"`
	Attachments    []Attachment `json:"attachments"`
	Size           int          `json:"size"`
	DownloadUrl    string       `json:"downloadUrl"`
	SourceUrl      string       `json:"sourceUrl"`
	CreatedAt      time.Time    `json:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt"`
	AccountID      string       `json:"accountId"`
}

type MessageConfig struct {
	Page int
}

type MessageOption func(config *MessageConfig)

func WithPage(page int) MessageOption {
	return func(config *MessageConfig) {
		config.Page = page
	}
}

func (c *Client) GetMessages(ctx context.Context, opts ...MessageOption) ([]*SimpleMessage, error) {
	config := &MessageConfig{Page: 1}
	for _, opt := range opts {
		opt(config)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", c.baseUrl+"/messages?page="+strconv.Itoa(config.Page), nil)
	if err != nil {
		return nil, err
	}

	var result []*SimpleMessage
	err = c.authenticatedRequest(req, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) GetMessage(ctx context.Context, id string) (*FullMessage, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.baseUrl+"/messages/"+id, nil)
	if err != nil {
		return nil, err
	}

	var result FullMessage
	err = c.authenticatedRequest(req, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteMessage(ctx context.Context, id string) error {
	req, err := http.NewRequestWithContext(ctx, "DELETE", c.baseUrl+"/messages/"+id, nil)
	if err != nil {
		return err
	}

	err = c.authenticatedRequest(req, struct{}{}) //TODO: check if this works with an empty struct
	return err
}

func (c *Client) SetMessageToSeen(ctx context.Context, id string) error {
	reqBody := struct {
		Seen bool `json:"seen"`
	}{
		Seen: true,
	}

	reqBodyBuffer := new(bytes.Buffer)
	err := json.NewEncoder(reqBodyBuffer).Encode(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "PATCH", c.baseUrl+"/messages/"+id, reqBodyBuffer)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/merge-patch+json")

	err = c.authenticatedRequest(req, struct{}{}) //TODO: check if this works with an empty struct
	return err
}

func (c *Client) GetAttachment(ctx context.Context, messageId string, attachmentId string) error {
	return nil //TODO: implement attachments, source
}
