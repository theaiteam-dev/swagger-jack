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
	adminBroadcastsCreatePostCmdBody string
	adminBroadcastsCreatePostCmdBodyFile string
	adminBroadcastsCreatePostCmd_broadcastId string
	adminBroadcastsCreatePostCmd_workspaceId string
)

var adminBroadcastsCreatePostCmd = &cobra.Command{
	Use: "create-post",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if adminBroadcastsCreatePostCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminBroadcastsCreatePostCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminBroadcastsCreatePostCmdBody = string(fileData)
		}
		if adminBroadcastsCreatePostCmdBody != "" {
			if !json.Valid([]byte(adminBroadcastsCreatePostCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminBroadcastsCreatePostCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/admin/broadcasts/cancel", pathParams, queryParams, bodyObj)
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
		bodyMap["broadcastId"] = adminBroadcastsCreatePostCmd_broadcastId
		bodyMap["workspaceId"] = adminBroadcastsCreatePostCmd_workspaceId
		resp, err := c.Do("POST", "/api/admin/broadcasts/cancel", pathParams, queryParams, bodyMap)
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
	adminBroadcastsCmd.AddCommand(adminBroadcastsCreatePostCmd)
	adminBroadcastsCreatePostCmd.Flags().StringVar(&adminBroadcastsCreatePostCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminBroadcastsCreatePostCmd.Flags().StringVar(&adminBroadcastsCreatePostCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	adminBroadcastsCreatePostCmd.Flags().StringVar(&adminBroadcastsCreatePostCmd_broadcastId, "broadcastId", "", "")
	adminBroadcastsCreatePostCmd.Flags().StringVar(&adminBroadcastsCreatePostCmd_workspaceId, "workspaceId", "", "")
	adminBroadcastsCreatePostCmd.MarkFlagRequired("broadcastId")
	adminBroadcastsCreatePostCmd.MarkFlagRequired("workspaceId")
}
