package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alejoca7/geo-back/handler"
)

func main() {
	// Obtiene el puerto desde las variables de entorno, útil en producción
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Usa 8080 por defecto en local
	}

	// Inicia el servidor HTTP local usando VercelHandler
	log.Printf("Server running on port %s", port)
	err := http.ListenAndServe(":"+port, http.HandlerFunc(handler.VercelHandler))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
