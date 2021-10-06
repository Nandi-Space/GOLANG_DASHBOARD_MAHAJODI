package store

import (
	"fmt"
	"log"
	"net/smtp"
)

func (state *State) SendOTP(email, name, otpCode string) error {
	from := state.Config.Mail.Username
	pass := state.Config.Mail.Password

	to := email

	msg := "From: MahaJodi < namaste@mahajodi.space >\r\n" +
		"To:" + email + "\r\n" +
		"Subject: One Time Password \r\n\r\n" +
		"Welcome To the MahaJodi\r\n" +
		"Dear, " + name + "\r\n" +
		"Your Login OTP is : " + otpCode + "\r\n\r\n" +
		"Please DO NOT SHARE your otp with anyone" + "\n\n" +

		"Warm regards,\r\n" +
		"MahaJodi"

	err := smtp.SendMail(state.Config.Mail.Host+":"+state.Config.Mail.Port,
		smtp.PlainAuth("", from, pass, state.Config.Mail.Host),
		from, []string{to}, []byte(msg))

	// handling the errors
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)

	}
	return nil
}
