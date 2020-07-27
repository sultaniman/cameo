package cameo

import (
	"strings"
)

type Message struct {
	Subject string
	Body    string
	RawBody []byte
}

func (m *Message) IsValid() bool {
	if len(m.Subject) < 3 {
		return false
	}

	if len(strings.Split(m.Body, " ")) < 2 {
		return false
	}

	return true
}

func (m *Message) Encrypt(gpg *GPG) ([]byte, error) {
	message, err := gpg.Encrypt([]byte(m.Body))
	if err != nil {
		return nil, err
	}

	return message, nil
}
