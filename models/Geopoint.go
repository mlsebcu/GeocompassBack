package models

import "gorm.io/gorm"

type Geopoint struct {
	gorm.Model
	BeneficiaryID   int     `json:"beneficiary_id" gorm:"not null"`
	Nombre          string  `json:"nombre" gorm:"not null"`
	Latitude        float64 `json:"latitude" gorm:"not null"`
	Longitude       float64 `json:"longitude" gorm:"not null"`
	Address         string  `json:"address"`
	FechaNacimiento string  `json:"fecha_nacimiento"` // Nuevo campo
	Edad            int     `json:"edad"`             // Nuevo campo
	Telefono        string  `json:"telefono"`         // Nuevo campo
	ImageURL        string  `json:"image_url"`        // URL para la imagen de la vivienda
}
