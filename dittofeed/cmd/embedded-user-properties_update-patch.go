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
	embeddedUserPropertiesUpdatePatchCmdBody string
	embeddedUserPropertiesUpdatePatchCmdBodyFile string
	embeddedUserPropertiesUpdatePatchCmd_id string
	embeddedUserPropertiesUpdatePatchCmd_status string
	embeddedUserPropertiesUpdatePatchCmd_workspaceId string
)

var embeddedUserPropertiesUpdatePatchCmd = &cobra.Command{
	Use: "update-patch",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if embeddedUserPropertiesUpdatePatchCmdBodyFile != "" {
			fileData, err := os.ReadFile(embeddedUserPropertiesUpdatePatchCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			embeddedUserPropertiesUpdatePatchCmdBody = string(fileData)
		}
		if embeddedUserPropertiesUpdatePatchCmdBody != "" {
			if !json.Valid([]byte(embeddedUserPropertiesUpdatePatchCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(embeddedUserPropertiesUpdatePatchCmdBody), &bodyObj)
			resp, err := c.Do("PATCH", "/api-l/embedded/user-properties/status", pathParams, queryParams, bodyObj)
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
		bodyMap["id"] = embeddedUserPropertiesUpdatePatchCmd_id
		bodyMap["status"] = embeddedUserPropertiesUpdatePatchCmd_status
		bodyMap["workspaceId"] = embeddedUserPropertiesUpdatePatchCmd_workspaceId
		resp, err := c.Do("PATCH", "/api-l/embedded/user-properties/status", pathParams, queryParams, bodyMap)
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
	embeddedUserPropertiesCmd.AddCommand(embeddedUserPropertiesUpdatePatchCmd)
	embeddedUserPropertiesUpdatePatchCmd.Flags().StringVar(&embeddedUserPropertiesUpdatePatchCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	embeddedUserPropertiesUpdatePatchCmd.Flags().StringVar(&embeddedUserPropertiesUpdatePatchCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	embeddedUserPropertiesUpdatePatchCmd.Flags().StringVar(&embeddedUserPropertiesUpdatePatchCmd_id, "id", "", "")
	embeddedUserPropertiesUpdatePatchCmd.Flags().StringVar(&embeddedUserPropertiesUpdatePatchCmd_status, "status", "", "")
	embeddedUserPropertiesUpdatePatchCmd.Flags().StringVar(&embeddedUserPropertiesUpdatePatchCmd_workspaceId, "workspaceId", "", "")
	embeddedUserPropertiesUpdatePatchCmd.MarkFlagRequired("id")
	embeddedUserPropertiesUpdatePatchCmd.MarkFlagRequired("status")
	embeddedUserPropertiesUpdatePatchCmd.MarkFlagRequired("workspaceId")
}
