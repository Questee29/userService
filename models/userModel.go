package models

type User struct {
	ID       uint   `json:"id"`
	Phone    string `json:"phone_number"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
