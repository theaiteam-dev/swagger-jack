package cmd

import "github.com/spf13/cobra"

var groupsUsersCmd = &cobra.Command{
	Use: "groups-users",
	Short: "groups-users",
}

func init() {
	rootCmd.AddCommand(groupsUsersCmd)
}
