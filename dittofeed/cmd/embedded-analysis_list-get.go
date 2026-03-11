package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	embeddedAnalysisListGetCmd_workspaceId string
	embeddedAnalysisListGetCmd_startDate string
	embeddedAnalysisListGetCmd_endDate string
	embeddedAnalysisListGetCmd_filters string
)

var embeddedAnalysisListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = embeddedAnalysisListGetCmd_workspaceId
		queryParams["startDate"] = embeddedAnalysisListGetCmd_startDate
		queryParams["endDate"] = embeddedAnalysisListGetCmd_endDate
		queryParams["filters"] = embeddedAnalysisListGetCmd_filters
		resp, err := c.Do("GET", "/api-l/embedded/analysis/summary", pathParams, queryParams, nil)
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
	embeddedAnalysisCmd.AddCommand(embeddedAnalysisListGetCmd)
	embeddedAnalysisListGetCmd.Flags().StringVar(&embeddedAnalysisListGetCmd_workspaceId, "workspaceId", "", "")
	embeddedAnalysisListGetCmd.Flags().StringVar(&embeddedAnalysisListGetCmd_startDate, "startDate", "", "")
	embeddedAnalysisListGetCmd.Flags().StringVar(&embeddedAnalysisListGetCmd_endDate, "endDate", "", "")
	embeddedAnalysisListGetCmd.Flags().StringVar(&embeddedAnalysisListGetCmd_filters, "filters", "", "")
	embeddedAnalysisListGetCmd.MarkFlagRequired("workspaceId")
	embeddedAnalysisListGetCmd.MarkFlagRequired("startDate")
	embeddedAnalysisListGetCmd.MarkFlagRequired("endDate")
}
