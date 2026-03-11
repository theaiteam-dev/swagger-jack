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
	internalApiLWorkspacesUpdatePutCmdBody string
	internalApiLWorkspacesUpdatePutCmdBodyFile string
	internalApiLWorkspacesUpdatePutCmd_domain string
	internalApiLWorkspacesUpdatePutCmd_externalId string
	internalApiLWorkspacesUpdatePutCmd_name string
	internalApiLWorkspacesUpdatePutCmd_workspaceId string
)

var internalApiLWorkspacesUpdatePutCmd = &cobra.Command{
	Use: "update-put",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if internalApiLWorkspacesUpdatePutCmdBodyFile != "" {
			fileData, err := os.ReadFile(internalApiLWorkspacesUpdatePutCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			internalApiLWorkspacesUpdatePutCmdBody = string(fileData)
		}
		if internalApiLWorkspacesUpdatePutCmdBody != "" {
			if !json.Valid([]byte(internalApiLWorkspacesUpdatePutCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(internalApiLWorkspacesUpdatePutCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/internal-api-l/workspaces/child", pathParams, queryParams, bodyObj)
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
		bodyMap["domain"] = internalApiLWorkspacesUpdatePutCmd_domain
		bodyMap["externalId"] = internalApiLWorkspacesUpdatePutCmd_externalId
		bodyMap["name"] = internalApiLWorkspacesUpdatePutCmd_name
		bodyMap["workspaceId"] = internalApiLWorkspacesUpdatePutCmd_workspaceId
		resp, err := c.Do("PUT", "/internal-api-l/workspaces/child", pathParams, queryParams, bodyMap)
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
	internalApiLWorkspacesCmd.AddCommand(internalApiLWorkspacesUpdatePutCmd)
	internalApiLWorkspacesUpdatePutCmd.Flags().StringVar(&internalApiLWorkspacesUpdatePutCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	internalApiLWorkspacesUpdatePutCmd.Flags().StringVar(&internalApiLWorkspacesUpdatePutCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	internalApiLWorkspacesUpdatePutCmd.Flags().StringVar(&internalApiLWorkspacesUpdatePutCmd_domain, "domain", "", "")
	internalApiLWorkspacesUpdatePutCmd.Flags().StringVar(&internalApiLWorkspacesUpdatePutCmd_externalId, "externalId", "", "")
	internalApiLWorkspacesUpdatePutCmd.Flags().StringVar(&internalApiLWorkspacesUpdatePutCmd_name, "name", "", "")
	internalApiLWorkspacesUpdatePutCmd.Flags().StringVar(&internalApiLWorkspacesUpdatePutCmd_workspaceId, "workspaceId", "", "")
	internalApiLWorkspacesUpdatePutCmd.MarkFlagRequired("externalId")
	internalApiLWorkspacesUpdatePutCmd.MarkFlagRequired("name")
	internalApiLWorkspacesUpdatePutCmd.MarkFlagRequired("workspaceId")
}
