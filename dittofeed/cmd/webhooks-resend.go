package cmd

import "github.com/spf13/cobra"

var webhooksResendCmd = &cobra.Command{
	Use: "webhooks-resend",
	Short: "webhooks-resend",
}

func init() {
	rootCmd.AddCommand(webhooksResendCmd)
}
