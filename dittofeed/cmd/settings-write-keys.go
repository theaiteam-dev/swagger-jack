package cmd

import "github.com/spf13/cobra"

var settingsWriteKeysCmd = &cobra.Command{
	Use: "settings-write-keys",
	Short: "settings-write-keys",
}

func init() {
	rootCmd.AddCommand(settingsWriteKeysCmd)
}
