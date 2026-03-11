package cmd

import "github.com/spf13/cobra"

var adminUsersCmd = &cobra.Command{
	Use: "admin-users",
	Short: "admin-users",
}

func init() {
	rootCmd.AddCommand(adminUsersCmd)
}
