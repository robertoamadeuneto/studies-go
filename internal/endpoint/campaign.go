package endpoint

import (
	"emailn/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (handler *Handler) PostCampaign(writer http.ResponseWriter, request *http.Request) (interface{}, int, error) {
	var newCampaignDto contract.NewCampaignDto
	render.DecodeJSON(request.Body, &newCampaignDto)

	id, err := handler.CampaignService.Create(newCampaignDto)

	return map[string]string{"id": id}, 201, err
}

func (handler *Handler) GetCampaigns(writer http.ResponseWriter, request *http.Request) (interface{}, int, error) {
	campaign, err := handler.CampaignService.Get()

	return campaign, 200, err
}
