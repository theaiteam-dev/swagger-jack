package cmd

import "github.com/spf13/cobra"

var embeddedBroadcastsCmd = &cobra.Command{
	Use: "embedded-broadcasts",
	Short: "embedded-broadcasts",
}

func init() {
	rootCmd.AddCommand(embeddedBroadcastsCmd)
}
