package mail

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/smtp"
	"net/textproto"
	"strings"
)

type Attachment struct {
	Name        string
	ContentType string
	Content     []byte
}

type EmailClient struct {
	url  string
	auth smtp.Auth
	from string
	to   []string
}

func New(host string, port string, username string, password string, from string, to []string) *EmailClient {
	return &EmailClient{
		url:  host + ":" + port,
		auth: smtp.PlainAuth("", username, password, host),
		from: from,
		to:   to,
	}
}

func (c EmailClient) Send(body string, attachment Attachment) error {
	// Set up the message
	message := &bytes.Buffer{}
	writer := multipart.NewWriter(message)

	// Add the email body
	part, err := writer.CreatePart(textproto.MIMEHeader{
		"Content-Type": []string{"text/html; charset=utf-8"},
	})
	if err != nil {
		return err
	}
	_, err = part.Write([]byte(body))
	if err != nil {
		return err
	}

	// Add the attachment
	part, err = writer.CreatePart(textproto.MIMEHeader{
		"Content-Type":        []string{attachment.ContentType},
		"Content-Disposition": []string{fmt.Sprintf(`attachment; filename="%s"`, attachment.Name)},
	})
	if err != nil {
		return err
	}
	_, err = io.Copy(part, bytes.NewReader(attachment.Content))
	if err != nil {
		return err
	}

	// Finalize the message
	err = writer.Close()
	if err != nil {
		return err
	}

	// Set up the email message
	email := []byte("From: " + c.from + "\r\n" +
		"To: " + strings.Join(c.to, ",") + "\r\n" +
		"Subject: Stealcam Drop\r\n" +
		"MIME-version: 1.0\r\n" +
		"Content-Type: multipart/mixed; boundary=" + writer.Boundary() + "\r\n" +
		"\r\n" +
		message.String())

	// Send the email
	return smtp.SendMail(c.url, c.auth, c.from, c.to, email)
}
