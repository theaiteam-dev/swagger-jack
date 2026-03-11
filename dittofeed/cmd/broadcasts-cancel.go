package cmd

import "github.com/spf13/cobra"

var broadcastsCancelCmd = &cobra.Command{
	Use: "broadcasts-cancel",
	Short: "broadcasts-cancel",
}

func init() {
	rootCmd.AddCommand(broadcastsCancelCmd)
}
