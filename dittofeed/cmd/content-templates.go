package cmd

import "github.com/spf13/cobra"

var contentTemplatesCmd = &cobra.Command{
	Use: "content-templates",
	Short: "content-templates",
}

func init() {
	rootCmd.AddCommand(contentTemplatesCmd)
}
