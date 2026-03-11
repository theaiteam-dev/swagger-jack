package cmd

import "github.com/spf13/cobra"

var embeddedJourneysCmd = &cobra.Command{
	Use: "embedded-journeys",
	Short: "embedded-journeys",
}

func init() {
	rootCmd.AddCommand(embeddedJourneysCmd)
}
