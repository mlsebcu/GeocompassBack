package handler

import (
	"fmt"
	"net/http"
	"os"

	"log"

	"github.com/alejoca7/geo-back/db"
	"github.com/alejoca7/geo-back/models"
	"github.com/alejoca7/geo-back/routes"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
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

	// Inicio rutas
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	router.HandleFunc("/login", routes.LoginHandler).Methods("POST")

	// Rutas para los geopoints
	router.HandleFunc("/geopoints", routes.GetGeopointsHandler).Methods("GET")
	router.HandleFunc("/geopoints", routes.PostGeopointHandler).Methods("POST")
	router.HandleFunc("/geopoints/{id}", routes.GetGeopointHandler).Methods("GET")
	router.HandleFunc("/geopoints/{id}", routes.DeleteGeopointHandler).Methods("DELETE")
	router.HandleFunc("/geopoints/{id}", routes.UpdateGeopointHandler).Methods("PUT")

	// Rutas para las geovisitas
	router.HandleFunc("/geovisitas", routes.GetGeovisitasHandler).Methods("GET")
	router.HandleFunc("/geovisitas", routes.PostGeovisitaHandler).Methods("POST")
	router.HandleFunc("/geovisitas/{id}", routes.GetGeovisitaHandler).Methods("GET")
	router.HandleFunc("/geovisitas/{id}", routes.DeleteGeovisitaHandler).Methods("DELETE")
	router.HandleFunc("/geovisitas/{id}", routes.UpdateGeovisitaHandler).Methods("PUT")

	// Rutas para geodatos que simplemente reutilizan los manejadores de geovisitas
	router.HandleFunc("/geodatos", routes.GetGeovisitasHandler).Methods("GET")
	router.HandleFunc("/geodatos/{id}", routes.GetGeovisitaHandler).Methods("GET")

	// Ruta para subir imágenes
	router.HandleFunc("/upload", UploadImageHandler).Methods("POST")

	// Fin rutas

	// Envuelve el router con CORS
	handler := c.Handler(router)

	// Maneja la petición
	handler.ServeHTTP(w, r)
}

// UploadImageHandler maneja la subida de imágenes a Cloudinary
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	// Limitar el tamaño del archivo a 10 MB
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error al leer el archivo", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Generar un nombre único para la imagen
	imageID := uuid.New().String() + "-" + handler.Filename

	// Inicializar Cloudinary
	cld, err := cloudinary.NewFromParams(
		os.Getenv("CLOUDINARY_CLOUD_NAME"),
		os.Getenv("CLOUDINARY_API_KEY"),
		os.Getenv("CLOUDINARY_API_SECRET"),
	)
	if err != nil {
		log.Fatalf("Error al inicializar Cloudinary: %v", err)
	}

	// Subir la imagen a Cloudinary
	resp, err := cld.Upload.Upload(r.Context(), file, uploader.UploadParams{
		PublicID: imageID,
	})
	if err != nil {
		http.Error(w, "Error al subir la imagen a Cloudinary", http.StatusInternalServerError)
		return
	}

	// Responder con la URL de la imagen
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"image_url": "%s"}`, resp.SecureURL)))
}
