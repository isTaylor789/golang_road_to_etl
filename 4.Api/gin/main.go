package main

import (
	"shopping_car/config"
	"shopping_car/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa la conexión a la base de datos
	db := config.InitDB()

	// Migra los modelos
	config.MigrateDB(db)

	// Configura el router
	router := gin.Default()
	routes.RegisterRoutes(router, db)

	// Inicia el servidor
	router.Run(":8080")
}
