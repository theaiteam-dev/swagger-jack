package cmd

import "github.com/spf13/cobra"

var adminUserPropertiesCmd = &cobra.Command{
	Use: "admin-user-properties",
	Short: "admin-user-properties",
}

func init() {
	rootCmd.AddCommand(adminUserPropertiesCmd)
}
