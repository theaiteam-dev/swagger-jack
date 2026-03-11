package cmd

import "github.com/spf13/cobra"

var adminComponentConfigurationsCmd = &cobra.Command{
	Use: "admin-component-configurations",
	Short: "admin-component-configurations",
}

func init() {
	rootCmd.AddCommand(adminComponentConfigurationsCmd)
}
