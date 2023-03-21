package database

import "emailn/internal/domain/campaign"

type CampaignRepositoryImpl struct {
	campaigns []campaign.Campaign
}

func (repository *CampaignRepositoryImpl) Save(campaign *campaign.Campaign) error {
	repository.campaigns = append(repository.campaigns, *campaign)

	return nil
}

func (repository *CampaignRepositoryImpl) FindAll() ([]campaign.Campaign, error) {
	if repository.campaigns == nil {
		return []campaign.Campaign{}, nil
	}

	return repository.campaigns, nil
}
