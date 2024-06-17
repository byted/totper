/*
Copyright © 2024 Stefan Selent
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFilePath string
)

var rootCmd = &cobra.Command{
	Use:   "totper",
	Short: "A simple TOTP authenticator",
	Long:  `A tool to manage TOTP secrets. Secrets are stored in your systems secrets storage.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnFinalize(func() {
		cobra.CheckErr(viper.WriteConfig())
	})

	rootCmd.PersistentFlags().StringVar(&configFilePath, "config", "", "config file (default is $HOME/.config/totper.yaml)")
	viper.SetConfigName("totper")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config")
}

func initConfig() {
	if configFilePath != "" {
		viper.SetConfigFile(configFilePath)
		cobra.CheckErr(viper.ReadInConfig())
		fmt.Println("Using custom config file:", viper.ConfigFileUsed())
	} else {
		homeDir, err := os.UserHomeDir()
		cobra.CheckErr(err)
		defaultPath := path.Join(homeDir, ".config")

		viper.AddConfigPath(defaultPath)
		viper.SetConfigType("yaml")
		viper.SetConfigName("totper")

		err = viper.ReadInConfig()
		notFoundErr := viper.ConfigFileNotFoundError{}

		if errors.As(err, &notFoundErr) {
			cobra.CheckErr(viper.SafeWriteConfig())
		}
	}
}
