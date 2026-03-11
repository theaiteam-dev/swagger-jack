package cmd

import "github.com/spf13/cobra"

var adminDeliveriesCmd = &cobra.Command{
	Use: "admin-deliveries",
	Short: "admin-deliveries",
}

func init() {
	rootCmd.AddCommand(adminDeliveriesCmd)
}
