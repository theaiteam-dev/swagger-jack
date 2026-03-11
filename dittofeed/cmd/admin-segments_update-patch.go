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
	adminSegmentsUpdatePatchCmdBody string
	adminSegmentsUpdatePatchCmdBodyFile string
	adminSegmentsUpdatePatchCmd_id string
	adminSegmentsUpdatePatchCmd_status string
	adminSegmentsUpdatePatchCmd_workspaceId string
)

var adminSegmentsUpdatePatchCmd = &cobra.Command{
	Use: "update-patch",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if adminSegmentsUpdatePatchCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminSegmentsUpdatePatchCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminSegmentsUpdatePatchCmdBody = string(fileData)
		}
		if adminSegmentsUpdatePatchCmdBody != "" {
			if !json.Valid([]byte(adminSegmentsUpdatePatchCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminSegmentsUpdatePatchCmdBody), &bodyObj)
			resp, err := c.Do("PATCH", "/api/admin/segments/status", pathParams, queryParams, bodyObj)
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
		bodyMap["id"] = adminSegmentsUpdatePatchCmd_id
		bodyMap["status"] = adminSegmentsUpdatePatchCmd_status
		bodyMap["workspaceId"] = adminSegmentsUpdatePatchCmd_workspaceId
		resp, err := c.Do("PATCH", "/api/admin/segments/status", pathParams, queryParams, bodyMap)
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
	adminSegmentsCmd.AddCommand(adminSegmentsUpdatePatchCmd)
	adminSegmentsUpdatePatchCmd.Flags().StringVar(&adminSegmentsUpdatePatchCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminSegmentsUpdatePatchCmd.Flags().StringVar(&adminSegmentsUpdatePatchCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	adminSegmentsUpdatePatchCmd.Flags().StringVar(&adminSegmentsUpdatePatchCmd_id, "id", "", "")
	adminSegmentsUpdatePatchCmd.Flags().StringVar(&adminSegmentsUpdatePatchCmd_status, "status", "", "")
	adminSegmentsUpdatePatchCmd.Flags().StringVar(&adminSegmentsUpdatePatchCmd_workspaceId, "workspaceId", "", "")
	adminSegmentsUpdatePatchCmd.MarkFlagRequired("id")
	adminSegmentsUpdatePatchCmd.MarkFlagRequired("status")
	adminSegmentsUpdatePatchCmd.MarkFlagRequired("workspaceId")
}
