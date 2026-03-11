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
	adminComponentConfigurationsUpdateCmdBody string
	adminComponentConfigurationsUpdateCmdBodyFile string
	adminComponentConfigurationsUpdateCmd_definition string
	adminComponentConfigurationsUpdateCmd_id string
	adminComponentConfigurationsUpdateCmd_name string
	adminComponentConfigurationsUpdateCmd_workspaceId string
)

var adminComponentConfigurationsUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if adminComponentConfigurationsUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(adminComponentConfigurationsUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			adminComponentConfigurationsUpdateCmdBody = string(fileData)
		}
		if adminComponentConfigurationsUpdateCmdBody != "" {
			if !json.Valid([]byte(adminComponentConfigurationsUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(adminComponentConfigurationsUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/admin/component-configurations/", pathParams, queryParams, bodyObj)
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
		bodyMap["definition"] = adminComponentConfigurationsUpdateCmd_definition
		bodyMap["id"] = adminComponentConfigurationsUpdateCmd_id
		bodyMap["name"] = adminComponentConfigurationsUpdateCmd_name
		bodyMap["workspaceId"] = adminComponentConfigurationsUpdateCmd_workspaceId
		resp, err := c.Do("PUT", "/api/admin/component-configurations/", pathParams, queryParams, bodyMap)
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
	adminComponentConfigurationsCmd.AddCommand(adminComponentConfigurationsUpdateCmd)
	adminComponentConfigurationsUpdateCmd.Flags().StringVar(&adminComponentConfigurationsUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	adminComponentConfigurationsUpdateCmd.Flags().StringVar(&adminComponentConfigurationsUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	adminComponentConfigurationsUpdateCmd.Flags().StringVar(&adminComponentConfigurationsUpdateCmd_definition, "definition", "", "")
	adminComponentConfigurationsUpdateCmd.Flags().StringVar(&adminComponentConfigurationsUpdateCmd_id, "id", "", "")
	adminComponentConfigurationsUpdateCmd.Flags().StringVar(&adminComponentConfigurationsUpdateCmd_name, "name", "", "")
	adminComponentConfigurationsUpdateCmd.Flags().StringVar(&adminComponentConfigurationsUpdateCmd_workspaceId, "workspaceId", "", "")
	adminComponentConfigurationsUpdateCmd.MarkFlagRequired("name")
	adminComponentConfigurationsUpdateCmd.MarkFlagRequired("workspaceId")
}
