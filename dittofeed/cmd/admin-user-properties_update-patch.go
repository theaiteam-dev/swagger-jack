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
	adminUserPropertiesUpdatePatchCmdBody string
	adminUserPropertiesUpdatePatchCmdBodyFile string
	adminUserPropertiesUpdatePatchCmd_id string
	adminUserPropertiesUpdatePatchCmd_status string
	adminUserPropertiesUpdatePatchCmd_workspaceId string
)

var adminUserPropertiesUpdatePatchCmd = &cobra.Command{
	Use: "update-patch",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if adminUserPropertiesUpdatePatchCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminUserPropertiesUpdatePatchCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminUserPropertiesUpdatePatchCmdBody = string(fileData)
		}
		if adminUserPropertiesUpdatePatchCmdBody != "" {
			if !json.Valid([]byte(adminUserPropertiesUpdatePatchCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminUserPropertiesUpdatePatchCmdBody), &bodyObj)
			resp, err := c.Do("PATCH", "/api/admin/user-properties/status", pathParams, queryParams, bodyObj)
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
		bodyMap["id"] = adminUserPropertiesUpdatePatchCmd_id
		bodyMap["status"] = adminUserPropertiesUpdatePatchCmd_status
		bodyMap["workspaceId"] = adminUserPropertiesUpdatePatchCmd_workspaceId
		resp, err := c.Do("PATCH", "/api/admin/user-properties/status", pathParams, queryParams, bodyMap)
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
	adminUserPropertiesCmd.AddCommand(adminUserPropertiesUpdatePatchCmd)
	adminUserPropertiesUpdatePatchCmd.Flags().StringVar(&adminUserPropertiesUpdatePatchCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminUserPropertiesUpdatePatchCmd.Flags().StringVar(&adminUserPropertiesUpdatePatchCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	adminUserPropertiesUpdatePatchCmd.Flags().StringVar(&adminUserPropertiesUpdatePatchCmd_id, "id", "", "")
	adminUserPropertiesUpdatePatchCmd.Flags().StringVar(&adminUserPropertiesUpdatePatchCmd_status, "status", "", "")
	adminUserPropertiesUpdatePatchCmd.Flags().StringVar(&adminUserPropertiesUpdatePatchCmd_workspaceId, "workspaceId", "", "")
	adminUserPropertiesUpdatePatchCmd.MarkFlagRequired("id")
	adminUserPropertiesUpdatePatchCmd.MarkFlagRequired("status")
	adminUserPropertiesUpdatePatchCmd.MarkFlagRequired("workspaceId")
}
