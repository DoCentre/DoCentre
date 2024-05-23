package services

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/docentre/docentre/models"
	"github.com/docentre/docentre/repositories"
)

type Config struct {
	Host         string
	Port         string
	AuthEmail    string
	AuthPassword string
	SenderEmail  string
	DocUrlTmpl   string
}

func loadConfig() (Config, error) {
	var config Config
	fields := []*string{&config.Host, &config.Port, &config.AuthEmail, &config.AuthPassword, &config.SenderEmail, &config.DocUrlTmpl}
	fieldNames := []string{"SMTP_HOST", "SMTP_PORT", "AUTH_EMAIL", "AUTH_PASSWORD", "SENDER_EMAIL", "DOC_URL_TMPL"}
	for i, field := range fields {
		var ext bool
		*field, ext = os.LookupEnv(fieldNames[i])
		if !ext {
			return Config{}, fmt.Errorf("missing environment variable %s", fieldNames[i])
		}
	}
	return config, nil
}

func getEmailOf(userID uint) (string, error) {
	var user models.User
	if err := repositories.DB.First(&user, userID).Error; err != nil {
		return "", err
	}
	return user.Email, nil
}

func sendEmailToApprover(documentID uint, approverID uint) error {
	config, err := loadConfig()
	if err != nil {
		return err
	}

	// Prepare the necessary fields for sending an email.
	rcpt, err := getEmailOf(approverID)
	if err != nil {
		return err
	}
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	auth := smtp.PlainAuth("", config.AuthEmail, config.AuthPassword, config.Host)
	subject := "[DoCentre] Document Approval Request"
	header := fmt.Sprintf(
		"To: %s\r\nFrom: %s\r\nSubject: %s\r\n\r\n",
		rcpt, config.SenderEmail, subject)
	contentTmpl := "You're assigned as an approver of the document %s, please review it."
	url := fmt.Sprintf(config.DocUrlTmpl, documentID)
	content := fmt.Sprintf(contentTmpl, url)
	// Send the email.
	if err := smtp.SendMail(addr, auth, config.SenderEmail, []string{rcpt}, []byte(header+content)); err != nil {
		return err
	}
	return nil
}
