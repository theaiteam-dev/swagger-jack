package cmd

import "github.com/spf13/cobra"

var internalApiDebugCmd = &cobra.Command{
	Use: "internal-api-debug",
	Short: "internal-api-debug",
}

func init() {
	rootCmd.AddCommand(internalApiDebugCmd)
}
