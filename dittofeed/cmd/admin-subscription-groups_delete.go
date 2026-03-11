package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	adminSubscriptionGroupsDeleteCmd_id string
	adminSubscriptionGroupsDeleteCmd_workspaceId string
)

var adminSubscriptionGroupsDeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		bodyMap := map[string]interface{}{}
		bodyMap["id"] = adminSubscriptionGroupsDeleteCmd_id
		bodyMap["workspaceId"] = adminSubscriptionGroupsDeleteCmd_workspaceId
		resp, err := c.Do("DELETE", "/api/admin/subscription-groups/", pathParams, queryParams, bodyMap)
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
	adminSubscriptionGroupsCmd.AddCommand(adminSubscriptionGroupsDeleteCmd)
	adminSubscriptionGroupsDeleteCmd.Flags().StringVar(&adminSubscriptionGroupsDeleteCmd_id, "id", "", "")
	adminSubscriptionGroupsDeleteCmd.Flags().StringVar(&adminSubscriptionGroupsDeleteCmd_workspaceId, "workspaceId", "", "")
	adminSubscriptionGroupsDeleteCmd.MarkFlagRequired("id")
	adminSubscriptionGroupsDeleteCmd.MarkFlagRequired("workspaceId")
}
