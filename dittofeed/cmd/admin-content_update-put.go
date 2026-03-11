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
	adminContentUpdatePutCmdBody string
	adminContentUpdatePutCmdBodyFile string
	adminContentUpdatePutCmd_resourceType string
	adminContentUpdatePutCmd_workspaceId string
	adminContentUpdatePutCmd_definition string
	adminContentUpdatePutCmd_draft string
	adminContentUpdatePutCmd_id string
	adminContentUpdatePutCmd_name string
)

var adminContentUpdatePutCmd = &cobra.Command{
	Use: "update-put",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if adminContentUpdatePutCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminContentUpdatePutCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminContentUpdatePutCmdBody = string(fileData)
		}
		if adminContentUpdatePutCmdBody != "" {
			if !json.Valid([]byte(adminContentUpdatePutCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminContentUpdatePutCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/admin/content/templates", pathParams, queryParams, bodyObj)
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
		bodyMap["resourceType"] = adminContentUpdatePutCmd_resourceType
		bodyMap["workspaceId"] = adminContentUpdatePutCmd_workspaceId
		bodyMap["definition"] = adminContentUpdatePutCmd_definition
		bodyMap["draft"] = adminContentUpdatePutCmd_draft
		bodyMap["id"] = adminContentUpdatePutCmd_id
		bodyMap["name"] = adminContentUpdatePutCmd_name
		resp, err := c.Do("PUT", "/api/admin/content/templates", pathParams, queryParams, bodyMap)
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
	adminContentCmd.AddCommand(adminContentUpdatePutCmd)
	adminContentUpdatePutCmd.Flags().StringVar(&adminContentUpdatePutCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminContentUpdatePutCmd.Flags().StringVar(&adminContentUpdatePutCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	adminContentUpdatePutCmd.Flags().StringVar(&adminContentUpdatePutCmd_resourceType, "resourceType", "", "")
	adminContentUpdatePutCmd.Flags().StringVar(&adminContentUpdatePutCmd_workspaceId, "workspaceId", "", "")
	adminContentUpdatePutCmd.Flags().StringVar(&adminContentUpdatePutCmd_definition, "definition", "", "")
	adminContentUpdatePutCmd.Flags().StringVar(&adminContentUpdatePutCmd_draft, "draft", "", "")
	adminContentUpdatePutCmd.Flags().StringVar(&adminContentUpdatePutCmd_id, "id", "", "")
	adminContentUpdatePutCmd.Flags().StringVar(&adminContentUpdatePutCmd_name, "name", "", "")
	adminContentUpdatePutCmd.MarkFlagRequired("workspaceId")
	adminContentUpdatePutCmd.MarkFlagRequired("name")
}
