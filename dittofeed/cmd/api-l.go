package cmd

import "github.com/spf13/cobra"

var apiLCmd = &cobra.Command{
	Use: "api-l",
	Short: "api-l",
}

func init() {
	rootCmd.AddCommand(apiLCmd)
}
