package cmd

import "github.com/spf13/cobra"

var subscriptionGroupsCmd = &cobra.Command{
	Use: "subscription-groups",
	Short: "subscription-groups",
}

func init() {
	rootCmd.AddCommand(subscriptionGroupsCmd)
}
