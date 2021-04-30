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
		models.WithMatchModel(),
		models.WithPredictionModel(),
	)
	must(err)

	// Initiate controllers
	authController := controllers.NewAuthController(services.User)
	userController := controllers.NewUserController(services.User)
	matchController := controllers.NewMatchController(services.Match)
	predictionController := controllers.NewPredictionController(services.Prediction)

	// Initiate all midlewares
	//userMW := middleware.RequireUser{
	//	UserService:        services.User,
	//}

	// Initiate all server routes
	r := mux.NewRouter()
	r.HandleFunc("/register", authController.Register).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/login", authController.Login).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/activation/{token}", authController.Activation).Methods("GET")

	//r.HandleFunc("/users/{id}", userMW.Required(userController.Index)).Methods("GET")
	r.HandleFunc("/users/top", userController.GetTop).Methods("GET")
	//r.HandleFunc("/users/", authController.Update).Methods("PUT")
	//r.HandleFunc("/users/", authController.Delete).Methods("DELETE")

	r.HandleFunc("/matches", matchController.GetAll).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/matches", matchController.Create).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/matches/{id}", matchController.Update).Methods(http.MethodPut, http.MethodOptions)
	r.HandleFunc("/matches/{id}", matchController.Delete).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/predictions/{id}", predictionController.GetAllByUserId).Methods(http.MethodGet, http.MethodOptions)

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
