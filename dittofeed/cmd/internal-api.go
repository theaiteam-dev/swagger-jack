package cmd

import "github.com/spf13/cobra"

var internalApiCmd = &cobra.Command{
	Use: "internal-api",
	Short: "internal-api",
}

func init() {
	rootCmd.AddCommand(internalApiCmd)
}
