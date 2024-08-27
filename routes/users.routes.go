package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/alejoca7/geo-back/db"
	"github.com/alejoca7/geo-back/models"
	"github.com/gorilla/mux"
)

// GetUsersHandler obtiene todos los usuarios
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

// GetUserHandler obtiene un usuario específico por ID
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	json.NewEncoder(w).Encode(&user)
}

// PostUserHandler crea un nuevo usuario y encripta su contraseña antes de guardarla
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decodificar el JSON recibido
	json.NewDecoder(r.Body).Decode(&user)

	// Crear el usuario en la base de datos
	createdUser := db.DB.Create(&user)
	if createdUser.Error != nil {
		w.WriteHeader(http.StatusBadRequest) // error 400
		w.Write([]byte(createdUser.Error.Error()))
		return
	}

	json.NewEncoder(w).Encode(&user)
}

// DeleteUserHandler elimina un usuario por ID
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}

// LoginHandler maneja la autenticación del usuario
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input models.User
	var user models.User

	// Decodificar el JSON recibido
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Solicitud inválida"))
		log.Println("Error al decodificar el JSON:", err)
		return
	}

	log.Println("Datos de entrada recibidos:", input)

	// Asegurarse de que se está recibiendo una contraseña no vacía
	if input.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Contraseña no proporcionada"))
		log.Println("Contraseña no proporcionada en la solicitud")
		return
	}

	// Buscar el usuario por nombre de usuario (ignorar mayúsculas y minúsculas)
	if err := db.DB.Where("LOWER(username) = LOWER(?)", input.Username).First(&user).Error; err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Usuario no encontrado"))
		log.Println("Usuario no encontrado:", input.Username)
		return
	}

	// Log para verificar que se está comparando la contraseña
	log.Println("Contraseña ingresada:", input.Password)
	log.Println("Contraseña en la base de datos:", user.Password)

	// Comparar la contraseña recibida con la contraseña encriptada almacenada
	if !user.ComparePassword(input.Password) {
		log.Println("ComparePassword devolvió false, contraseña incorrecta.")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Contraseña incorrecta"))
		return
	}
	log.Println("ComparePassword devolvió true, contraseña correcta.")

	// Si la autenticación es exitosa
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Inicio de sesión exitoso"))
	log.Println("Usuario autenticado correctamente:", input.Username)
}
