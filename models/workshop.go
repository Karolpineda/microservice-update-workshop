package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Workshop struct {
	ID      uuid.UUID `gorm:"primary_key;type:uuid" json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
}

// UpdateWorkshop - Actualiza un workshop en la base de datos
func UpdateWorkshop(db *gorm.DB, workshop *Workshop) (*Workshop, error) {
	var existingWorkshop Workshop
	if err := db.First(&existingWorkshop, "id = ?", workshop.ID).Error; err != nil {
		return nil, errors.New("Workshop not found")
	}

	// Actualizar valores permitidos
	existingWorkshop.Name = workshop.Name
	existingWorkshop.Address = workshop.Address

	if err := db.Save(&existingWorkshop).Error; err != nil {
		return nil, err
	}
	return &existingWorkshop, nil
}

// Obtener todos los workshops
func GetAllWorkshops(db *gorm.DB) []Workshop {
	var workshops []Workshop
	db.Find(&workshops)
	return workshops
}
