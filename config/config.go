package config

import (
	"fmt"
	"log"
	"os"

	"microserviceupdateworkshops/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func SetupDatabase() *gorm.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		GetEnvVariable("DB_HOST"),
		GetEnvVariable("DB_PORT"),
		GetEnvVariable("DB_USER"),
		GetEnvVariable("DB_PASSWORD"),
		GetEnvVariable("DB_NAME"),
	)

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Migraciones
	db.AutoMigrate(&models.Workshop{})
	return db
}

// GetEnvVariable obtiene una variable de entorno
func GetEnvVariable(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("No environment variable found for %s", key)
	}
	return val
}
