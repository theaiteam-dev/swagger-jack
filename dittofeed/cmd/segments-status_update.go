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
	segmentsStatusUpdateCmdBody string
	segmentsStatusUpdateCmdBodyFile string
	segmentsStatusUpdateCmd_id string
	segmentsStatusUpdateCmd_status string
	segmentsStatusUpdateCmd_workspaceId string
)

var segmentsStatusUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if segmentsStatusUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(segmentsStatusUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			segmentsStatusUpdateCmdBody = string(fileData)
		}
		if segmentsStatusUpdateCmdBody != "" {
			if !json.Valid([]byte(segmentsStatusUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(segmentsStatusUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PATCH", "/api/segments/status", pathParams, queryParams, bodyObj)
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
		bodyMap["id"] = segmentsStatusUpdateCmd_id
		bodyMap["status"] = segmentsStatusUpdateCmd_status
		bodyMap["workspaceId"] = segmentsStatusUpdateCmd_workspaceId
		resp, err := c.Do("PATCH", "/api/segments/status", pathParams, queryParams, bodyMap)
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
	segmentsStatusCmd.AddCommand(segmentsStatusUpdateCmd)
	segmentsStatusUpdateCmd.Flags().StringVar(&segmentsStatusUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	segmentsStatusUpdateCmd.Flags().StringVar(&segmentsStatusUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	segmentsStatusUpdateCmd.Flags().StringVar(&segmentsStatusUpdateCmd_id, "id", "", "")
	segmentsStatusUpdateCmd.Flags().StringVar(&segmentsStatusUpdateCmd_status, "status", "", "")
	segmentsStatusUpdateCmd.Flags().StringVar(&segmentsStatusUpdateCmd_workspaceId, "workspaceId", "", "")
	segmentsStatusUpdateCmd.MarkFlagRequired("id")
	segmentsStatusUpdateCmd.MarkFlagRequired("status")
	segmentsStatusUpdateCmd.MarkFlagRequired("workspaceId")
}
