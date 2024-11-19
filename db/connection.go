package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	// Usa POSTGRES_URL directamente si está disponible
	dsn := os.Getenv("POSTGRES_URL")

	// Verifica si la variable de entorno está configurada
	if dsn == "" {
		log.Fatal("Falta la variable de entorno POSTGRES_URL")
	}

	// Probar la conexión a la base de datos usando Gorm y PostgreSQL
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	} else {
		log.Println("Conexión exitosa a la base de datos")
	}
}
