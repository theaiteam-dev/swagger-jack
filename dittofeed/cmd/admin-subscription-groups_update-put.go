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
	adminSubscriptionGroupsUpdatePutCmdBody string
	adminSubscriptionGroupsUpdatePutCmdBodyFile string
	adminSubscriptionGroupsUpdatePutCmd_type string
	adminSubscriptionGroupsUpdatePutCmd_workspaceId string
	adminSubscriptionGroupsUpdatePutCmd_channel string
	adminSubscriptionGroupsUpdatePutCmd_id string
	adminSubscriptionGroupsUpdatePutCmd_name string
)

var adminSubscriptionGroupsUpdatePutCmd = &cobra.Command{
	Use: "update-put",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if adminSubscriptionGroupsUpdatePutCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminSubscriptionGroupsUpdatePutCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminSubscriptionGroupsUpdatePutCmdBody = string(fileData)
		}
		if adminSubscriptionGroupsUpdatePutCmdBody != "" {
			if !json.Valid([]byte(adminSubscriptionGroupsUpdatePutCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminSubscriptionGroupsUpdatePutCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/admin/subscription-groups/", pathParams, queryParams, bodyObj)
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
		bodyMap["type"] = adminSubscriptionGroupsUpdatePutCmd_type
		bodyMap["workspaceId"] = adminSubscriptionGroupsUpdatePutCmd_workspaceId
		bodyMap["channel"] = adminSubscriptionGroupsUpdatePutCmd_channel
		bodyMap["id"] = adminSubscriptionGroupsUpdatePutCmd_id
		bodyMap["name"] = adminSubscriptionGroupsUpdatePutCmd_name
		resp, err := c.Do("PUT", "/api/admin/subscription-groups/", pathParams, queryParams, bodyMap)
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
	adminSubscriptionGroupsCmd.AddCommand(adminSubscriptionGroupsUpdatePutCmd)
	adminSubscriptionGroupsUpdatePutCmd.Flags().StringVar(&adminSubscriptionGroupsUpdatePutCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminSubscriptionGroupsUpdatePutCmd.Flags().StringVar(&adminSubscriptionGroupsUpdatePutCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	adminSubscriptionGroupsUpdatePutCmd.Flags().StringVar(&adminSubscriptionGroupsUpdatePutCmd_type, "type", "", "")
	adminSubscriptionGroupsUpdatePutCmd.Flags().StringVar(&adminSubscriptionGroupsUpdatePutCmd_workspaceId, "workspaceId", "", "")
	adminSubscriptionGroupsUpdatePutCmd.Flags().StringVar(&adminSubscriptionGroupsUpdatePutCmd_channel, "channel", "", "")
	adminSubscriptionGroupsUpdatePutCmd.Flags().StringVar(&adminSubscriptionGroupsUpdatePutCmd_id, "id", "", "")
	adminSubscriptionGroupsUpdatePutCmd.Flags().StringVar(&adminSubscriptionGroupsUpdatePutCmd_name, "name", "", "")
	adminSubscriptionGroupsUpdatePutCmd.MarkFlagRequired("type")
	adminSubscriptionGroupsUpdatePutCmd.MarkFlagRequired("workspaceId")
	adminSubscriptionGroupsUpdatePutCmd.MarkFlagRequired("channel")
}
