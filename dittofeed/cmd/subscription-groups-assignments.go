package cmd

import "github.com/spf13/cobra"

var subscriptionGroupsAssignmentsCmd = &cobra.Command{
	Use: "subscription-groups-assignments",
	Short: "subscription-groups-assignments",
}

func init() {
	rootCmd.AddCommand(subscriptionGroupsAssignmentsCmd)
}
