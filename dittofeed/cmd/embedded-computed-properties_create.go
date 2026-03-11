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
	embeddedComputedPropertiesCreateCmdBody string
	embeddedComputedPropertiesCreateCmdBodyFile string
	embeddedComputedPropertiesCreateCmd_workspaceId string
)

var embeddedComputedPropertiesCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if embeddedComputedPropertiesCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(embeddedComputedPropertiesCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			embeddedComputedPropertiesCreateCmdBody = string(fileData)
		}
		if embeddedComputedPropertiesCreateCmdBody != "" {
			if !json.Valid([]byte(embeddedComputedPropertiesCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(embeddedComputedPropertiesCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api-l/embedded/computed-properties/trigger-recompute", pathParams, queryParams, bodyObj)
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
		bodyMap["workspaceId"] = embeddedComputedPropertiesCreateCmd_workspaceId
		resp, err := c.Do("POST", "/api-l/embedded/computed-properties/trigger-recompute", pathParams, queryParams, bodyMap)
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
	embeddedComputedPropertiesCmd.AddCommand(embeddedComputedPropertiesCreateCmd)
	embeddedComputedPropertiesCreateCmd.Flags().StringVar(&embeddedComputedPropertiesCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	embeddedComputedPropertiesCreateCmd.Flags().StringVar(&embeddedComputedPropertiesCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	embeddedComputedPropertiesCreateCmd.Flags().StringVar(&embeddedComputedPropertiesCreateCmd_workspaceId, "workspaceId", "", "")
	embeddedComputedPropertiesCreateCmd.MarkFlagRequired("workspaceId")
}
