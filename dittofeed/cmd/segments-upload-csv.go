package cmd

import "github.com/spf13/cobra"

var segmentsUploadCsvCmd = &cobra.Command{
	Use: "segments-upload-csv",
	Short: "segments-upload-csv",
}

func init() {
	rootCmd.AddCommand(segmentsUploadCsvCmd)
}
