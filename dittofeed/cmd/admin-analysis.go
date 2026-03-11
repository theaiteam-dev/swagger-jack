package cmd

import "github.com/spf13/cobra"

var adminAnalysisCmd = &cobra.Command{
	Use: "admin-analysis",
	Short: "admin-analysis",
}

func init() {
	rootCmd.AddCommand(adminAnalysisCmd)
}
