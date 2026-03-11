package cmd

import "github.com/spf13/cobra"

var analysisJourneyStatsCmd = &cobra.Command{
	Use: "analysis-journey-stats",
	Short: "analysis-journey-stats",
}

func init() {
	rootCmd.AddCommand(analysisJourneyStatsCmd)
}
