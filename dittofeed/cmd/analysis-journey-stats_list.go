package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	analysisJourneyStatsListCmd_workspaceId string
	analysisJourneyStatsListCmd_journeyId string
	analysisJourneyStatsListCmd_startDate string
	analysisJourneyStatsListCmd_endDate string
)

var analysisJourneyStatsListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = analysisJourneyStatsListCmd_workspaceId
		queryParams["journeyId"] = analysisJourneyStatsListCmd_journeyId
		queryParams["startDate"] = analysisJourneyStatsListCmd_startDate
		queryParams["endDate"] = analysisJourneyStatsListCmd_endDate
		resp, err := c.Do("GET", "/api/analysis/journey-stats", pathParams, queryParams, nil)
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
	analysisJourneyStatsCmd.AddCommand(analysisJourneyStatsListCmd)
	analysisJourneyStatsListCmd.Flags().StringVar(&analysisJourneyStatsListCmd_workspaceId, "workspaceId", "", "")
	analysisJourneyStatsListCmd.Flags().StringVar(&analysisJourneyStatsListCmd_journeyId, "journeyId", "", "")
	analysisJourneyStatsListCmd.Flags().StringVar(&analysisJourneyStatsListCmd_startDate, "startDate", "", "")
	analysisJourneyStatsListCmd.Flags().StringVar(&analysisJourneyStatsListCmd_endDate, "endDate", "", "")
	analysisJourneyStatsListCmd.MarkFlagRequired("workspaceId")
	analysisJourneyStatsListCmd.MarkFlagRequired("journeyId")
	analysisJourneyStatsListCmd.MarkFlagRequired("startDate")
	analysisJourneyStatsListCmd.MarkFlagRequired("endDate")
}
