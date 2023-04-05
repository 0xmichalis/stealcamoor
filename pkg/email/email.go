package mail

import (
	"net/smtp"
)

type EmailClient struct {
	url  string
	auth smtp.Auth
	from string
}

func New(password string, from string) *EmailClient {
	smtpHost := "live.smtp.mailtrap.io"
	auth := smtp.PlainAuth("", "api", password, smtpHost)

	return &EmailClient{
		url:  smtpHost + ":25",
		auth: auth,
		from: from,
	}
}

func (c EmailClient) Send(to []string, message string) error {
	return smtp.SendMail(c.url, c.auth, c.from, to, []byte(message))
}
