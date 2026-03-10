package cmd

import "github.com/spf13/cobra"

var internalApiLCmd = &cobra.Command{
	Use: "internal-api-l",
	Short: "internal-api-l",
}

func init() {
	rootCmd.AddCommand(internalApiLCmd)
}
