package mailtm

import (
	"net"
	"net/http"
	"time"
)

type Service string

const (
	MAIL_TM Service = "https://api.mail.tm"
)

type MailClient struct {
	http    *http.Client
	service Service
}

func New() (*MailClient, error) {
	t := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		// We use ABSURDLY large keys, and should probably not.
		TLSHandshakeTimeout: 60 * time.Second,
	}
	c := &http.Client{
		Transport: t,
	}

	return &MailClient{service: MAIL_TM, http: c}, nil
}
