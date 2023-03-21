package service

import (
	"emailn/internal/core/command"
	"emailn/internal/core/entity"
	internalerror "emailn/internal/core/error"
	"emailn/internal/core/repository"
)

type CampaignService interface {
	Create(command command.NewCampaignCommand) (string, error)
	FindAll() ([]entity.Campaign, error)
}

type CampaignServiceImpl struct {
	CampaignRepository repository.CampaignRepository
}

func (service *CampaignServiceImpl) Create(command command.NewCampaignCommand) (string, error) {
	campaign, err := entity.NewCampaign(command.Name, command.Content, command.Contacts)

	if err != nil {
		return "", err
	}

	err = service.CampaignRepository.Save(campaign)

	if err != nil {
		return "", internalerror.InternalServerError
	}

	return campaign.Id, nil
}

func (service *CampaignServiceImpl) FindAll() ([]entity.Campaign, error) {
	campaign, err := service.CampaignRepository.FindAll()

	if err != nil {
		return nil, internalerror.InternalServerError
	}

	return campaign, nil
}
