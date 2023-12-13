package mail

import (
	"crypto/tls"
	"strconv"

	"gopkg.in/gomail.v2"
)

type MailService struct {
	password string
	from     string
	server   string
	port     string
}

type MailServiceParams struct {
	MailService
}

func (m *MailService) Send(to string, subject string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", m.from)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", "test.")
	port, _ := strconv.Atoi(m.port)
	d := gomail.NewDialer(m.server, port, m.from, m.password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(mail); err != nil {
		return err
	}
	return nil
}

func NewMailService(params MailServiceParams) *MailService {
	return &MailService{
		password: params.password,
		from:     params.from,
		server:   params.server,
		port:     params.port,
	}
}
