package campaign

import (
	"emailn/internal"
	"emailn/internal/command"
)

type Service interface {
	Create(command command.NewCampaignCommand) (string, error)
	FindAll() ([]Campaign, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (service *ServiceImpl) Create(command command.NewCampaignCommand) (string, error) {
	campaign, err := NewCampaign(command.Name, command.Content, command.Contacts)

	if err != nil {
		return "", err
	}

	err = service.Repository.Save(campaign)

	if err != nil {
		return "", internal.InternalServerError
	}

	return campaign.Id, nil
}

func (service *ServiceImpl) FindAll() ([]Campaign, error) {
	campaign, err := service.Repository.FindAll()

	if err != nil {
		return nil, internal.InternalServerError
	}

	return campaign, nil
}
