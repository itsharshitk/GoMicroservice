package routes

import (
	"auth-service/config"
	"auth-service/models"
	"auth-service/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	hashedPwd, _ := utils.HashPassword(user.Password)

	_, err := config.DB.Exec("INSERT INTO users (email, password) VALUES (?, ?)", user.Email, hashedPwd)

	if err != nil {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	var storedPwd string
	err := config.DB.QueryRow("SELECT password FROM users WHERE email = ?", user.Email).Scan(&storedPwd)
	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPasswordHash(user.Password, storedPwd) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, "JWT error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
