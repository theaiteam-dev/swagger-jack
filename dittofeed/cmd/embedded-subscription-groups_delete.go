package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	embeddedSubscriptionGroupsDeleteCmd_workspaceId string
	embeddedSubscriptionGroupsDeleteCmd_id string
)

var embeddedSubscriptionGroupsDeleteCmd = &cobra.Command{
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
		bodyMap["workspaceId"] = embeddedSubscriptionGroupsDeleteCmd_workspaceId
		bodyMap["id"] = embeddedSubscriptionGroupsDeleteCmd_id
		resp, err := c.Do("DELETE", "/api-l/embedded/subscription-groups/", pathParams, queryParams, bodyMap)
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
	embeddedSubscriptionGroupsCmd.AddCommand(embeddedSubscriptionGroupsDeleteCmd)
	embeddedSubscriptionGroupsDeleteCmd.Flags().StringVar(&embeddedSubscriptionGroupsDeleteCmd_workspaceId, "workspaceId", "", "")
	embeddedSubscriptionGroupsDeleteCmd.Flags().StringVar(&embeddedSubscriptionGroupsDeleteCmd_id, "id", "", "")
	embeddedSubscriptionGroupsDeleteCmd.MarkFlagRequired("workspaceId")
	embeddedSubscriptionGroupsDeleteCmd.MarkFlagRequired("id")
}
