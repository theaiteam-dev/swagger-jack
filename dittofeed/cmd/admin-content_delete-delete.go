package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	adminContentDeleteDeleteCmd_type string
	adminContentDeleteDeleteCmd_workspaceId string
	adminContentDeleteDeleteCmd_id string
)

var adminContentDeleteDeleteCmd = &cobra.Command{
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
		bodyMap["type"] = adminContentDeleteDeleteCmd_type
		bodyMap["workspaceId"] = adminContentDeleteDeleteCmd_workspaceId
		bodyMap["id"] = adminContentDeleteDeleteCmd_id
		resp, err := c.Do("DELETE", "/api/admin/content/templates", pathParams, queryParams, bodyMap)
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
	adminContentCmd.AddCommand(adminContentDeleteDeleteCmd)
	adminContentDeleteDeleteCmd.Flags().StringVar(&adminContentDeleteDeleteCmd_type, "type", "", "")
	adminContentDeleteDeleteCmd.Flags().StringVar(&adminContentDeleteDeleteCmd_workspaceId, "workspaceId", "", "")
	adminContentDeleteDeleteCmd.Flags().StringVar(&adminContentDeleteDeleteCmd_id, "id", "", "")
	adminContentDeleteDeleteCmd.MarkFlagRequired("type")
	adminContentDeleteDeleteCmd.MarkFlagRequired("workspaceId")
	adminContentDeleteDeleteCmd.MarkFlagRequired("id")
}
