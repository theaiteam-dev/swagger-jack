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
	secretsUpdateCmdBody string
	secretsUpdateCmdBodyFile string
	secretsUpdateCmd_value string
	secretsUpdateCmd_workspaceId string
	secretsUpdateCmd_configValue string
	secretsUpdateCmd_name string
)

var secretsUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if secretsUpdateCmdBodyFile != "" {
			fileData, err := os.ReadFile(secretsUpdateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			secretsUpdateCmdBody = string(fileData)
		}
		if secretsUpdateCmdBody != "" {
			if !json.Valid([]byte(secretsUpdateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(secretsUpdateCmdBody), &bodyObj)
			resp, err := c.Do("PUT", "/api/secrets/", pathParams, queryParams, bodyObj)
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
		bodyMap["value"] = secretsUpdateCmd_value
		bodyMap["workspaceId"] = secretsUpdateCmd_workspaceId
		bodyMap["configValue"] = secretsUpdateCmd_configValue
		bodyMap["name"] = secretsUpdateCmd_name
		resp, err := c.Do("PUT", "/api/secrets/", pathParams, queryParams, bodyMap)
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
	secretsCmd.AddCommand(secretsUpdateCmd)
	secretsUpdateCmd.Flags().StringVar(&secretsUpdateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	secretsUpdateCmd.Flags().StringVar(&secretsUpdateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	secretsUpdateCmd.Flags().StringVar(&secretsUpdateCmd_value, "value", "", "")
	secretsUpdateCmd.Flags().StringVar(&secretsUpdateCmd_workspaceId, "workspaceId", "", "")
	secretsUpdateCmd.Flags().StringVar(&secretsUpdateCmd_configValue, "configValue", "", "")
	secretsUpdateCmd.Flags().StringVar(&secretsUpdateCmd_name, "name", "", "")
	secretsUpdateCmd.MarkFlagRequired("workspaceId")
	secretsUpdateCmd.MarkFlagRequired("name")
}
