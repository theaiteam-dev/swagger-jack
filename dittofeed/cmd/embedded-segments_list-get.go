package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	embeddedSegmentsListGetCmd_workspaceId string
	embeddedSegmentsListGetCmd_segmentId string
)

var embeddedSegmentsListGetCmd = &cobra.Command{
	Use: "list-get",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		queryParams["workspaceId"] = embeddedSegmentsListGetCmd_workspaceId
		queryParams["segmentId"] = embeddedSegmentsListGetCmd_segmentId
		resp, err := c.Do("GET", "/api-l/embedded/segments/manual-segment/status", pathParams, queryParams, nil)
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
	embeddedSegmentsCmd.AddCommand(embeddedSegmentsListGetCmd)
	embeddedSegmentsListGetCmd.Flags().StringVar(&embeddedSegmentsListGetCmd_workspaceId, "workspaceId", "", "")
	embeddedSegmentsListGetCmd.Flags().StringVar(&embeddedSegmentsListGetCmd_segmentId, "segmentId", "", "")
	embeddedSegmentsListGetCmd.MarkFlagRequired("workspaceId")
	embeddedSegmentsListGetCmd.MarkFlagRequired("segmentId")
}
