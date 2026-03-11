package cmd

import "github.com/spf13/cobra"

var deliveriesCountCmd = &cobra.Command{
	Use: "deliveries-count",
	Short: "deliveries-count",
}

func init() {
	rootCmd.AddCommand(deliveriesCountCmd)
}
