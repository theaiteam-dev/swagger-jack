package cmd

import "github.com/spf13/cobra"

var usersCountCmd = &cobra.Command{
	Use: "users-count",
	Short: "users-count",
}

func init() {
	rootCmd.AddCommand(usersCountCmd)
}
