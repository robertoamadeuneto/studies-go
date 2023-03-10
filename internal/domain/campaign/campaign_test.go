package compaign

import (
	"strconv"
	"testing"
)

func TestNewCampaign(t *testing.T) {
	name := "New Campaign"
	content := "This is a test Campaign"
	rawContacts := []string{"test.campaign@email.com", "+551234567890"}

	campaign := NewCampaign(name, content, rawContacts)

	if campaign.Id != "1" {
		t.Errorf("Expected Id: 1")
	}
	if campaign.Name != name {
		t.Errorf("Expected Name: %s", name)
	}
	if campaign.Content != content {
		t.Errorf("Expected Content: %s", content)
	}
	if len(campaign.Contacts) != len(rawContacts) {
		t.Errorf("Expected Contacts length: %s", strconv.Itoa(len(rawContacts)))
	}
}
