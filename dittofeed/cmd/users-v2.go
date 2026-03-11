package cmd

import "github.com/spf13/cobra"

var usersV2Cmd = &cobra.Command{
	Use: "users-v2",
	Short: "users-v2",
}

func init() {
	rootCmd.AddCommand(usersV2Cmd)
}
