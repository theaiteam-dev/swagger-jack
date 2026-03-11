package cmd

import "github.com/spf13/cobra"

var embeddedAnalysisCmd = &cobra.Command{
	Use: "embedded-analysis",
	Short: "embedded-analysis",
}

func init() {
	rootCmd.AddCommand(embeddedAnalysisCmd)
}
