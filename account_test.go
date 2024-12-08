package mailtm

import (
	"context"
	"testing"
)

func TestClient_CreateAccount(t *testing.T) {
	ctx := context.Background()
	client := New()

	domains, _ := client.GetDomains(ctx)
	username := "test@" + domains[0].Domain
	password := "test"

	acc, err := client.CreateAccount(ctx, username, password)
	if err != nil {
		t.Error(err)
	}
	println(acc.ID)
}
