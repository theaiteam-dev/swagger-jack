package cmd

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	adminJourneysListGetCmd_workspaceId string
	adminJourneysListGetCmd_getPartial bool
	adminJourneysListGetCmd_resourceType string
)

var adminJourneysListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminJourneysListGetCmd_workspaceId
		queryParams["getPartial"] = strconv.FormatBool(adminJourneysListGetCmd_getPartial)
		adminJourneysListGetCmd_ids_vals, _ := cmd.Flags().GetStringArray("ids")
		queryParams["ids"] = strings.Join(adminJourneysListGetCmd_ids_vals, ",")
		queryParams["resourceType"] = adminJourneysListGetCmd_resourceType
		resp, err := c.Do("GET", "/api/admin/journeys/", pathParams, queryParams, nil)
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
	adminJourneysCmd.AddCommand(adminJourneysListGetCmd)
	adminJourneysListGetCmd.Flags().StringVar(&adminJourneysListGetCmd_workspaceId, "workspaceId", "", "")
	adminJourneysListGetCmd.Flags().BoolVar(&adminJourneysListGetCmd_getPartial, "getPartial", false, "")
	adminJourneysListGetCmd.Flags().StringArray("ids", nil, "")
	adminJourneysListGetCmd.Flags().StringVar(&adminJourneysListGetCmd_resourceType, "resourceType", "", "")
	adminJourneysListGetCmd.MarkFlagRequired("workspaceId")
}
