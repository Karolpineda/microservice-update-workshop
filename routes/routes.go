package routes

import (
	"microserviceupdateworkshops/controllers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	// Ruta Health
	r.HandleFunc("/health", controllers.HealthCheck).Methods("GET")

	// Ruta para actualizar un workshop
	r.HandleFunc("/workshops/{id}", controllers.UpdateWorkshop).Methods("PUT")

}
