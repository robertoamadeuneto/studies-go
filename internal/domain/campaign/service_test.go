package campaign

import (
	"emailn/internal/contract"
	"errors"
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
	repository     = new(repositoryMock)
	service        = Service{repository}
	newCampaignDto = contract.NewCampaignDto{
		Name:     "New Campaign",
		Content:  "This is a test Campaign",
		Contacts: []string{"test.campaign@email.com", "+551234567890"},
	}
)

func Test_Should_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaignDto.Name ||
			campaign.Content != newCampaignDto.Content ||
			len(campaign.Contacts) != len(newCampaignDto.Contacts) {
			return false
		}

		return true
	})).Return(nil)

	id, err := service.Create(newCampaignDto)

	assert.NotEmpty(id)
	assert.Nil(err)
	repository.AssertNumberOfCalls(t, "Save", 1)
	repository.AssertExpectations(t)
}

func Test_Should_Not_Create_Campaign_When_NewCampaign_Returns_Error(t *testing.T) {
	assert := assert.New(t)
	newCampaignDto.Name = ""

	_, err := service.Create(newCampaignDto)

	assert.NotNil(err)
	repository.AssertNotCalled(t, "Save")
}
func Test_Should_Not_Create_Campaign_When_Repository_Save_Returns_Error(t *testing.T) {
	assert := assert.New(t)
	mockedErrorMessage := "Error trying to communicate with database!"
	repository.On("Save", mock.Anything).Return(errors.New(mockedErrorMessage))

	_, err := service.Create(newCampaignDto)

	assert.NotNil(err)
	assert.Equal(mockedErrorMessage, err.Error())
	repository.AssertNumberOfCalls(t, "Save", 1)
}
