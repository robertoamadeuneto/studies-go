package entity

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name        = "New Campaign"
	content     = "This is a test Campaign"
	rawContacts = []string{"test.campaign@email.com", "+551234567890"}

	fakerInstance = faker.New()
)

func Test_NewCampaign_WithValidData_ReturnsCampaign(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, rawContacts)

	assert.NotEmpty(campaign.Id)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(rawContacts))
	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_WithEmptyName_ReturnsError(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign("", content, rawContacts)

	assert.Equal("Name should have a min size of 5", err.Error())
}

func Test_NewCampaign_WithNameThatDoesNotHaveTheMinSize_ReturnsError(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign("Alde", content, rawContacts)

	assert.Equal("Name should have a min size of 5", err.Error())
}

func Test_NewCampaign_WithNameThatExceedsTheMaxSize_ReturnsError(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(fakerInstance.Lorem().Text(30), content, rawContacts)

	assert.Equal("Name should have a max size of 24", err.Error())
}

func Test_NewCampaign_WithEmptyContent_ReturnsError(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, "", rawContacts)

	assert.Equal("Content should have a min size of 5", err.Error())
}

func Test_NewCampaign_WithContentThatDoesNotHaveTheMinSize_ReturnsError(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, "This", rawContacts)

	assert.Equal("Content should have a min size of 5", err.Error())
}

func Test_NewCampaign_WithContentThatExceedsTheMaxSize_ReturnsError(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, fakerInstance.Lorem().Text(1040), rawContacts)

	assert.Equal("Content should have a max size of 1024", err.Error())
}

func Test_NewCampaign_WithEmptyContacts_ReturnsError(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, content, []string{})

	assert.Equal("Contacts should have a min size of 1", err.Error())
}

func Test_NewCampaign_WithInvalidEmailAsContact_ReturnsError(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, content, []string{"invalid_email.com", "+551234567890"})

	assert.Equal("Value is an invalid e-mail or an invalid phone number", err.Error())
}

func Test_NewCampaign_WithInvalidPhoneNumberAsContact_ReturnsError(t *testing.T) {
	assert := assert.New(t)
	_, err := NewCampaign(name, content, []string{"test.campaign@email.com", "1234++"})

	assert.Equal("Value is an invalid e-mail or an invalid phone number", err.Error())
}
