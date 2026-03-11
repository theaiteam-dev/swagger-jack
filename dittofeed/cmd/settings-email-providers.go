package cmd

import "github.com/spf13/cobra"

var settingsEmailProvidersCmd = &cobra.Command{
	Use: "settings-email-providers",
	Short: "settings-email-providers",
}

func init() {
	rootCmd.AddCommand(settingsEmailProvidersCmd)
}
