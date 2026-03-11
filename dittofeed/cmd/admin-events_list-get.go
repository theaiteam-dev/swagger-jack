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
	adminEventsListGetCmd_workspaceId string
	adminEventsListGetCmd_searchTerm string
	adminEventsListGetCmd_userId string
	adminEventsListGetCmd_messageId string
	adminEventsListGetCmd_startDate string
	adminEventsListGetCmd_endDate string
	adminEventsListGetCmd_broadcastId string
	adminEventsListGetCmd_journeyId string
	adminEventsListGetCmd_eventType string
	adminEventsListGetCmd_includeContext bool
)

var adminEventsListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminEventsListGetCmd_workspaceId
		queryParams["searchTerm"] = adminEventsListGetCmd_searchTerm
		queryParams["userId"] = adminEventsListGetCmd_userId
		queryParams["messageId"] = adminEventsListGetCmd_messageId
		queryParams["startDate"] = adminEventsListGetCmd_startDate
		queryParams["endDate"] = adminEventsListGetCmd_endDate
		adminEventsListGetCmd_event_vals, _ := cmd.Flags().GetStringArray("event")
		queryParams["event"] = strings.Join(adminEventsListGetCmd_event_vals, ",")
		queryParams["broadcastId"] = adminEventsListGetCmd_broadcastId
		queryParams["journeyId"] = adminEventsListGetCmd_journeyId
		queryParams["eventType"] = adminEventsListGetCmd_eventType
		queryParams["includeContext"] = strconv.FormatBool(adminEventsListGetCmd_includeContext)
		resp, err := c.Do("GET", "/api/admin/events/download", pathParams, queryParams, nil)
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
	adminEventsCmd.AddCommand(adminEventsListGetCmd)
	adminEventsListGetCmd.Flags().StringVar(&adminEventsListGetCmd_workspaceId, "workspaceId", "", "")
	adminEventsListGetCmd.Flags().StringVar(&adminEventsListGetCmd_searchTerm, "searchTerm", "", "")
	adminEventsListGetCmd.Flags().StringVar(&adminEventsListGetCmd_userId, "userId", "", "Unique identifier for the user. Should be the id of the user in your system. Only applicable to logged in users.")
	adminEventsListGetCmd.Flags().StringVar(&adminEventsListGetCmd_messageId, "messageId", "", "")
	adminEventsListGetCmd.Flags().StringVar(&adminEventsListGetCmd_startDate, "startDate", "", "")
	adminEventsListGetCmd.Flags().StringVar(&adminEventsListGetCmd_endDate, "endDate", "", "")
	adminEventsListGetCmd.Flags().StringArray("event", nil, "")
	adminEventsListGetCmd.Flags().StringVar(&adminEventsListGetCmd_broadcastId, "broadcastId", "", "")
	adminEventsListGetCmd.Flags().StringVar(&adminEventsListGetCmd_journeyId, "journeyId", "", "")
	adminEventsListGetCmd.Flags().StringVar(&adminEventsListGetCmd_eventType, "eventType", "", "")
	adminEventsListGetCmd.Flags().BoolVar(&adminEventsListGetCmd_includeContext, "includeContext", false, "")
	adminEventsListGetCmd.MarkFlagRequired("workspaceId")
}
