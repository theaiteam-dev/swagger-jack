package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	embeddedContentDeleteDeleteCmd_workspaceId string
	embeddedContentDeleteDeleteCmd_id string
	embeddedContentDeleteDeleteCmd_type string
)

var embeddedContentDeleteDeleteCmd = &cobra.Command{
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
		bodyMap["workspaceId"] = embeddedContentDeleteDeleteCmd_workspaceId
		bodyMap["id"] = embeddedContentDeleteDeleteCmd_id
		bodyMap["type"] = embeddedContentDeleteDeleteCmd_type
		resp, err := c.Do("DELETE", "/api-l/embedded/content/templates", pathParams, queryParams, bodyMap)
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
	embeddedContentCmd.AddCommand(embeddedContentDeleteDeleteCmd)
	embeddedContentDeleteDeleteCmd.Flags().StringVar(&embeddedContentDeleteDeleteCmd_workspaceId, "workspaceId", "", "")
	embeddedContentDeleteDeleteCmd.Flags().StringVar(&embeddedContentDeleteDeleteCmd_id, "id", "", "")
	embeddedContentDeleteDeleteCmd.Flags().StringVar(&embeddedContentDeleteDeleteCmd_type, "type", "", "")
	embeddedContentDeleteDeleteCmd.MarkFlagRequired("workspaceId")
	embeddedContentDeleteDeleteCmd.MarkFlagRequired("id")
	embeddedContentDeleteDeleteCmd.MarkFlagRequired("type")
}
