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
	adminBroadcastsUpdatePutCmdBody string
	adminBroadcastsUpdatePutCmdBodyFile string
	adminBroadcastsUpdatePutCmd_id string
	adminBroadcastsUpdatePutCmd_name string
	adminBroadcastsUpdatePutCmd_workspaceId string
)

var adminBroadcastsUpdatePutCmd = &cobra.Command{
	Use: "update-put",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if adminBroadcastsUpdatePutCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminBroadcastsUpdatePutCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminBroadcastsUpdatePutCmdBody = string(fileData)
		}
		if adminBroadcastsUpdatePutCmdBody != "" {
			if !json.Valid([]byte(adminBroadcastsUpdatePutCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminBroadcastsUpdatePutCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/admin/broadcasts/", pathParams, queryParams, bodyObj)
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
		bodyMap["id"] = adminBroadcastsUpdatePutCmd_id
		bodyMap["name"] = adminBroadcastsUpdatePutCmd_name
		bodyMap["workspaceId"] = adminBroadcastsUpdatePutCmd_workspaceId
		resp, err := c.Do("PUT", "/api/admin/broadcasts/", pathParams, queryParams, bodyMap)
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
	adminBroadcastsCmd.AddCommand(adminBroadcastsUpdatePutCmd)
	adminBroadcastsUpdatePutCmd.Flags().StringVar(&adminBroadcastsUpdatePutCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminBroadcastsUpdatePutCmd.Flags().StringVar(&adminBroadcastsUpdatePutCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	adminBroadcastsUpdatePutCmd.Flags().StringVar(&adminBroadcastsUpdatePutCmd_id, "id", "", "")
	adminBroadcastsUpdatePutCmd.Flags().StringVar(&adminBroadcastsUpdatePutCmd_name, "name", "", "")
	adminBroadcastsUpdatePutCmd.Flags().StringVar(&adminBroadcastsUpdatePutCmd_workspaceId, "workspaceId", "", "")
	adminBroadcastsUpdatePutCmd.MarkFlagRequired("id")
	adminBroadcastsUpdatePutCmd.MarkFlagRequired("workspaceId")
}
