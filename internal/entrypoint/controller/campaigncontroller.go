package controller

import (
	"emailn/internal/core/command"
	"emailn/internal/core/service"
	"net/http"

	"github.com/go-chi/render"
)

type CampaignController struct {
	CampaignService service.CampaignService
}

func (controller *CampaignController) Create(writer http.ResponseWriter, request *http.Request) (interface{}, int, error) {
	var command command.NewCampaignCommand
	render.DecodeJSON(request.Body, &command)

	id, err := controller.CampaignService.Create(command)

	return map[string]string{"id": id}, 201, err
}

func (controller *CampaignController) FindAll(writer http.ResponseWriter, request *http.Request) (interface{}, int, error) {
	campaign, err := controller.CampaignService.FindAll()

	return campaign, 200, err
}
