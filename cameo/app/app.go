package app

import (
	"golang.org/x/crypto/openpgp"
)

type Config struct {
	Domains   []string
	FormTitle string `mapstructure:"form_title"`
	Port      int
	Version   string
	Mailer    Mailer
	Logs      Logs
	GPG       GPG
}

type Mailer struct {
	Host      string
	Port      *int
	User      string
	Pass      string
	Retries   int
	SendTo    *string `mapstructure:"send_to"`
	FromEmail string  `mapstructure:"from_email"`
}

type GPG struct {
	PubKey   string `mapstructure:"pub_key"`
	Entities []*openpgp.Entity
}

type Logs struct {
	Level string
}
