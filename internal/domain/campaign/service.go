package campaign

import "emailn/internal/contract"

type Service struct {
	Repository Repository
}

func (service *Service) Create(dto contract.NewCampaignDto) error {
	return nil
}
