package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// Importar el paquete docs (para Swagger) y el godotenv
	_ "microserviceupdateworkshops/docs"
	"microserviceupdateworkshops/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Workshop API
// @version 1.0
// @description API de microservicio para gestión de workshops
// @BasePath /

func main() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("No se pudo cargar el archivo .env (o no existe).")
	}

	r := mux.NewRouter()

	// Registrar rutas
	routes.RegisterRoutes(r)

	// Usar el handler de Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Tomar el puerto desde la variable de entorno
	port := os.Getenv("PORT")
	if port == "" {
		port = "8098" // Valor por defecto si no existe en .env
	}

	fmt.Printf("API listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
