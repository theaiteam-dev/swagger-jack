package cmd

import "github.com/spf13/cobra"

var broadcastsStartCmd = &cobra.Command{
	Use: "broadcasts-start",
	Short: "broadcasts-start",
}

func init() {
	rootCmd.AddCommand(broadcastsStartCmd)
}
