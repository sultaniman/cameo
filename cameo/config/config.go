package config

import (
	"bytes"
	"fmt"
	"github.com/imanhodjaev/cameo/cameo/app"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/openpgp"
	"io/ioutil"
	"os"
)

const ConfigPath = "/etc/cameo/config.yml"

const PORT = 4000

func LoadConfig(configPath string, loadGPG bool) *app.Config {
	if &configPath == nil {
		fmt.Println("Using default configuration")
		viper.SetConfigFile(ConfigPath)
	} else {
		fmt.Printf("Using provided configuration: %s", configPath)
		viper.SetConfigFile(configPath)
	}

	viper.SetConfigType("yaml")
	viper.SetDefault("port", PORT)
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

	var config app.Config
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

	if loadGPG {
		config.GPG.Entities = readGPGKey(config.GPG.PubKey)
	}

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
