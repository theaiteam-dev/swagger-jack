package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	internalApiLWorkspacesDeleteCmd_workspaceId string
	internalApiLWorkspacesDeleteCmd_externalId string
)

var internalApiLWorkspacesDeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = internalApiLWorkspacesDeleteCmd_workspaceId
		queryParams["externalId"] = internalApiLWorkspacesDeleteCmd_externalId
		resp, err := c.Do("DELETE", "/internal-api-l/workspaces/", pathParams, queryParams, nil)
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
	internalApiLWorkspacesCmd.AddCommand(internalApiLWorkspacesDeleteCmd)
	internalApiLWorkspacesDeleteCmd.Flags().StringVar(&internalApiLWorkspacesDeleteCmd_workspaceId, "workspaceId", "", "")
	internalApiLWorkspacesDeleteCmd.Flags().StringVar(&internalApiLWorkspacesDeleteCmd_externalId, "externalId", "", "")
	internalApiLWorkspacesDeleteCmd.MarkFlagRequired("workspaceId")
	internalApiLWorkspacesDeleteCmd.MarkFlagRequired("externalId")
}
