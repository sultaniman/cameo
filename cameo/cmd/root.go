package cmd

import (
	"fmt"
	"github.com/imanhodjaev/cameo/cameo"
	"github.com/imanhodjaev/cameo/cameo/config"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Short: "Cameo allows you to create GPG encrypted contact forms",
	Long:  "Generate your GPG key and deploy your contact form with Cameo",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cameo.Serve(appConfig); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var (
	configFile string
	gpgPubKey  string
	port       int
	appConfig  *cameo.App
)

const defaultConfigPath = "/etc/cameo/config.yml"

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(prepareConfig)
	rootCmd.Flags().StringVar(&configFile, "config", defaultConfigPath, "Config file")
	rootCmd.Flags().StringVar(&gpgPubKey, "pub-key", "", "GPG public key")
	rootCmd.Flags().IntVar(&port, "port", 4000, "Server port")
}

func prepareConfig() {
	conf := config.LoadConfig(configFile, true)
	conf.Port = port

	// If GPG key set then replace
	if strings.TrimSpace(gpgPubKey) != "" {
		conf.GPG.PubKey = gpgPubKey
	}

	appConfig = &cameo.App{
		Path:   configFile,
		Config: conf,
	}
}
