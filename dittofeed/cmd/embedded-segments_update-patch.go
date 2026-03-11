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
	embeddedSegmentsUpdatePatchCmdBody string
	embeddedSegmentsUpdatePatchCmdBodyFile string
	embeddedSegmentsUpdatePatchCmd_id string
	embeddedSegmentsUpdatePatchCmd_status string
	embeddedSegmentsUpdatePatchCmd_workspaceId string
)

var embeddedSegmentsUpdatePatchCmd = &cobra.Command{
	Use: "update-patch",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if embeddedSegmentsUpdatePatchCmdBodyFile != "" {
			fileData, err := os.ReadFile(embeddedSegmentsUpdatePatchCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			embeddedSegmentsUpdatePatchCmdBody = string(fileData)
		}
		if embeddedSegmentsUpdatePatchCmdBody != "" {
			if !json.Valid([]byte(embeddedSegmentsUpdatePatchCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(embeddedSegmentsUpdatePatchCmdBody), &bodyObj)
			resp, err := c.Do("PATCH", "/api-l/embedded/segments/status", pathParams, queryParams, bodyObj)
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
		bodyMap["id"] = embeddedSegmentsUpdatePatchCmd_id
		bodyMap["status"] = embeddedSegmentsUpdatePatchCmd_status
		bodyMap["workspaceId"] = embeddedSegmentsUpdatePatchCmd_workspaceId
		resp, err := c.Do("PATCH", "/api-l/embedded/segments/status", pathParams, queryParams, bodyMap)
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
	embeddedSegmentsCmd.AddCommand(embeddedSegmentsUpdatePatchCmd)
	embeddedSegmentsUpdatePatchCmd.Flags().StringVar(&embeddedSegmentsUpdatePatchCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	embeddedSegmentsUpdatePatchCmd.Flags().StringVar(&embeddedSegmentsUpdatePatchCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	embeddedSegmentsUpdatePatchCmd.Flags().StringVar(&embeddedSegmentsUpdatePatchCmd_id, "id", "", "")
	embeddedSegmentsUpdatePatchCmd.Flags().StringVar(&embeddedSegmentsUpdatePatchCmd_status, "status", "", "")
	embeddedSegmentsUpdatePatchCmd.Flags().StringVar(&embeddedSegmentsUpdatePatchCmd_workspaceId, "workspaceId", "", "")
	embeddedSegmentsUpdatePatchCmd.MarkFlagRequired("id")
	embeddedSegmentsUpdatePatchCmd.MarkFlagRequired("status")
	embeddedSegmentsUpdatePatchCmd.MarkFlagRequired("workspaceId")
}
