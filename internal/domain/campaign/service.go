package campaign

import (
	"emailn/internal/dto"
	"emailn/internal/internalerror"
)

type Service struct {
	Repository Repository
}

func (service *Service) Create(dto dto.NewCampaignDto) (string, error) {
	campaign, err := NewCampaign(dto.Name, dto.Content, dto.Contacts)

	if err != nil {
		return "", err
	}

	err = service.Repository.Save(campaign)

	if err != nil {
		return "", internalerror.InternalServerError
	}

	return campaign.Id, nil
}

func (service *Service) Get() ([]Campaign, error) {
	campaign, err := service.Repository.Get()

	if err != nil {
		return nil, internalerror.InternalServerError
	}

	return campaign, nil
}
