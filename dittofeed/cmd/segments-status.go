package cmd

import "github.com/spf13/cobra"

var segmentsStatusCmd = &cobra.Command{
	Use: "segments-status",
	Short: "segments-status",
}

func init() {
	rootCmd.AddCommand(segmentsStatusCmd)
}
