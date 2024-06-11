/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"syscall"

	"github.com/byted/totper/internal/keystore"
	"github.com/byted/totper/internal/totp"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var addCmd = &cobra.Command{
	Use:   "add ACCOUNT_NAME",
	Short: "Add a new TOTP account",
	Long: `
Stores a new secret for generating TOTP under the name ACCOUNT_NAME.
Secret is expected to be Base32 encoded.
	`,
	RunE: addAccount,
	Args: cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addAccount(cmd *cobra.Command, args []string) error {
	fmt.Print("Enter Secret: ")
	secret, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		return fmt.Errorf("unable to read secret: %v", err)
	}
	fmt.Println()

	if err := keystore.StoreSecretIfNotExists(args[0], string(secret)); err != nil {
		return fmt.Errorf("unable to store TOTP secret: %v", err)
	}

	totper, err := totp.NewTOTPer(string(secret))
	if err != nil {
		return fmt.Errorf("unable to create TOTP generator: %v", err)
	}

	fmt.Printf("Added TOPT for %s. First TOTP: %d\n", args[0], totper.GetOTP())
	return nil
}
