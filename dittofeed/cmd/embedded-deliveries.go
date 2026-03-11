package cmd

import "github.com/spf13/cobra"

var embeddedDeliveriesCmd = &cobra.Command{
	Use: "embedded-deliveries",
	Short: "embedded-deliveries",
}

func init() {
	rootCmd.AddCommand(embeddedDeliveriesCmd)
}
