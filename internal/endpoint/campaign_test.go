package endpoint

import (
	"bytes"
	"emailn/internal/domain/campaign"
	"emailn/internal/dto"
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

func (service *serviceMock) Create(dto dto.NewCampaignDto) (string, error) {
	args := service.Called(dto)

	return args.String(0), args.Error(1)
}

func (service *serviceMock) Get() ([]campaign.Campaign, error) {
	args := service.Called()

	return args.Get(0).([]campaign.Campaign), args.Error(1)
}

func Test_CampaignPost_WhenRequestIsValid_CreatesNewCampaign(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)
	handler := Handler{CampaignService: service}
	body := dto.NewCampaignDto{
		Name:     "Name",
		Content:  "Content",
		Contacts: []string{"test@email.com", "+55987458565"},
	}
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(body)
	request, _ := http.NewRequest("POST", "/", &buffer)
	recorder := httptest.NewRecorder()
	service.On("Create", mock.MatchedBy(func(request dto.NewCampaignDto) bool {
		if request.Name == body.Name && request.Content == body.Content {
			return true
		}

		return false
	})).Return("123", nil)

	json, status, err := handler.PostCampaign(recorder, request)
	assert.Equal(201, status)
	assert.NotNil(json)
	assert.Nil(err)
}
