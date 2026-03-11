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
	adminDeliveriesListGetCmd_workspaceId string
	adminDeliveriesListGetCmd_fromIdentifier string
	adminDeliveriesListGetCmd_toIdentifier string
	adminDeliveriesListGetCmd_journeyId string
	adminDeliveriesListGetCmd_userId string
	adminDeliveriesListGetCmd_startDate string
	adminDeliveriesListGetCmd_endDate string
	adminDeliveriesListGetCmd_sortBy string
	adminDeliveriesListGetCmd_sortDirection string
	adminDeliveriesListGetCmd_broadcastId string
	adminDeliveriesListGetCmd_groupId string
)

var adminDeliveriesListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminDeliveriesListGetCmd_workspaceId
		queryParams["fromIdentifier"] = adminDeliveriesListGetCmd_fromIdentifier
		queryParams["toIdentifier"] = adminDeliveriesListGetCmd_toIdentifier
		queryParams["journeyId"] = adminDeliveriesListGetCmd_journeyId
		queryParams["userId"] = adminDeliveriesListGetCmd_userId
		adminDeliveriesListGetCmd_channels_vals, _ := cmd.Flags().GetStringArray("channels")
		queryParams["channels"] = strings.Join(adminDeliveriesListGetCmd_channels_vals, ",")
		adminDeliveriesListGetCmd_from_vals, _ := cmd.Flags().GetStringArray("from")
		queryParams["from"] = strings.Join(adminDeliveriesListGetCmd_from_vals, ",")
		adminDeliveriesListGetCmd_to_vals, _ := cmd.Flags().GetStringArray("to")
		queryParams["to"] = strings.Join(adminDeliveriesListGetCmd_to_vals, ",")
		adminDeliveriesListGetCmd_statuses_vals, _ := cmd.Flags().GetStringArray("statuses")
		queryParams["statuses"] = strings.Join(adminDeliveriesListGetCmd_statuses_vals, ",")
		adminDeliveriesListGetCmd_templateIds_vals, _ := cmd.Flags().GetStringArray("templateIds")
		queryParams["templateIds"] = strings.Join(adminDeliveriesListGetCmd_templateIds_vals, ",")
		queryParams["startDate"] = adminDeliveriesListGetCmd_startDate
		queryParams["endDate"] = adminDeliveriesListGetCmd_endDate
		queryParams["sortBy"] = adminDeliveriesListGetCmd_sortBy
		queryParams["sortDirection"] = adminDeliveriesListGetCmd_sortDirection
		queryParams["broadcastId"] = adminDeliveriesListGetCmd_broadcastId
		adminDeliveriesListGetCmd_triggeringProperties_vals, _ := cmd.Flags().GetStringArray("triggeringProperties")
		queryParams["triggeringProperties"] = strings.Join(adminDeliveriesListGetCmd_triggeringProperties_vals, ",")
		queryParams["groupId"] = adminDeliveriesListGetCmd_groupId
		adminDeliveriesListGetCmd_contextValues_vals, _ := cmd.Flags().GetStringArray("contextValues")
		queryParams["contextValues"] = strings.Join(adminDeliveriesListGetCmd_contextValues_vals, ",")
		resp, err := c.Do("GET", "/api/admin/deliveries/download", pathParams, queryParams, nil)
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
	adminDeliveriesCmd.AddCommand(adminDeliveriesListGetCmd)
	adminDeliveriesListGetCmd.Flags().StringVar(&adminDeliveriesListGetCmd_workspaceId, "workspaceId", "", "")
	adminDeliveriesListGetCmd.Flags().StringVar(&adminDeliveriesListGetCmd_fromIdentifier, "fromIdentifier", "", "")
	adminDeliveriesListGetCmd.Flags().StringVar(&adminDeliveriesListGetCmd_toIdentifier, "toIdentifier", "", "")
	adminDeliveriesListGetCmd.Flags().StringVar(&adminDeliveriesListGetCmd_journeyId, "journeyId", "", "")
	adminDeliveriesListGetCmd.Flags().StringVar(&adminDeliveriesListGetCmd_userId, "userId", "", "")
	adminDeliveriesListGetCmd.Flags().StringArray("channels", nil, "")
	adminDeliveriesListGetCmd.Flags().StringArray("from", nil, "")
	adminDeliveriesListGetCmd.Flags().StringArray("to", nil, "")
	adminDeliveriesListGetCmd.Flags().StringArray("statuses", nil, "")
	adminDeliveriesListGetCmd.Flags().StringArray("templateIds", nil, "")
	adminDeliveriesListGetCmd.Flags().StringVar(&adminDeliveriesListGetCmd_startDate, "startDate", "", "")
	adminDeliveriesListGetCmd.Flags().StringVar(&adminDeliveriesListGetCmd_endDate, "endDate", "", "")
	adminDeliveriesListGetCmd.Flags().StringVar(&adminDeliveriesListGetCmd_sortBy, "sortBy", "", "")
	adminDeliveriesListGetCmd.Flags().StringVar(&adminDeliveriesListGetCmd_sortDirection, "sortDirection", "", "")
	adminDeliveriesListGetCmd.Flags().StringVar(&adminDeliveriesListGetCmd_broadcastId, "broadcastId", "", "")
	adminDeliveriesListGetCmd.Flags().StringArray("triggeringProperties", nil, "")
	adminDeliveriesListGetCmd.Flags().StringVar(&adminDeliveriesListGetCmd_groupId, "groupId", "", "")
	adminDeliveriesListGetCmd.Flags().StringArray("contextValues", nil, "")
	adminDeliveriesListGetCmd.MarkFlagRequired("workspaceId")
}
