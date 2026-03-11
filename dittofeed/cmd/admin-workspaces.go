package cmd

import "github.com/spf13/cobra"

var adminWorkspacesCmd = &cobra.Command{
	Use: "admin-workspaces",
	Short: "admin-workspaces",
}

func init() {
	rootCmd.AddCommand(adminWorkspacesCmd)
}
