package cmd

import "github.com/spf13/cobra"

var broadcastsResumeCmd = &cobra.Command{
	Use: "broadcasts-resume",
	Short: "broadcasts-resume",
}

func init() {
	rootCmd.AddCommand(broadcastsResumeCmd)
}
