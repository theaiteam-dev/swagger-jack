package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	adminGroupsListGetCmd_workspaceId string
	adminGroupsListGetCmd_userId string
	adminGroupsListGetCmd_limit string
	adminGroupsListGetCmd_offset string
)

var adminGroupsListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminGroupsListGetCmd_workspaceId
		queryParams["userId"] = adminGroupsListGetCmd_userId
		queryParams["limit"] = adminGroupsListGetCmd_limit
		queryParams["offset"] = adminGroupsListGetCmd_offset
		resp, err := c.Do("GET", "/api/admin/groups/user-groups", pathParams, queryParams, nil)
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
	adminGroupsCmd.AddCommand(adminGroupsListGetCmd)
	adminGroupsListGetCmd.Flags().StringVar(&adminGroupsListGetCmd_workspaceId, "workspaceId", "", "")
	adminGroupsListGetCmd.Flags().StringVar(&adminGroupsListGetCmd_userId, "userId", "", "")
	adminGroupsListGetCmd.Flags().StringVar(&adminGroupsListGetCmd_limit, "limit", "", "")
	adminGroupsListGetCmd.Flags().StringVar(&adminGroupsListGetCmd_offset, "offset", "", "")
	adminGroupsListGetCmd.MarkFlagRequired("workspaceId")
	adminGroupsListGetCmd.MarkFlagRequired("userId")
}
