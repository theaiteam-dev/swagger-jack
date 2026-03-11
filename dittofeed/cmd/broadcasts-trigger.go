package cmd

import "github.com/spf13/cobra"

var broadcastsTriggerCmd = &cobra.Command{
	Use: "broadcasts-trigger",
	Short: "broadcasts-trigger",
}

func init() {
	rootCmd.AddCommand(broadcastsTriggerCmd)
}
