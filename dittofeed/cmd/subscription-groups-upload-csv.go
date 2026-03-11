package cmd

import "github.com/spf13/cobra"

var subscriptionGroupsUploadCsvCmd = &cobra.Command{
	Use: "subscription-groups-upload-csv",
	Short: "subscription-groups-upload-csv",
}

func init() {
	rootCmd.AddCommand(subscriptionGroupsUploadCsvCmd)
}
