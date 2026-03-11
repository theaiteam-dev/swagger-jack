package cmd

import "github.com/spf13/cobra"

var journeysV2Cmd = &cobra.Command{
	Use: "journeys-v2",
	Short: "journeys-v2",
}

func init() {
	rootCmd.AddCommand(journeysV2Cmd)
}
