package compaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "New Campaign"
	content := "This is a test Campaign"
	rawContacts := []string{"test.campaign@email.com", "+551234567890"}

	campaign := NewCampaign(name, content, rawContacts)

	assert.Equal(campaign.Id, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(rawContacts))
}
