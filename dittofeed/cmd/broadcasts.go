package cmd

import "github.com/spf13/cobra"

var broadcastsCmd = &cobra.Command{
	Use: "broadcasts",
	Short: "broadcasts",
}

func init() {
	rootCmd.AddCommand(broadcastsCmd)
}
