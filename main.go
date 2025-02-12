package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "microserviceupdateworkshops/docs"
	"microserviceupdateworkshops/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title... (tu swagger docs aquí)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error cargando .env")
	}

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Configuración CORS
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // En producción restringe esto
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8097"
	}

	fmt.Printf("Servidor iniciado en puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, cors(r))) // <--- Aplica el middleware aquí
}
