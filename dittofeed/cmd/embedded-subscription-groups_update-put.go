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
	embeddedSubscriptionGroupsUpdatePutCmdBody string
	embeddedSubscriptionGroupsUpdatePutCmdBodyFile string
	embeddedSubscriptionGroupsUpdatePutCmd_userUpdates []string
	embeddedSubscriptionGroupsUpdatePutCmd_workspaceId string
)

var embeddedSubscriptionGroupsUpdatePutCmd = &cobra.Command{
	Use: "update-put",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if embeddedSubscriptionGroupsUpdatePutCmdBodyFile != "" {
			fileData, err := os.ReadFile(embeddedSubscriptionGroupsUpdatePutCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			embeddedSubscriptionGroupsUpdatePutCmdBody = string(fileData)
		}
		if embeddedSubscriptionGroupsUpdatePutCmdBody != "" {
			if !json.Valid([]byte(embeddedSubscriptionGroupsUpdatePutCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(embeddedSubscriptionGroupsUpdatePutCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api-l/embedded/subscription-groups/assignments", pathParams, queryParams, bodyObj)
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
		bodyMap["userUpdates"] = embeddedSubscriptionGroupsUpdatePutCmd_userUpdates
		bodyMap["workspaceId"] = embeddedSubscriptionGroupsUpdatePutCmd_workspaceId
		resp, err := c.Do("PUT", "/api-l/embedded/subscription-groups/assignments", pathParams, queryParams, bodyMap)
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
	embeddedSubscriptionGroupsCmd.AddCommand(embeddedSubscriptionGroupsUpdatePutCmd)
	embeddedSubscriptionGroupsUpdatePutCmd.Flags().StringVar(&embeddedSubscriptionGroupsUpdatePutCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	embeddedSubscriptionGroupsUpdatePutCmd.Flags().StringVar(&embeddedSubscriptionGroupsUpdatePutCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	embeddedSubscriptionGroupsUpdatePutCmd.Flags().StringArrayVar(&embeddedSubscriptionGroupsUpdatePutCmd_userUpdates, "userUpdates", nil, "")
	embeddedSubscriptionGroupsUpdatePutCmd.Flags().StringVar(&embeddedSubscriptionGroupsUpdatePutCmd_workspaceId, "workspaceId", "", "")
	embeddedSubscriptionGroupsUpdatePutCmd.MarkFlagRequired("userUpdates")
	embeddedSubscriptionGroupsUpdatePutCmd.MarkFlagRequired("workspaceId")
}
