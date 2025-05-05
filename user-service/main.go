package main

import (
	"log"
	"net/http"
	"user-service/config"
	"user-service/middleware"
	"user-service/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()

	r := mux.NewRouter()

	protected := r.PathPrefix("/profile").Subrouter()
	protected.Use(middleware.JWTMiddleware)

	protected.HandleFunc("/{email}", routes.GetProfile).Methods("GET")
	protected.HandleFunc("/", routes.CreateOrUpdateProfile).Methods("POST")

	log.Println("User Service running on :8001")
	log.Fatal(http.ListenAndServe(":8001", r))
}
