package cmd

import "github.com/spf13/cobra"

var usersSubscriptionsCmd = &cobra.Command{
	Use: "users-subscriptions",
	Short: "users-subscriptions",
}

func init() {
	rootCmd.AddCommand(usersSubscriptionsCmd)
}
