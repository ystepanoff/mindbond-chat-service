package db

import (
	"log"

	"flotta-home/mindbond/chat-service/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(
		&models.Contact{},
		&models.Chat{},
		&models.Message{},
	); err != nil {
		log.Fatalln(err)
	}

	return Handler{db}
}
