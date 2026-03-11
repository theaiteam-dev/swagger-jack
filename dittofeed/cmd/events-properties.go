package cmd

import "github.com/spf13/cobra"

var eventsPropertiesCmd = &cobra.Command{
	Use: "events-properties",
	Short: "events-properties",
}

func init() {
	rootCmd.AddCommand(eventsPropertiesCmd)
}
