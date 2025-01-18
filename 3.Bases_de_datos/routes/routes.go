package routes

import (
	"invoice_app/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", handlers.CreateUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods(http.MethodGet)

	return r
}
