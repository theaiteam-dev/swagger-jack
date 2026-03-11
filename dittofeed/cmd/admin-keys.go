package cmd

import "github.com/spf13/cobra"

var adminKeysCmd = &cobra.Command{
	Use: "admin-keys",
	Short: "admin-keys",
}

func init() {
	rootCmd.AddCommand(adminKeysCmd)
}
