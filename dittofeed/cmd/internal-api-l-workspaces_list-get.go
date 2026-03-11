package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	internalApiLWorkspacesListGetCmd_workspaceId string
	internalApiLWorkspacesListGetCmd_externalId string
	internalApiLWorkspacesListGetCmd_limit string
	internalApiLWorkspacesListGetCmd_offset string
)

var internalApiLWorkspacesListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = internalApiLWorkspacesListGetCmd_workspaceId
		queryParams["externalId"] = internalApiLWorkspacesListGetCmd_externalId
		queryParams["limit"] = internalApiLWorkspacesListGetCmd_limit
		queryParams["offset"] = internalApiLWorkspacesListGetCmd_offset
		resp, err := c.Do("GET", "/internal-api-l/workspaces/child", pathParams, queryParams, nil)
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
	internalApiLWorkspacesCmd.AddCommand(internalApiLWorkspacesListGetCmd)
	internalApiLWorkspacesListGetCmd.Flags().StringVar(&internalApiLWorkspacesListGetCmd_workspaceId, "workspaceId", "", "")
	internalApiLWorkspacesListGetCmd.Flags().StringVar(&internalApiLWorkspacesListGetCmd_externalId, "externalId", "", "")
	internalApiLWorkspacesListGetCmd.Flags().StringVar(&internalApiLWorkspacesListGetCmd_limit, "limit", "", "")
	internalApiLWorkspacesListGetCmd.Flags().StringVar(&internalApiLWorkspacesListGetCmd_offset, "offset", "", "")
	internalApiLWorkspacesListGetCmd.MarkFlagRequired("workspaceId")
	internalApiLWorkspacesListGetCmd.MarkFlagRequired("externalId")
}
