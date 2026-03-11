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
	subscriptionGroupsUpdateCmdBody string
	subscriptionGroupsUpdateCmdBodyFile string
	subscriptionGroupsUpdateCmd_channel string
	subscriptionGroupsUpdateCmd_id string
	subscriptionGroupsUpdateCmd_name string
	subscriptionGroupsUpdateCmd_type string
	subscriptionGroupsUpdateCmd_workspaceId string
)

var subscriptionGroupsUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if subscriptionGroupsUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(subscriptionGroupsUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			subscriptionGroupsUpdateCmdBody = string(fileData)
		}
		if subscriptionGroupsUpdateCmdBody != "" {
			if !json.Valid([]byte(subscriptionGroupsUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(subscriptionGroupsUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/subscription-groups/", pathParams, queryParams, bodyObj)
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
		bodyMap["channel"] = subscriptionGroupsUpdateCmd_channel
		bodyMap["id"] = subscriptionGroupsUpdateCmd_id
		bodyMap["name"] = subscriptionGroupsUpdateCmd_name
		bodyMap["type"] = subscriptionGroupsUpdateCmd_type
		bodyMap["workspaceId"] = subscriptionGroupsUpdateCmd_workspaceId
		resp, err := c.Do("PUT", "/api/subscription-groups/", pathParams, queryParams, bodyMap)
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
	subscriptionGroupsCmd.AddCommand(subscriptionGroupsUpdateCmd)
	subscriptionGroupsUpdateCmd.Flags().StringVar(&subscriptionGroupsUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	subscriptionGroupsUpdateCmd.Flags().StringVar(&subscriptionGroupsUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	subscriptionGroupsUpdateCmd.Flags().StringVar(&subscriptionGroupsUpdateCmd_channel, "channel", "", "")
	subscriptionGroupsUpdateCmd.Flags().StringVar(&subscriptionGroupsUpdateCmd_id, "id", "", "")
	subscriptionGroupsUpdateCmd.Flags().StringVar(&subscriptionGroupsUpdateCmd_name, "name", "", "")
	subscriptionGroupsUpdateCmd.Flags().StringVar(&subscriptionGroupsUpdateCmd_type, "type", "", "")
	subscriptionGroupsUpdateCmd.Flags().StringVar(&subscriptionGroupsUpdateCmd_workspaceId, "workspaceId", "", "")
	subscriptionGroupsUpdateCmd.MarkFlagRequired("channel")
	subscriptionGroupsUpdateCmd.MarkFlagRequired("type")
	subscriptionGroupsUpdateCmd.MarkFlagRequired("workspaceId")
}
