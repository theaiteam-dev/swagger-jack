package cmd

import "github.com/spf13/cobra"

var embeddedContentCmd = &cobra.Command{
	Use: "embedded-content",
	Short: "embedded-content",
}

func init() {
	rootCmd.AddCommand(embeddedContentCmd)
}
