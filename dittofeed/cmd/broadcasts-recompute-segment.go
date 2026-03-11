package cmd

import "github.com/spf13/cobra"

var broadcastsRecomputeSegmentCmd = &cobra.Command{
	Use: "broadcasts-recompute-segment",
	Short: "broadcasts-recompute-segment",
}

func init() {
	rootCmd.AddCommand(broadcastsRecomputeSegmentCmd)
}
