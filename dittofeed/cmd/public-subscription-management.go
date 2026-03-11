package cmd

import "github.com/spf13/cobra"

var publicSubscriptionManagementCmd = &cobra.Command{
	Use: "public-subscription-management",
	Short: "public-subscription-management",
}

func init() {
	rootCmd.AddCommand(publicSubscriptionManagementCmd)
}
