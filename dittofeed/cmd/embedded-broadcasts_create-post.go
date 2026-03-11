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
	embeddedBroadcastsCreatePostCmdBody string
	embeddedBroadcastsCreatePostCmdBodyFile string
	embeddedBroadcastsCreatePostCmd_broadcastId string
	embeddedBroadcastsCreatePostCmd_workspaceId string
)

var embeddedBroadcastsCreatePostCmd = &cobra.Command{
	Use: "create-post",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if embeddedBroadcastsCreatePostCmdBodyFile != "" {
			fileData, err := os.ReadFile(embeddedBroadcastsCreatePostCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			embeddedBroadcastsCreatePostCmdBody = string(fileData)
		}
		if embeddedBroadcastsCreatePostCmdBody != "" {
			if !json.Valid([]byte(embeddedBroadcastsCreatePostCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(embeddedBroadcastsCreatePostCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api-l/embedded/broadcasts/start", pathParams, queryParams, bodyObj)
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
		bodyMap["broadcastId"] = embeddedBroadcastsCreatePostCmd_broadcastId
		bodyMap["workspaceId"] = embeddedBroadcastsCreatePostCmd_workspaceId
		resp, err := c.Do("POST", "/api-l/embedded/broadcasts/start", pathParams, queryParams, bodyMap)
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
	embeddedBroadcastsCmd.AddCommand(embeddedBroadcastsCreatePostCmd)
	embeddedBroadcastsCreatePostCmd.Flags().StringVar(&embeddedBroadcastsCreatePostCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	embeddedBroadcastsCreatePostCmd.Flags().StringVar(&embeddedBroadcastsCreatePostCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	embeddedBroadcastsCreatePostCmd.Flags().StringVar(&embeddedBroadcastsCreatePostCmd_broadcastId, "broadcastId", "", "")
	embeddedBroadcastsCreatePostCmd.Flags().StringVar(&embeddedBroadcastsCreatePostCmd_workspaceId, "workspaceId", "", "")
	embeddedBroadcastsCreatePostCmd.MarkFlagRequired("broadcastId")
	embeddedBroadcastsCreatePostCmd.MarkFlagRequired("workspaceId")
}
