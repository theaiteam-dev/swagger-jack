package cmd

import "github.com/spf13/cobra"

var publicAppsCmd = &cobra.Command{
	Use: "public-apps",
	Short: "public-apps",
}

func init() {
	rootCmd.AddCommand(publicAppsCmd)
}
