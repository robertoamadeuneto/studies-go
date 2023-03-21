package repository

import "emailn/internal/core/entity"

type CampaignRepository interface {
	Save(campaign *entity.Campaign) error
	FindAll() ([]entity.Campaign, error)
}
