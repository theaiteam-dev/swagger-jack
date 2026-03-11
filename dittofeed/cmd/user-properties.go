package cmd

import "github.com/spf13/cobra"

var userPropertiesCmd = &cobra.Command{
	Use: "user-properties",
	Short: "user-properties",
}

func init() {
	rootCmd.AddCommand(userPropertiesCmd)
}
