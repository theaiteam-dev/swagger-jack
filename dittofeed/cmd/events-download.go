package cmd

import "github.com/spf13/cobra"

var eventsDownloadCmd = &cobra.Command{
	Use: "events-download",
	Short: "events-download",
}

func init() {
	rootCmd.AddCommand(eventsDownloadCmd)
}
