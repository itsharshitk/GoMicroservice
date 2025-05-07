// File: auth-service/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"auth-service/config"
	"auth-service/routes"

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
	r.HandleFunc("/register", routes.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", routes.LoginHandler).Methods("POST")

	fmt.Println("Auth Service running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
