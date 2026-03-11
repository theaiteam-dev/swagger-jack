package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	adminWorkspacesDeleteCmd_workspaceId string
	adminWorkspacesDeleteCmd_externalId string
)

var adminWorkspacesDeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminWorkspacesDeleteCmd_workspaceId
		queryParams["externalId"] = adminWorkspacesDeleteCmd_externalId
		resp, err := c.Do("DELETE", "/api-l/admin/workspaces/", pathParams, queryParams, nil)
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
	adminWorkspacesCmd.AddCommand(adminWorkspacesDeleteCmd)
	adminWorkspacesDeleteCmd.Flags().StringVar(&adminWorkspacesDeleteCmd_workspaceId, "workspaceId", "", "")
	adminWorkspacesDeleteCmd.Flags().StringVar(&adminWorkspacesDeleteCmd_externalId, "externalId", "", "")
	adminWorkspacesDeleteCmd.MarkFlagRequired("workspaceId")
	adminWorkspacesDeleteCmd.MarkFlagRequired("externalId")
}
