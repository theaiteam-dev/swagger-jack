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
	adminKeysCreateCmdBody string
	adminKeysCreateCmdBodyFile string
	adminKeysCreateCmd_workspaceId string
	adminKeysCreateCmd_name string
)

var adminKeysCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if adminKeysCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminKeysCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminKeysCreateCmdBody = string(fileData)
		}
		if adminKeysCreateCmdBody != "" {
			if !json.Valid([]byte(adminKeysCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminKeysCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/admin-keys/", pathParams, queryParams, bodyObj)
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
		bodyMap["workspaceId"] = adminKeysCreateCmd_workspaceId
		bodyMap["name"] = adminKeysCreateCmd_name
		resp, err := c.Do("POST", "/api/admin-keys/", pathParams, queryParams, bodyMap)
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
	adminKeysCmd.AddCommand(adminKeysCreateCmd)
	adminKeysCreateCmd.Flags().StringVar(&adminKeysCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminKeysCreateCmd.Flags().StringVar(&adminKeysCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	adminKeysCreateCmd.Flags().StringVar(&adminKeysCreateCmd_workspaceId, "workspaceId", "", "")
	adminKeysCreateCmd.Flags().StringVar(&adminKeysCreateCmd_name, "name", "", "")
	adminKeysCreateCmd.MarkFlagRequired("workspaceId")
	adminKeysCreateCmd.MarkFlagRequired("name")
}
