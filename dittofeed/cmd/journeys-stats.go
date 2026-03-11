package cmd

import "github.com/spf13/cobra"

var journeysStatsCmd = &cobra.Command{
	Use: "journeys-stats",
	Short: "journeys-stats",
}

func init() {
	rootCmd.AddCommand(journeysStatsCmd)
}
