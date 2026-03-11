package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	groupsUsersListCmd_workspaceId string
	groupsUsersListCmd_groupId string
	groupsUsersListCmd_limit string
	groupsUsersListCmd_offset string
)

var groupsUsersListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = groupsUsersListCmd_workspaceId
		queryParams["groupId"] = groupsUsersListCmd_groupId
		queryParams["limit"] = groupsUsersListCmd_limit
		queryParams["offset"] = groupsUsersListCmd_offset
		resp, err := c.Do("GET", "/api/groups/users", pathParams, queryParams, nil)
		if err != nil {
			return err
		}
		jsonMode, _ := cmd.Root().PersistentFlags().GetBool("json")
		noColor, _ := cmd.Root().PersistentFlags().GetBool("no-color")
		if jsonMode {
			fmt.Printf("%s\n", string(resp))
		} else {
			if err := output.PrintTable(resp, noColor); err != nil {
				fmt.Println(string(resp))
			}
		}
		return nil
	},
}

func init() {
	groupsUsersCmd.AddCommand(groupsUsersListCmd)
	groupsUsersListCmd.Flags().StringVar(&groupsUsersListCmd_workspaceId, "workspaceId", "", "")
	groupsUsersListCmd.Flags().StringVar(&groupsUsersListCmd_groupId, "groupId", "", "")
	groupsUsersListCmd.Flags().StringVar(&groupsUsersListCmd_limit, "limit", "", "")
	groupsUsersListCmd.Flags().StringVar(&groupsUsersListCmd_offset, "offset", "", "")
	groupsUsersListCmd.MarkFlagRequired("workspaceId")
	groupsUsersListCmd.MarkFlagRequired("groupId")
}
