package main

import (
	"net/http"

	"github.com/alejoca7/geo-back/db"
	"github.com/alejoca7/geo-back/models"
	"github.com/alejoca7/geo-back/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	db.DBConnection()

	// AutoMigrate se asegura de que las tablas para los modelos existan
	db.DB.AutoMigrate(&models.User{}, &models.Geopoint{}) // Incluimos la migración del modelo Geopoint

	r := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/login", routes.LoginHandler).Methods("POST")

	// Nuevas rutas para los geopoints
	r.HandleFunc("/geopoints", routes.GetGeopointsHandler).Methods("GET")
	r.HandleFunc("/geopoints", routes.PostGeopointHandler).Methods("POST")
	r.HandleFunc("/geopoints/{id}", routes.GetGeopointHandler).Methods("GET")
	r.HandleFunc("/geopoints/{id}", routes.DeleteGeopointHandler).Methods("DELETE")

	// Ruta para subir imágenes
	r.HandleFunc("/upload", routes.UploadImageHandler).Methods("POST")

	// Ruta para servir archivos estáticos
	r.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))

	handler := c.Handler(r)
	http.ListenAndServe(":8080", handler)
}
