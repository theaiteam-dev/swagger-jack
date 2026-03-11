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
	deliveriesDownloadListCmd_workspaceId string
	deliveriesDownloadListCmd_fromIdentifier string
	deliveriesDownloadListCmd_toIdentifier string
	deliveriesDownloadListCmd_journeyId string
	deliveriesDownloadListCmd_userId string
	deliveriesDownloadListCmd_startDate string
	deliveriesDownloadListCmd_endDate string
	deliveriesDownloadListCmd_sortBy string
	deliveriesDownloadListCmd_sortDirection string
	deliveriesDownloadListCmd_broadcastId string
	deliveriesDownloadListCmd_groupId string
)

var deliveriesDownloadListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = deliveriesDownloadListCmd_workspaceId
		queryParams["fromIdentifier"] = deliveriesDownloadListCmd_fromIdentifier
		queryParams["toIdentifier"] = deliveriesDownloadListCmd_toIdentifier
		queryParams["journeyId"] = deliveriesDownloadListCmd_journeyId
		queryParams["userId"] = deliveriesDownloadListCmd_userId
		deliveriesDownloadListCmd_channels_vals, _ := cmd.Flags().GetStringArray("channels")
		queryParams["channels"] = strings.Join(deliveriesDownloadListCmd_channels_vals, ",")
		deliveriesDownloadListCmd_from_vals, _ := cmd.Flags().GetStringArray("from")
		queryParams["from"] = strings.Join(deliveriesDownloadListCmd_from_vals, ",")
		deliveriesDownloadListCmd_to_vals, _ := cmd.Flags().GetStringArray("to")
		queryParams["to"] = strings.Join(deliveriesDownloadListCmd_to_vals, ",")
		deliveriesDownloadListCmd_statuses_vals, _ := cmd.Flags().GetStringArray("statuses")
		queryParams["statuses"] = strings.Join(deliveriesDownloadListCmd_statuses_vals, ",")
		deliveriesDownloadListCmd_templateIds_vals, _ := cmd.Flags().GetStringArray("templateIds")
		queryParams["templateIds"] = strings.Join(deliveriesDownloadListCmd_templateIds_vals, ",")
		queryParams["startDate"] = deliveriesDownloadListCmd_startDate
		queryParams["endDate"] = deliveriesDownloadListCmd_endDate
		queryParams["sortBy"] = deliveriesDownloadListCmd_sortBy
		queryParams["sortDirection"] = deliveriesDownloadListCmd_sortDirection
		queryParams["broadcastId"] = deliveriesDownloadListCmd_broadcastId
		deliveriesDownloadListCmd_triggeringProperties_vals, _ := cmd.Flags().GetStringArray("triggeringProperties")
		queryParams["triggeringProperties"] = strings.Join(deliveriesDownloadListCmd_triggeringProperties_vals, ",")
		queryParams["groupId"] = deliveriesDownloadListCmd_groupId
		deliveriesDownloadListCmd_contextValues_vals, _ := cmd.Flags().GetStringArray("contextValues")
		queryParams["contextValues"] = strings.Join(deliveriesDownloadListCmd_contextValues_vals, ",")
		resp, err := c.Do("GET", "/api/deliveries/download", pathParams, queryParams, nil)
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
	deliveriesDownloadCmd.AddCommand(deliveriesDownloadListCmd)
	deliveriesDownloadListCmd.Flags().StringVar(&deliveriesDownloadListCmd_workspaceId, "workspaceId", "", "")
	deliveriesDownloadListCmd.Flags().StringVar(&deliveriesDownloadListCmd_fromIdentifier, "fromIdentifier", "", "")
	deliveriesDownloadListCmd.Flags().StringVar(&deliveriesDownloadListCmd_toIdentifier, "toIdentifier", "", "")
	deliveriesDownloadListCmd.Flags().StringVar(&deliveriesDownloadListCmd_journeyId, "journeyId", "", "")
	deliveriesDownloadListCmd.Flags().StringVar(&deliveriesDownloadListCmd_userId, "userId", "", "")
	deliveriesDownloadListCmd.Flags().StringArray("channels", nil, "")
	deliveriesDownloadListCmd.Flags().StringArray("from", nil, "")
	deliveriesDownloadListCmd.Flags().StringArray("to", nil, "")
	deliveriesDownloadListCmd.Flags().StringArray("statuses", nil, "")
	deliveriesDownloadListCmd.Flags().StringArray("templateIds", nil, "")
	deliveriesDownloadListCmd.Flags().StringVar(&deliveriesDownloadListCmd_startDate, "startDate", "", "")
	deliveriesDownloadListCmd.Flags().StringVar(&deliveriesDownloadListCmd_endDate, "endDate", "", "")
	deliveriesDownloadListCmd.Flags().StringVar(&deliveriesDownloadListCmd_sortBy, "sortBy", "", "")
	deliveriesDownloadListCmd.Flags().StringVar(&deliveriesDownloadListCmd_sortDirection, "sortDirection", "", "")
	deliveriesDownloadListCmd.Flags().StringVar(&deliveriesDownloadListCmd_broadcastId, "broadcastId", "", "")
	deliveriesDownloadListCmd.Flags().StringArray("triggeringProperties", nil, "")
	deliveriesDownloadListCmd.Flags().StringVar(&deliveriesDownloadListCmd_groupId, "groupId", "", "")
	deliveriesDownloadListCmd.Flags().StringArray("contextValues", nil, "")
	deliveriesDownloadListCmd.MarkFlagRequired("workspaceId")
}
