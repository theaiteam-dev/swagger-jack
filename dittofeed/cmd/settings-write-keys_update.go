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
	settingsWriteKeysUpdateCmdBody string
	settingsWriteKeysUpdateCmdBodyFile string
	settingsWriteKeysUpdateCmd_workspaceId string
	settingsWriteKeysUpdateCmd_writeKeyName string
)

var settingsWriteKeysUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if settingsWriteKeysUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(settingsWriteKeysUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			settingsWriteKeysUpdateCmdBody = string(fileData)
		}
		if settingsWriteKeysUpdateCmdBody != "" {
			if !json.Valid([]byte(settingsWriteKeysUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(settingsWriteKeysUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/settings/write-keys", pathParams, queryParams, bodyObj)
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
		bodyMap["workspaceId"] = settingsWriteKeysUpdateCmd_workspaceId
		bodyMap["writeKeyName"] = settingsWriteKeysUpdateCmd_writeKeyName
		resp, err := c.Do("PUT", "/api/settings/write-keys", pathParams, queryParams, bodyMap)
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
	settingsWriteKeysCmd.AddCommand(settingsWriteKeysUpdateCmd)
	settingsWriteKeysUpdateCmd.Flags().StringVar(&settingsWriteKeysUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	settingsWriteKeysUpdateCmd.Flags().StringVar(&settingsWriteKeysUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	settingsWriteKeysUpdateCmd.Flags().StringVar(&settingsWriteKeysUpdateCmd_workspaceId, "workspaceId", "", "")
	settingsWriteKeysUpdateCmd.Flags().StringVar(&settingsWriteKeysUpdateCmd_writeKeyName, "writeKeyName", "", "")
	settingsWriteKeysUpdateCmd.MarkFlagRequired("workspaceId")
	settingsWriteKeysUpdateCmd.MarkFlagRequired("writeKeyName")
}
