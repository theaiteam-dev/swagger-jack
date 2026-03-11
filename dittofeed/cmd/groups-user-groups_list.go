package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	groupsUserGroupsListCmd_workspaceId string
	groupsUserGroupsListCmd_userId string
	groupsUserGroupsListCmd_limit string
	groupsUserGroupsListCmd_offset string
)

var groupsUserGroupsListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = groupsUserGroupsListCmd_workspaceId
		queryParams["userId"] = groupsUserGroupsListCmd_userId
		queryParams["limit"] = groupsUserGroupsListCmd_limit
		queryParams["offset"] = groupsUserGroupsListCmd_offset
		resp, err := c.Do("GET", "/api/groups/user-groups", pathParams, queryParams, nil)
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
	groupsUserGroupsCmd.AddCommand(groupsUserGroupsListCmd)
	groupsUserGroupsListCmd.Flags().StringVar(&groupsUserGroupsListCmd_workspaceId, "workspaceId", "", "")
	groupsUserGroupsListCmd.Flags().StringVar(&groupsUserGroupsListCmd_userId, "userId", "", "")
	groupsUserGroupsListCmd.Flags().StringVar(&groupsUserGroupsListCmd_limit, "limit", "", "")
	groupsUserGroupsListCmd.Flags().StringVar(&groupsUserGroupsListCmd_offset, "offset", "", "")
	groupsUserGroupsListCmd.MarkFlagRequired("workspaceId")
	groupsUserGroupsListCmd.MarkFlagRequired("userId")
}
