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
	adminContentListCmd_workspaceId string
	adminContentListCmd_resourceType string
)

var adminContentListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminContentListCmd_workspaceId
		adminContentListCmd_ids_vals, _ := cmd.Flags().GetStringArray("ids")
		queryParams["ids"] = strings.Join(adminContentListCmd_ids_vals, ",")
		queryParams["resourceType"] = adminContentListCmd_resourceType
		resp, err := c.Do("GET", "/api/admin/content/templates", pathParams, queryParams, nil)
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
	adminContentCmd.AddCommand(adminContentListCmd)
	adminContentListCmd.Flags().StringVar(&adminContentListCmd_workspaceId, "workspaceId", "", "")
	adminContentListCmd.Flags().StringArray("ids", nil, "")
	adminContentListCmd.Flags().StringVar(&adminContentListCmd_resourceType, "resourceType", "", "")
	adminContentListCmd.MarkFlagRequired("workspaceId")
}
