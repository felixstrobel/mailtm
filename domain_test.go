package mailtm

import (
	"context"
	"regexp"
	"testing"
)

func TestClient_GetDomains(t *testing.T) {
	ctx := context.Background()
	tldPattern := regexp.MustCompile("([a-z0-9A-Z]\\.)*[a-z0-9-]+\\.([a-z0-9]{2,24})+(\\.co\\.([a-z0-9]{2,24})|\\.([a-z0-9]{2,24}))*")

	client := New()
	var domains, err = client.GetDomains(ctx)
	if err != nil {
		t.Fatal(err)
	}

	for _, domain := range domains {
		println(domain.Domain)
		if !tldPattern.MatchString(domain.Domain) {
			t.Error(domain.Domain)
		}
	}
}

func TestClient_GetDomain(t *testing.T) {
	ctx := context.Background()

	client := New()
	var domains, _ = client.GetDomains(ctx)

	domainOne := domains[0]
	domainTwo, err := client.GetDomain(ctx, domainOne.ID)
	if err != nil {
		t.Error(err)
	}
	if domainOne.ID != domainTwo.ID {
		t.Error("domains not matching")
	}
}
