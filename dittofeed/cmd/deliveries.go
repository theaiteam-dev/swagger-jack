package cmd

import "github.com/spf13/cobra"

var deliveriesCmd = &cobra.Command{
	Use: "deliveries",
	Short: "deliveries",
}

func init() {
	rootCmd.AddCommand(deliveriesCmd)
}
