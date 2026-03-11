package cmd

import "github.com/spf13/cobra"

var integrationsCmd = &cobra.Command{
	Use: "integrations",
	Short: "integrations",
}

func init() {
	rootCmd.AddCommand(integrationsCmd)
}
