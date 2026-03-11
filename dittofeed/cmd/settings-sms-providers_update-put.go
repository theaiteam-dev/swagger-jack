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
	settingsSmsProvidersUpdatePutCmdBody string
	settingsSmsProvidersUpdatePutCmdBodyFile string
	settingsSmsProvidersUpdatePutCmd_config string
	settingsSmsProvidersUpdatePutCmd_setDefault bool
	settingsSmsProvidersUpdatePutCmd_workspaceId string
)

var settingsSmsProvidersUpdatePutCmd = &cobra.Command{
	Use: "update-put",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if settingsSmsProvidersUpdatePutCmdBodyFile != "" {
			fileData, err := os.ReadFile(settingsSmsProvidersUpdatePutCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			settingsSmsProvidersUpdatePutCmdBody = string(fileData)
		}
		if settingsSmsProvidersUpdatePutCmdBody != "" {
			if !json.Valid([]byte(settingsSmsProvidersUpdatePutCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(settingsSmsProvidersUpdatePutCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/settings/sms-providers", pathParams, queryParams, bodyObj)
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
		bodyMap["config"] = settingsSmsProvidersUpdatePutCmd_config
		bodyMap["setDefault"] = settingsSmsProvidersUpdatePutCmd_setDefault
		bodyMap["workspaceId"] = settingsSmsProvidersUpdatePutCmd_workspaceId
		resp, err := c.Do("PUT", "/api/settings/sms-providers", pathParams, queryParams, bodyMap)
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
	settingsSmsProvidersCmd.AddCommand(settingsSmsProvidersUpdatePutCmd)
	settingsSmsProvidersUpdatePutCmd.Flags().StringVar(&settingsSmsProvidersUpdatePutCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	settingsSmsProvidersUpdatePutCmd.Flags().StringVar(&settingsSmsProvidersUpdatePutCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	settingsSmsProvidersUpdatePutCmd.Flags().StringVar(&settingsSmsProvidersUpdatePutCmd_config, "config", "", "")
	settingsSmsProvidersUpdatePutCmd.Flags().BoolVar(&settingsSmsProvidersUpdatePutCmd_setDefault, "setDefault", false, "")
	settingsSmsProvidersUpdatePutCmd.Flags().StringVar(&settingsSmsProvidersUpdatePutCmd_workspaceId, "workspaceId", "", "")
	settingsSmsProvidersUpdatePutCmd.MarkFlagRequired("config")
	settingsSmsProvidersUpdatePutCmd.MarkFlagRequired("workspaceId")
}
