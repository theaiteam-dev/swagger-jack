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
	adminWorkspacesUpdateCmdBody string
	adminWorkspacesUpdateCmdBodyFile string
	adminWorkspacesUpdateCmd_domain string
	adminWorkspacesUpdateCmd_externalId string
	adminWorkspacesUpdateCmd_name string
	adminWorkspacesUpdateCmd_workspaceId string
)

var adminWorkspacesUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if adminWorkspacesUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminWorkspacesUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminWorkspacesUpdateCmdBody = string(fileData)
		}
		if adminWorkspacesUpdateCmdBody != "" {
			if !json.Valid([]byte(adminWorkspacesUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminWorkspacesUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api-l/admin/workspaces/child", pathParams, queryParams, bodyObj)
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
		bodyMap["domain"] = adminWorkspacesUpdateCmd_domain
		bodyMap["externalId"] = adminWorkspacesUpdateCmd_externalId
		bodyMap["name"] = adminWorkspacesUpdateCmd_name
		bodyMap["workspaceId"] = adminWorkspacesUpdateCmd_workspaceId
		resp, err := c.Do("PUT", "/api-l/admin/workspaces/child", pathParams, queryParams, bodyMap)
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
	adminWorkspacesCmd.AddCommand(adminWorkspacesUpdateCmd)
	adminWorkspacesUpdateCmd.Flags().StringVar(&adminWorkspacesUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminWorkspacesUpdateCmd.Flags().StringVar(&adminWorkspacesUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	adminWorkspacesUpdateCmd.Flags().StringVar(&adminWorkspacesUpdateCmd_domain, "domain", "", "")
	adminWorkspacesUpdateCmd.Flags().StringVar(&adminWorkspacesUpdateCmd_externalId, "externalId", "", "")
	adminWorkspacesUpdateCmd.Flags().StringVar(&adminWorkspacesUpdateCmd_name, "name", "", "")
	adminWorkspacesUpdateCmd.Flags().StringVar(&adminWorkspacesUpdateCmd_workspaceId, "workspaceId", "", "")
	adminWorkspacesUpdateCmd.MarkFlagRequired("externalId")
	adminWorkspacesUpdateCmd.MarkFlagRequired("name")
	adminWorkspacesUpdateCmd.MarkFlagRequired("workspaceId")
}
