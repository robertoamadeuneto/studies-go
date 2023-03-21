package campaign

import (
	"emailn/internal"
	"emailn/internal/command"
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

func (repository *repositoryMock) FindAll() ([]Campaign, error) {
	args := repository.Called()

	return args.Get(0).([]Campaign), args.Error(1)
}

var (
	newCampaigncommand = command.NewCampaignCommand{
		Name:     "New Campaign",
		Content:  "This is a test Campaign",
		Contacts: []string{"test.campaign@email.com", "+551234567890"},
	}
)

func Test_Create_WithValidData_ReturnsCampaignId(t *testing.T) {
	assert := assert.New(t)
	repository := new(repositoryMock)
	service := ServiceImpl{repository}
	repository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaigncommand.Name ||
			campaign.Content != newCampaigncommand.Content ||
			len(campaign.Contacts) != len(newCampaigncommand.Contacts) {
			return false
		}

		return true
	})).Return(nil)

	id, err := service.Create(newCampaigncommand)

	assert.NotEmpty(id)
	assert.Nil(err)
	repository.AssertNumberOfCalls(t, "Save", 1)
	repository.AssertExpectations(t)
}

func Test_Create_WithInvalidNewCampaigncommand_ReturnsError(t *testing.T) {
	assert := assert.New(t)
	repository := new(repositoryMock)
	service := ServiceImpl{repository}

	_, err := service.Create(command.NewCampaignCommand{})

	assert.NotNil(err)
	repository.AssertNotCalled(t, "Save")
}

func Test_Create_WhenRepositorySaveFails_ReturnsError(t *testing.T) {
	assert := assert.New(t)
	repository := new(repositoryMock)
	service := ServiceImpl{repository}
	repository.On("Save", mock.Anything).Return(errors.New("Error trying to communicate with database!"))

	_, err := service.Create(newCampaigncommand)

	assert.NotNil(err)
	assert.True(errors.Is(internal.InternalServerError, err))
	repository.AssertNumberOfCalls(t, "Save", 1)
}
