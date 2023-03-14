package campaign

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

func Test_Should_Build_NewCampaign(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, rawContacts)

	assert.NotEmpty(campaign.Id)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(rawContacts))
	assert.Greater(campaign.CreatedOn, now)
}

func Test_Should_Not_Build_NewCampaign_When_Name_Is_Empty(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign("", content, rawContacts)

	assert.Equal("Name is required", err.Error())
}

func Test_Should_Not_Build_NewCampaign_When_Content_Is_Empty(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, "", rawContacts)

	assert.Equal("Content is required", err.Error())
}

func Test_Should_Not_Build_NewCampaign_When_Contacts_Are_Empty(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, content, []string{})

	assert.Equal("Contacts are required", err.Error())
}
