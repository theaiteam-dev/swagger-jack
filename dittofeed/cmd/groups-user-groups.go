package cmd

import "github.com/spf13/cobra"

var groupsUserGroupsCmd = &cobra.Command{
	Use: "groups-user-groups",
	Short: "groups-user-groups",
}

func init() {
	rootCmd.AddCommand(groupsUserGroupsCmd)
}
