package lib

import (
	"crypto/tls"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	host := os.Getenv("DEV_EMAIL_HOST")
	port, _ := strconv.Atoi(os.Getenv("DEV_EMAIL_PORT"))
	user := os.Getenv("DEV_EMAIL_USER")
	password := os.Getenv("DEV_EMAIL_PASS")

	d := gomail.NewDialer(
		host,
		port,
		user,
		password,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", "Ikan Cupang <7V4t9@example.com>")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return d.DialAndSend(m)
}
