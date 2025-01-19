package handlers

import (
	"encoding/json"
	"net/http"

	"shopping_car/models"

	"gorm.io/gorm"
)

func GetCars(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var cars []models.Car
	if err := db.Preload("Items").Find(&cars).Error; err != nil {
		http.Error(w, "Error al obtener los autos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}
