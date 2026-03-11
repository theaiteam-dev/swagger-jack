package cmd

import "github.com/spf13/cobra"

var userPropertyIndicesCmd = &cobra.Command{
	Use: "user-property-indices",
	Short: "user-property-indices",
}

func init() {
	rootCmd.AddCommand(userPropertyIndicesCmd)
}
