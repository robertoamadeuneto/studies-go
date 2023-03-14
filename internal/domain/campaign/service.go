package campaign

import (
	"emailn/internal/contract"
	"emailn/internal/internalerrors"
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
		return "", internalerrors.InternalServerError
	}

	return campaign.Id, nil
}
