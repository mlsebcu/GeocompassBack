package handler

import (
	"net/http"

	"github.com/alejoca7/geo-back/db"
	"github.com/alejoca7/geo-back/models"
	"github.com/alejoca7/geo-back/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Función que maneja las peticiones HTTP
func VercelHandler(w http.ResponseWriter, r *http.Request) {
	db.DBConnection()

	// AutoMigrate asegura que las tablas para los modelos existan
	db.DB.AutoMigrate(&models.User{}, &models.Geopoint{}, &models.Geovisitas{})

	// Crea un nuevo enrutador
	router := mux.NewRouter()

	// Configura CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Definir rutas
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	// (Más rutas...)

	// Envuelve el router con CORS
	handler := c.Handler(router)

	// Maneja la petición
	handler.ServeHTTP(w, r)
}
