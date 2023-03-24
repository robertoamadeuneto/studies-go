package repository

import (
	"emailn/internal/core/entity"

	"gorm.io/gorm"
)

type CampaignRepositoryImpl struct {
	DatabaseConnection *gorm.DB
}

func (repository *CampaignRepositoryImpl) Save(campaign *entity.Campaign) error {
	if result := repository.DatabaseConnection.Create(campaign); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *CampaignRepositoryImpl) FindAll() ([]entity.Campaign, error) {
	var campaigns []entity.Campaign

	if result := repository.DatabaseConnection.Find(&campaigns); result.Error != nil {
		return nil, result.Error
	}

	return campaigns, nil
}
