package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	adminKeysDeleteCmd_workspaceId string
	adminKeysDeleteCmd_id string
)

var adminKeysDeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminKeysDeleteCmd_workspaceId
		queryParams["id"] = adminKeysDeleteCmd_id
		resp, err := c.Do("DELETE", "/api/admin-keys/", pathParams, queryParams, nil)
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
	adminKeysCmd.AddCommand(adminKeysDeleteCmd)
	adminKeysDeleteCmd.Flags().StringVar(&adminKeysDeleteCmd_workspaceId, "workspaceId", "", "")
	adminKeysDeleteCmd.Flags().StringVar(&adminKeysDeleteCmd_id, "id", "", "")
	adminKeysDeleteCmd.MarkFlagRequired("workspaceId")
	adminKeysDeleteCmd.MarkFlagRequired("id")
}
