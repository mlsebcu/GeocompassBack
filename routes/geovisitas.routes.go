package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alejoca7/geo-back/db"
	"github.com/alejoca7/geo-back/models"
	"github.com/gorilla/mux"
)

// GetGeovisitasHandler obtiene todas las geovisitas
func GetGeovisitasHandler(w http.ResponseWriter, r *http.Request) {
	var geovisitas []models.Geovisitas

	// Buscar las geovisitas en la base de datos
	result := db.DB.Find(&geovisitas)
	if result.Error != nil {
		fmt.Println("Error obteniendo geovisitas:", result.Error)
	} else {
		fmt.Println("Geovisitas obtenidas correctamente:", len(geovisitas))
	}

	// Devolver las geovisitas al frontend
	json.NewEncoder(w).Encode(&geovisitas)
}

// GetGeovisitaHandler obtiene una geovisita específica por ID
func GetGeovisitaHandler(w http.ResponseWriter, r *http.Request) {
	var geovisita models.Geovisitas
	params := mux.Vars(r)
	result := db.DB.First(&geovisita, params["id"])

	if result.Error != nil {
		fmt.Println("Error buscando geovisita:", result.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Geovisita not found"))
		return
	}

	json.NewEncoder(w).Encode(&geovisita)
}

// PostGeovisitaHandler crea una nueva geovisita
func PostGeovisitaHandler(w http.ResponseWriter, r *http.Request) {
	var geovisita models.Geovisitas
	json.NewDecoder(r.Body).Decode(&geovisita)

	createdGeovisita := db.DB.Create(&geovisita)
	if createdGeovisita.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(createdGeovisita.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&geovisita)
}

// DeleteGeovisitaHandler elimina una geovisita por ID
func DeleteGeovisitaHandler(w http.ResponseWriter, r *http.Request) {
	var geovisita models.Geovisitas
	params := mux.Vars(r)
	result := db.DB.First(&geovisita, params["id"])

	if result.Error != nil {
		fmt.Println("Error buscando geovisita para eliminar:", result.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Geovisita not found"))
		return
	}

	db.DB.Unscoped().Delete(&geovisita)
	w.WriteHeader(http.StatusOK)
}

// UpdateGeovisitaHandler actualiza una geovisita por ID
func UpdateGeovisitaHandler(w http.ResponseWriter, r *http.Request) {
	var geovisita models.Geovisitas
	params := mux.Vars(r)

	// Buscar la geovisita existente en la base de datos
	result := db.DB.First(&geovisita, params["id"])

	if result.Error != nil {
		fmt.Println("Error buscando geovisita para actualizar:", result.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Geovisita not found"))
		return
	}

	// Actualizar los campos de la geovisita
	var updatedGeovisita models.Geovisitas
	json.NewDecoder(r.Body).Decode(&updatedGeovisita)

	geovisita.Nombre = updatedGeovisita.Nombre
	geovisita.FechaNacimiento = updatedGeovisita.FechaNacimiento
	geovisita.Edad = updatedGeovisita.Edad
	geovisita.FechaVisita = updatedGeovisita.FechaVisita
	geovisita.Address = updatedGeovisita.Address
	geovisita.Telefono = updatedGeovisita.Telefono
	geovisita.NombreMadre = updatedGeovisita.NombreMadre
	geovisita.NombrePadre = updatedGeovisita.NombrePadre
	geovisita.NombreEncargado = updatedGeovisita.NombreEncargado
	geovisita.Hombres = updatedGeovisita.Hombres
	geovisita.Mujeres = updatedGeovisita.Mujeres
	geovisita.InscritosCDI = updatedGeovisita.InscritosCDI
	geovisita.ConQuienVive = updatedGeovisita.ConQuienVive
	geovisita.ComoVive = updatedGeovisita.ComoVive
	geovisita.TipoCasa = updatedGeovisita.TipoCasa
	geovisita.QuienesTrabajan = updatedGeovisita.QuienesTrabajan
	geovisita.TrabajaNino = updatedGeovisita.TrabajaNino
	geovisita.Observaciones = updatedGeovisita.Observaciones

	// Guardar los cambios en la base de datos
	db.DB.Save(&geovisita)

	json.NewEncoder(w).Encode(&geovisita)
}
