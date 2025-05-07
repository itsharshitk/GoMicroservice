package main

import (
	"log"
	"net/http"
	"user-service/config"
	"user-service/middleware"
	"user-service/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()

	r := mux.NewRouter()

	protected := r.PathPrefix("/profile").Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/{email}", routes.GetProfile).Methods("GET")
	protected.HandleFunc("/", routes.CreateOrUpdateProfile).Methods("POST")

	log.Println("User Service running on :8001")
	log.Fatal(http.ListenAndServe(":8001", r))
}
