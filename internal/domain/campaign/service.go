package campaign

import (
	"emailn/internal/contract"
)

type Service struct {
	Repository Repository
}

func (service *Service) Create(dto contract.NewCampaignDto) (string, error) {
	campaign, _ := NewCampaign(dto.Name, dto.Content, dto.Contacts)

	service.Repository.Save(campaign)

	return campaign.Id, nil
}
