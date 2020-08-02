package app

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/domodwyer/mailyak"
	"github.com/sirupsen/logrus"
	"net/smtp"
	"time"
)

type Sender struct {
	Retries int
}

// Send encrypted email with GPG and retry with exponential backoff
// until retry limit reached will return and error if it failed.
func (m Sender) SendMessage(message *Message, mailer *Mailer) error {
	mail := mailyak.New(
		fmt.Sprintf("%s:%d", mailer.Host, *mailer.Port),
		smtp.PlainAuth("", mailer.User, mailer.Pass, mailer.Host),
	)

	mail.To(*mailer.SendTo)
	mail.From(mailer.FromEmail)
	mail.Subject(message.Subject)
	reader := bytes.NewReader(message.RawBody)
	mail.Plain().Set(encodeMessage(string(message.RawBody)))
	mail.Attach("message.gpg", reader)
	if err := mail.Send(); err != nil {
		if m.Retries <= mailer.Retries {
			wait := m.Retries << 2
			logrus.Info(fmt.Sprintf("Waiting %d seconds before next retry...", wait))
			time.Sleep(time.Duration(wait))
			m.Retries += 1
			return m.SendMessage(message, mailer)
		}

		return err
	}

	return nil
}

func SendAsync(message *Message, config *Config) {
	go func() {
		body, err := message.Encrypt(&config.GPG)
		if err != nil {
			logrus.Error("Unable to encrypt message")
			logrus.Error(err)
		}

		message.RawBody = body
		sender := Sender{Retries: 0}
		if err := sender.SendMessage(message, &config.Mailer); err != nil {
			logrus.Error("Unable to send message")
			logrus.Error(err)
		}
	}()
}

const GpgMessage = `
-----BEGIN PGP MESSAGE-----

%s
-----END PGP MESSAGE-----
`

func encodeMessage(body string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(body))
	return fmt.Sprintf(GpgMessage, encoded)
}
