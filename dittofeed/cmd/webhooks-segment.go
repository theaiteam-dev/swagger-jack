package cmd

import "github.com/spf13/cobra"

var webhooksSegmentCmd = &cobra.Command{
	Use: "webhooks-segment",
	Short: "webhooks-segment",
}

func init() {
	rootCmd.AddCommand(webhooksSegmentCmd)
}
