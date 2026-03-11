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
	broadcastsTriggerUpdateCmdBody string
	broadcastsTriggerUpdateCmdBodyFile string
	broadcastsTriggerUpdateCmd_workspaceId string
	broadcastsTriggerUpdateCmd_id string
)

var broadcastsTriggerUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if broadcastsTriggerUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(broadcastsTriggerUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			broadcastsTriggerUpdateCmdBody = string(fileData)
		}
		if broadcastsTriggerUpdateCmdBody != "" {
			if !json.Valid([]byte(broadcastsTriggerUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(broadcastsTriggerUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/broadcasts/trigger", pathParams, queryParams, bodyObj)
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
		bodyMap["workspaceId"] = broadcastsTriggerUpdateCmd_workspaceId
		bodyMap["id"] = broadcastsTriggerUpdateCmd_id
		resp, err := c.Do("PUT", "/api/broadcasts/trigger", pathParams, queryParams, bodyMap)
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
	broadcastsTriggerCmd.AddCommand(broadcastsTriggerUpdateCmd)
	broadcastsTriggerUpdateCmd.Flags().StringVar(&broadcastsTriggerUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	broadcastsTriggerUpdateCmd.Flags().StringVar(&broadcastsTriggerUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	broadcastsTriggerUpdateCmd.Flags().StringVar(&broadcastsTriggerUpdateCmd_workspaceId, "workspaceId", "", "")
	broadcastsTriggerUpdateCmd.Flags().StringVar(&broadcastsTriggerUpdateCmd_id, "id", "", "")
	broadcastsTriggerUpdateCmd.MarkFlagRequired("workspaceId")
	broadcastsTriggerUpdateCmd.MarkFlagRequired("id")
}
