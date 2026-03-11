package cmd

import "github.com/spf13/cobra"

var journeysCmd = &cobra.Command{
	Use: "journeys",
	Short: "journeys",
}

func init() {
	rootCmd.AddCommand(journeysCmd)
}
