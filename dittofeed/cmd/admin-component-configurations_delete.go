package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	adminComponentConfigurationsDeleteCmd_workspaceId string
	adminComponentConfigurationsDeleteCmd_id string
)

var adminComponentConfigurationsDeleteCmd = &cobra.Command{
	Use: "delete",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminComponentConfigurationsDeleteCmd_workspaceId
		queryParams["id"] = adminComponentConfigurationsDeleteCmd_id
		resp, err := c.Do("DELETE", "/api/admin/component-configurations/", pathParams, queryParams, nil)
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
	adminComponentConfigurationsCmd.AddCommand(adminComponentConfigurationsDeleteCmd)
	adminComponentConfigurationsDeleteCmd.Flags().StringVar(&adminComponentConfigurationsDeleteCmd_workspaceId, "workspaceId", "", "")
	adminComponentConfigurationsDeleteCmd.Flags().StringVar(&adminComponentConfigurationsDeleteCmd_id, "id", "", "")
	adminComponentConfigurationsDeleteCmd.MarkFlagRequired("workspaceId")
	adminComponentConfigurationsDeleteCmd.MarkFlagRequired("id")
}
