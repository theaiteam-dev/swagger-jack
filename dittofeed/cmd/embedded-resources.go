package cmd

import "github.com/spf13/cobra"

var embeddedResourcesCmd = &cobra.Command{
	Use: "embedded-resources",
	Short: "embedded-resources",
}

func init() {
	rootCmd.AddCommand(embeddedResourcesCmd)
}
