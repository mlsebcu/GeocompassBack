package models

import "gorm.io/gorm"

type Geovisitas struct {
	gorm.Model
	BeneficiaryID   int    `json:"beneficiary_id" gorm:"not null"`
	Nombre          string `json:"nombre" gorm:"not null"`
	FechaNacimiento string `json:"fecha_nacimiento" gorm:"not null"` // Puede ser time.Time
	Edad            int    `json:"edad" gorm:"not null"`
	FechaVisita     string `json:"fecha_visita" gorm:"not null"` // Puede ser time.Time
	Direccion       string `json:"direccion" gorm:"not null"`
	Telefono        string `json:"telefono" gorm:"not null"`
	NombreMadre     string `json:"nombre_madre"`
	NombrePadre     string `json:"nombre_padre"`
	NombreEncargado string `json:"nombre_encargado"`
	Hombres         int    `json:"hombres"`
	Mujeres         int    `json:"mujeres"`
	InscritosCDI    int    `json:"inscritos_cdi"`
	ConQuienVive    string `json:"con_quien_vive"` // 'Padre', 'Madre', 'Padre/Madre'
	ComoVive        string `json:"como_vive"`      // 'Casa propia', 'Alquilada', 'Prestada'
	TipoCasa        string `json:"tipo_casa"`      // 'Madera', 'Block'
	QuienesTrabajan string `json:"quienes_trabajan"`
	TrabajaNino     string `json:"trabaja_nino"` // 'Si', 'No'
	Observaciones   string `json:"observaciones"`
}
