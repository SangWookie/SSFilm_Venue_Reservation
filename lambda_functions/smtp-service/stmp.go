package main

import (
	"fmt"
	"net/smtp"
	"os"
)

var (
	smtpServer = "smtp.gmail.com"
	smtpPort   = "587"
	username   = os.Getenv("GMAIL_USERNAME") // 환경 변수에서 가져오기
	password   = os.Getenv("GMAIL_PASSWORD") // 앱 비밀번호 사용
)

type EmailRequest struct {
	To      string
	Subject string
	Body    string
}

// EmailSender is a function type for sending emails
type EmailSender func(to, subject, body string) error

type SMTPManagerIFace interface {
	SendEmailWithGoogle(to, subject, body string) error
}

type SMTPManager struct{}

func (r *SMTPManager) SendEmailWithGoogle(to, subject, body string) error {
	auth := smtp.PlainAuth("", username, password, smtpServer)
	log.Infof("SMTP Auth created with username: %s, server: %s", username, smtpServer)

	log.Infof("Preparing to send email")
	log.Infof("To: %s", to)
	log.Infof("Subject: %s", subject)
	log.Infof("Body: %s", body)

	// 이메일 포맷
	msg := []byte(fmt.Sprintf(
		"To: %s\r\n"+
			"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=\"UTF-8\"\r\n"+
			"\r\n"+
			"%s\r\n", to, subject, body,
	))

	log.Infof("Full SMTP message:\n%s", msg)

	// SMTP 서버에 연결하여 이메일 전송
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, username, []string{to}, msg)
	if err != nil {
		return err
	}
	return nil
}

func SendEmail(sender SMTPManagerIFace, to, subject, body string) error {
	return sender.SendEmailWithGoogle(to, subject, body)
}
