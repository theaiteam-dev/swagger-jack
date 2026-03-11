package cmd

import "github.com/spf13/cobra"

var embeddedComputedPropertiesCmd = &cobra.Command{
	Use: "embedded-computed-properties",
	Short: "embedded-computed-properties",
}

func init() {
	rootCmd.AddCommand(embeddedComputedPropertiesCmd)
}
