package cmd

import "github.com/spf13/cobra"

var usersCmd = &cobra.Command{
	Use: "users",
	Short: "users",
}

func init() {
	rootCmd.AddCommand(usersCmd)
}
