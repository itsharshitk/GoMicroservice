package models

type Profile struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
