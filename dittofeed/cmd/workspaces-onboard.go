package cmd

import "github.com/spf13/cobra"

var workspacesOnboardCmd = &cobra.Command{
	Use: "workspaces-onboard",
	Short: "workspaces-onboard",
}

func init() {
	rootCmd.AddCommand(workspacesOnboardCmd)
}
