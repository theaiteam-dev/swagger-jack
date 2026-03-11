package cmd

import "github.com/spf13/cobra"

var adminContentCmd = &cobra.Command{
	Use: "admin-content",
	Short: "admin-content",
}

func init() {
	rootCmd.AddCommand(adminContentCmd)
}
