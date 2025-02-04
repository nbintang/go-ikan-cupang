package lib

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {

	d := gomail.NewDialer(
		"sandbox.smtp.mailtrap.io",
		2525,
		"34764a45f991b6",
		"4cda87c97faab2",
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", "Ikan Cupang <7V4t9@example.com>")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return d.DialAndSend(m)
}
