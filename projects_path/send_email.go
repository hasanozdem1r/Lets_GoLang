package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	// Sender data.
	from := "<sender_address>"
	password := "<sender_password>"

	// Receiver email address.
	to := []string{
		"<receiver_address>",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
