package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/alejoca7/geo-back/db"
	"github.com/alejoca7/geo-back/models"
	"github.com/gorilla/mux"
)

// UploadImageHandler maneja la subida de imágenes
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Crear el archivo en el servidor
	fileName := filepath.Join("uploads", header.Filename)
	out, err := os.Create(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// Guardar el archivo
	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Devolver la URL de la imagen guardada
	imageURL := fmt.Sprintf("http://localhost:8080/%s", fileName)
	json.NewEncoder(w).Encode(map[string]string{"image_url": imageURL})
}

// GetGeopointsHandler obtiene todos los geopoints
func GetGeopointsHandler(w http.ResponseWriter, r *http.Request) {
	var geopoints []models.Geopoint

	// Imprimir un mensaje en la consola cuando se recibe una solicitud
	fmt.Println("Solicitud para obtener geopoints recibida")

	// Buscar los geopoints en la base de datos
	result := db.DB.Find(&geopoints)
	if result.Error != nil {
		fmt.Println("Error obteniendo geopoints:", result.Error)
	} else {
		fmt.Println("Geopoints obtenidos correctamente:", len(geopoints))
	}

	// Imprimir los geopoints en la consola
	for _, geopoint := range geopoints {
		fmt.Printf("ID: %d, Nombre: %s, Latitud: %f, Longitud: %f\n", geopoint.ID, geopoint.Nombre, geopoint.Latitude, geopoint.Longitude)
	}

	// Devolver los geopoints al frontend
	json.NewEncoder(w).Encode(&geopoints)
}

// GetGeopointHandler obtiene un geopoint específico por ID
func GetGeopointHandler(w http.ResponseWriter, r *http.Request) {
	var geopoint models.Geopoint
	params := mux.Vars(r)
	db.DB.First(&geopoint, params["id"])

	if geopoint.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Geopoint not found"))
		return
	}
	json.NewEncoder(w).Encode(&geopoint)
}

// PostGeopointHandler crea un nuevo geopoint
func PostGeopointHandler(w http.ResponseWriter, r *http.Request) {
	var geopoint models.Geopoint
	json.NewDecoder(r.Body).Decode(&geopoint)

	createdGeopoint := db.DB.Create(&geopoint)
	if createdGeopoint.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(createdGeopoint.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&geopoint)
}

// DeleteGeopointHandler elimina un geopoint por ID
func DeleteGeopointHandler(w http.ResponseWriter, r *http.Request) {
	var geopoint models.Geopoint
	params := mux.Vars(r)
	db.DB.First(&geopoint, params["id"])

	if geopoint.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Geopoint not found"))
		return
	}

	db.DB.Unscoped().Delete(&geopoint)
	w.WriteHeader(http.StatusOK)
}
