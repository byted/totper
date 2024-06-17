/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"

	"github.com/byted/totper/internal/keystore"
	"github.com/byted/totper/internal/totp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getOtpCmd represents the getOtp command
var getOtpCmd = &cobra.Command{
	Use:   "get-totp ACCOUNT_NAME",
	Short: "Gets the current TOTP for the account",
	RunE:  getTotp,
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(getOtpCmd)
}

func getTotp(cmd *cobra.Command, args []string) error {
	accountName := args[0]

	ax := viper.GetStringSlice("accounts")
	if !slices.Contains(ax, accountName) {
		return fmt.Errorf("TOTP account %s does not exist", accountName)
	}

	secret, err := keystore.RetrieveSecret(accountName)
	if err != nil {
		return fmt.Errorf("unable to generate TOTP: %v", err)
	}

	totper, err := totp.NewTOTPer(string(secret))
	if err != nil {
		return fmt.Errorf("unable to generate TOTP: %v", err)
	}

	fmt.Printf("TOTP: %d\n", totper.GetOTP())
	return nil
}
