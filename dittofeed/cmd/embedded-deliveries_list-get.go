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
	embeddedDeliveriesListGetCmd_workspaceId string
	embeddedDeliveriesListGetCmd_fromIdentifier string
	embeddedDeliveriesListGetCmd_toIdentifier string
	embeddedDeliveriesListGetCmd_journeyId string
	embeddedDeliveriesListGetCmd_userId string
	embeddedDeliveriesListGetCmd_limit string
	embeddedDeliveriesListGetCmd_cursor string
	embeddedDeliveriesListGetCmd_startDate string
	embeddedDeliveriesListGetCmd_endDate string
	embeddedDeliveriesListGetCmd_sortBy string
	embeddedDeliveriesListGetCmd_sortDirection string
	embeddedDeliveriesListGetCmd_broadcastId string
	embeddedDeliveriesListGetCmd_groupId string
)

var embeddedDeliveriesListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = embeddedDeliveriesListGetCmd_workspaceId
		queryParams["fromIdentifier"] = embeddedDeliveriesListGetCmd_fromIdentifier
		queryParams["toIdentifier"] = embeddedDeliveriesListGetCmd_toIdentifier
		queryParams["journeyId"] = embeddedDeliveriesListGetCmd_journeyId
		queryParams["userId"] = embeddedDeliveriesListGetCmd_userId
		embeddedDeliveriesListGetCmd_channels_vals, _ := cmd.Flags().GetStringArray("channels")
		queryParams["channels"] = strings.Join(embeddedDeliveriesListGetCmd_channels_vals, ",")
		queryParams["limit"] = embeddedDeliveriesListGetCmd_limit
		queryParams["cursor"] = embeddedDeliveriesListGetCmd_cursor
		embeddedDeliveriesListGetCmd_from_vals, _ := cmd.Flags().GetStringArray("from")
		queryParams["from"] = strings.Join(embeddedDeliveriesListGetCmd_from_vals, ",")
		embeddedDeliveriesListGetCmd_to_vals, _ := cmd.Flags().GetStringArray("to")
		queryParams["to"] = strings.Join(embeddedDeliveriesListGetCmd_to_vals, ",")
		embeddedDeliveriesListGetCmd_statuses_vals, _ := cmd.Flags().GetStringArray("statuses")
		queryParams["statuses"] = strings.Join(embeddedDeliveriesListGetCmd_statuses_vals, ",")
		embeddedDeliveriesListGetCmd_templateIds_vals, _ := cmd.Flags().GetStringArray("templateIds")
		queryParams["templateIds"] = strings.Join(embeddedDeliveriesListGetCmd_templateIds_vals, ",")
		queryParams["startDate"] = embeddedDeliveriesListGetCmd_startDate
		queryParams["endDate"] = embeddedDeliveriesListGetCmd_endDate
		queryParams["sortBy"] = embeddedDeliveriesListGetCmd_sortBy
		queryParams["sortDirection"] = embeddedDeliveriesListGetCmd_sortDirection
		queryParams["broadcastId"] = embeddedDeliveriesListGetCmd_broadcastId
		embeddedDeliveriesListGetCmd_triggeringProperties_vals, _ := cmd.Flags().GetStringArray("triggeringProperties")
		queryParams["triggeringProperties"] = strings.Join(embeddedDeliveriesListGetCmd_triggeringProperties_vals, ",")
		queryParams["groupId"] = embeddedDeliveriesListGetCmd_groupId
		embeddedDeliveriesListGetCmd_contextValues_vals, _ := cmd.Flags().GetStringArray("contextValues")
		queryParams["contextValues"] = strings.Join(embeddedDeliveriesListGetCmd_contextValues_vals, ",")
		resp, err := c.Do("GET", "/api-l/embedded/deliveries/", pathParams, queryParams, nil)
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
	embeddedDeliveriesCmd.AddCommand(embeddedDeliveriesListGetCmd)
	embeddedDeliveriesListGetCmd.Flags().StringVar(&embeddedDeliveriesListGetCmd_workspaceId, "workspaceId", "", "")
	embeddedDeliveriesListGetCmd.Flags().StringVar(&embeddedDeliveriesListGetCmd_fromIdentifier, "fromIdentifier", "", "")
	embeddedDeliveriesListGetCmd.Flags().StringVar(&embeddedDeliveriesListGetCmd_toIdentifier, "toIdentifier", "", "")
	embeddedDeliveriesListGetCmd.Flags().StringVar(&embeddedDeliveriesListGetCmd_journeyId, "journeyId", "", "")
	embeddedDeliveriesListGetCmd.Flags().StringVar(&embeddedDeliveriesListGetCmd_userId, "userId", "", "")
	embeddedDeliveriesListGetCmd.Flags().StringArray("channels", nil, "")
	embeddedDeliveriesListGetCmd.Flags().StringVar(&embeddedDeliveriesListGetCmd_limit, "limit", "", "")
	embeddedDeliveriesListGetCmd.Flags().StringVar(&embeddedDeliveriesListGetCmd_cursor, "cursor", "", "")
	embeddedDeliveriesListGetCmd.Flags().StringArray("from", nil, "")
	embeddedDeliveriesListGetCmd.Flags().StringArray("to", nil, "")
	embeddedDeliveriesListGetCmd.Flags().StringArray("statuses", nil, "")
	embeddedDeliveriesListGetCmd.Flags().StringArray("templateIds", nil, "")
	embeddedDeliveriesListGetCmd.Flags().StringVar(&embeddedDeliveriesListGetCmd_startDate, "startDate", "", "")
	embeddedDeliveriesListGetCmd.Flags().StringVar(&embeddedDeliveriesListGetCmd_endDate, "endDate", "", "")
	embeddedDeliveriesListGetCmd.Flags().StringVar(&embeddedDeliveriesListGetCmd_sortBy, "sortBy", "", "")
	embeddedDeliveriesListGetCmd.Flags().StringVar(&embeddedDeliveriesListGetCmd_sortDirection, "sortDirection", "", "")
	embeddedDeliveriesListGetCmd.Flags().StringVar(&embeddedDeliveriesListGetCmd_broadcastId, "broadcastId", "", "")
	embeddedDeliveriesListGetCmd.Flags().StringArray("triggeringProperties", nil, "")
	embeddedDeliveriesListGetCmd.Flags().StringVar(&embeddedDeliveriesListGetCmd_groupId, "groupId", "", "")
	embeddedDeliveriesListGetCmd.Flags().StringArray("contextValues", nil, "")
	embeddedDeliveriesListGetCmd.MarkFlagRequired("workspaceId")
}
