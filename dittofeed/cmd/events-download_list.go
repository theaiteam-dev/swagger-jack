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
	eventsDownloadListCmd_workspaceId string
	eventsDownloadListCmd_searchTerm string
	eventsDownloadListCmd_userId string
	eventsDownloadListCmd_messageId string
	eventsDownloadListCmd_startDate string
	eventsDownloadListCmd_endDate string
	eventsDownloadListCmd_broadcastId string
	eventsDownloadListCmd_journeyId string
	eventsDownloadListCmd_eventType string
	eventsDownloadListCmd_includeContext bool
)

var eventsDownloadListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = eventsDownloadListCmd_workspaceId
		queryParams["searchTerm"] = eventsDownloadListCmd_searchTerm
		queryParams["userId"] = eventsDownloadListCmd_userId
		queryParams["messageId"] = eventsDownloadListCmd_messageId
		queryParams["startDate"] = eventsDownloadListCmd_startDate
		queryParams["endDate"] = eventsDownloadListCmd_endDate
		eventsDownloadListCmd_event_vals, _ := cmd.Flags().GetStringArray("event")
		queryParams["event"] = strings.Join(eventsDownloadListCmd_event_vals, ",")
		queryParams["broadcastId"] = eventsDownloadListCmd_broadcastId
		queryParams["journeyId"] = eventsDownloadListCmd_journeyId
		queryParams["eventType"] = eventsDownloadListCmd_eventType
		queryParams["includeContext"] = strconv.FormatBool(eventsDownloadListCmd_includeContext)
		resp, err := c.Do("GET", "/api/events/download", pathParams, queryParams, nil)
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
	eventsDownloadCmd.AddCommand(eventsDownloadListCmd)
	eventsDownloadListCmd.Flags().StringVar(&eventsDownloadListCmd_workspaceId, "workspaceId", "", "")
	eventsDownloadListCmd.Flags().StringVar(&eventsDownloadListCmd_searchTerm, "searchTerm", "", "")
	eventsDownloadListCmd.Flags().StringVar(&eventsDownloadListCmd_userId, "userId", "", "Unique identifier for the user. Should be the id of the user in your system. Only applicable to logged in users.")
	eventsDownloadListCmd.Flags().StringVar(&eventsDownloadListCmd_messageId, "messageId", "", "")
	eventsDownloadListCmd.Flags().StringVar(&eventsDownloadListCmd_startDate, "startDate", "", "")
	eventsDownloadListCmd.Flags().StringVar(&eventsDownloadListCmd_endDate, "endDate", "", "")
	eventsDownloadListCmd.Flags().StringArray("event", nil, "")
	eventsDownloadListCmd.Flags().StringVar(&eventsDownloadListCmd_broadcastId, "broadcastId", "", "")
	eventsDownloadListCmd.Flags().StringVar(&eventsDownloadListCmd_journeyId, "journeyId", "", "")
	eventsDownloadListCmd.Flags().StringVar(&eventsDownloadListCmd_eventType, "eventType", "", "")
	eventsDownloadListCmd.Flags().BoolVar(&eventsDownloadListCmd_includeContext, "includeContext", false, "")
	eventsDownloadListCmd.MarkFlagRequired("workspaceId")
}
