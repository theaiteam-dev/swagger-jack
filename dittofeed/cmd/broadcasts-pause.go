package cmd

import "github.com/spf13/cobra"

var broadcastsPauseCmd = &cobra.Command{
	Use: "broadcasts-pause",
	Short: "broadcasts-pause",
}

func init() {
	rootCmd.AddCommand(broadcastsPauseCmd)
}
