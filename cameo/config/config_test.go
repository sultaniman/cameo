package config

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	//LoadConfig()
	configFile, _ := filepath.Abs("./../../etc/cameo/config.yml")
	config := LoadConfig(configFile, false)
	assert.Equal(t, 4000, config.Port )
	assert.Equal(t, "0.0.1", config.Version )
	assert.Equal(t, "A Message To Unicorn", config.FormTitle )
	assert.Equal(t, "/path/to/pub.key", config.GPG.PubKey)
	assert.Equal(t, "INFO", config.Logs.Level)
	assert.Equal(t, []string{"example.com", "second-example.com"}, config.Domains)
	assert.Equal(t, "smtp.google.com", config.Mailer.Host)
	assert.Equal(t, 587, *config.Mailer.Port)
	assert.Equal(t, "user", config.Mailer.User)
	assert.Equal(t, "pass", config.Mailer.Pass)
	assert.Equal(t, 3, config.Mailer.Retries)
	assert.Equal(t, "secret@secret.mail", *config.Mailer.SendTo)
	assert.Equal(t, "contact@myform.com", config.Mailer.FromEmail)

}
