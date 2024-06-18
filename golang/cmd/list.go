package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all TOTP accounts",
	RunE:  listAccounts,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listAccounts(cmd *cobra.Command, args []string) error {
	ax := viper.GetStringSlice("accounts")

	fmt.Println("Available accounts:")
	for _, a := range ax {
		fmt.Printf("- %s\n", a)
	}
	return nil
}
