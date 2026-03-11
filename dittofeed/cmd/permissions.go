package cmd

import "github.com/spf13/cobra"

var permissionsCmd = &cobra.Command{
	Use: "permissions",
	Short: "permissions",
}

func init() {
	rootCmd.AddCommand(permissionsCmd)
}
