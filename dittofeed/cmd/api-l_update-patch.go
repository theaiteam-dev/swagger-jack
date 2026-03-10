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
	apiLUpdatePatchCmdBody string
	apiLUpdatePatchCmdBodyFile string
	apiLUpdatePatchCmd_id string
	apiLUpdatePatchCmd_status string
	apiLUpdatePatchCmd_workspaceId string
)

var apiLUpdatePatchCmd = &cobra.Command{
	Use: "update-patch",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if apiLUpdatePatchCmdBodyFile != "" {
			fileData, err := os.ReadFile(apiLUpdatePatchCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			apiLUpdatePatchCmdBody = string(fileData)
		}
		if apiLUpdatePatchCmdBody != "" {
			if !json.Valid([]byte(apiLUpdatePatchCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(apiLUpdatePatchCmdBody), &bodyObj)
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
		bodyMap["id"] = apiLUpdatePatchCmd_id
		bodyMap["status"] = apiLUpdatePatchCmd_status
		bodyMap["workspaceId"] = apiLUpdatePatchCmd_workspaceId
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
	apiLCmd.AddCommand(apiLUpdatePatchCmd)
	apiLUpdatePatchCmd.Flags().StringVar(&apiLUpdatePatchCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	apiLUpdatePatchCmd.Flags().StringVar(&apiLUpdatePatchCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	apiLUpdatePatchCmd.Flags().StringVar(&apiLUpdatePatchCmd_id, "id", "", "")
	apiLUpdatePatchCmd.Flags().StringVar(&apiLUpdatePatchCmd_status, "status", "", "")
	apiLUpdatePatchCmd.Flags().StringVar(&apiLUpdatePatchCmd_workspaceId, "workspaceId", "", "")
	apiLUpdatePatchCmd.MarkFlagRequired("id")
	apiLUpdatePatchCmd.MarkFlagRequired("status")
	apiLUpdatePatchCmd.MarkFlagRequired("workspaceId")
}
