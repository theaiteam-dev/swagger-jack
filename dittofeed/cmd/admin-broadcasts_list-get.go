package cmd

import (
	"fmt"
	"os"
	"strings"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	adminBroadcastsListGetCmd_workspaceId string
)

var adminBroadcastsListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminBroadcastsListGetCmd_workspaceId
		adminBroadcastsListGetCmd_ids_vals, _ := cmd.Flags().GetStringArray("ids")
		queryParams["ids"] = strings.Join(adminBroadcastsListGetCmd_ids_vals, ",")
		resp, err := c.Do("GET", "/api/admin/broadcasts/", pathParams, queryParams, nil)
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
	adminBroadcastsCmd.AddCommand(adminBroadcastsListGetCmd)
	adminBroadcastsListGetCmd.Flags().StringVar(&adminBroadcastsListGetCmd_workspaceId, "workspaceId", "", "")
	adminBroadcastsListGetCmd.Flags().StringArray("ids", nil, "")
	adminBroadcastsListGetCmd.MarkFlagRequired("workspaceId")
}
