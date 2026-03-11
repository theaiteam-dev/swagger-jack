package cmd

import "github.com/spf13/cobra"

var adminComputedPropertiesCmd = &cobra.Command{
	Use: "admin-computed-properties",
	Short: "admin-computed-properties",
}

func init() {
	rootCmd.AddCommand(adminComputedPropertiesCmd)
}
