package database

import "gorm.io/gorm"

type Handler struct {
	connection *gorm.DB
}

func New(connection *gorm.DB) Handler {
	return Handler{connection}
}
