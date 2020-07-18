package cameo

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

const ConfigFile = "/etc/cameo/config.yml"

type Mailer struct {
	Host      string
	Port      *int
	User      string
	Pass      string
	AppPass   string
	Retries   int
	SendTo    string
	FromEmail string
}

type Domain struct {
	Address string
}

type Log struct {
	Level string
}

type GPG struct {
	PubKey string
}

type Config struct {
	Version string
	Mailer
	Log
	GPG
}

func LoadConfig() {
	viper.SetConfigFile(ConfigFile)
	viper.SetConfigType("yaml")

	if !configExists() {
		panic(fmt.Sprintf("Config file not found: %s", ConfigFile))
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Unable to read config file: %s", ConfigFile))
	}
}

func configExists() bool {
	if _, err := os.Stat(ConfigFile); os.IsExist(err) {
		return true
	}

	return false
}
