package cameo

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/openpgp"
	"io/ioutil"
	"os"
)

const ConfigPath = "/etc/cameo/config.yml"

type Mailer struct {
	Host      string
	Port      *int
	User      string
	Pass      string
	Retries   int
	SendTo    *string `mapstructure:"send_to"`
	FromEmail string  `mapstructure:"from_email"`
}

type Logs struct {
	Level string
}

type GPG struct {
	PubKey   string `mapstructure:"pub_key"`
	Entities []*openpgp.Entity
}

type Config struct {
	Domains   []string
	Views     string
	FormTitle string `mapstructure:"form_title"`
	Port      int
	Version   string
	Mailer    Mailer
	Logs      Logs
	GPG       GPG
}

func LoadConfig(configPath string) *Config {
	if &configPath == nil {
		fmt.Println("Using default configuration")
		viper.SetConfigFile(ConfigPath)
	} else {
		fmt.Printf("Using provided configuration: %s", configPath)
		viper.SetConfigFile(configPath)
	}

	viper.SetConfigType("yaml")
	viper.SetDefault("mailer.retries", 3)
	viper.SetDefault("logs.level", "INFO")
	viper.SetEnvPrefix("cameo")
	viper.AutomaticEnv()

	if !configExists(configPath) {
		panic(fmt.Sprintf("Config file not found: %s", configPath))
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Unable to read config file: %s", configPath))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	if config.Mailer.SendTo == nil {
		logrus.Error("Please configure mailer.send_to address")
		os.Exit(1)
	}

	logLevel, err := logrus.ParseLevel(viper.GetString("logs.level"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	logrus.SetLevel(logLevel)

	config.GPG.Entities = readGPGKey(config.GPG.PubKey)

	return &config
}

func configExists(configPath string) bool {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return false
	}

	return true
}

func readGPGKey(keyPath string) []*openpgp.Entity {
	pubKey, err := ioutil.ReadFile(keyPath)
	if err != nil {
		logrus.Error("Unable to read public key", err)
		logrus.Error(err)
		os.Exit(1)
	}

	entityList, err := openpgp.ReadArmoredKeyRing(bytes.NewBufferString(string(pubKey)))
	if err != nil {
		logrus.Error("Unable to read openpgp armored key ring")
		logrus.Error(err)
		os.Exit(1)
	}

	var entities []*openpgp.Entity
	for _, entity := range entityList {
		entities = append(entities, entity)
	}

	return entities
}
