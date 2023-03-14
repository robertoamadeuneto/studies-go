package campaign

import (
	"emailn/internal/contract"
)

type Service struct {
	Repository Repository
}

func (service *Service) Create(dto contract.NewCampaignDto) (string, error) {
	campaign, err := NewCampaign(dto.Name, dto.Content, dto.Contacts)

	if err != nil {
		return "", err
	}

	err = service.Repository.Save(campaign)

	if err != nil {
		return "", err
	}

	return campaign.Id, nil
}
