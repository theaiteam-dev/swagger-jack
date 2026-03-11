package cmd

import "github.com/spf13/cobra"

var computedPropertiesTriggerRecomputeCmd = &cobra.Command{
	Use: "computed-properties-trigger-recompute",
	Short: "computed-properties-trigger-recompute",
}

func init() {
	rootCmd.AddCommand(computedPropertiesTriggerRecomputeCmd)
}
