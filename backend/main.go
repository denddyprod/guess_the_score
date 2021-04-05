package main

import (
	"github.com/gorilla/mux"

	"guess_the_score/backend/controllers"
	"guess_the_score/backend/models"

	"log"
	"net/http"
	"time"
)

func main() {
	var cfg *Config
	cfg, err := LoadConfig()
	must(err)

	// Initiate services with all parameters
	services, err := models.NewServices(
		models.WithDatabase(cfg.MongoDBConfig.ConnectionInfo(), cfg.MongoDBConfig.Name()),
		models.WithUserService(),
	)
	must(err)

	// Initiate controllers
	authController := controllers.NewAuthController(services.User)

	// Initiate all server routes
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", authController.Index).Methods("GET")
	r.HandleFunc("/users/", authController.Register).Methods("POST")
	r.HandleFunc("/users/", authController.Update).Methods("PUT")
	r.HandleFunc("/users/", authController.Delete).Methods("DELETE")

	defer func() {
		// Close Database connection
		err = services.CloseDatabase()
		must(err)
	}()

	// Initiate server configuration
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
		Addr:         ":" + cfg.HTTPServer.Port,
	}

	log.Println("Starting backend application on " + cfg.HTTPServer.Port)
	log.Fatal(server.ListenAndServe())

}

func must(err error) {
	if err != nil {
		log.Println(err)
	}
}
