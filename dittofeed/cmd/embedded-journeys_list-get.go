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
	embeddedJourneysListGetCmd_workspaceId string
	embeddedJourneysListGetCmd_getPartial bool
	embeddedJourneysListGetCmd_resourceType string
)

var embeddedJourneysListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = embeddedJourneysListGetCmd_workspaceId
		queryParams["getPartial"] = strconv.FormatBool(embeddedJourneysListGetCmd_getPartial)
		embeddedJourneysListGetCmd_ids_vals, _ := cmd.Flags().GetStringArray("ids")
		queryParams["ids"] = strings.Join(embeddedJourneysListGetCmd_ids_vals, ",")
		queryParams["resourceType"] = embeddedJourneysListGetCmd_resourceType
		resp, err := c.Do("GET", "/api-l/embedded/journeys/", pathParams, queryParams, nil)
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
	embeddedJourneysCmd.AddCommand(embeddedJourneysListGetCmd)
	embeddedJourneysListGetCmd.Flags().StringVar(&embeddedJourneysListGetCmd_workspaceId, "workspaceId", "", "")
	embeddedJourneysListGetCmd.Flags().BoolVar(&embeddedJourneysListGetCmd_getPartial, "getPartial", false, "")
	embeddedJourneysListGetCmd.Flags().StringArray("ids", nil, "")
	embeddedJourneysListGetCmd.Flags().StringVar(&embeddedJourneysListGetCmd_resourceType, "resourceType", "", "")
	embeddedJourneysListGetCmd.MarkFlagRequired("workspaceId")
}
