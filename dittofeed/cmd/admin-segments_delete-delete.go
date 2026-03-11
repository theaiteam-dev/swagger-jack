package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	adminSegmentsDeleteDeleteCmd_id string
	adminSegmentsDeleteDeleteCmd_workspaceId string
)

var adminSegmentsDeleteDeleteCmd = &cobra.Command{
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
		bodyMap["id"] = adminSegmentsDeleteDeleteCmd_id
		bodyMap["workspaceId"] = adminSegmentsDeleteDeleteCmd_workspaceId
		resp, err := c.Do("DELETE", "/api/admin/segments/", pathParams, queryParams, bodyMap)
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
	adminSegmentsCmd.AddCommand(adminSegmentsDeleteDeleteCmd)
	adminSegmentsDeleteDeleteCmd.Flags().StringVar(&adminSegmentsDeleteDeleteCmd_id, "id", "", "")
	adminSegmentsDeleteDeleteCmd.Flags().StringVar(&adminSegmentsDeleteDeleteCmd_workspaceId, "workspaceId", "", "")
	adminSegmentsDeleteDeleteCmd.MarkFlagRequired("id")
	adminSegmentsDeleteDeleteCmd.MarkFlagRequired("workspaceId")
}
