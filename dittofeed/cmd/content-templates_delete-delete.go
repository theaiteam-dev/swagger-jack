package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	contentTemplatesDeleteDeleteCmd_id string
	contentTemplatesDeleteDeleteCmd_type string
	contentTemplatesDeleteDeleteCmd_workspaceId string
)

var contentTemplatesDeleteDeleteCmd = &cobra.Command{
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
		bodyMap["id"] = contentTemplatesDeleteDeleteCmd_id
		bodyMap["type"] = contentTemplatesDeleteDeleteCmd_type
		bodyMap["workspaceId"] = contentTemplatesDeleteDeleteCmd_workspaceId
		resp, err := c.Do("DELETE", "/api/content/templates", pathParams, queryParams, bodyMap)
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
	contentTemplatesCmd.AddCommand(contentTemplatesDeleteDeleteCmd)
	contentTemplatesDeleteDeleteCmd.Flags().StringVar(&contentTemplatesDeleteDeleteCmd_id, "id", "", "")
	contentTemplatesDeleteDeleteCmd.Flags().StringVar(&contentTemplatesDeleteDeleteCmd_type, "type", "", "")
	contentTemplatesDeleteDeleteCmd.Flags().StringVar(&contentTemplatesDeleteDeleteCmd_workspaceId, "workspaceId", "", "")
	contentTemplatesDeleteDeleteCmd.MarkFlagRequired("id")
	contentTemplatesDeleteDeleteCmd.MarkFlagRequired("type")
	contentTemplatesDeleteDeleteCmd.MarkFlagRequired("workspaceId")
}
