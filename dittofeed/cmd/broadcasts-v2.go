package cmd

import "github.com/spf13/cobra"

var broadcastsV2Cmd = &cobra.Command{
	Use: "broadcasts-v2",
	Short: "broadcasts-v2",
}

func init() {
	rootCmd.AddCommand(broadcastsV2Cmd)
}
