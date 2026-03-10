package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	apiDeleteDeleteCmd_id string
	apiDeleteDeleteCmd_workspaceId string
)

var apiDeleteDeleteCmd = &cobra.Command{
	Use: "delete-delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		bodyMap := map[string]interface{}{}
		bodyMap["id"] = apiDeleteDeleteCmd_id
		bodyMap["workspaceId"] = apiDeleteDeleteCmd_workspaceId
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
	apiCmd.AddCommand(apiDeleteDeleteCmd)
	apiDeleteDeleteCmd.Flags().StringVar(&apiDeleteDeleteCmd_id, "id", "", "")
	apiDeleteDeleteCmd.Flags().StringVar(&apiDeleteDeleteCmd_workspaceId, "workspaceId", "", "")
	apiDeleteDeleteCmd.MarkFlagRequired("id")
	apiDeleteDeleteCmd.MarkFlagRequired("workspaceId")
}
