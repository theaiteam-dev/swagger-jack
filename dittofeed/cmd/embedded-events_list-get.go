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
	embeddedEventsListGetCmd_workspaceId string
	embeddedEventsListGetCmd_searchTerm string
	embeddedEventsListGetCmd_userId string
	embeddedEventsListGetCmd_offset string
	embeddedEventsListGetCmd_limit string
	embeddedEventsListGetCmd_messageId string
	embeddedEventsListGetCmd_startDate string
	embeddedEventsListGetCmd_endDate string
	embeddedEventsListGetCmd_broadcastId string
	embeddedEventsListGetCmd_journeyId string
	embeddedEventsListGetCmd_eventType string
	embeddedEventsListGetCmd_includeContext bool
)

var embeddedEventsListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = embeddedEventsListGetCmd_workspaceId
		queryParams["searchTerm"] = embeddedEventsListGetCmd_searchTerm
		queryParams["userId"] = embeddedEventsListGetCmd_userId
		queryParams["offset"] = embeddedEventsListGetCmd_offset
		queryParams["limit"] = embeddedEventsListGetCmd_limit
		queryParams["messageId"] = embeddedEventsListGetCmd_messageId
		queryParams["startDate"] = embeddedEventsListGetCmd_startDate
		queryParams["endDate"] = embeddedEventsListGetCmd_endDate
		embeddedEventsListGetCmd_event_vals, _ := cmd.Flags().GetStringArray("event")
		queryParams["event"] = strings.Join(embeddedEventsListGetCmd_event_vals, ",")
		queryParams["broadcastId"] = embeddedEventsListGetCmd_broadcastId
		queryParams["journeyId"] = embeddedEventsListGetCmd_journeyId
		queryParams["eventType"] = embeddedEventsListGetCmd_eventType
		queryParams["includeContext"] = strconv.FormatBool(embeddedEventsListGetCmd_includeContext)
		resp, err := c.Do("GET", "/api-l/embedded/events/", pathParams, queryParams, nil)
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
	embeddedEventsCmd.AddCommand(embeddedEventsListGetCmd)
	embeddedEventsListGetCmd.Flags().StringVar(&embeddedEventsListGetCmd_workspaceId, "workspaceId", "", "")
	embeddedEventsListGetCmd.Flags().StringVar(&embeddedEventsListGetCmd_searchTerm, "searchTerm", "", "")
	embeddedEventsListGetCmd.Flags().StringVar(&embeddedEventsListGetCmd_userId, "userId", "", "Unique identifier for the user. Should be the id of the user in your system. Only applicable to logged in users.")
	embeddedEventsListGetCmd.Flags().StringVar(&embeddedEventsListGetCmd_offset, "offset", "", "")
	embeddedEventsListGetCmd.Flags().StringVar(&embeddedEventsListGetCmd_limit, "limit", "", "")
	embeddedEventsListGetCmd.Flags().StringVar(&embeddedEventsListGetCmd_messageId, "messageId", "", "")
	embeddedEventsListGetCmd.Flags().StringVar(&embeddedEventsListGetCmd_startDate, "startDate", "", "")
	embeddedEventsListGetCmd.Flags().StringVar(&embeddedEventsListGetCmd_endDate, "endDate", "", "")
	embeddedEventsListGetCmd.Flags().StringArray("event", nil, "")
	embeddedEventsListGetCmd.Flags().StringVar(&embeddedEventsListGetCmd_broadcastId, "broadcastId", "", "")
	embeddedEventsListGetCmd.Flags().StringVar(&embeddedEventsListGetCmd_journeyId, "journeyId", "", "")
	embeddedEventsListGetCmd.Flags().StringVar(&embeddedEventsListGetCmd_eventType, "eventType", "", "")
	embeddedEventsListGetCmd.Flags().BoolVar(&embeddedEventsListGetCmd_includeContext, "includeContext", false, "")
	embeddedEventsListGetCmd.MarkFlagRequired("workspaceId")
}
