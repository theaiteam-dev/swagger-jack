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
	userPropertiesStatusUpdateCmdBody string
	userPropertiesStatusUpdateCmdBodyFile string
	userPropertiesStatusUpdateCmd_status string
	userPropertiesStatusUpdateCmd_workspaceId string
	userPropertiesStatusUpdateCmd_id string
)

var userPropertiesStatusUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if userPropertiesStatusUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(userPropertiesStatusUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			userPropertiesStatusUpdateCmdBody = string(fileData)
		}
		if userPropertiesStatusUpdateCmdBody != "" {
			if !json.Valid([]byte(userPropertiesStatusUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(userPropertiesStatusUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PATCH", "/api/user-properties/status", pathParams, queryParams, bodyObj)
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
		bodyMap["status"] = userPropertiesStatusUpdateCmd_status
		bodyMap["workspaceId"] = userPropertiesStatusUpdateCmd_workspaceId
		bodyMap["id"] = userPropertiesStatusUpdateCmd_id
		resp, err := c.Do("PATCH", "/api/user-properties/status", pathParams, queryParams, bodyMap)
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
	userPropertiesStatusCmd.AddCommand(userPropertiesStatusUpdateCmd)
	userPropertiesStatusUpdateCmd.Flags().StringVar(&userPropertiesStatusUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	userPropertiesStatusUpdateCmd.Flags().StringVar(&userPropertiesStatusUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	userPropertiesStatusUpdateCmd.Flags().StringVar(&userPropertiesStatusUpdateCmd_status, "status", "", "")
	userPropertiesStatusUpdateCmd.Flags().StringVar(&userPropertiesStatusUpdateCmd_workspaceId, "workspaceId", "", "")
	userPropertiesStatusUpdateCmd.Flags().StringVar(&userPropertiesStatusUpdateCmd_id, "id", "", "")
	userPropertiesStatusUpdateCmd.MarkFlagRequired("status")
	userPropertiesStatusUpdateCmd.MarkFlagRequired("workspaceId")
	userPropertiesStatusUpdateCmd.MarkFlagRequired("id")
}
