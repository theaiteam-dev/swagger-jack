package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	analysisSummaryListCmd_workspaceId string
	analysisSummaryListCmd_startDate string
	analysisSummaryListCmd_endDate string
	analysisSummaryListCmd_filters string
)

var analysisSummaryListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = analysisSummaryListCmd_workspaceId
		queryParams["startDate"] = analysisSummaryListCmd_startDate
		queryParams["endDate"] = analysisSummaryListCmd_endDate
		queryParams["filters"] = analysisSummaryListCmd_filters
		resp, err := c.Do("GET", "/api/analysis/summary", pathParams, queryParams, nil)
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
	analysisSummaryCmd.AddCommand(analysisSummaryListCmd)
	analysisSummaryListCmd.Flags().StringVar(&analysisSummaryListCmd_workspaceId, "workspaceId", "", "")
	analysisSummaryListCmd.Flags().StringVar(&analysisSummaryListCmd_startDate, "startDate", "", "")
	analysisSummaryListCmd.Flags().StringVar(&analysisSummaryListCmd_endDate, "endDate", "", "")
	analysisSummaryListCmd.Flags().StringVar(&analysisSummaryListCmd_filters, "filters", "", "")
	analysisSummaryListCmd.MarkFlagRequired("workspaceId")
	analysisSummaryListCmd.MarkFlagRequired("startDate")
	analysisSummaryListCmd.MarkFlagRequired("endDate")
}
