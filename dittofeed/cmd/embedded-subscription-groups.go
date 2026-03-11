package cmd

import "github.com/spf13/cobra"

var embeddedSubscriptionGroupsCmd = &cobra.Command{
	Use: "embedded-subscription-groups",
	Short: "embedded-subscription-groups",
}

func init() {
	rootCmd.AddCommand(embeddedSubscriptionGroupsCmd)
}
