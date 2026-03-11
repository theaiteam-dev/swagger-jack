package cmd

import "github.com/spf13/cobra"

var publicWebhooksCmd = &cobra.Command{
	Use: "public-webhooks",
	Short: "public-webhooks",
}

func init() {
	rootCmd.AddCommand(publicWebhooksCmd)
}
