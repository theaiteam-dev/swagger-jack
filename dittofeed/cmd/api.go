package cmd

import "github.com/spf13/cobra"

var apiCmd = &cobra.Command{
	Use: "api",
	Short: "api",
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
