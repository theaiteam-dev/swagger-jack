package cmd

import "github.com/spf13/cobra"

var userPropertiesStatusCmd = &cobra.Command{
	Use: "user-properties-status",
	Short: "user-properties-status",
}

func init() {
	rootCmd.AddCommand(userPropertiesStatusCmd)
}
