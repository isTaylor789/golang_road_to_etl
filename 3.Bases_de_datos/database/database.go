package database

import (
	"invoice_app/config"
	"invoice_app/models"
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	DB = config.ConnectDB()

	// Migrar los modelos
	err := DB.AutoMigrate(&models.User{}, &models.Order{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
