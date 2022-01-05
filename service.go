package mailtm

import (
	"errors"
	"regexp"
)

type Service struct {
	Url string
}

func NewService(url string) (*Service, error) {
	urlPattern := regexp.MustCompile("([a-z0-9A-Z]\\.)*[a-z0-9-]+\\.([a-z0-9]{2,24})+(\\.co\\.([a-z0-9]{2,24})|\\.([a-z0-9]{2,24}))*")

	if !urlPattern.MatchString(url) {
		return nil, errors.New("given string does not match a pattern for a valid url")
	}

	return &Service{Url: url}, nil
}

func (c *MailClient) SetService(s Service) {
	c.Service = s
}
