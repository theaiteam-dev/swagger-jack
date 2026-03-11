package cmd

import "github.com/spf13/cobra"

var adminEventsCmd = &cobra.Command{
	Use: "admin-events",
	Short: "admin-events",
}

func init() {
	rootCmd.AddCommand(adminEventsCmd)
}
