package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	// Gmail SMTP server settings
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"

	// Your Gmail account credentials
	from := "rahulp229@gmail.com"
	password := "R@hul1991"

	// Recipient's email address
	to := "rahupneo@gmail.com"

	// Email subject and body
	subject := "Test email from Go"
	body := "Hello from Go!"

	// Set up authentication information
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// Set up TLS configuration
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer,
	}

	// Connect to SMTP server
	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%s", smtpServer, smtpPort), tlsConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Create a new SMTP client
	client, err := smtp.NewClient(conn, smtpServer)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Quit()

	// Authenticate with the SMTP server
	if err := client.Auth(auth); err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient
	if err := client.Mail(from); err != nil {
		log.Fatal(err)
	}
	if err := client.Rcpt(to); err != nil {
		log.Fatal(err)
	}

	// Send the email body
	w, err := client.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write([]byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)))
	if err != nil {
		log.Fatal(err)
	}
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully!")
}
