package cmd

import "github.com/spf13/cobra"

var analysisSummaryCmd = &cobra.Command{
	Use: "analysis-summary",
	Short: "analysis-summary",
}

func init() {
	rootCmd.AddCommand(analysisSummaryCmd)
}
