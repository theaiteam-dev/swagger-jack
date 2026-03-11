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
	contentTemplatesListCmd_workspaceId string
	contentTemplatesListCmd_resourceType string
)

var contentTemplatesListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = contentTemplatesListCmd_workspaceId
		contentTemplatesListCmd_ids_vals, _ := cmd.Flags().GetStringArray("ids")
		queryParams["ids"] = strings.Join(contentTemplatesListCmd_ids_vals, ",")
		queryParams["resourceType"] = contentTemplatesListCmd_resourceType
		resp, err := c.Do("GET", "/api/content/templates", pathParams, queryParams, nil)
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
	contentTemplatesCmd.AddCommand(contentTemplatesListCmd)
	contentTemplatesListCmd.Flags().StringVar(&contentTemplatesListCmd_workspaceId, "workspaceId", "", "")
	contentTemplatesListCmd.Flags().StringArray("ids", nil, "")
	contentTemplatesListCmd.Flags().StringVar(&contentTemplatesListCmd_resourceType, "resourceType", "", "")
	contentTemplatesListCmd.MarkFlagRequired("workspaceId")
}
