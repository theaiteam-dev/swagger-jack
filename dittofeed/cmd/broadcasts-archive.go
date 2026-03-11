package cmd

import "github.com/spf13/cobra"

var broadcastsArchiveCmd = &cobra.Command{
	Use: "broadcasts-archive",
	Short: "broadcasts-archive",
}

func init() {
	rootCmd.AddCommand(broadcastsArchiveCmd)
}
