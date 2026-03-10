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
	apiUpdatePatchCmdBody string
	apiUpdatePatchCmdBodyFile string
	apiUpdatePatchCmd_id string
	apiUpdatePatchCmd_status string
	apiUpdatePatchCmd_workspaceId string
)

var apiUpdatePatchCmd = &cobra.Command{
	Use: "update-patch",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if apiUpdatePatchCmdBodyFile != "" {
			fileData, err := os.ReadFile(apiUpdatePatchCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			apiUpdatePatchCmdBody = string(fileData)
		}
		if apiUpdatePatchCmdBody != "" {
			if !json.Valid([]byte(apiUpdatePatchCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(apiUpdatePatchCmdBody), &bodyObj)
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
		bodyMap["id"] = apiUpdatePatchCmd_id
		bodyMap["status"] = apiUpdatePatchCmd_status
		bodyMap["workspaceId"] = apiUpdatePatchCmd_workspaceId
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
	apiCmd.AddCommand(apiUpdatePatchCmd)
	apiUpdatePatchCmd.Flags().StringVar(&apiUpdatePatchCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	apiUpdatePatchCmd.Flags().StringVar(&apiUpdatePatchCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	apiUpdatePatchCmd.Flags().StringVar(&apiUpdatePatchCmd_id, "id", "", "")
	apiUpdatePatchCmd.Flags().StringVar(&apiUpdatePatchCmd_status, "status", "", "")
	apiUpdatePatchCmd.Flags().StringVar(&apiUpdatePatchCmd_workspaceId, "workspaceId", "", "")
	apiUpdatePatchCmd.MarkFlagRequired("id")
	apiUpdatePatchCmd.MarkFlagRequired("status")
	apiUpdatePatchCmd.MarkFlagRequired("workspaceId")
}
