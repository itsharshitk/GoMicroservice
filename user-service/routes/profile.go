package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"user-service/config"
	"user-service/models"

	"github.com/gorilla/mux"
)

// Get profile
func GetProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	var profile models.Profile
	err := config.DB.QueryRow("SELECT id, email, name, phone, address FROM profiles WHERE email = ?", email).Scan(
		&profile.ID, &profile.Email, &profile.Name, &profile.Phone, &profile.Address,
	)

	if err == sql.ErrNoRows {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(profile)
}

func CreateOrUpdateProfile(w http.ResponseWriter, r *http.Request) {
	var profile models.Profile
	_ = json.NewDecoder(r.Body).Decode(&profile)

	_, err := config.DB.Exec(`
		INSERT INTO profiles (email, name, phone, address) 
		VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE name=?, phone=?, address=?
	`, profile.Email, profile.Name, profile.Phone, profile.Address,
		profile.Name, profile.Phone, profile.Address)

	if err != nil {
		http.Error(w, "Insert/Update failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Profile saved"})
}
