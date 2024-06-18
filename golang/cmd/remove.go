package cmd

import (
	"fmt"
	"slices"

	"github.com/byted/totper/internal/keystore"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
	accountName := args[0]
	ax := viper.GetStringSlice("accounts")
	lenBefore := len(ax)
	ax = slices.DeleteFunc(ax, func(a string) bool { return a == accountName })
	viper.Set("accounts", ax)

	if lenBefore == len(ax) {
		return fmt.Errorf("TOTP account %s does not exist", accountName)
	}

	if err := keystore.RemoveSecret(accountName); err != nil {
		return fmt.Errorf("unable to remove TOTP secret: %s", err)
	}

	fmt.Printf("Removed TOPT for %s\n", accountName)
	return nil
}
