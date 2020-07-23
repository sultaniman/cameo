package cameo

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
	"os"
)

const ConfigPath = "/etc/cameo/config.yml"

type Mailer struct {
	Host      string
	Port      *int
	User      string
	Pass      string
	AppPass   string `mapstructure:"app_pass"`
	Retries   int
	SendTo    string `mapstructure:"send_to"`
	FromEmail string `mapstructure:"from_email"`
}

type Domain []string

type Logs struct {
	Level string
}

type GPG struct {
	PubKey string `mapstructure:"pub_key"`
	Entity *openpgp.Entity
}

type Config struct {
	Domains        []Domain
	FormTitle      string `mapstructure:"form_title"`
	Port           int
	Version        string
	Mailer
	Logs
	GPG
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

	logLevel, err := logrus.ParseLevel(viper.GetString("logs.level"))
	if err != nil {
		logLevel = logrus.InfoLevel
	}

	logrus.SetLevel(logLevel)

	config.GPG.Entity = readGPGKey(config.GPG.PubKey)

	return &config
}

func configExists(configPath string) bool {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return false
	}

	return true
}

func readGPGKey(keyPath string) *openpgp.Entity {
	file, err := os.Open(keyPath)
	if err != nil {
		logrus.Error("Unable to read public key", err)
		os.Exit(1)
	}
	defer file.Close()

	block, err := armor.Decode(file)
	if err != nil {
		logrus.Error("Unable to decode public key", err)
		os.Exit(1)
	}

	entity, err := openpgp.ReadEntity(packet.NewReader(block.Body))
	if err != nil {
		logrus.Error("Unable to read openpgp Entity", err)
		os.Exit(1)
	}

	return entity
}
