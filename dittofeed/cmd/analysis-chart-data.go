package cmd

import "github.com/spf13/cobra"

var analysisChartDataCmd = &cobra.Command{
	Use: "analysis-chart-data",
	Short: "analysis-chart-data",
}

func init() {
	rootCmd.AddCommand(analysisChartDataCmd)
}
