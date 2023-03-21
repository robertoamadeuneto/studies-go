package repository

import "emailn/internal/core/entity"

type CampaignRepositoryImpl struct {
	campaigns []entity.Campaign
}

func (repository *CampaignRepositoryImpl) Save(campaign *entity.Campaign) error {
	repository.campaigns = append(repository.campaigns, *campaign)

	return nil
}

func (repository *CampaignRepositoryImpl) FindAll() ([]entity.Campaign, error) {
	if repository.campaigns == nil {
		return []entity.Campaign{}, nil
	}

	return repository.campaigns, nil
}
