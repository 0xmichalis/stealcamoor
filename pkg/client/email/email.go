package mail

import (
	"net/smtp"
)

type EmailClient struct {
	url  string
	auth smtp.Auth
	from string
}

func New(host string, port string, username string, password string, from string) *EmailClient {
	return &EmailClient{
		url:  host + ":" + port,
		auth: smtp.PlainAuth("", username, password, host),
		from: from,
	}
}

func (c EmailClient) Send(to []string, message string) error {
	return smtp.SendMail(c.url, c.auth, c.from, to, []byte(message))
}
