package cmd

import "github.com/spf13/cobra"

var stripeCheckoutSessionCmd = &cobra.Command{
	Use: "stripe-checkout-session",
	Short: "stripe-checkout-session",
}

func init() {
	rootCmd.AddCommand(stripeCheckoutSessionCmd)
}
