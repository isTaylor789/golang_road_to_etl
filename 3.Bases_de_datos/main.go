package main

import (
	"invoice_app/database"
	"invoice_app/routes"
	"log"
	"net/http"
)

func main() {
	database.InitDB()

	r := routes.SetupRoutes()

	log.Println("Server running on port 10080")
	log.Fatal(http.ListenAndServe(":10080", r))
}
