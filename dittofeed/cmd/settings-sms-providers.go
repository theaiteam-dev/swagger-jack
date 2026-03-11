package cmd

import "github.com/spf13/cobra"

var settingsSmsProvidersCmd = &cobra.Command{
	Use: "settings-sms-providers",
	Short: "settings-sms-providers",
}

func init() {
	rootCmd.AddCommand(settingsSmsProvidersCmd)
}
