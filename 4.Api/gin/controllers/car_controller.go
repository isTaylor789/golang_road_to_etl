package controllers

import (
	"net/http"

	"shopping_car/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCars(c *gin.Context, db *gorm.DB) {
	var cars []models.Car
	if err := db.Preload("Items").Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cars)
}
