package mailtm

import (
	"net/http"
)

type Service string

const (
	MAIL_TM Service = "https://api.mail.tm"
)

type MailClient struct {
	http    *http.Client
	Service Service
}

func New() (*MailClient, error) {
	return &MailClient{Service: MAIL_TM, http: &http.Client{}}, nil
}
