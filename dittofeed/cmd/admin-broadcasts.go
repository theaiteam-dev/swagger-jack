package cmd

import "github.com/spf13/cobra"

var adminBroadcastsCmd = &cobra.Command{
	Use: "admin-broadcasts",
	Short: "admin-broadcasts",
}

func init() {
	rootCmd.AddCommand(adminBroadcastsCmd)
}
