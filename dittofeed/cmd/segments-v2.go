package cmd

import "github.com/spf13/cobra"

var segmentsV2Cmd = &cobra.Command{
	Use: "segments-v2",
	Short: "segments-v2",
}

func init() {
	rootCmd.AddCommand(segmentsV2Cmd)
}
