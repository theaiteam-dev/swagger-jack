package cmd

import "github.com/spf13/cobra"

var embeddedUserPropertiesCmd = &cobra.Command{
	Use: "embedded-user-properties",
	Short: "embedded-user-properties",
}

func init() {
	rootCmd.AddCommand(embeddedUserPropertiesCmd)
}
