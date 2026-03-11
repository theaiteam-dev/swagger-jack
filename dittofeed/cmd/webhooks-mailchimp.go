package cmd

import "github.com/spf13/cobra"

var webhooksMailchimpCmd = &cobra.Command{
	Use: "webhooks-mailchimp",
	Short: "webhooks-mailchimp",
}

func init() {
	rootCmd.AddCommand(webhooksMailchimpCmd)
}
