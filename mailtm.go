package mailtm

import (
	"crypto/rand"
	"github.com/google/uuid"
	"math/big"
	"net/http"
)

type MailClient struct {
	ID         uuid.UUID
	Service    Service
	HttpClient *http.Client
	Domain     Domain
	Account    Account
	Token      string
}

func NewMailClient() (*MailClient, error) {
	service, err := NewService("https://api.mail.tm")
	if err != nil {
		return nil, err
	}

	return &MailClient{ID: uuid.New(), Service: *service, HttpClient: &http.Client{}}, nil
}

func NewMailClientService(s *Service) (*MailClient, error) {
	return &MailClient{ID: uuid.New(), Service: *s}, nil
}

func RandomString(n int) (string, error) {
	const letters = "0123456789abcdefghijklmnopqrstuvwxyz"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
