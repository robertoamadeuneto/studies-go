package compaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign(t *testing.T) {
	assert := assert.New(t)
	name := "New Campaign"
	content := "This is a test Campaign"
	rawContacts := []string{"test.campaign@email.com", "+551234567890"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, rawContacts)

	assert.NotEmpty(campaign.Id)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(rawContacts))
	assert.Greater(campaign.CreatedOn, now)
}
