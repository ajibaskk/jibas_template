package email

import (
	"fmt"
	"jibas-template/internal/domain"
	"net/smtp"
	"os"
)

type EmailServiceImpl struct {
	smtpHost     string
	smtpPort     string
	smtpUsername string
	smtpPassword string
}

// NewEmailService creates a new instance of EmailServiceImpl
func NewEmailService() domain.EmailService {
	return &EmailServiceImpl{
		smtpHost:     os.Getenv("SMTP_HOST"),
		smtpPort:     os.Getenv("SMTP_PORT"),
		smtpUsername: os.Getenv("SMTP_USERNAME"),
		smtpPassword: os.Getenv("SMTP_PASSWORD"),
	}
}

// SendEmail sends an email to the specified recipient
func (e *EmailServiceImpl) SendEmail(to string, subject string, body string) error {
	// Set up authentication information.
	auth := smtp.PlainAuth("", e.smtpUsername, e.smtpPassword, e.smtpHost)

	// Format email content
	msg := "From: " + e.smtpUsername + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Send email
	addr := fmt.Sprintf("%s:%s", e.smtpHost, e.smtpPort)
	err := smtp.SendMail(addr, auth, e.smtpUsername, []string{to}, []byte(msg))
	if err != nil {
		fmt.Printf("Failed to send email: %v\n", err)
		return err
	}

	fmt.Printf("Email successfully sent to %s with subject: %s\n", to, subject)
	return nil
}
