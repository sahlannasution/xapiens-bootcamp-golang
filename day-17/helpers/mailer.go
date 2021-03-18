package helpers

import (
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func RegisterMailer(email, message string) {
	to := []string{email}
	// cc := []string{"emanuelpras4@gmail.com"}

	subject := "Registration Success!"
	err := SendEmailConfig(to, subject, message)
	if err != nil { // kalau ada error saat send email
		log.Println(err.Error())
	}

	log.Println("Mail Sent!")
}

func SendEmailConfig(to []string, subject, message string) error {
	// configuration for sending emails
	// smptp_hostname, email, password, smptp_port

	err := godotenv.Load(".env")

	if err == nil { // if there's no error finding env file.

		body := "From: " + os.Getenv("MAILER_EMAIL") + "\n" +
			"To: " + strings.Join(to, ",") + "\n" +
			"Subject: " + subject + "\n\n" +
			message

		auth := smtp.PlainAuth("", os.Getenv("MAILER_EMAIL"), os.Getenv("MAILER_PASSWORD"), os.Getenv("MAILER_HOST"))
		smtpAdd := os.Getenv("MAILER_HOST") + ":" + os.Getenv("MAILER_PORT")
		err := smtp.SendMail(smtpAdd, auth, os.Getenv("MAILER_EMAIL"), to, []byte(body))

		if err == nil {
			return nil
		}
		return err
	}
	return err
}
