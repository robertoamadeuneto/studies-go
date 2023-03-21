package service

import (
	"emailn/internal/core/command"
	"emailn/internal/core/entity"
	internalerror "emailn/internal/core/error"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (repository *repositoryMock) Save(campaign *entity.Campaign) error {
	args := repository.Called(campaign)

	return args.Error(0)
}

func (repository *repositoryMock) FindAll() ([]entity.Campaign, error) {
	args := repository.Called()

	return args.Get(0).([]entity.Campaign), args.Error(1)
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
	service := CampaignServiceImpl{repository}
	repository.On("Save", mock.MatchedBy(func(campaign *entity.Campaign) bool {
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
	service := CampaignServiceImpl{repository}

	_, err := service.Create(command.NewCampaignCommand{})

	assert.NotNil(err)
	repository.AssertNotCalled(t, "Save")
}

func Test_Create_WhenRepositorySaveFails_ReturnsError(t *testing.T) {
	assert := assert.New(t)
	repository := new(repositoryMock)
	service := CampaignServiceImpl{repository}
	repository.On("Save", mock.Anything).Return(errors.New("Error trying to communicate with database!"))

	_, err := service.Create(newCampaigncommand)

	assert.NotNil(err)
	assert.True(errors.Is(internalerror.InternalServerError, err))
	repository.AssertNumberOfCalls(t, "Save", 1)
}
