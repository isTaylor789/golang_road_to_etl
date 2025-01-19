package db

import (
	"log"

	"shopping_car/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=dev password=123456 dbname=astra-db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}
	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&models.Car{}, &models.Item{})
}
