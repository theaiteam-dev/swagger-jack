package cmd

import "github.com/spf13/cobra"

var webhooksTwilioCmd = &cobra.Command{
	Use: "webhooks-twilio",
	Short: "webhooks-twilio",
}

func init() {
	rootCmd.AddCommand(webhooksTwilioCmd)
}
