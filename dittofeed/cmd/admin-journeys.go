package cmd

import "github.com/spf13/cobra"

var adminJourneysCmd = &cobra.Command{
	Use: "admin-journeys",
	Short: "admin-journeys",
}

func init() {
	rootCmd.AddCommand(adminJourneysCmd)
}
