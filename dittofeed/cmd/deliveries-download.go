package cmd

import "github.com/spf13/cobra"

var deliveriesDownloadCmd = &cobra.Command{
	Use: "deliveries-download",
	Short: "deliveries-download",
}

func init() {
	rootCmd.AddCommand(deliveriesDownloadCmd)
}
