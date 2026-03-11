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
	eventsListCmd_workspaceId string
	eventsListCmd_searchTerm string
	eventsListCmd_userId string
	eventsListCmd_offset string
	eventsListCmd_limit string
	eventsListCmd_messageId string
	eventsListCmd_startDate string
	eventsListCmd_endDate string
	eventsListCmd_broadcastId string
	eventsListCmd_journeyId string
	eventsListCmd_eventType string
	eventsListCmd_includeContext bool
)

var eventsListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = eventsListCmd_workspaceId
		queryParams["searchTerm"] = eventsListCmd_searchTerm
		queryParams["userId"] = eventsListCmd_userId
		queryParams["offset"] = eventsListCmd_offset
		queryParams["limit"] = eventsListCmd_limit
		queryParams["messageId"] = eventsListCmd_messageId
		queryParams["startDate"] = eventsListCmd_startDate
		queryParams["endDate"] = eventsListCmd_endDate
		eventsListCmd_event_vals, _ := cmd.Flags().GetStringArray("event")
		queryParams["event"] = strings.Join(eventsListCmd_event_vals, ",")
		queryParams["broadcastId"] = eventsListCmd_broadcastId
		queryParams["journeyId"] = eventsListCmd_journeyId
		queryParams["eventType"] = eventsListCmd_eventType
		queryParams["includeContext"] = strconv.FormatBool(eventsListCmd_includeContext)
		resp, err := c.Do("GET", "/api/events/", pathParams, queryParams, nil)
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
	eventsCmd.AddCommand(eventsListCmd)
	eventsListCmd.Flags().StringVar(&eventsListCmd_workspaceId, "workspaceId", "", "")
	eventsListCmd.Flags().StringVar(&eventsListCmd_searchTerm, "searchTerm", "", "")
	eventsListCmd.Flags().StringVar(&eventsListCmd_userId, "userId", "", "Unique identifier for the user. Should be the id of the user in your system. Only applicable to logged in users.")
	eventsListCmd.Flags().StringVar(&eventsListCmd_offset, "offset", "", "")
	eventsListCmd.Flags().StringVar(&eventsListCmd_limit, "limit", "", "")
	eventsListCmd.Flags().StringVar(&eventsListCmd_messageId, "messageId", "", "")
	eventsListCmd.Flags().StringVar(&eventsListCmd_startDate, "startDate", "", "")
	eventsListCmd.Flags().StringVar(&eventsListCmd_endDate, "endDate", "", "")
	eventsListCmd.Flags().StringArray("event", nil, "")
	eventsListCmd.Flags().StringVar(&eventsListCmd_broadcastId, "broadcastId", "", "")
	eventsListCmd.Flags().StringVar(&eventsListCmd_journeyId, "journeyId", "", "")
	eventsListCmd.Flags().StringVar(&eventsListCmd_eventType, "eventType", "", "")
	eventsListCmd.Flags().BoolVar(&eventsListCmd_includeContext, "includeContext", false, "")
	eventsListCmd.MarkFlagRequired("workspaceId")
}
