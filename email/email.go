package email

import (
	"crypto/tls"
	"fmt"
	"net/mail"
	"net/smtp"
	"strings"
)

// SMTP represents an SMTP connection
type SMTP struct {
	Host     string
	User     string
	Port     int
	Password string
	From     mail.Address
	Auth     smtp.Auth
	TLS      *tls.Config
	DSN      string
}

// New returns a new SMTP struct
func New(host string, port int, user, pass, email, name string) *SMTP {
	return &SMTP{
		Host:     host,
		Port:     port,
		User:     user,
		Password: pass,
		From:     mail.Address{Name: name, Address: email},
		Auth:     smtp.PlainAuth("", user, pass, host),
		TLS:      &tls.Config{InsecureSkipVerify: true, ServerName: host},
		DSN:      fmt.Sprintf("%v:%v", host, port),
	}
}

// Send will send an email using the SMTP parameters
func (s SMTP) Send(subject, body string, emails ...string) error {
	to := mail.Address{Name: "", Address: strings.Join(emails, ",")}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = s.From.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	return s.sendTLS(to, message)
}

func (s SMTP) sendTLS(to mail.Address, message string) error {
	// Connect to the SMTP Server

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", s.DSN, s.TLS)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, s.Host)
	if err != nil {
		return err
	}

	// Auth
	if err = c.Auth(s.Auth); err != nil {
		return err
	}

	// To && From
	if err = c.Mail(s.From.Address); err != nil {
		return err
	}

	if err = c.Rcpt(to.Address); err != nil {
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()
	return nil
}
