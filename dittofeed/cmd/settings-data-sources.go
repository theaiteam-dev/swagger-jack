package cmd

import "github.com/spf13/cobra"

var settingsDataSourcesCmd = &cobra.Command{
	Use: "settings-data-sources",
	Short: "settings-data-sources",
}

func init() {
	rootCmd.AddCommand(settingsDataSourcesCmd)
}
