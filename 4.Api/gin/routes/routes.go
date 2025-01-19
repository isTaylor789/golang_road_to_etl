package routes

import (
	"shopping_car/controllers"
	"shopping_car/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	// Ruta para generar el token
	router.POST("/auth/token", controllers.GenerateToken)

	// Rutas protegidas por el middleware JWT
	router.GET("/cars", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		controllers.GetCars(c, db)
	})
}
