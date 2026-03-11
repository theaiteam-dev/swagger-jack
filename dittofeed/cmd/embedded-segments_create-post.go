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
	embeddedSegmentsCreatePostCmdBody string
	embeddedSegmentsCreatePostCmdBodyFile string
	embeddedSegmentsCreatePostCmd_segmentId string
	embeddedSegmentsCreatePostCmd_sync bool
	embeddedSegmentsCreatePostCmd_userIds []string
	embeddedSegmentsCreatePostCmd_workspaceId string
	embeddedSegmentsCreatePostCmd_append bool
)

var embeddedSegmentsCreatePostCmd = &cobra.Command{
	Use: "create-post",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if embeddedSegmentsCreatePostCmdBodyFile != "" {
			fileData, err := os.ReadFile(embeddedSegmentsCreatePostCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			embeddedSegmentsCreatePostCmdBody = string(fileData)
		}
		if embeddedSegmentsCreatePostCmdBody != "" {
			if !json.Valid([]byte(embeddedSegmentsCreatePostCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(embeddedSegmentsCreatePostCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api-l/embedded/segments/manual-segment/update", pathParams, queryParams, bodyObj)
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
		bodyMap["segmentId"] = embeddedSegmentsCreatePostCmd_segmentId
		bodyMap["sync"] = embeddedSegmentsCreatePostCmd_sync
		bodyMap["userIds"] = embeddedSegmentsCreatePostCmd_userIds
		bodyMap["workspaceId"] = embeddedSegmentsCreatePostCmd_workspaceId
		bodyMap["append"] = embeddedSegmentsCreatePostCmd_append
		resp, err := c.Do("POST", "/api-l/embedded/segments/manual-segment/update", pathParams, queryParams, bodyMap)
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
	embeddedSegmentsCmd.AddCommand(embeddedSegmentsCreatePostCmd)
	embeddedSegmentsCreatePostCmd.Flags().StringVar(&embeddedSegmentsCreatePostCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	embeddedSegmentsCreatePostCmd.Flags().StringVar(&embeddedSegmentsCreatePostCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	embeddedSegmentsCreatePostCmd.Flags().StringVar(&embeddedSegmentsCreatePostCmd_segmentId, "segmentId", "", "")
	embeddedSegmentsCreatePostCmd.Flags().BoolVar(&embeddedSegmentsCreatePostCmd_sync, "sync", false, "")
	embeddedSegmentsCreatePostCmd.Flags().StringArrayVar(&embeddedSegmentsCreatePostCmd_userIds, "userIds", nil, "")
	embeddedSegmentsCreatePostCmd.Flags().StringVar(&embeddedSegmentsCreatePostCmd_workspaceId, "workspaceId", "", "")
	embeddedSegmentsCreatePostCmd.Flags().BoolVar(&embeddedSegmentsCreatePostCmd_append, "append", false, "")
	embeddedSegmentsCreatePostCmd.MarkFlagRequired("segmentId")
	embeddedSegmentsCreatePostCmd.MarkFlagRequired("userIds")
	embeddedSegmentsCreatePostCmd.MarkFlagRequired("workspaceId")
}
