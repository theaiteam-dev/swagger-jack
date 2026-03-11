package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	adminAnalysisListGetCmd_workspaceId string
	adminAnalysisListGetCmd_startDate string
	adminAnalysisListGetCmd_endDate string
	adminAnalysisListGetCmd_filters string
)

var adminAnalysisListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = adminAnalysisListGetCmd_workspaceId
		queryParams["startDate"] = adminAnalysisListGetCmd_startDate
		queryParams["endDate"] = adminAnalysisListGetCmd_endDate
		queryParams["filters"] = adminAnalysisListGetCmd_filters
		resp, err := c.Do("GET", "/api/admin/analysis/summary", pathParams, queryParams, nil)
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
	adminAnalysisCmd.AddCommand(adminAnalysisListGetCmd)
	adminAnalysisListGetCmd.Flags().StringVar(&adminAnalysisListGetCmd_workspaceId, "workspaceId", "", "")
	adminAnalysisListGetCmd.Flags().StringVar(&adminAnalysisListGetCmd_startDate, "startDate", "", "")
	adminAnalysisListGetCmd.Flags().StringVar(&adminAnalysisListGetCmd_endDate, "endDate", "", "")
	adminAnalysisListGetCmd.Flags().StringVar(&adminAnalysisListGetCmd_filters, "filters", "", "")
	adminAnalysisListGetCmd.MarkFlagRequired("workspaceId")
	adminAnalysisListGetCmd.MarkFlagRequired("startDate")
	adminAnalysisListGetCmd.MarkFlagRequired("endDate")
}
