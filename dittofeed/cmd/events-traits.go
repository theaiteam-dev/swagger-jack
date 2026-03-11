package cmd

import "github.com/spf13/cobra"

var eventsTraitsCmd = &cobra.Command{
	Use: "events-traits",
	Short: "events-traits",
}

func init() {
	rootCmd.AddCommand(eventsTraitsCmd)
}
