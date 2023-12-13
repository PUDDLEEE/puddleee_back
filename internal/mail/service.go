package mail

import (
	"crypto/tls"
	"strconv"

	"gopkg.in/gomail.v2"
)

type MailService struct {
	Password string
	From     string
	Server   string
	Port     string
}

type MailServiceParams struct {
	MailService
}

func (m *MailService) Send(to string, subject string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", m.From)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", "test.")
	port, _ := strconv.Atoi(m.Port)
	d := gomail.NewDialer(m.Server, port, m.From, m.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}

func NewMailService(params MailServiceParams) *MailService {
	return &MailService{
		Password: params.Password,
		From:     params.From,
		Server:   params.Server,
		Port:     params.Port,
	}
}
