package campaign

import (
	"emailn/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (repository *repositoryMock) Save(campaign *Campaign) error {
	args := repository.Called(campaign)

	return args.Error(0)
}

var (
	repository = new(repositoryMock)
	service    = Service{repository}
)

func TestCreate(t *testing.T) {
	assert := assert.New(t)
	dto := contract.NewCampaignDto{
		Name:     "New Campaign",
		Content:  "This is a test Campaign",
		Contacts: []string{"test.campaign@email.com", "+551234567890"},
	}
	repository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != dto.Name ||
			campaign.Content != dto.Content ||
			len(campaign.Contacts) != len(dto.Contacts) {
			return false
		}

		return true
	})).Return(nil)

	id, err := service.Create(dto)

	assert.NotEmpty(id)
	assert.Nil(err)
	repository.AssertExpectations(t)
}
