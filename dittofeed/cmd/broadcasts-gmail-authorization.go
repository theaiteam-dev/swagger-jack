package cmd

import "github.com/spf13/cobra"

var broadcastsGmailAuthorizationCmd = &cobra.Command{
	Use: "broadcasts-gmail-authorization",
	Short: "broadcasts-gmail-authorization",
}

func init() {
	rootCmd.AddCommand(broadcastsGmailAuthorizationCmd)
}
