package main

import (
	"log"
	"net/http"

	"shopping_car/db"
	"shopping_car/handlers"
	"shopping_car/middleware"
)

func main() {
	// Conexión a la base de datos
	database := db.InitDB()
	db.MigrateDB(database)

	// Rutas
	http.HandleFunc("/auth/token", handlers.GenerateToken)

	http.Handle("/cars", middleware.JWTAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetCars(w, r, database)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	// Iniciar servidor
	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
