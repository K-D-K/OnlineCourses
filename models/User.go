package models

// User Struct
type User struct {
	Model
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Role    int    `json:"role"`
}
