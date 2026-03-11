package cmd

import "github.com/spf13/cobra"

var secretsCmd = &cobra.Command{
	Use: "secrets",
	Short: "secrets",
}

func init() {
	rootCmd.AddCommand(secretsCmd)
}
