package cmd

import "github.com/spf13/cobra"

var adminSegmentsCmd = &cobra.Command{
	Use: "admin-segments",
	Short: "admin-segments",
}

func init() {
	rootCmd.AddCommand(adminSegmentsCmd)
}
