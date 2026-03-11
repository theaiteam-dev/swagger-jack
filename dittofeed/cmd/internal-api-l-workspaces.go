package cmd

import "github.com/spf13/cobra"

var internalApiLWorkspacesCmd = &cobra.Command{
	Use: "internal-api-l-workspaces",
	Short: "internal-api-l-workspaces",
}

func init() {
	rootCmd.AddCommand(internalApiLWorkspacesCmd)
}
