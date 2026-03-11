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
	computedPropertiesTriggerRecomputeCreateCmdBody string
	computedPropertiesTriggerRecomputeCreateCmdBodyFile string
	computedPropertiesTriggerRecomputeCreateCmd_workspaceId string
)

var computedPropertiesTriggerRecomputeCreateCmd = &cobra.Command{
	Use: "create",
	Short: "",
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		baseURL, _ := cmd.Root().PersistentFlags().GetString("base-url")
		token := os.Getenv("DITTOFEED_TOKEN")
		c := client.NewClient(baseURL, token)
		pathParams := map[string]string{}
		queryParams := map[string]string{}
		if computedPropertiesTriggerRecomputeCreateCmdBodyFile != "" {
			fileData, err := os.ReadFile(computedPropertiesTriggerRecomputeCreateCmdBodyFile)
			if err != nil {
				return fmt.Errorf("reading body-file: %w", err)
			}
			if !json.Valid(fileData) {
				return fmt.Errorf("body-file does not contain valid JSON")
			}
			computedPropertiesTriggerRecomputeCreateCmdBody = string(fileData)
		}
		if computedPropertiesTriggerRecomputeCreateCmdBody != "" {
			if !json.Valid([]byte(computedPropertiesTriggerRecomputeCreateCmdBody)) {
				return fmt.Errorf("--body does not contain valid JSON")
			}
			var bodyObj interface{}
			_ = json.Unmarshal([]byte(computedPropertiesTriggerRecomputeCreateCmdBody), &bodyObj)
			resp, err := c.Do("POST", "/api/computed-properties/trigger-recompute", pathParams, queryParams, bodyObj)
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
		bodyMap["workspaceId"] = computedPropertiesTriggerRecomputeCreateCmd_workspaceId
		resp, err := c.Do("POST", "/api/computed-properties/trigger-recompute", pathParams, queryParams, bodyMap)
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
	computedPropertiesTriggerRecomputeCmd.AddCommand(computedPropertiesTriggerRecomputeCreateCmd)
	computedPropertiesTriggerRecomputeCreateCmd.Flags().StringVar(&computedPropertiesTriggerRecomputeCreateCmdBody, "body", "", "Raw JSON body (overrides individual flags)")
	computedPropertiesTriggerRecomputeCreateCmd.Flags().StringVar(&computedPropertiesTriggerRecomputeCreateCmdBodyFile, "body-file", "", "Path to JSON file to use as request body")
	computedPropertiesTriggerRecomputeCreateCmd.Flags().StringVar(&computedPropertiesTriggerRecomputeCreateCmd_workspaceId, "workspaceId", "", "")
	computedPropertiesTriggerRecomputeCreateCmd.MarkFlagRequired("workspaceId")
}
