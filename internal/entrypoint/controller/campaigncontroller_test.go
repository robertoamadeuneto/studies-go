package controller

import (
	"bytes"
	"emailn/internal/core/command"
	"emailn/internal/core/entity"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (service *serviceMock) Create(command command.NewCampaignCommand) (string, error) {
	args := service.Called(command)

	return args.String(0), args.Error(1)
}

func (service *serviceMock) FindAll() ([]entity.Campaign, error) {
	args := service.Called()

	return args.Get(0).([]entity.Campaign), args.Error(1)
}

func Test_CampaignPost_WhenRequestIsValid_CreatesNewCampaign(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)
	handler := CampaignController{CampaignService: service}
	body := command.NewCampaignCommand{
		Name:     "Name",
		Content:  "Content",
		Contacts: []string{"test@email.com", "+55987458565"},
	}
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(body)
	request, _ := http.NewRequest("POST", "/", &buffer)
	recorder := httptest.NewRecorder()
	service.On("Create", mock.MatchedBy(func(request command.NewCampaignCommand) bool {
		if request.Name == body.Name && request.Content == body.Content {
			return true
		}

		return false
	})).Return("123", nil)

	json, status, err := handler.Create(recorder, request)
	assert.Equal(201, status)
	assert.NotNil(json)
	assert.Nil(err)
}
