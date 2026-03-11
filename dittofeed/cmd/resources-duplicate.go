package cmd

import "github.com/spf13/cobra"

var resourcesDuplicateCmd = &cobra.Command{
	Use: "resources-duplicate",
	Short: "resources-duplicate",
}

func init() {
	rootCmd.AddCommand(resourcesDuplicateCmd)
}
