package cmd

import "github.com/spf13/cobra"

var embeddedUsersCmd = &cobra.Command{
	Use: "embedded-users",
	Short: "embedded-users",
}

func init() {
	rootCmd.AddCommand(embeddedUsersCmd)
}
