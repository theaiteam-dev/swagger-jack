package cmd

import "github.com/spf13/cobra"

var webhooksAmazonSesCmd = &cobra.Command{
	Use: "webhooks-amazon-ses",
	Short: "webhooks-amazon-ses",
}

func init() {
	rootCmd.AddCommand(webhooksAmazonSesCmd)
}
