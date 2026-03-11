package cmd

import "github.com/spf13/cobra"

var stripeBillingPortalCmd = &cobra.Command{
	Use: "stripe-billing-portal",
	Short: "stripe-billing-portal",
}

func init() {
	rootCmd.AddCommand(stripeBillingPortalCmd)
}
