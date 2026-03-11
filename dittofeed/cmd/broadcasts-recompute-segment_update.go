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
	broadcastsRecomputeSegmentUpdateCmdBody string
	broadcastsRecomputeSegmentUpdateCmdBodyFile string
	broadcastsRecomputeSegmentUpdateCmd_broadcastId string
	broadcastsRecomputeSegmentUpdateCmd_workspaceId string
)

var broadcastsRecomputeSegmentUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if broadcastsRecomputeSegmentUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(broadcastsRecomputeSegmentUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			broadcastsRecomputeSegmentUpdateCmdBody = string(fileData)
		}
		if broadcastsRecomputeSegmentUpdateCmdBody != "" {
			if !json.Valid([]byte(broadcastsRecomputeSegmentUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(broadcastsRecomputeSegmentUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/broadcasts/recompute-segment", pathParams, queryParams, bodyObj)
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
		bodyMap["broadcastId"] = broadcastsRecomputeSegmentUpdateCmd_broadcastId
		bodyMap["workspaceId"] = broadcastsRecomputeSegmentUpdateCmd_workspaceId
		resp, err := c.Do("PUT", "/api/broadcasts/recompute-segment", pathParams, queryParams, bodyMap)
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
	broadcastsRecomputeSegmentCmd.AddCommand(broadcastsRecomputeSegmentUpdateCmd)
	broadcastsRecomputeSegmentUpdateCmd.Flags().StringVar(&broadcastsRecomputeSegmentUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	broadcastsRecomputeSegmentUpdateCmd.Flags().StringVar(&broadcastsRecomputeSegmentUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	broadcastsRecomputeSegmentUpdateCmd.Flags().StringVar(&broadcastsRecomputeSegmentUpdateCmd_broadcastId, "broadcastId", "", "")
	broadcastsRecomputeSegmentUpdateCmd.Flags().StringVar(&broadcastsRecomputeSegmentUpdateCmd_workspaceId, "workspaceId", "", "")
	broadcastsRecomputeSegmentUpdateCmd.MarkFlagRequired("broadcastId")
	broadcastsRecomputeSegmentUpdateCmd.MarkFlagRequired("workspaceId")
}
