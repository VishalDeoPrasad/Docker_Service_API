package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	// Sender's email address and password
	from := "vishal.prasad2009@gmail.com"
	password := "oxwr hxxl uegs oqyt"

	// Recipient's email address
	to := "tyagivikalp99@gmail.com"

	// SMTP server details
	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	// Message content
	message := []byte("Subject: Test Email\n\nThis is a test email body.")

	// Authentication information
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// SMTP connection
	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	err := smtp.SendMail(smtpAddr, auth, from, []string{to}, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}
