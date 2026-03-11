package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"dittofeed/internal/client"
	"dittofeed/internal/output"
)

var (
	segmentsManualSegmentCreatePostCmdBody string
	segmentsManualSegmentCreatePostCmdBodyFile string
	segmentsManualSegmentCreatePostCmd_append bool
	segmentsManualSegmentCreatePostCmd_segmentId string
	segmentsManualSegmentCreatePostCmd_sync bool
	segmentsManualSegmentCreatePostCmd_userIds []string
	segmentsManualSegmentCreatePostCmd_workspaceId string
)

var segmentsManualSegmentCreatePostCmd = &cobra.Command{
	Use: "create-post",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if segmentsManualSegmentCreatePostCmdBodyFile != "" {
			fileData, err := os.ReadFile(segmentsManualSegmentCreatePostCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			segmentsManualSegmentCreatePostCmdBody = string(fileData)
		}
		if segmentsManualSegmentCreatePostCmdBody != "" {
			if !json.Valid([]byte(segmentsManualSegmentCreatePostCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(segmentsManualSegmentCreatePostCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/segments/manual-segment/update", pathParams, queryParams, bodyObj)
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
		}
		bodyMap := map[string]interface{}{}
		bodyMap["append"] = segmentsManualSegmentCreatePostCmd_append
		bodyMap["segmentId"] = segmentsManualSegmentCreatePostCmd_segmentId
		bodyMap["sync"] = segmentsManualSegmentCreatePostCmd_sync
		bodyMap["userIds"] = segmentsManualSegmentCreatePostCmd_userIds
		bodyMap["workspaceId"] = segmentsManualSegmentCreatePostCmd_workspaceId
		resp, err := c.Do("POST", "/api/segments/manual-segment/update", pathParams, queryParams, bodyMap)
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
	segmentsManualSegmentCmd.AddCommand(segmentsManualSegmentCreatePostCmd)
	segmentsManualSegmentCreatePostCmd.Flags().StringVar(&segmentsManualSegmentCreatePostCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	segmentsManualSegmentCreatePostCmd.Flags().StringVar(&segmentsManualSegmentCreatePostCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	segmentsManualSegmentCreatePostCmd.Flags().BoolVar(&segmentsManualSegmentCreatePostCmd_append, "append", false, "")
	segmentsManualSegmentCreatePostCmd.Flags().StringVar(&segmentsManualSegmentCreatePostCmd_segmentId, "segmentId", "", "")
	segmentsManualSegmentCreatePostCmd.Flags().BoolVar(&segmentsManualSegmentCreatePostCmd_sync, "sync", false, "")
	segmentsManualSegmentCreatePostCmd.Flags().StringArrayVar(&segmentsManualSegmentCreatePostCmd_userIds, "userIds", nil, "")
	segmentsManualSegmentCreatePostCmd.Flags().StringVar(&segmentsManualSegmentCreatePostCmd_workspaceId, "workspaceId", "", "")
	segmentsManualSegmentCreatePostCmd.MarkFlagRequired("segmentId")
	segmentsManualSegmentCreatePostCmd.MarkFlagRequired("userIds")
	segmentsManualSegmentCreatePostCmd.MarkFlagRequired("workspaceId")
}
