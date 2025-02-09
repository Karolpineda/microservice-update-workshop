package controllers

import (
	"encoding/json"
	"microserviceupdateworkshops/config"
	"microserviceupdateworkshops/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// UpdateWorkshop
// @Summary Actualiza un workshop existente
// @Description Modifica un workshop basado en su ID
// @Tags Workshops
// @Accept  json
// @Produce  json
// @Param   id  path  string  true  "ID del Workshop"
// @Param   workshop  body  models.Workshop  true  "Datos actualizados del Workshop"
// @Success 200 {object} models.Workshop
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Workshop not found"
// @Router /workshops/{id} [put]
func UpdateWorkshop(w http.ResponseWriter, r *http.Request) {
	// Obtiene el ID desde la URL con Gorilla Mux
	vars := mux.Vars(r)
	id, exists := vars["id"]
	if !exists {
		http.Error(w, "Missing workshop ID", http.StatusBadRequest)
		return
	}

	// Convierte el ID a UUID
	workshopID, err := uuid.Parse(id)
	if err != nil {

		http.Error(w, "Invalid workshop ID format", http.StatusBadRequest)
		return
	}

	// Decodifica el cuerpo de la solicitud
	var updatedWorkshop models.Workshop
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&updatedWorkshop); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Asigna el ID al objeto actualizado
	updatedWorkshop.ID = workshopID

	// Conectar con la base de datos y actualizar
	db := config.SetupDatabase()
	workshop, err := models.UpdateWorkshop(db, &updatedWorkshop)
	if err != nil {
		http.Error(w, "Workshop not found", http.StatusNotFound)
		return
	}

	// Responde con el objeto actualizado
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workshop)
}

// HealthCheck
// @Summary Verifica el estado del microservicio
// @Description Retorna un mensaje que indica que el microservicio está en funcionamiento
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "Microservice is up and running"})
}

// GetAllWorkshops
// @Summary Obtiene todos los Workshops
// @Description Devuelve una lista JSON con todos los Workshops existentes en la base de datos
// @Tags Workshops
// @Produce json
// @Success 200 {array} models.Workshop
// @Router /workshops [get]
func GetAllWorkshops(w http.ResponseWriter, r *http.Request) {
	db := config.SetupDatabase()
	workshops := models.GetAllWorkshops(db)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workshops)
}
