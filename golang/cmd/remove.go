/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/byted/totper/internal/keystore"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove ACCOUNT_NAME",
	Short: "Remove a TOTP account",
	RunE:  removeAccount,
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

func removeAccount(cmd *cobra.Command, args []string) error {
	if err := keystore.RemoveSecret(args[0]); err != nil {
		return fmt.Errorf("unable to remove TOTP secret: %s", err)
	}

	fmt.Printf("Removed TOPT for %s\n", args[0])
	return nil
}
