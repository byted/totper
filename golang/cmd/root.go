/*
Copyright © 2024 Stefan Selent
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
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

func init() {}
