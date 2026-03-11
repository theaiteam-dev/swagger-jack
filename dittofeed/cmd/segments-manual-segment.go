package cmd

import "github.com/spf13/cobra"

var segmentsManualSegmentCmd = &cobra.Command{
	Use: "segments-manual-segment",
	Short: "segments-manual-segment",
}

func init() {
	rootCmd.AddCommand(segmentsManualSegmentCmd)
}
