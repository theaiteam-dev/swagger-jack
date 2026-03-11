package cmd

import "github.com/spf13/cobra"

var computedPropertiesPeriodsCmd = &cobra.Command{
	Use: "computed-properties-periods",
	Short: "computed-properties-periods",
}

func init() {
	rootCmd.AddCommand(computedPropertiesPeriodsCmd)
}
