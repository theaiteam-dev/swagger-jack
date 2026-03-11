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
	adminSegmentsCreatePostCmdBody string
	adminSegmentsCreatePostCmdBodyFile string
	adminSegmentsCreatePostCmd_userIds []string
	adminSegmentsCreatePostCmd_workspaceId string
	adminSegmentsCreatePostCmd_append bool
	adminSegmentsCreatePostCmd_segmentId string
	adminSegmentsCreatePostCmd_sync bool
)

var adminSegmentsCreatePostCmd = &cobra.Command{
	Use: "create-post",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if adminSegmentsCreatePostCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminSegmentsCreatePostCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminSegmentsCreatePostCmdBody = string(fileData)
		}
		if adminSegmentsCreatePostCmdBody != "" {
			if !json.Valid([]byte(adminSegmentsCreatePostCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminSegmentsCreatePostCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/admin/segments/manual-segment/update", pathParams, queryParams, bodyObj)
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
		bodyMap["userIds"] = adminSegmentsCreatePostCmd_userIds
		bodyMap["workspaceId"] = adminSegmentsCreatePostCmd_workspaceId
		bodyMap["append"] = adminSegmentsCreatePostCmd_append
		bodyMap["segmentId"] = adminSegmentsCreatePostCmd_segmentId
		bodyMap["sync"] = adminSegmentsCreatePostCmd_sync
		resp, err := c.Do("POST", "/api/admin/segments/manual-segment/update", pathParams, queryParams, bodyMap)
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
	adminSegmentsCmd.AddCommand(adminSegmentsCreatePostCmd)
	adminSegmentsCreatePostCmd.Flags().StringVar(&adminSegmentsCreatePostCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminSegmentsCreatePostCmd.Flags().StringVar(&adminSegmentsCreatePostCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	adminSegmentsCreatePostCmd.Flags().StringArrayVar(&adminSegmentsCreatePostCmd_userIds, "userIds", nil, "")
	adminSegmentsCreatePostCmd.Flags().StringVar(&adminSegmentsCreatePostCmd_workspaceId, "workspaceId", "", "")
	adminSegmentsCreatePostCmd.Flags().BoolVar(&adminSegmentsCreatePostCmd_append, "append", false, "")
	adminSegmentsCreatePostCmd.Flags().StringVar(&adminSegmentsCreatePostCmd_segmentId, "segmentId", "", "")
	adminSegmentsCreatePostCmd.Flags().BoolVar(&adminSegmentsCreatePostCmd_sync, "sync", false, "")
	adminSegmentsCreatePostCmd.MarkFlagRequired("userIds")
	adminSegmentsCreatePostCmd.MarkFlagRequired("workspaceId")
	adminSegmentsCreatePostCmd.MarkFlagRequired("segmentId")
}
