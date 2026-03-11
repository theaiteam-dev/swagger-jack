package cmd

import "github.com/spf13/cobra"

var resourcesCmd = &cobra.Command{
	Use: "resources",
	Short: "resources",
}

func init() {
	rootCmd.AddCommand(resourcesCmd)
}
