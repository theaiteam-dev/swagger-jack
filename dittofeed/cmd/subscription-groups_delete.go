package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	subscriptionGroupsDeleteCmd_id string
	subscriptionGroupsDeleteCmd_workspaceId string
)

var subscriptionGroupsDeleteCmd = &cobra.Command{
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
		bodyMap["id"] = subscriptionGroupsDeleteCmd_id
		bodyMap["workspaceId"] = subscriptionGroupsDeleteCmd_workspaceId
		resp, err := c.Do("DELETE", "/api/subscription-groups/", pathParams, queryParams, bodyMap)
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
	subscriptionGroupsCmd.AddCommand(subscriptionGroupsDeleteCmd)
	subscriptionGroupsDeleteCmd.Flags().StringVar(&subscriptionGroupsDeleteCmd_id, "id", "", "")
	subscriptionGroupsDeleteCmd.Flags().StringVar(&subscriptionGroupsDeleteCmd_workspaceId, "workspaceId", "", "")
	subscriptionGroupsDeleteCmd.MarkFlagRequired("id")
	subscriptionGroupsDeleteCmd.MarkFlagRequired("workspaceId")
}
