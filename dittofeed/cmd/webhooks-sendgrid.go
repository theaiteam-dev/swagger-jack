package cmd

import "github.com/spf13/cobra"

var webhooksSendgridCmd = &cobra.Command{
	Use: "webhooks-sendgrid",
	Short: "webhooks-sendgrid",
}

func init() {
	rootCmd.AddCommand(webhooksSendgridCmd)
}
