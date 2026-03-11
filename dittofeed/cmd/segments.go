package cmd

import "github.com/spf13/cobra"

var segmentsCmd = &cobra.Command{
	Use: "segments",
	Short: "segments",
}

func init() {
	rootCmd.AddCommand(segmentsCmd)
}
