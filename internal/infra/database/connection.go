package database

import (
	"emailn/internal/core/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	databaseUrl := "postgres://emailn:emailn@localhost:5432/emailn"

	connection, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	connection.AutoMigrate(&entity.Campaign{}, &entity.Contact{})

	return connection
}
