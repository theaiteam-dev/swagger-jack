package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	analysisChartDataListCmd_workspaceId string
	analysisChartDataListCmd_startDate string
	analysisChartDataListCmd_endDate string
	analysisChartDataListCmd_granularity string
	analysisChartDataListCmd_groupBy string
	analysisChartDataListCmd_filters string
)

var analysisChartDataListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = analysisChartDataListCmd_workspaceId
		queryParams["startDate"] = analysisChartDataListCmd_startDate
		queryParams["endDate"] = analysisChartDataListCmd_endDate
		queryParams["granularity"] = analysisChartDataListCmd_granularity
		queryParams["groupBy"] = analysisChartDataListCmd_groupBy
		queryParams["filters"] = analysisChartDataListCmd_filters
		resp, err := c.Do("GET", "/api/analysis/chart-data", pathParams, queryParams, nil)
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
	analysisChartDataCmd.AddCommand(analysisChartDataListCmd)
	analysisChartDataListCmd.Flags().StringVar(&analysisChartDataListCmd_workspaceId, "workspaceId", "", "")
	analysisChartDataListCmd.Flags().StringVar(&analysisChartDataListCmd_startDate, "startDate", "", "")
	analysisChartDataListCmd.Flags().StringVar(&analysisChartDataListCmd_endDate, "endDate", "", "")
	analysisChartDataListCmd.Flags().StringVar(&analysisChartDataListCmd_granularity, "granularity", "", "")
	analysisChartDataListCmd.Flags().StringVar(&analysisChartDataListCmd_groupBy, "groupBy", "", "")
	analysisChartDataListCmd.Flags().StringVar(&analysisChartDataListCmd_filters, "filters", "", "")
	analysisChartDataListCmd.MarkFlagRequired("workspaceId")
	analysisChartDataListCmd.MarkFlagRequired("startDate")
	analysisChartDataListCmd.MarkFlagRequired("endDate")
}
