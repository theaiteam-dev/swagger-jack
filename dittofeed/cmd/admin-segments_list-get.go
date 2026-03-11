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
	adminSegmentsListGetCmd_workspaceId string
	adminSegmentsListGetCmd_resourceType string
)

var adminSegmentsListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminSegmentsListGetCmd_workspaceId
		adminSegmentsListGetCmd_ids_vals, _ := cmd.Flags().GetStringArray("ids")
		queryParams["ids"] = strings.Join(adminSegmentsListGetCmd_ids_vals, ",")
		queryParams["resourceType"] = adminSegmentsListGetCmd_resourceType
		resp, err := c.Do("GET", "/api/admin/segments/", pathParams, queryParams, nil)
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
	adminSegmentsCmd.AddCommand(adminSegmentsListGetCmd)
	adminSegmentsListGetCmd.Flags().StringVar(&adminSegmentsListGetCmd_workspaceId, "workspaceId", "", "")
	adminSegmentsListGetCmd.Flags().StringArray("ids", nil, "")
	adminSegmentsListGetCmd.Flags().StringVar(&adminSegmentsListGetCmd_resourceType, "resourceType", "", "")
	adminSegmentsListGetCmd.MarkFlagRequired("workspaceId")
}
