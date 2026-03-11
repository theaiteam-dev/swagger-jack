package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	adminJourneysDeleteDeleteCmd_id string
	adminJourneysDeleteDeleteCmd_workspaceId string
)

var adminJourneysDeleteDeleteCmd = &cobra.Command{
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
		bodyMap["id"] = adminJourneysDeleteDeleteCmd_id
		bodyMap["workspaceId"] = adminJourneysDeleteDeleteCmd_workspaceId
		resp, err := c.Do("DELETE", "/api/admin/journeys/", pathParams, queryParams, bodyMap)
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
	adminJourneysCmd.AddCommand(adminJourneysDeleteDeleteCmd)
	adminJourneysDeleteDeleteCmd.Flags().StringVar(&adminJourneysDeleteDeleteCmd_id, "id", "", "")
	adminJourneysDeleteDeleteCmd.Flags().StringVar(&adminJourneysDeleteDeleteCmd_workspaceId, "workspaceId", "", "")
	adminJourneysDeleteDeleteCmd.MarkFlagRequired("id")
	adminJourneysDeleteDeleteCmd.MarkFlagRequired("workspaceId")
}
