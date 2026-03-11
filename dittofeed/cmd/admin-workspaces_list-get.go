package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	adminWorkspacesListGetCmd_workspaceId string
	adminWorkspacesListGetCmd_externalId string
	adminWorkspacesListGetCmd_limit string
	adminWorkspacesListGetCmd_offset string
)

var adminWorkspacesListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminWorkspacesListGetCmd_workspaceId
		queryParams["externalId"] = adminWorkspacesListGetCmd_externalId
		queryParams["limit"] = adminWorkspacesListGetCmd_limit
		queryParams["offset"] = adminWorkspacesListGetCmd_offset
		resp, err := c.Do("GET", "/api-l/admin/workspaces/child", pathParams, queryParams, nil)
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
	adminWorkspacesCmd.AddCommand(adminWorkspacesListGetCmd)
	adminWorkspacesListGetCmd.Flags().StringVar(&adminWorkspacesListGetCmd_workspaceId, "workspaceId", "", "")
	adminWorkspacesListGetCmd.Flags().StringVar(&adminWorkspacesListGetCmd_externalId, "externalId", "", "")
	adminWorkspacesListGetCmd.Flags().StringVar(&adminWorkspacesListGetCmd_limit, "limit", "", "")
	adminWorkspacesListGetCmd.Flags().StringVar(&adminWorkspacesListGetCmd_offset, "offset", "", "")
	adminWorkspacesListGetCmd.MarkFlagRequired("workspaceId")
	adminWorkspacesListGetCmd.MarkFlagRequired("externalId")
}
