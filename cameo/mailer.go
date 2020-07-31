package cameo

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/domodwyer/mailyak"
	"github.com/imanhodjaev/cameo/cameo/config"
	"github.com/sirupsen/logrus"
	"net/smtp"
	"time"
)

// Send encrypted email with GPG and retry with exponential backoff
// until retry limit reached will return and error if it failed.
func (m *config.Mailer) SendMessage(message *Message, gpg *config.GPG, retries int) error {
	mail := mailyak.New(
		fmt.Sprintf("%s:%d", m.Host, *m.Port),
		smtp.PlainAuth("", m.User, m.Pass, m.Host),
	)

	mail.To(*m.SendTo)
	mail.From(m.FromEmail)
	mail.Subject(message.Subject)
	reader := bytes.NewReader(message.RawBody)
	mail.Plain().Set(encodeMessage(string(message.RawBody)))
	mail.Attach("message.gpg", reader)
	if err := mail.Send(); err != nil {
		if retries <= m.Retries {
			wait := retries << 2
			logrus.Info(fmt.Sprintf("Waiting %d seconds before next retry...", wait))
			time.Sleep(time.Duration(wait))
			return m.SendMessage(message, gpg, retries+1)
		}

		return err
	}

	return nil
}

func (m *Mailer) SendAsync(message *Message, gpg *GPG) {
	go func() {
		body, err := message.Encrypt(gpg)
		if err != nil {
			logrus.Error("Unable to encrypt message")
			logrus.Error(err)
		}

		message.RawBody = body
		if err := m.SendMessage(message, gpg, 0); err != nil {
			logrus.Error("Unable to send message")
			logrus.Error(err)
		}
	}()
}

const GPG_MESSAGE = `
-----BEGIN PGP MESSAGE-----

%s
-----END PGP MESSAGE-----
`

func encodeMessage(body string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(body))
	return fmt.Sprintf(GPG_MESSAGE, encoded)
}
