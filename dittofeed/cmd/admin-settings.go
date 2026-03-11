package cmd

import "github.com/spf13/cobra"

var adminSettingsCmd = &cobra.Command{
	Use: "admin-settings",
	Short: "admin-settings",
}

func init() {
	rootCmd.AddCommand(adminSettingsCmd)
}
