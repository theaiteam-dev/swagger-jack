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
	deliveriesListCmd_workspaceId string
	deliveriesListCmd_fromIdentifier string
	deliveriesListCmd_toIdentifier string
	deliveriesListCmd_journeyId string
	deliveriesListCmd_userId string
	deliveriesListCmd_limit string
	deliveriesListCmd_cursor string
	deliveriesListCmd_startDate string
	deliveriesListCmd_endDate string
	deliveriesListCmd_sortBy string
	deliveriesListCmd_sortDirection string
	deliveriesListCmd_broadcastId string
	deliveriesListCmd_groupId string
)

var deliveriesListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = deliveriesListCmd_workspaceId
		queryParams["fromIdentifier"] = deliveriesListCmd_fromIdentifier
		queryParams["toIdentifier"] = deliveriesListCmd_toIdentifier
		queryParams["journeyId"] = deliveriesListCmd_journeyId
		queryParams["userId"] = deliveriesListCmd_userId
		deliveriesListCmd_channels_vals, _ := cmd.Flags().GetStringArray("channels")
		queryParams["channels"] = strings.Join(deliveriesListCmd_channels_vals, ",")
		queryParams["limit"] = deliveriesListCmd_limit
		queryParams["cursor"] = deliveriesListCmd_cursor
		deliveriesListCmd_from_vals, _ := cmd.Flags().GetStringArray("from")
		queryParams["from"] = strings.Join(deliveriesListCmd_from_vals, ",")
		deliveriesListCmd_to_vals, _ := cmd.Flags().GetStringArray("to")
		queryParams["to"] = strings.Join(deliveriesListCmd_to_vals, ",")
		deliveriesListCmd_statuses_vals, _ := cmd.Flags().GetStringArray("statuses")
		queryParams["statuses"] = strings.Join(deliveriesListCmd_statuses_vals, ",")
		deliveriesListCmd_templateIds_vals, _ := cmd.Flags().GetStringArray("templateIds")
		queryParams["templateIds"] = strings.Join(deliveriesListCmd_templateIds_vals, ",")
		queryParams["startDate"] = deliveriesListCmd_startDate
		queryParams["endDate"] = deliveriesListCmd_endDate
		queryParams["sortBy"] = deliveriesListCmd_sortBy
		queryParams["sortDirection"] = deliveriesListCmd_sortDirection
		queryParams["broadcastId"] = deliveriesListCmd_broadcastId
		deliveriesListCmd_triggeringProperties_vals, _ := cmd.Flags().GetStringArray("triggeringProperties")
		queryParams["triggeringProperties"] = strings.Join(deliveriesListCmd_triggeringProperties_vals, ",")
		queryParams["groupId"] = deliveriesListCmd_groupId
		deliveriesListCmd_contextValues_vals, _ := cmd.Flags().GetStringArray("contextValues")
		queryParams["contextValues"] = strings.Join(deliveriesListCmd_contextValues_vals, ",")
		resp, err := c.Do("GET", "/api/deliveries/", pathParams, queryParams, nil)
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
	deliveriesCmd.AddCommand(deliveriesListCmd)
	deliveriesListCmd.Flags().StringVar(&deliveriesListCmd_workspaceId, "workspaceId", "", "")
	deliveriesListCmd.Flags().StringVar(&deliveriesListCmd_fromIdentifier, "fromIdentifier", "", "")
	deliveriesListCmd.Flags().StringVar(&deliveriesListCmd_toIdentifier, "toIdentifier", "", "")
	deliveriesListCmd.Flags().StringVar(&deliveriesListCmd_journeyId, "journeyId", "", "")
	deliveriesListCmd.Flags().StringVar(&deliveriesListCmd_userId, "userId", "", "")
	deliveriesListCmd.Flags().StringArray("channels", nil, "")
	deliveriesListCmd.Flags().StringVar(&deliveriesListCmd_limit, "limit", "", "")
	deliveriesListCmd.Flags().StringVar(&deliveriesListCmd_cursor, "cursor", "", "")
	deliveriesListCmd.Flags().StringArray("from", nil, "")
	deliveriesListCmd.Flags().StringArray("to", nil, "")
	deliveriesListCmd.Flags().StringArray("statuses", nil, "")
	deliveriesListCmd.Flags().StringArray("templateIds", nil, "")
	deliveriesListCmd.Flags().StringVar(&deliveriesListCmd_startDate, "startDate", "", "")
	deliveriesListCmd.Flags().StringVar(&deliveriesListCmd_endDate, "endDate", "", "")
	deliveriesListCmd.Flags().StringVar(&deliveriesListCmd_sortBy, "sortBy", "", "")
	deliveriesListCmd.Flags().StringVar(&deliveriesListCmd_sortDirection, "sortDirection", "", "")
	deliveriesListCmd.Flags().StringVar(&deliveriesListCmd_broadcastId, "broadcastId", "", "")
	deliveriesListCmd.Flags().StringArray("triggeringProperties", nil, "")
	deliveriesListCmd.Flags().StringVar(&deliveriesListCmd_groupId, "groupId", "", "")
	deliveriesListCmd.Flags().StringArray("contextValues", nil, "")
	deliveriesListCmd.MarkFlagRequired("workspaceId")
}
