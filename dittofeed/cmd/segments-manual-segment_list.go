package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	segmentsManualSegmentListCmd_workspaceId string
	segmentsManualSegmentListCmd_segmentId string
)

var segmentsManualSegmentListCmd = &cobra.Command{
	Use: "list",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = segmentsManualSegmentListCmd_workspaceId
		queryParams["segmentId"] = segmentsManualSegmentListCmd_segmentId
		resp, err := c.Do("GET", "/api/segments/manual-segment/status", pathParams, queryParams, nil)
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
	segmentsManualSegmentCmd.AddCommand(segmentsManualSegmentListCmd)
	segmentsManualSegmentListCmd.Flags().StringVar(&segmentsManualSegmentListCmd_workspaceId, "workspaceId", "", "")
	segmentsManualSegmentListCmd.Flags().StringVar(&segmentsManualSegmentListCmd_segmentId, "segmentId", "", "")
	segmentsManualSegmentListCmd.MarkFlagRequired("workspaceId")
	segmentsManualSegmentListCmd.MarkFlagRequired("segmentId")
}
