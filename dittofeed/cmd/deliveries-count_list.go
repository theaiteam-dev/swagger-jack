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
	deliveriesCountListCmd_workspaceId string
	deliveriesCountListCmd_fromIdentifier string
	deliveriesCountListCmd_toIdentifier string
	deliveriesCountListCmd_journeyId string
	deliveriesCountListCmd_userId string
	deliveriesCountListCmd_limit string
	deliveriesCountListCmd_cursor string
	deliveriesCountListCmd_startDate string
	deliveriesCountListCmd_endDate string
	deliveriesCountListCmd_sortBy string
	deliveriesCountListCmd_sortDirection string
	deliveriesCountListCmd_broadcastId string
	deliveriesCountListCmd_groupId string
)

var deliveriesCountListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = deliveriesCountListCmd_workspaceId
		queryParams["fromIdentifier"] = deliveriesCountListCmd_fromIdentifier
		queryParams["toIdentifier"] = deliveriesCountListCmd_toIdentifier
		queryParams["journeyId"] = deliveriesCountListCmd_journeyId
		queryParams["userId"] = deliveriesCountListCmd_userId
		deliveriesCountListCmd_channels_vals, _ := cmd.Flags().GetStringArray("channels")
		queryParams["channels"] = strings.Join(deliveriesCountListCmd_channels_vals, ",")
		queryParams["limit"] = deliveriesCountListCmd_limit
		queryParams["cursor"] = deliveriesCountListCmd_cursor
		deliveriesCountListCmd_from_vals, _ := cmd.Flags().GetStringArray("from")
		queryParams["from"] = strings.Join(deliveriesCountListCmd_from_vals, ",")
		deliveriesCountListCmd_to_vals, _ := cmd.Flags().GetStringArray("to")
		queryParams["to"] = strings.Join(deliveriesCountListCmd_to_vals, ",")
		deliveriesCountListCmd_statuses_vals, _ := cmd.Flags().GetStringArray("statuses")
		queryParams["statuses"] = strings.Join(deliveriesCountListCmd_statuses_vals, ",")
		deliveriesCountListCmd_templateIds_vals, _ := cmd.Flags().GetStringArray("templateIds")
		queryParams["templateIds"] = strings.Join(deliveriesCountListCmd_templateIds_vals, ",")
		queryParams["startDate"] = deliveriesCountListCmd_startDate
		queryParams["endDate"] = deliveriesCountListCmd_endDate
		queryParams["sortBy"] = deliveriesCountListCmd_sortBy
		queryParams["sortDirection"] = deliveriesCountListCmd_sortDirection
		queryParams["broadcastId"] = deliveriesCountListCmd_broadcastId
		deliveriesCountListCmd_triggeringProperties_vals, _ := cmd.Flags().GetStringArray("triggeringProperties")
		queryParams["triggeringProperties"] = strings.Join(deliveriesCountListCmd_triggeringProperties_vals, ",")
		queryParams["groupId"] = deliveriesCountListCmd_groupId
		deliveriesCountListCmd_contextValues_vals, _ := cmd.Flags().GetStringArray("contextValues")
		queryParams["contextValues"] = strings.Join(deliveriesCountListCmd_contextValues_vals, ",")
		resp, err := c.Do("GET", "/api/deliveries/count", pathParams, queryParams, nil)
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
	deliveriesCountCmd.AddCommand(deliveriesCountListCmd)
	deliveriesCountListCmd.Flags().StringVar(&deliveriesCountListCmd_workspaceId, "workspaceId", "", "")
	deliveriesCountListCmd.Flags().StringVar(&deliveriesCountListCmd_fromIdentifier, "fromIdentifier", "", "")
	deliveriesCountListCmd.Flags().StringVar(&deliveriesCountListCmd_toIdentifier, "toIdentifier", "", "")
	deliveriesCountListCmd.Flags().StringVar(&deliveriesCountListCmd_journeyId, "journeyId", "", "")
	deliveriesCountListCmd.Flags().StringVar(&deliveriesCountListCmd_userId, "userId", "", "")
	deliveriesCountListCmd.Flags().StringArray("channels", nil, "")
	deliveriesCountListCmd.Flags().StringVar(&deliveriesCountListCmd_limit, "limit", "", "")
	deliveriesCountListCmd.Flags().StringVar(&deliveriesCountListCmd_cursor, "cursor", "", "")
	deliveriesCountListCmd.Flags().StringArray("from", nil, "")
	deliveriesCountListCmd.Flags().StringArray("to", nil, "")
	deliveriesCountListCmd.Flags().StringArray("statuses", nil, "")
	deliveriesCountListCmd.Flags().StringArray("templateIds", nil, "")
	deliveriesCountListCmd.Flags().StringVar(&deliveriesCountListCmd_startDate, "startDate", "", "")
	deliveriesCountListCmd.Flags().StringVar(&deliveriesCountListCmd_endDate, "endDate", "", "")
	deliveriesCountListCmd.Flags().StringVar(&deliveriesCountListCmd_sortBy, "sortBy", "", "")
	deliveriesCountListCmd.Flags().StringVar(&deliveriesCountListCmd_sortDirection, "sortDirection", "", "")
	deliveriesCountListCmd.Flags().StringVar(&deliveriesCountListCmd_broadcastId, "broadcastId", "", "")
	deliveriesCountListCmd.Flags().StringArray("triggeringProperties", nil, "")
	deliveriesCountListCmd.Flags().StringVar(&deliveriesCountListCmd_groupId, "groupId", "", "")
	deliveriesCountListCmd.Flags().StringArray("contextValues", nil, "")
	deliveriesCountListCmd.MarkFlagRequired("workspaceId")
}
