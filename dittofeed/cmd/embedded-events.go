package cmd

import "github.com/spf13/cobra"

var embeddedEventsCmd = &cobra.Command{
	Use: "embedded-events",
	Short: "embedded-events",
}

func init() {
	rootCmd.AddCommand(embeddedEventsCmd)
}
