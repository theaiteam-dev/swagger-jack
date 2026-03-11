package cmd

import "github.com/spf13/cobra"

var adminGroupsCmd = &cobra.Command{
	Use: "admin-groups",
	Short: "admin-groups",
}

func init() {
	rootCmd.AddCommand(adminGroupsCmd)
}
