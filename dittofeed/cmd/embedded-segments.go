package cmd

import "github.com/spf13/cobra"

var embeddedSegmentsCmd = &cobra.Command{
	Use: "embedded-segments",
	Short: "embedded-segments",
}

func init() {
	rootCmd.AddCommand(embeddedSegmentsCmd)
}
