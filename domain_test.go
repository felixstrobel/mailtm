package mailtm

import (
	"regexp"
	"testing"
)

func TestMailClient_GetDomains(t *testing.T) {
	tldPattern := regexp.MustCompile("([a-z0-9A-Z]\\.)*[a-z0-9-]+\\.([a-z0-9]{2,24})+(\\.co\\.([a-z0-9]{2,24})|\\.([a-z0-9]{2,24}))*")

	client, _ := New()
	var domains, err = client.GetDomains()
	if err != nil {
		t.Fatal(err)
	}

	for _, domain := range domains {
		if !tldPattern.MatchString(domain.TLD) {
			t.Error("")
		}
	}
}

func TestMailClient_GetDomainByID(t *testing.T) {
	client, _ := New()
	var domains, _ = client.GetDomains()

	domainOne := domains[0]
	domainTwo, _ := client.GetDomainByID(domainOne.ID)
	if domainOne != *domainTwo {
		t.Error("domains not matching")
	}
}
