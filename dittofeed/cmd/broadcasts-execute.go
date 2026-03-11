package cmd

import "github.com/spf13/cobra"

var broadcastsExecuteCmd = &cobra.Command{
	Use: "broadcasts-execute",
	Short: "broadcasts-execute",
}

func init() {
	rootCmd.AddCommand(broadcastsExecuteCmd)
}
