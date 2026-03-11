package cmd

import "github.com/spf13/cobra"

var adminSubscriptionGroupsCmd = &cobra.Command{
	Use: "admin-subscription-groups",
	Short: "admin-subscription-groups",
}

func init() {
	rootCmd.AddCommand(adminSubscriptionGroupsCmd)
}
