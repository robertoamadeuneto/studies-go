package compaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name        = "New Campaign"
	content     = "This is a test Campaign"
	rawContacts = []string{"test.campaign@email.com", "+551234567890"}
)

func Test_NewCampaign(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, rawContacts)

	assert.NotEmpty(campaign.Id)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(rawContacts))
	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_NameIsRequired(t *testing.T) {
	assert := assert.New(t)
	_, _error := NewCampaign("", content, rawContacts)

	assert.Equal("Name is required", _error.Error())
}

func Test_NewCampaign_ContentIsRequired(t *testing.T) {
	assert := assert.New(t)
	_, _error := NewCampaign(name, "", rawContacts)

	assert.Equal("Content is required", _error.Error())
}

func Test_NewCampaign_ContactsAreRequired(t *testing.T) {
	assert := assert.New(t)
	_, _error := NewCampaign(name, content, []string{})

	assert.Equal("Contacts are required", _error.Error())
}
