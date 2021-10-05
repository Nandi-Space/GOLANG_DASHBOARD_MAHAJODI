package models

type Admin struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Phone    string `json:"phone"`
}
