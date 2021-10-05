package models

type Admin struct {
	ID         int64  `json:"id"`
	UserName   string `json:"username"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	Phone     string `json:"phone"`
	OTP       string    `json:"otp"`
}
