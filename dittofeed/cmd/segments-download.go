package cmd

import "github.com/spf13/cobra"

var segmentsDownloadCmd = &cobra.Command{
	Use: "segments-download",
	Short: "segments-download",
}

func init() {
	rootCmd.AddCommand(segmentsDownloadCmd)
}
