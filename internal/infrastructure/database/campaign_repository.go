package database

import "emailn/internal/domain/campaign"

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (repository *CampaignRepository) Save(campaign *campaign.Campaign) error {
	repository.campaigns = append(repository.campaigns, *campaign)

	return nil
}

func (repository *CampaignRepository) Get() ([]campaign.Campaign, error) {
	if repository.campaigns == nil {
		return []campaign.Campaign{}, nil
	}

	return repository.campaigns, nil
}
